package main

import (
	"log"

	"github.com/oneplus1000/goph"
)

func main() {
	pserv := goph.Serv{
		WwwRoot: "/Users/oneplus/go/src/github.com/oneplus1000/goph/testing/hello",
		PhpBin:  "/usr/local/bin/php-cgi",
	}
	err := pserv.Start()
	if err != nil {
		panicErr(err)
	}
	defer pserv.Close()
}

func panicErr(err error) {
	log.Panicf("%v", err)
}
