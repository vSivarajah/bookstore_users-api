package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vSivarajah/bookstore_users-api/domain/users"
	"github.com/vSivarajah/bookstore_users-api/services"
	"github.com/vSivarajah/bookstore_users-api/utils/errors"
)

func GetUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "implement me")
}

func CreateUser(c *gin.Context) {
	user := users.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
