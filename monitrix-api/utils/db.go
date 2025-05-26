package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := "root:Jj32631122.@tcp(localhost:3306)/monitrix?parseTime=true"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Erro ao validar conexão com o banco: %v", err)
	}

	fmt.Println("Banco de dados conectado com sucesso!")
}
