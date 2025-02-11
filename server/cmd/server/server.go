package main

import (
	"database/sql"
	"github.com/DiegoJCordeiro/golang-study/activity/server/cfg"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/webserver/handlers"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

func main() {

	var db *sql.DB

	configuration, err := cfg.LoadConfiguration("./cmd/server")

	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	if db, err = sql.Open(configuration.DBDriver, configuration.DBHost); err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	server := handlers.NewWebServer(":8080")

	configureHandlersOnWebServer(server, &httpClient, db)

	err = server.Start()

	if err != nil {
		panic("error to start the server.")
	}
}

func configureHandlersOnWebServer(server *handlers.WebServer, httpClient *http.Client, db *sql.DB) {

	queryQuotationHandler := NewQueryQuotationHandler(db, httpClient)
	updateQuotationHandler := NewUpdateQuotationHandler(db)
	createQuotationHandler := NewCreateQuotationHandler(db)
	deleteQuotationHandler := NewDeleteQuotationHandler(db)

	server.AddHandler("GET /quotation", queryQuotationHandler.Handler)
	server.AddHandler("PUT /quotation", updateQuotationHandler.Handler)
	server.AddHandler("POST /quotation", createQuotationHandler.Handler)
	server.AddHandler("DELETE /quotation", deleteQuotationHandler.Handler)
}
