package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)
var router *gin.Engine
var connection *sqlx.DB
var achivmentList []Achivment
var (
	runes  = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890")
)
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
	router.GET("/exit", exit)
	router.GET("/register", regPageHandler)
	router.POST("/registerPost", regPagePostHandler)
	router.GET("/login", loginPageHandler)
	router.POST("/loginPost", loginPostPageHandler)
	router.GET("/", handlerIndex)
	router.GET("/admin", handlerAdminIndex)
	router.GET("/createAchivment", handlerCreateAchivment)
	router.POST("/create", handlerCreateAchivmentCreate)
	router.POST("/upload", handlerUploadImage)
	router.GET("/r", handlerRemoveAchivment)
	router.GET("/getAchivment", handlerGetAchivment)
	_ = router.Run(":8080")
}
func exit(c *gin.Context){
	var a User
	login := getCookie(c.Request,"login")
	token := getCookie(c.Request,"session_token")
	if(a.checkToken(login,token)){
		a.deleteSession(login,token)
	}else{

	}
	deleteCookie(c.Writer, "login")
	deleteCookie(c.Writer, "session_token")
	handlerIndex(c)
}
func regPageHandler(c *gin.Context){
	c.HTML(200, "register.html", gin.H{})
}
func regPagePostHandler(c *gin.Context){
	var a User
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
func loginPageHandler(c *gin.Context){
	c.HTML(200, "login.html", gin.H{})
}
func loginPostPageHandler(c *gin.Context){
	var a User
	e := c.BindJSON(&a)
	if e != nil {
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}
	e = a.login()
	if e != nil {
		c.JSON(200, gin.H{
			"Error": "Не удалось зарегистрировать пользователя",
		})
		return
	}
	token := genToken()
	setCookie(c,"login",a.Login,30)
	setCookie(c,"session_token",token,30)
	a.newSession(a.Login,token)
	c.JSON(200, gin.H{
		"Error": nil,
	})
}
func setCookie(c *gin.Context, name, value string,d int) {
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
	}
	if d != 0{
		expires := time.Now().AddDate(0,0,d)
		cookie.Expires = expires
	}
	http.SetCookie(c.Writer, &cookie)
}
func genToken() string {
	s := make([]rune, 15)
	for i := range s {
		s[i] = runes[rand.Intn(len(runes))]
	}
	return string(s)
}
func deleteCookie(w http.ResponseWriter,name string){
	cookie := http.Cookie{
		Name: name,
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
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
}
func handlerRemoveAchivment(c *gin.Context) {
	id := c.Query("id")
	connection.Exec("DELETE FROM achivments WHERE id = $1", id)
}
func getCookie(r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		return ""
	}
	return c.Value
}
func handlerCreateAchivment(c *gin.Context) {
	login := getCookie(c.Request,"login")
	var u User
	u.Login = login
	token := getCookie(c.Request,"session_token")
	if(!u.checkToken(login,token)){
		fmt.Fprint(c.Writer,"no authorization");
	}else{
		c.HTML(200, "createAchivment.html", gin.H{"user" : u.Login})
	}

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
func handlerAdminIndex(c *gin.Context) {
	login := getCookie(c.Request,"login")
	var u User
	u.Login = login
	token := getCookie(c.Request,"session_token")
	if(!u.checkToken(login,token)){
		fmt.Fprint(c.Writer,"no authorization");
	}else{
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

}
func handlerIndex(c *gin.Context) {
	fmt.Println("asdasd");
	login := getCookie(c.Request,"login")
	var u User
	u.Login = login;
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
	if login == "" {
		c.HTML(200, "index.html", gin.H{"Achivments": achivmentList})
	}else{
		c.HTML(200, "index.html", gin.H{"Achivments": achivmentList,"User": u})
	}
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
