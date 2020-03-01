package main

import (
	"io/ioutil"
	"log"

	"github.com/jszwec/csvutil"
)

// User data from csv file
type User struct {
	ID   string `csv:"id"`
	Name string `csv:"name"`
}

func main() {
	file, err := ioutil.ReadFile("user.csv")
	if err != nil {
		log.Fatal(err)
	}
	var users []User
	csverr := csvutil.Unmarshal(file, &users)
	if csverr != nil {
		log.Fatal(csverr)
	}

	log.Println(users)
}
