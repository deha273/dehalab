// mybackup
package main

import (
	"fmt"
	"os"

	"github.com/yaliv/go-pkg/copydir"
)

func main() {
	err := copydir.Copy("sample data", "backup", false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
