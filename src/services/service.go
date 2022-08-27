package services

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDbConnection() (*sql.DB, error) {
	connectionString := "root:root@(127.0.0.1:3306)/test-db?parseTime=true"
	conn, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Falha em conectar ao db")
		panic(0)
	}
	createUser(conn)
	return conn, err
}
func createUser(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS movies(
		id int auto_increment,
		title text not null,
		genre text not null,
		year text not null,
		Primary key(id)
	);`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("impossivel criar tabela movies")
		return
	}
}
