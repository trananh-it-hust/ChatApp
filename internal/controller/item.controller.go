package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"main.go/internal/service"

	"main.go/pkg/response"
)

type ItemController struct {
	itemService *service.ItemService
}

func Create() *ItemController {
	return &ItemController{
		itemService: service.Create(),
	}
}

type ItemService interface {
	CreateItem() int
}

func (i *ItemController) CreateItem(c *gin.Context) {
	// input c
	// get body
	var x interface{}
	if err := c.ShouldBindJSON(&x); err != nil {
		c.JSON(400, response.NewAppError(err, "Bad Request", "Body is empty", "bad_request"))
		return
	}
	fmt.Println("x", x)
	if x == nil {
		c.JSON(400, response.NewAppError(nil, "Bad Request", "Body is empty", "bad_request"))
		return
	}
	// Call service
	y := i.itemService.CreateItem()

	// output
	c.JSON(200, response.NewResponse(200, "success", y))
}
