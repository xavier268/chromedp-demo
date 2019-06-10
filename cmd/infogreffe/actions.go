package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

// Login with the provided credentials in cliParam
func Login(ctx context.Context) error {

	c, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Second)
	defer cancel()

	selForm := "//form[@name='loginForm']"
	selEmail := selForm + "//input[@name='email']"
	selPass := selForm + "//input[@name='pwd']"
	selCheckbox := selForm + "//input[@type='checkbox']"
	selValider := selForm + "//input[@type='image']"

	var location string

	err := chromedp.Run(c,
		chromedp.Navigate(cliParam.url),
		chromedp.WaitVisible(selEmail),
		chromedp.SendKeys(selEmail, cliParam.user),
		chromedp.SendKeys(selPass, cliParam.passwd),
		chromedp.Click(selCheckbox),
		chromedp.Click(selValider),
		chromedp.WaitNotPresent(selForm),
		chromedp.Location(&location),
	)

	if err != nil {
		fmt.Println("Error while authenticating ?! ", err)
	}

	if err == nil {
		fmt.Println("Authenticated as ", cliParam.user)
		fmt.Println("Now on : ", location)
	}

	return err
}

// RequestMails all documents from the given file by mail
// We must be authenticated first.
// Role is the greffe number
func RequestMails(ctx context.Context, role string) error {

	// Get the page related to that file,
	// then click on each mail link to trigger sending

	c, cancel := context.WithTimeout(ctx, time.Duration(30)*time.Second)
	defer cancel()

	var location string

	selRole := "//input[@name='role']"

	if err := chromedp.Run(c, chromedp.Location(&location)); err != nil {
		fmt.Println("Unable to read page location ?!", err)
		return err
	}

	if location != "http://www2.infogreffe.fr/infogreffe/judIndex.do" {
		err := errors.New("Unexpected location : " + location + " Did you authenticate first ?")
		fmt.Println(err)
		return err
	}

	err := chromedp.Run(c,
		chromedp.Navigate("/infogreffe/selectRechercheAffaireAction.do?search=0"),
		chromedp.WaitVisible(selRole),
		chromedp.SendKeys(selRole, role),
		chromedp.Submit(selRole),
		chromedp.WaitNotPresent(selRole),
		chromedp.Navigate("/infogreffe/listeAffaireDocsNumerises.do?indiceaff=0"),
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// TO DO - find all related links and click them ...

	return err
}
