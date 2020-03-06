package main

import (
	"fmt"
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

// Get User from CSV
func GetUser(csvPath, userID string) (User, error) {
	file, err := ioutil.ReadFile(csvPath)
	if err != nil {
		return User{}, err
	}
	var users []User
	csverr := csvutil.Unmarshal(file, &users)
	if csverr != nil {
		return User{}, csverr
	}

	for _, user := range users {
		if user.ID == userID {
			return user, nil
		}
	}

	return User{}, fmt.Errorf("No such user with ID: %s", userID)
}

// Add User to CSV
func AddUser(csvPath string, target User) error {
	file, err := ioutil.ReadFile(csvPath)
	if err != nil {
		return err
	}
	var users []User
	csverr := csvutil.Unmarshal(file, &users)
	if csverr != nil {
		return csverr
	}

	for _, user := range users {
		if user.ID == target.ID {
			return fmt.Errorf("User `%s` already exists", target.ID)
		}
	}
	users = append(users, target)

	data, merr := csvutil.Marshal(users)
	if merr != nil {
		return merr
	}
	werr := ioutil.WriteFile(csvPath, data, 0644)
	if werr != nil {
		return werr
	}
	return nil
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
