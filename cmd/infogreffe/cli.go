package main

import (
	"flag"
	"fmt"
)

// Cli flags
type cliParamType struct {
	user     string
	passwd   string
	url      string
	reset    bool
	help     bool
	minutes  int
	headless bool
}

var cliParam cliParamType

const (
	myurl = "https://www2.juges.infogreffe.fr"
)

// Initialize command line
func init() {
	flag.StringVar(&cliParam.user, "u", "", "set user name")
	flag.StringVar(&cliParam.user, "user", "", "set user name")

	flag.StringVar(&cliParam.passwd, "p", "", "set password")
	flag.StringVar(&cliParam.passwd, "password", "", "set password")

	flag.StringVar(&cliParam.url, "url", "http://juges.infogreffe.fr", "change default url")

	flag.IntVar(&cliParam.minutes, "minutes", 1, "Time out in minutes")

	flag.BoolVar(&cliParam.reset, "reset", true, "clear after processing")

	flag.BoolVar(&cliParam.headless, "headless", false, "run in headless mode")

	flag.BoolVar(&cliParam.help, "h", false, "display this help menu and quit")
	flag.BoolVar(&cliParam.help, "help", false, "display this help menu and quit")

}

func processCli() bool {
	flag.Parse()
	fmt.Printf("\nStarting parsing with %v\n", cliParam)
	if cliParam.help {
		fmt.Println("Available options :")
		flag.PrintDefaults()
		return false
	}

	if cliParam.user == "" || cliParam.passwd == "" {
		fmt.Println("****  You did not set a valid user name and password ?!")
		flag.PrintDefaults()
		return false
	}

	return true

}
