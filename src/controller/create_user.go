package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/guinichetti/crud-go/src/configuration/rest_error"
	"github.com/guinichetti/crud-go/src/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_error.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
