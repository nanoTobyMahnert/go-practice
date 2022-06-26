package main

import (
	"fmt"
	"strconv"
)

type book struct {
	product
	published interface{}
}

func (b *book) print() {
	b.product.print()
	p := format(b.published)
	fmt.Printf("\t - (%v)\n", p)
}

func format(v interface{}) string {
	var t int
	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	default:
		return "unknown"
	}
	return fmt.Sprintf("%d-%02d-%02d", t/10000, t%10000/100, t%100)
}
