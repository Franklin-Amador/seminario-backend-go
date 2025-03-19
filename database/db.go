package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var DB *sql.DB

// InitDB inicializa la conexión a la base de datos
func InitDB() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Obtener la URL de la base de datos
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL no está configurada en el archivo .env")
	}

	// Conectar a la base de datos
	DB, err = sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	// Verificar la conexión
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error al hacer ping a la base de datos: %v", err)
	}

	fmt.Println("Conexión a la base de datos exitosa")
}

// GetUsers obtiene la lista de usuarios con id y username
func GetUsers() ([]map[string]interface{}, error) {
	rows, err := DB.Query("SELECT id, username FROM mdl_user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []map[string]interface{}

	for rows.Next() {
		var id int
		var username string

		if err := rows.Scan(&id, &username); err != nil {
			return nil, err
		}

		users = append(users, map[string]interface{}{
			"id":       id,
			"username": username,
		})
	}

	return users, nil
}

// GetDB devuelve la instancia de la base de datos
func GetDB() *sql.DB {
	return DB
}
