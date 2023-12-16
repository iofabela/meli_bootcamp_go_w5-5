package main

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w5-5/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// NO MODIFICAR
	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/melisprint")
	if err := db.Ping(); err != nil {
		panic(err)
	}
	r := gin.Default()

	router := routes.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
