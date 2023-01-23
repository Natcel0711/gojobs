package jobs

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/playwright-community/playwright-go"
)

func MiBanco() {
	err := godotenv.Load(".env")
	errCheck(err)
	pw, err := playwright.Run()
	errCheck(err)
	browser, err := pw.Chromium.Launch()
	errCheck(err)
	useragent := "Mozilla/5.0 (Windows Phone 10.0; Android 6.0.1; Microsoft; Lumia 950) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Mobile Safari/537.36 Edge/15.14977"
	page, err := browser.NewPage(playwright.BrowserNewContextOptions{
		UserAgent: &useragent,
	})
	errCheck(err)
	if _, err = page.Goto("https://www.bancopopular.com/cibp-web/actions/login#_home", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	takeScreenshot(&page)
	err = page.Fill("#username", os.Getenv("USERNAME"))
	errCheck(err)
	time.Sleep(1 * time.Second)
	err = page.Keyboard().Press("Enter")
	errCheck(err)
	takeScreenshot(&page)
	err = page.Fill("#answer", os.Getenv("PASSWORD"))
	errCheck(err)
	err = page.Keyboard().Press("Enter")
	errCheck(err)
	time.Sleep(5 * time.Second)
	takeScreenshot(&page)
	if _, err = page.Goto("https://www.bancopopular.com/cibp-web/actions/login#_accounts", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	balanceLabel, err := page.Locator(".balance")
	errCheck(err)
	balance, err := balanceLabel.InnerText()
	errCheck(err)
	log.Println(balance)
	err = balanceLabel.Click()
	errCheck(err)
	balanceAnchor, err := page.Locator("a[text=\"Current Statement\"]")
	errCheck(err)
	err = balanceAnchor.Click()
	errCheck(err)
	takeScreenshot(&page)
	if err = browser.Close(); err != nil {
		errCheck(err)
	}
	if err = pw.Stop(); err != nil {
		errCheck(err)
	}
}

func errCheck(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

func takeScreenshot(pageP *playwright.Page) {
	page := *pageP
	log.Println("Screenshotting...")
	if _, err := page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("foo.png"),
	}); err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	}
}
