package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igorsouzaccruz/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/igorsouzaccruz/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/igorsouzaccruz/meu-primeiro-crud-go/src/controller/model/response"
)

func CreateUser(c *gin.Context) {
	log.Println("Init Create User controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("There are some incorrect filds, erros=%s", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(http.StatusOK, response)
}
