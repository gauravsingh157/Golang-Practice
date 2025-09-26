package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Employee struct
type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

// In-memory employee storage
var employees []Employee

func main() {
	r := gin.Default()

	// POST API - Add Employee
	r.POST("/employee", func(c *gin.Context) {
		var newEmployee Employee

		// Bind JSON body to struct
		if err := c.ShouldBindJSON(&newEmployee); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Generate unique ID for employee
		newEmployee.ID = uuid.New().String()

		// Add to employees slice
		employees = append(employees, newEmployee)

		c.JSON(http.StatusCreated, gin.H{
			"message":  "Employee created successfully",
			"employee": newEmployee,
		})
	})

	// GET API - Get All Employees
	r.GET("/employees", func(c *gin.Context) {
		c.JSON(http.StatusOK, employees)
	})

	// Start server
	r.Run(":8080")
}
