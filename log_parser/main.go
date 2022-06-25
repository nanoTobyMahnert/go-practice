package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	p := newParser()

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {

		parsed := parse(p, in.Text())

		update(p, parsed)

	}

	summarize(p)
	dumErrs([]error{in.Err(), err(p)})
}

func dumErrs(errs []error) {
	for _, err := range errs {
		if err != nil {
			fmt.Println("> Err:", err)
		}
	}
}

func summarize(p *parser) {

	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "Domain", "Visits")
	fmt.Println(strings.Repeat("-", 42))

	sort.Strings(p.domains)
	for _, domain := range p.domains {
		fmt.Printf("%-30s %10d\n", domain, p.sum[domain].visits)
	}
	fmt.Printf("\n%-30s %10d\n", "Total", p.total)
}
