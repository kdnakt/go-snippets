package main

import (
	"io/ioutil"
	"log"
	"strings"

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

	target := "foobar"
	for _, user := range users {
		if user.ID == target {
			log.Println(user)
		}
	}

	for i, user := range users {
		if strings.HasPrefix(user.ID, "bar") {
			users = delete(users, i)
		}
	}
	log.Println(users)

	newUser := User{
		ID:   "buz",
		Name: "no name",
	}
	users = append(users, newUser)
	log.Println(users)

	data, merr := csvutil.Marshal(users)
	if merr != nil {
		log.Fatal(merr)
	}
	werr := ioutil.WriteFile("out.csv", data, 0644)
	if werr != nil {
		log.Fatal(werr)
	}
}

func delete(users []User, i int) []User {
	users = append(users[:i], users[i+1:]...)
	res := make([]User, len(users))
	copy(res, users)
	return res
}
