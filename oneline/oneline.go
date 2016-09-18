// oneline
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Print("now: ")
	for i := 0; i <= 15; i++ {
		fmt.Print(i)
		time.Sleep(100 * time.Millisecond)
		bs := strings.Repeat("\b", 1+i/10)
		fmt.Print(bs)
	}
}
