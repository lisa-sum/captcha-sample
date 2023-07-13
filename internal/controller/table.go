package controller

import (
	"echarts_example/internal/schema"
	"echarts_example/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TableInter interface {
	GetOneTable(c *gin.Context)
	GetAllTable(c *gin.Context)
	CreateTable(c *gin.Context)
}

type Table struct{}

func (t Table) GetOneTable(c *gin.Context) {
	table, err := service.GetOneTable()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"body": table,
	})
}

func (t Table) GetAllTable(c *gin.Context) {
	table, err := service.GetAllTable()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"body": table,
	})
}

func (t Table) CreateTable(c *gin.Context) {
	var createData schema.Echarts
	if err := c.ShouldBindJSON(&createData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Printf("%v", createData)
	createTable, err := service.CreateTable(createData)
	//createTable, err := service.CreateTable()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"body": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"body": createTable,
	})
}
