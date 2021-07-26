package Controllers

import (
	"fmt"
	Models2 "freshers-bootcamp/day3/exercise2/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)
//GetUsers ... Get all users
func GetStudents(c *gin.Context) {
	var student []Models2.Student
	err := Models2.GetAllStudents(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
//CreateUser ... Create User
func CreateStudent(c *gin.Context) {
	var student Models2.Student
	c.BindJSON(&student)
	err := Models2.CreateStudent(&student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
//GetUserByID ... Get the user by id
func GetStudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var student Models2.Student
	err := Models2.GetStudentByID(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
//UpdateUser ... Update the user information
func UpdateStudent(c *gin.Context) {
	var student Models2.Student
	id := c.Params.ByName("id")
	err := Models2.GetStudentByID(&student, id)
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	c.BindJSON(&student)
	err = Models2.UpdateStudent(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
//DeleteUser ... Delete the user
func DeleteStudent(c *gin.Context) {
	var student Models2.Student
	id := c.Params.ByName("id")
	err := Models2.DeleteStudent(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}