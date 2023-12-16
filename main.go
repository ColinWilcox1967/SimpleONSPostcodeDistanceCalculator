package main

import (
	"fmt"
	"os"
	"flag"
	distance "./distance"
)

const (
	VERSION = "1.0"
)

var (
	postcode1, postcode2 string
	showHelp bool
)

func showBanner () {
	fmt.Printf ("Simple Postcode Distance Calculator (Version %s).\n", VERSION)
	fmt.Printf ("(c) Colin Wilcox 2023.\n")
}

func showSyntax() {
	fmt.Println ("Arguments: -post1=<postcode 1> -post2=<postcode2>")
}

func getArguments() {

	flag.StringVar(&postcode1, "post1", "", "First UK postcode.")
	flag.StringVar(&postcode2, "post2", "", "First UK postcode.")
	flag.BoolVar(&showHelp,	   "help", false, "Show the command line syntax.")
	
	flag.Parse ()
}

func main() {
	
	showBanner()
	getArguments()
	

	if showHelp {
		showSyntax()
		os.Exit(0)
	}

		
	if postcode1 == "" || postcode2 == "" {
		fmt.Printf("You need two UK postcodes to calculate their distance apart.\n")
		os.Exit(-1)
	}

	fmt.Printf("\nThe distance between '%s' and '%s' is %f miles.\n", postcode1, postcode2, distance.GeoDistance(postcode1, postcode2))
}
