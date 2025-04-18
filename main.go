/*
Hands-on exercise #67 - interfaces & mock testing a db

Interfaces in Go are a powerful tool for abstraction and can be especially useful when you
want to write unit tests for functions that interact with a database. By creating an interface
that describes the behavior of the database interactions your code needs, you can swap out
the real database for a mock one in your tests.
*/
package main

import (
	"fmt"
	"log"
)

// User represents a user with an id and first name.
type User struct {
	ID    int
	First string
}

// DummyDataStore is a dummy service which does nothing
type DummyDataStore struct {
}

func (dd DummyDataStore) GetUser(Id int) (User, error) {
	return User{}, nil
}

func (dd DummyDataStore) SaveUser(u User) error {
	return nil
}

// MockDatastore is a temporary service that stores retrievable data.
type MockDatastore struct {
	Users map[int]User
}

func (md MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User %v not found", id)
	}
	return user, nil
}

func (md MockDatastore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if ok {
		return fmt.Errorf("User %v already exists", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

// Datastore defines an interface for storing retrievable data.
// Any TYPE that implements this interface (has these two methods) is also of TYPE `Datastore`.
// This means any value of TYPE `MockDatastore` is also of TYPE `Datastore`
// This means we could have a value of TYPE `*sql.DB` and it can also be of TYPE `Datastore`
// This means we can write functions to take TYPE `Datastore` and use either of these:
// -- `MockDatastore`
// -- `*sql.DB`
// https://pkg.go.dev/database/sql#Open
type Datastore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

// Service represents a service that stores retrievable data.
// We will attach methods to `Service` so that we can use either of these:
// -- `MockDatastore`
// -- `*sql.DB`
type Service struct {
	ds Datastore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func main() {
	db := MockDatastore{
		Users: make(map[int]User),
	}

	srvc := Service{
		ds: db,
	}

	u1 := User{
		ID:    1,
		First: "James",
	}

	err := srvc.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	u1Returned, err := srvc.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(u1Returned)

	// use a empty dummy DataStore
	dd := DummyDataStore{}
	srvc2 := Service{
		ds: dd,
	}

	err = srvc2.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	u1Returned, err = srvc2.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	fmt.Println(u1Returned)

}
