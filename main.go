package main

/*
Este proyecto es de codigo abierto, espero que este repositorio sea de ayuda para cualquier persona.
*/

import (
	"context"
	"log"
	"net/http"
	"os"
	"pl-api/db"
	"pl-api/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/subosito/gotenv"
)

func main() {

	router := gin.Default()

	// Esto es para que desde cualquier parte se acepten solicitudes GET, POST, DELETE, etc.
	// Un poco peligroso pero para eso esta la API KEY.
	router.Use(cors.Default())

	// Cargamos las variables de entorno
	gotenv.Load()
	port := os.Getenv("PORT")

	// Es interesante pero conectarnos a MongoDB desde aqui hace que la API sea más rapida.
	client, err := db.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	// Asegurarnos de que se desconecte cuando se acabe la ejecución

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Este endpoint verifica que todo este funcionando, debido a que si se conecto a MongoDB esto se ejecuta.
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "API Working correctly",
		})
	})

	v1 := router.Group("/v1")

	v1.GET("/product", func(c *gin.Context) {
		products, err := db.ViewProducts(client)

		// Este es el manejo de errores, si hay un error que retorne algo
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if len(*products) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "no products found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"total":    len(*products), // El total de productos encontrados, que buena funcion lo de len
			"products": products,
		})
	})

	v1.DELETE("/product", func(c *gin.Context) {
		var deleteProduct models.DeleteProduct

		err := c.ShouldBindJSON(&deleteProduct)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Please enter a valid body",
			})
			return
		}

		product, err := db.DeleteProduct(client, deleteProduct.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if product.DeletedCount == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "there was an error trying to delete the product, please verify that the ID entered is correct.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":      deleteProduct.ID,
			"info":    product,
			"message": "The product was successfully disposed of.",
		})
	})

	v1.POST("/product", func(c *gin.Context) {

		var product models.Product

		err := c.ShouldBindJSON(&product)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Please enter a valid body",
			})
			return
		}

		id, err := gonanoid.New()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		err = db.NewProduct(client, product.Name, &id, product.Price, product.Amount)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "The product was added successfully",
			"id":      id,
		})
	})

	/* Endpoints de Ventas */

	v1.POST("/sale", func(c *gin.Context) {

		var sale models.Sale

		err := c.ShouldBindJSON(&sale)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Please enter a valid body",
			})
			return
		}

		id, err := gonanoid.New()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		err = db.NewSale(client, &id, sale.User, sale.Total, sale.Date, sale.Products)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "The sale was paid successfully",
			"id":      id,
		})

	})
	router.Run(port)

}
