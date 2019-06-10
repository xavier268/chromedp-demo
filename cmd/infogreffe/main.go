package main

import (
	"fmt"

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

}
