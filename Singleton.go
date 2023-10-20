package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type User struct {
	Id int
}

var singleInstanceOfUser *User

func getUserByID(id int) *User {
	if singleInstanceOfUser == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstanceOfUser == nil {
			fmt.Printf("User with %d ID has not found, creating User by this ID. \n", id)
			singleInstanceOfUser = &User{Id: id}
		} else {
			fmt.Printf("User with %d ID is exists. \n", id)
		}
	} else {
		fmt.Printf("User with %d ID is exists. \n", id)
	}

	return singleInstanceOfUser
}

func main() {
	userID := 111
	user := getUserByID(userID)
	fmt.Printf("User id: %d \n", user.Id)

	user = getUserByID(userID)
	fmt.Printf("User id: %d \n", user.Id)
}
