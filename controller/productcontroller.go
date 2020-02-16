package controller

import (
	"e-commenerce/model"

	"github.com/gin-gonic/gin"
)

func Showhomepage(c *gin.Context) {
	Products := []model.Product{{
		Name: "T SHİRT",
		Url:  "https://cdn.shopify.com/s/files/1/0066/9212/products/python-tshirt_grande.jpg?v=1562275988",
	}, {
		Name: "T SHİRT",
		Url:  "https://cdn.shopify.com/s/files/1/0066/9212/products/python-tshirt_grande.jpg?v=1562275988",
	}, {
		Name: "T SHİRT",
		Url:  "https://cdn.shopify.com/s/files/1/0066/9212/products/python-tshirt_grande.jpg?v=1562275988",
	},
		{
			Name: "T SHİRT",
			Url:  "https://cdn.shopify.com/s/files/1/0066/9212/products/python-tshirt_grande.jpg?v=1562275988",
		},
	}
	c.HTML(200, "welcome.gohtml", Products)
}
