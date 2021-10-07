package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var router *gin.Engine
var connection *sqlx.DB
var achivmentList []Achivment

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
	router.GET("/admin", handlerAdminIndex)
	router.GET("/createAchivment", handlerCreateAchivment)
	router.POST("/create", handlerCreateAchivmentCreate)
	router.POST("/upload", handlerUploadImage)
	router.GET("/r", handlerRemoveAchivment)
	router.GET("/getAchivment", handlerGetAchivment)
	_ = router.Run(":8080")
}

func handlerGetAchivment(c *gin.Context) {
	id := c.Query("id")
	var a Achivment
	rows, err := connection.DB.Query("select * FROM achivments WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&a.Id, &a.Name, &a.Description, &a.Image)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{"Achivments": a})
	fmt.Println(id)
}

func handlerRemoveAchivment(c *gin.Context) {
	id := c.Query("id")
	connection.Exec("DELETE FROM achivments WHERE id = $1", id)
}

func handlerCreateAchivment(c *gin.Context) {
	c.HTML(200, "createAchivment.html", gin.H{})
}

func handlerCreateAchivmentCreate(c *gin.Context) {
	var a Achivment
	e := c.BindJSON(&a)
	fmt.Println(e);
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

func handlerAdminIndex(c *gin.Context) {
	achivmentList = nil
	var a Achivment
	rows, err := connection.DB.Query("select id, name,image from achivments")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&a.Id, &a.Name, &a.Image)
		if err != nil {
			log.Fatal(err)
		}
		achivmentList = append(achivmentList, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(200, "admin.html", gin.H{"Achivments": achivmentList})
}

func handlerIndex(c *gin.Context) {
	achivmentList = nil
	var a Achivment
	rows, err := connection.DB.Query("select id, name,image from achivments")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&a.Id, &a.Name, &a.Image)
		if err != nil {
			log.Fatal(err)
		}
		achivmentList = append(achivmentList, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(200, "index.html", gin.H{"Achivments": achivmentList})
}

func handlerUploadImage(c *gin.Context) {
	form, _ := c.MultipartForm()
	_, e := os.Stat("front/img")
	if os.IsNotExist(e) {
		fmt.Println("ERROR: ", e.Error())
		e = os.MkdirAll("front/img", 0777)
		if e != nil {
			fmt.Println("ERROR: ", e.Error())
			c.JSON(200, gin.H{
				"Error": e.Error(),
			})
			return
		}
	}
	files := form.File["File"]
	file := files[0]
	if e := c.SaveUploadedFile(file, "front/img/"+file.Filename); e != nil {
		fmt.Println("ERROR: ", e.Error())
		c.JSON(200, gin.H{
			"Error": fmt.Sprintf("upload file e: %s", e.Error()),
		})
	}
	c.JSON(200, gin.H{
		"Error": nil,
	})
}
