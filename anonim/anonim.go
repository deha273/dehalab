// anonim
package main

import (
	"fmt"
)

func main() {
	for _, c := range [][]string{
		{"satu", "dua", "tiga", "empat", "lima"},
		{"merah", "kuning", "hijau"},
	} {
		fmt.Println(c)
	}
}
