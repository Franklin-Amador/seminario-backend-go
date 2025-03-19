package main

import (
	"log"
	"net/http"

	"seminario-backend-go/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la conexión a la base de datos
	database.InitDB()

	// Inicializar Gin
	r := gin.Default()

	// Ruta para indicar inicio exitoso
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Servidor iniciado exitosamente",
		})
	})

	// Ruta para obtener usuarios
	r.GET("/usuarios", func(c *gin.Context) {
		users, err := database.GetUsers()
		if err != nil {
			log.Println("Error al obtener usuarios:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error al obtener usuarios",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"usuarios": users})
	})

	// Ruta de prueba para verificar la conexión
	r.GET("/test", func(c *gin.Context) {
		db := database.GetDB()

		err := db.Ping()
		if err != nil {
			log.Println("Error en la conexión:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error en la conexión a la base de datos",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Conexión a la base de datos exitosa",
		})
	})

	// Iniciar el servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
