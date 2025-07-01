package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgres() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=yourdb sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error al conectar a PostgreSQL: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("No se pudo hacer ping a PostgreSQL: %v", err)
	}

	fmt.Println("✅ Conectado a PostgreSQL correctamente.")
}
