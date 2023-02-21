package main

import "fmt"

/*
	The main idea behind the Iterator pattern is to extract
	the iteration logic of a collection into a different object
	called iterator.

	This iterator provides a generic method of iterating over a
	collection independent of its type.
*/

// Collection

type Collection interface {
	createIterator() Iterator
}

// Concrete collection

type UserCollection struct {
	users []*User
}

func (u *UserCollection) createIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}

// Iterator

type Iterator interface {
	hasNext() bool
	getNext() *User
}

type UserIterator struct {
	index int
	users []*User
}

// Concrete Iterator

func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false

}
func (u *UserIterator) getNext() *User {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

// Client code

type User struct {
	name string
	age  int
}

func main() {

	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
