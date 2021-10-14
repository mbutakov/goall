package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	_ "strings"
)

// User ...
type User struct {
	Id        int    `json:"id,-"`
	Login     string `json:"login"`
	Password string `json:"password"`
}

func (u User) login() error {
	auth := md5.Sum([]byte(u.Login + u.Password));
	hash := md5.Sum([]byte(u.Password))
	passwordHash := hex.EncodeToString(hash[:])
	allHash := hex.EncodeToString(auth[:])
	fmt.Println(allHash)
	rows, err := connection.DB.Query("select * FROM users WHERE login = $1 and password = $2", u.Login,passwordHash)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		fmt.Println("succes")

	}else{
		return gin.Error{}
	}
	return nil
}

func(u User) checkToken(login, token string) bool {
	var rows, _ = connection.DB.Query("SELECT token FROM users WHERE login = $1 and token = $2",login,token)
	if rows.Next(){
		return true
	}
	rows.Close()
	return false
}
func (u User) newSession(login, token string) {
	connection.Exec("UPDATE users  SET token=$1 where login = $2",token,login)
}
func (u User) deleteSession(login, token string) {
	s := make([]rune, 45)
	for i := range s {
		s[i] = runes[rand.Intn(len(runes))]
	}
	connection.Exec("Update users set token = $3 WHERE login = $1 and token = $2",login,token,string(s))
}


func (u User) Create() error {
	hash := md5.Sum([]byte(u.Password))
	passwordHash := hex.EncodeToString(hash[:])
	rows, err := connection.DB.Query("select * FROM users WHERE login = $1", u.Login)
	if err != nil {
		log.Fatal(err)
	}
	 userNameHas := false
	if rows.Next() {
		userNameHas = true
		return gin.Error{}
	}else{
		userNameHas = false
	}
	if(!userNameHas){
		row := connection.QueryRow("INSERT INTO users(login,password) VALUES ($1,$2) RETURNING id",u.Login,passwordHash)
		e := row.Scan(&u.Id)
		if e != nil {
			return e
		}
		fmt.Println("Create new user with ID", u.Id)
		return nil
	}

	return nil
}



