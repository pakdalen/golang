package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type students struct {
	ID     string  `json:"id"`
	Gender string  `json:"Gender"`
	Name   string  `json:"name"`
	Age    float64 `json:"age"`
	Course float64 `json:"course"`
}

var studentss = []students{
	{ID: "1", Gender: "Male", Name: "John", Age: 18, Course: 1},
	{ID: "2", Gender: "Female", Name: "Raihan", Age: 19, Course: 2},
	{ID: "3", Gender: "Female", Name: "Niya", Age: 20, Course: 3},
	{ID: "4", Gender: "Male", Name: "Dalen", Age: 21, Course: 4},
}

func getStudentss(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}
func postStudentss(c *gin.Context) {
	var newStudents students

	if err := c.BindJSON(&newStudents); err != nil {
		return
	}

	studentss = append(studentss, newStudents)
	c.IndentedJSON(http.StatusCreated, newStudents)
}

func getStudentBYID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range studentss {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func deleteStudentBYID(c *gin.Context) {
	id := c.Param("id")
	for i, a := range studentss {
		if a.ID == id {
			studentss = append(studentss[:i], studentss[i+1:]...)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func main() {
	router := gin.Default()
	router.GET("/studentss", getStudentss)
	router.POST("/postStudentss", postStudentss)
	router.GET("/studentss/:id", getStudentBYID)
	router.DELETE("/deleteStudentss/:id", deleteStudentBYID)

	router.Run("localhost:8080")
}
