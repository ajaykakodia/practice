package main

import (
	"errors"
	"fmt"
	"log"
)

type User struct {
	ID   int
	Name string
}

type MockDatastore struct {
	Users map[int]User
}

func (md MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User not found: %d", id)
	}
	return user, nil
}

func (md MockDatastore) SaveUser(u User) error {
	if _, ok := md.Users[u.ID]; ok {
		return errors.New("User already exists")
	}
	md.Users[u.ID] = u
	return nil
}

type DataStore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

type Service struct {
	ds DataStore
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

	svc := Service{
		ds: db,
	}

	u1 := User{
		ID:   1,
		Name: "Ajay",
	}
	err := svc.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	u1Returned, err := svc.GetUser(1)
	if err != nil {
		log.Println("User not found")
	}
	fmt.Println(u1Returned)
}
