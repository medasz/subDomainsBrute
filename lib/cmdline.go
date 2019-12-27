package lib

import (
	"flag"
	"fmt"
	"os"
)

var (
	File           string
	Version        bool
	FullScan       bool
	Help           bool
	IgnoreIntranet bool
	Threads        int
	Process        int
	Output         string
	Target         string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage：subDomainsBrute [options]")
		println("\tsubDomainsBrute v0.0.1")
		println("\thttps://github.com/medasz/subDomainsBrute.git")
		flag.PrintDefaults()
	}
	flag.BoolVar(&Version, "version", false, "show program's version number and exit")
	flag.BoolVar(&Help, "help", false, "show this help message and exit")
	flag.StringVar(&File, "file", "subnames.txt", "File contains new line delimited subs.")
	flag.BoolVar(&FullScan, "full", false, "Full scan, NAMES FILE subnames_full.txt will be used to brute")
	flag.BoolVar(&IgnoreIntranet, "ignore-intranet", false, "Ignore domains pointed to private IPs")
	flag.IntVar(&Threads, "threads", 200, "Num of scan threads, 200 by default")
	flag.IntVar(&Process, "process", 6, "Num of scan Process, 6 by default")
	flag.StringVar(&Output, "output", Output, "Output file name. default is {target}.txt")
	flag.StringVar(&Target, "target", Target, "This is the scan target")
}

func Parse_args() {
	flag.Parse()
	if Version {
		println("subDomainsBrute v0.0.1")
		println("https://github.com/medasz/subDomainsBrute.git")
	} else if flag.NFlag() <= 0 {
		flag.Usage()
		os.Exit(2)
	} else if Help {
		flag.Usage()
	}
}
