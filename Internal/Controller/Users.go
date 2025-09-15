package Controller

import (
	"CSR/Internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (ctrl *Controller) GetAllUsers(c *gin.Context) {
	users, err := ctrl.service.GetAllUsers()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error while getting all users",
		})
		return
	}
	c.JSON(http.StatusOK, users)

}

func (ctrl *Controller) GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "param conversion error",
		})
		log.Fatal(err)
	}
	user, err := ctrl.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal error",
		})
		return
	}
	c.JSON(200, user)
}
func (ctrl *Controller) CreateNewUser(c *gin.Context) {
	var newUser models.User
	err := c.BindJSON(&newUser)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(newUser)
	err = ctrl.service.CreateNewUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error creating new user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
	})

}

func (ctrl *Controller) UpdateUserById(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error while getting request body",
		})
		log.Println(err)
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error while processing params",
		})
		log.Println(err)
		return
	}
	err = ctrl.service.UpdateUserById(id, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error while updating data",
		})
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
	})

}

func (ctrl *Controller) DeleteUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error while processing params",
		})
		log.Println(err)
		return
	}
	err = ctrl.service.DeleteUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error while deleting user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user successfully deleted",
	})

}
