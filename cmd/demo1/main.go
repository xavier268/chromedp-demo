package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	const (
		myurl = "https://news.google.com"
	)

	// Get a default (visible) allocator
	ctx, cancel := getVisibleAllocator()
	defer cancel()

	// Use allocator to get a browser context (MANDATORY - you cannot USE an allocator ctx directly )
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// Set timeout, just in case ...
	ctx, cancel = context.WithTimeout(ctx, 3*time.Minute)
	fmt.Println(time.Now(), " Timout started for 3 minutes")
	defer cancel()

	// ensure that the browser process is started
	checkBrowserStarted(ctx)

	// Actual processing done here

	getSelectedHeadlines(ctx, myurl)
	getAllHeadlines(ctx, myurl)
}
