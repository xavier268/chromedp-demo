package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

// Login with the provided credentials in cliParam
func Login(ctx context.Context) {

	selForm := "//form[@name='loginForm']"
	selEmail := selForm + "//input[@name='email']"
	selPass := selForm + "//input[@name='pwd']"
	selCheckbox := selForm + "//input[@type='checkbox']"
	selValider := selForm + "//input[@type='image']"

	chromedp.Run(ctx,
		chromedp.Navigate(cliParam.url),
		chromedp.WaitVisible(selEmail),
		chromedp.SendKeys(selEmail, cliParam.user),
		chromedp.SendKeys(selPass, cliParam.passwd),
		chromedp.Click(selCheckbox),
		chromedp.Click(selValider),
		chromedp.WaitNotPresent(selForm),
	)
	fmt.Println("Authenticated as ", cliParam.user)
	time.Sleep(time.Duration(cliParam.minutes) * time.Minute)

}
