// csv2mongo
package main

// Required packages import.
import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"gopkg.in/mgo.v2"
)

// It defines the type of order to take into MongoDB in the structure.
type Mongo struct {
	Name   string
	Date   string
	Number int
}

func main() {

	// Use mgo to connect to MongoDB.
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// The destination DB and Collection of MongoDB.
	c := session.DB("banana").C("test")

	// Open CSV.
	file, err := os.Open("test.csv")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read CSV.
	reader := csv.NewReader(file)

	for {
		record, err := reader.Read() // Read one line.
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		number, err := strconv.Atoi(record[2])

		// Perform Insert.
		err = c.Insert(&Mongo{record[0], record[1], number})

		if err != nil {
			panic(err)
		}
		log.Printf("%#v", record)
	}
}
