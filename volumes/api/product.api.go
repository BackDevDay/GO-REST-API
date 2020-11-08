package api

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"rest/db"
	"rest/interceptor"
	"rest/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//  SetupProductAPI - router
func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v1")
	{
		productAPI.GET("/product", interceptor.JwtVerify, getProduct)
		productAPI.GET("/product/:id", interceptor.JwtVerify, getProductByID)
		productAPI.POST("/product", interceptor.JwtVerify, createProduct)
		productAPI.PUT("/product", interceptor.JwtVerify, editProduct)
	}
}

func getProduct(context *gin.Context) {
	var product []model.Product

	keyword := context.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDatabase().Where("name like ?", keyword).Find(&product)
	} else {
		db.GetDatabase().Find(&product)
	}
	context.JSON(http.StatusOK, product)
}

func getProductByID(context *gin.Context) {
	var product model.Product
	db.GetDatabase().Where("id = ?", context.Param("id")).First(&product)
	context.JSON(http.StatusOK, product)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func saveImage(image *multipart.FileHeader, product *model.Product, context *gin.Context) {
	if image != nil {
		runningDir, _ := os.Getwd()
		product.Image = image.Filename
		extension := filepath.Ext(image.Filename)
		fileName := fmt.Sprintf("%d%s", product.ID, extension)
		filePath := fmt.Sprintf("%s/uploads/images/%s", runningDir, fileName)

		if fileExists(filePath) {
			os.Remove(filePath)
		}
		context.SaveUploadedFile(image, filePath)
		db.GetDatabase().Model(&product).Update("image", fileName)
	}
}

func createProduct(context *gin.Context) {

	product := model.Product{}
	product.Name = context.PostForm("name")
	product.Stock, _ = strconv.ParseInt(context.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(context.PostForm("price"), 64)
	product.CreatedAt = time.Now()
	db.GetDatabase().Create(&product)
	image, _ := context.FormFile("image")
	saveImage(image, &product, context)

	context.JSON(http.StatusOK, gin.H{"status": true, "message": "ok", "result": product})

}

func editProduct(context *gin.Context) {
	var product model.Product
	id, _ := strconv.ParseInt(context.PostForm("id"), 10, 32)
	product.ID = uint(id)
	product.Name = context.PostForm("name")
	product.Stock, _ = strconv.ParseInt(context.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(context.PostForm("price"), 64)

	db.GetDatabase().Save(&product)
	image, _ := context.FormFile("image")
	saveImage(image, &product, context)
	context.JSON(http.StatusOK, gin.H{"status": true, "message": "ok", "result": product})
}
