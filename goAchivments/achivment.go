package main

import "fmt"

// User ...
type Achivment struct {
	Id        int    `json:"id,-"`
	Name     string `json:"name"`
}

func (u Achivment) Create() error {
	row := connection.QueryRow("INSERT INTO achivments(name) VALUES ($1) RETURNING id",u.Name)
	fmt.Println(row)
	e := row.Scan(&u.Id)
	if e != nil {
		return e
	}
	fmt.Println("Create new user with ID", u.Id)

	return nil
}