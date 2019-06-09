package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

// Provides an allocator with default option, that opens a visible browser
// cookies etc are shared in a common dir that gets erased at the end
func getVisibleAllocator() (context.Context, context.CancelFunc) {
	dir, err := ioutil.TempDir("", "xavier-chromedp-demo")
	if err != nil {
		panic(err)
	}

	opts := []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		//chromedp.Headless,			  // Operate in headless mode
		chromedp.Flag("headless", false), // Display what's going on ...
		chromedp.DisableGPU,
		chromedp.UserDataDir(dir),
	}

	// Create allocator context
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)

	cancel2 := func() {
		cancel()
		os.RemoveAll(dir)
	}

	return ctx, cancel2
}

func checkBrowserStarted(ctx context.Context) {
	// ensure that the browser process is started
	if err := chromedp.Run(ctx); err != nil {
		panic(err)
	}
}

func getSelectedHeadlines(ctx context.Context, myurl string) {

	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(myurl),
		chromedp.Title(&res),
		myPrintAction{&res},
		// Dump outerHtml
		chromedp.OuterHTML("span ~ a", &res),
		myPrintAction{&res},
		// Search by css selector
		chromedp.Text("article h3", &res),
		myPrintAction{&res},
		// search by css selector
		chromedp.Text("article ~ div h4", &res),
		myPrintAction{&res},
		// Print text covering multiple tags
		chromedp.Text("article ~ div", &res),
		myPrintAction{&res},
		//Test xpath
		chromedp.Text("//h4//a", &res),
		myPrintAction{&res},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllHeadlines(ctx context.Context, myurl string) {

	var lnodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(myurl),
		chromedp.Nodes("//div//h4/a", &lnodes),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nlnodes : \t", lnodes)
	for i := range lnodes {
		fmt.Println("\n=================\tNode N° ", i, "\t====================")
		fmt.Println("href     \t", lnodes[i].AttributeValue("href"))
		fmt.Println("xpath    \t", lnodes[i].FullXPath())
		fmt.Println("nodevalue\t", lnodes[i].NodeValue)
		fmt.Println("nodename \t", lnodes[i].NodeName)
		fmt.Println()
	}
	fmt.Println("Done - getAllHeadlines")

}

// Define a printable action string
type myPrintAction struct {
	*string
}

func (s myPrintAction) Do(ctx context.Context) error {
	fmt.Printf("my log :\t%s\n", *(s.string))
	return nil
}
