package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/justincampbell/anybar"
)

var port int

func init() {
	flag.Usage = usage

	flag.IntVar(&port, "port", anybar.DefaultPort, "the port of an AnyBar instance")

	flag.Parse()
}

func main() {
	if flag.NArg() != 1 {
		usage()
		os.Exit(1)
	}

	icon := flag.Args()[0]
	anybar.SendTo(icon, port)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [options] color/icon\n", os.Args[0])
	flag.PrintDefaults()
}
