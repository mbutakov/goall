package main

import (
	"fmt"
	"strings"
)

// User ...
type Achivment struct {
	Id        int    `json:"id,-"`
	Name     string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
}

func (u Achivment) Create() error {
	u.Image = strings.TrimPrefix(u.Image, "C:\\fakepath\\")
	row := connection.QueryRow("INSERT INTO achivments(name,description,image) VALUES ($1,$2,$3) RETURNING id",u.Name,u.Description,u.Image)
	fmt.Println(row)
	e := row.Scan(&u.Id)
	if e != nil {
		return e
	}
	fmt.Println("Create new user with ID", u.Id)

	return nil
}
func (u Achivment) Remove() error {
	connection.Exec("DELETE FROM achivments WHERE id = $1 RETURNING id",u.Id)
	fmt.Println("Delete user with ID", u.Id)

	return nil
}