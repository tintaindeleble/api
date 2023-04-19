package controller

import (
	"api/api/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetById(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fmt.Sprintf("Get By Id: %s.", c.Param("ID")))
}

func GetRandom(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fmt.Sprintf("Get random %s quotes.", c.Param("limit")))
}

func New(c *gin.Context) {
	if err := service.New(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		c.IndentedJSON(http.StatusOK, "New quote added with ID: 10.")
	}
}

func Update(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fmt.Sprintf("Updated quote with ID: %s.", c.Param("ID")))
}

func Delete(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fmt.Sprintf("Deleted quote with ID: %s.", c.Param("ID")))
}
