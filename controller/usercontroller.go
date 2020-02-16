package controller

import (
	"e-commenerce/model"
	"e-commenerce/service"
	"e-commenerce/validation"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Loginpage(c *gin.Context) {
	c.HTML(200, "signin.gohtml", nil)
}

func Login(c *gin.Context) {
	signin := model.UserSignin{}
	if err := c.Bind(&signin); err != nil {
		log.Fatalf("binding problem %v", err)
	}
	b := service.Login(signin)
	if !b {
		c.HTML(200, "error.gohtml", nil)
		return

	}
	c.HTML(200, "welcome.gohtml", signin.Email)
}

func Signup(c *gin.Context) {
	s := model.UserSignup{}

	if err := c.Bind(&s); err != nil {
		c.Status(500) // todo create error page
		return
	}

	if err := validation.Validate.Struct(&s); err != nil {
		fmt.Println(err)
		c.Status(500) // return error
		return
	}

	if err := service.Createuser(s); err != nil {
		log.Fatal(err)
	}
	c.HTML(201, "signin.gohtml", gin.H{
		"email": s.Email,
	})
}
