package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/haohmaru3000/explicit_architecture/module/product/controller"
	productusecase "github.com/haohmaru3000/explicit_architecture/module/product/domain/usecase"
	productmysql "github.com/haohmaru3000/explicit_architecture/module/product/repository/mysql"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Setup dependencies
	repo := productmysql.NewMysqlRepository(db)
	useCase := productusecase.NewCreateProductUseCase(repo)
	api := controller.NewAPIController(useCase)

	v1 := r.Group("/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("", api.CreateProductAPI(db))
		}
	}

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
