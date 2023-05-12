package controllers

import (
	"net/http"

	"go-test-api/database"
	"go-test-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductRepo struct {
	Db *gorm.DB
}

func New() *ProductRepo {
	db := database.Initdb()
	db.AutoMigrate(&models.Products{})
	return &ProductRepo{Db: db}
}

func (db *ProductRepo) CreateProduct(c *gin.Context) {
	var product models.Products
	c.BindJSON(&product)
	err := models.CreateProduct(db.Db, &product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (db *ProductRepo) GetProducts(c *gin.Context) {
	var product []models.Products
	// c.BindJSON(&product)
	err := models.GetProducts(db.Db, &product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (db *ProductRepo) GetProductById(c *gin.Context) {
	//id,_:=strconv.Atoi(c.Param("id")) convert str to int
	id := c.Param("id")
	var product models.Products
	// c.BindJSON(&product)
	err := models.GetProductById(db.Db, &product, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, product)
}
