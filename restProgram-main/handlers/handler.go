package handlers

import (
	"net/http"
	"restProgram/models"

	"github.com/gin-gonic/gin"
)

type CreateStudentInput struct {
	Name    string `json:"Name" binding:"required"`
	Surname string `json:"Surname" binding:"required"`
	Id      int32  `json:"id" binding:"required"`
}

type UpdateStudentInput struct {
	Name    string `json:"name"`
	Surname string `json:"Surname" binding:"required"`
	Id      int32  `json:"id" binding:"required"`
}

func GetAllNames(context *gin.Context) {
	var students []models.Student
	models.DB.Find(&students)

	context.JSON(http.StatusOK, gin.H{"students": students})
}

func CreateStudent(context *gin.Context) {

	var input CreateStudentInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := models.Student{Name: input.Name, Surname: input.Surname, Id: input.id}
	models.DB.Create(&student)

	context.JSON(http.StatusOK, gin.H{"students": students})
}

func UpdateStudent(context *gin.Context) {
	var student models.Student

	if err := models.DB.Where("id=?", context.Param("id")).First(&student).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Student not exist"})
		return
	}

	var input UpdateStudentInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&student).Update(input)
	context.JSON(http.StatusOK, gin.H{"students": student})
}

func DeleteStudent(context *gin.Context) {
	var student models.Student

	if err := models.DB.Where("id=?", context.Param("id")).First(&student).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Student not exist"})
		return
	}

	models.DB.Delete(&student)
	context.JSON(http.StatusOK, gin.H{"Students": true})
}
