package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	s := flag.Arg(0)
	s = strings.ReplaceAll(s,"(","")
    s = strings.ReplaceAll(s,")","")
	s = strings.ReplaceAll(s,"-","")
	s = strings.ReplaceAll(s," ","")
	fmt.Println(s)

}
