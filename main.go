package main

import (
	"flag"
	"log"
)

func main() {
	debugFlag := flag.Bool("v", false, "sets log level to verbose/debug")
	//TODO helpFlag := flag.Bool("help", false, "display help message (TODO)")
	portFlag := flag.Int("p", 3000, "port to listen")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("Need one and only one arg for .mbtiles file path.")
	}
	c := tileClient(*debugFlag, args[0])

	web(*debugFlag, c, *portFlag)
}
