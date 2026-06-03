package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/igorsouzaccruz/meu-primeiro-crud-go/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadResquestError("Chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
