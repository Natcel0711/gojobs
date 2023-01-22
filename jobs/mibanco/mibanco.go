package jobs

import (
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
)

func MiBanco() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not launch playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch Chromium: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://www.popular.com/", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	err = page.Click("#btn-close-modalChooseRegion")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(3 * time.Second)
	err = page.Click("#btn-login")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(3 * time.Second)
	err = page.Click("#btn-login-mibanco")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(3 * time.Second)
	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("foo.png"),
	}); err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
