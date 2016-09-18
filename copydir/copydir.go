// copydir
package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(source, dest string, ch chan error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		ch <- err
		return
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		ch <- err
		return
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}
	}
	ch <- err
}

func copyDir(source, dest string, ch chan error) {

	// Get properties of source dir.
	sourceinfo, err := os.Stat(source)
	if err != nil {
		ch <- err
		return
	}

	// Create dest dir.

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		ch <- err
		return
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// Create sub-directories - recursively.
			ch := make(chan error)
			go copyDir(sourcefilepointer, destinationfilepointer, ch)
			if err = <-ch; err != nil {
				fmt.Println(err)
			}
		} else {
			// Perform copy.
			ch := make(chan error)
			go copyFile(sourcefilepointer, destinationfilepointer, ch)
			if err = <-ch; err != nil {
				fmt.Println(err)
			}
		}

	}
	ch <- err
}

func main() {
	source_dir := "sample data"

	dest_dir := "backup"

	overwrite := true

	fmt.Println("Source: " + source_dir)

	// Check if the source dir exist.
	src, err := os.Stat(source_dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if !src.IsDir() {
		fmt.Println("Source is not a directory.")
		os.Exit(1)
	}

	fmt.Println("Destination: " + dest_dir)

	/* _, err = os.Open(dest_dir)
	if !os.IsNotExist(err) {
		fmt.Println("Destination directory already exists. Abort!")
		os.Exit(1)
	} */

	// We will continue to copy if we meet either condition:
	// 1. The destination does not exist.
	// 2. The destination exists and it is a dir and overwrite is true.
	dest, err := os.Stat(dest_dir)
	if err == nil {
		if !dest.IsDir() {
			fmt.Println("Destination is not a directory.")
			os.Exit(1)
		}
		if !overwrite {
			fmt.Println("We will not overwrite the destination.")
			os.Exit(1)
		}
	}

	ch := make(chan error)
	go copyDir(source_dir, dest_dir, ch)
	if err = <-ch; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Directory copied.")
	}
}
