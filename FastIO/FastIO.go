// FastIO
package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	ints := getInts()
	for _, v := range ints {
		out.WriteString(strconv.Itoa(v) + "\n")
	}
	out.Flush()
}

func getInts() []int {
	//assumes POSITIVE INTEGERS. Check v for '-' if you have negative.
	var buf []byte
	buf, _ = ioutil.ReadAll(os.Stdin)
	var ints []int
	num := 0
	found := false
	for _, v := range buf {
		if '0' <= v && v <= '9' {
			num = 10*num + int(v-'0') //could use bitshifting here.
			found = true
		} else if found {
			ints = append(ints, num)
			found = false
			num = 0
		}
	}
	if found {
		ints = append(ints, num)
		found = false
		num = 0
	}
	return ints
}
