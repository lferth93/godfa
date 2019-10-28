package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	ins := flag.String("i", "", "Input file, deafult stdin")
	outs := flag.String("o", "", "Output file, default out.dfa")
	flag.Parse()
	in := os.Stdin
	out := os.Stdout
	if len(*ins) > 0 {
		inf, err := os.Open(*ins)
		if err != nil {
			log.Println(err)
		} else {
			in = inf
			defer in.Close()
		}
	}
	if len(*outs) > 0 {
		outf, err := os.Create(*outs)
		if err != nil {
			log.Println(err)
		} else {
			out = outf
			defer out.Close()
		}
	}

	nfa := newNfa(in)
	fmt.Println("NFA")
	fmt.Println(nfa)

	dfa := nfa.toDfa()
	fmt.Println("DFA")
	fmt.Println(dfa)
	dfa.write(out)
}
