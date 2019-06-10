package main

import (
	"context"
	"flag"
	"io/ioutil"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

// getBrowser gets a browser context to run on
func getBrowserCtx() (context.Context, context.CancelFunc) {

	if !flag.Parsed() {
		panic("Calling getBrowser befaore parsing of cli params are available ?!")
	}

	dir, err := ioutil.TempDir("", "zz-infogreffe")
	if err != nil {
		panic(err)
	}

	opts := []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", cliParam.headless), // headless or not ?
		chromedp.DisableGPU,
		chromedp.UserDataDir(dir),
	}

	// Create allocator context
	ctx, cancel1 := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel2 := chromedp.NewContext(ctx)
	ctx, cancel3 := context.WithTimeout(ctx, time.Duration(cliParam.minutes)*time.Minute)

	cancelAll := func() {
		cancel3()
		cancel2()
		cancel1()
		if cliParam.reset {
			os.RemoveAll(dir)
		}
	}

	return ctx, cancelAll
}
