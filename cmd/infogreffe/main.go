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

	Login(ctx)

	//debugging
	time.Sleep(time.Duration(3) * time.Minute)
}
