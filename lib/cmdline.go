package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	File     string
	Version  bool
	FullScan bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage：subDomainsBrute [options] target.com")
		println("\tsubDomainsBrute v0.0.1")
		println("\thttps://github.com/medasz/subDomainsBrute.git")
		flag.PrintDefaults()
	}
	flag.BoolVar(&Version, "version", false, "show program's version number and exit")
	flag.StringVar(&File, "file", "subnames.txt", "File contains new line delimited subs.")
	flag.BoolVar(&FullScan, "full", false, "Full scan, NAMES FILE subnames_full.txt will be used to brute")
	//flag.BoolVar()
}

func main() {
	flag.Parse()
	if Version {
		println("subDomainsBrute v0.0.1")
		println("https://github.com/medasz/subDomainsBrute.git")
	}else if len(flag.Args())<=0{
		flag.Usage()
	}
}
