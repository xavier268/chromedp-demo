package main

import (
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	if !processCli() {
		return
	}

	ctx, cancel := getBrowserCtx()
	defer cancel()

	if err := chromedp.Run(ctx); err != nil {
		fmt.Println("Error trying to access browser ?", err)
		return
	}

	if err := Login(ctx); err != nil {
		fmt.Println("Error trying to login ?", err)
		return
	}

	if err := RequestMails(ctx, "2019F00260"); err != nil {
		fmt.Println("Error trying to retirve mails ?", err)
		return
	}

	//debugging
	fmt.Println("Finished, waiting a little bit for debugging ...")
	time.Sleep(time.Duration(100) * time.Second)
	fmt.Println("Done.")
}
