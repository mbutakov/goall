package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var router *gin.Engine
var connection *sqlx.DB
var achivmentList []Achivment;

var connectionString = "host=127.0.0.1 port=5432 user=postgres password=12312345 dbname=test sslmode=disable"

func main() {
	var e error
	connection, e = sqlx.Open("postgres", connectionString)
	if e != nil {
		fmt.Println(e)
		return
	}

	router = gin.Default()
	router.Static("/assets/", "front/")
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", handlerIndex)
	router.GET("/createAchivment", handlerCreateAchivment)
	router.POST("/create", handlerCreateAchivmentCreate)
	router.GET("/r", handlerRemoveAchivment)

	_ = router.Run(":8080")
}


func handlerRemoveAchivment(c *gin.Context) {

	id := c.Query("id")
	connection.Exec("DELETE FROM achivments WHERE id = $1",id)

	//var a Achivment
	//e := c.BindJSON(&a)
	//if e != nil {
	//	c.JSON(200, gin.H{
	//		"Error": e.Error(),
	//	})
	//	return
	//}
	//e = a.Remove()
	//if e != nil {
	//	c.JSON(200, gin.H{
	//		"Error": "Не удалось удалить",
	//	})
	//	return
	//}
	//
	//c.JSON(200, gin.H{
	//	"Error": nil,
	//})
}

func handlerCreateAchivment(c *gin.Context) {
	c.HTML(200, "createAchivment.html", gin.H{})
}

func handlerCreateAchivmentCreate(c *gin.Context) {


	var a Achivment

	e := c.BindJSON(&a)
	if e != nil {
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}

	e = a.Create()
	if e != nil {
		c.JSON(200, gin.H{
			"Error": "Не удалось зарегистрировать пользователя",
		})
		return
	}

	c.JSON(200, gin.H{
		"Error": nil,
	})
}

// pkg.go.dev/text/template
func handlerIndex(c *gin.Context) {
	achivmentList = nil
	var a Achivment
	rows, err := connection.DB.Query("select id, name from achivments")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&a.Id, &a.Name)
		if err != nil {
			log.Fatal(err)
		}
		achivmentList = append(achivmentList,a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(200, "index.html", gin.H{"Achivments" : achivmentList})
}


