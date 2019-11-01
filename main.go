package main

import (
	"flag"
)

func main() {
	debugFlag := flag.Bool("debug", false, "sets log level to debug")
	//TODO helpFlag := flag.Bool("help", false, "display help message (TODO)")
	portFlag := flag.Int("p", 3000, "port to listen")
	flag.Parse()

	c := tileClient(*debugFlag)

	web(*debugFlag, c, *portFlag)
}

//Client client ot parse and access mbtile file
type Client struct {
}

func tileClient(debug bool) *Client {
	//TODO
	return &Client{}
}
