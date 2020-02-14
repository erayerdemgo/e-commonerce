package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"log"
)

func showproducts(c *gin.Context) {
	query := c.DefaultQuery("p", "1")
	fmt.Print(query)
}
func showhomepage(c *gin.Context) {

	c.HTML(200, "signup.gohtml", "sameteray")
}

type Signup struct {
	Name     string `form:"name" validate:"required,min=3,max=20"`
	Surname  string `form:"surname" validate:"required,min=2,max=20"`
	Eposta   string `form:"email" validate:"email"`
	Password string `form:"password"  validate:"min=3,max=15"` // todo regex here
}

func signup(c *gin.Context) {
	s := &Signup{}

	if err := c.Bind(s); err != nil {
		c.Status(500) // todo create error page
		return
	}
	fmt.Println(s)
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		fmt.Println(err)
		c.Status(500) // todo create error page
		return
	}

	// todo db insert

}
func main() {

	gin.SetMode("debug")
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/products/", )  // todo group here
	engine.GET("/", showhomepage)
	engine.POST("/signup", signup)
	log.Fatal(engine.Run(":8080"))

}
