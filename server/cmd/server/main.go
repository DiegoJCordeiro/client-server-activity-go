package main

import (
	"database/sql"
	"fmt"
	"github.com/DiegoJCordeiro/golang-study/activity/server/cfg"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/webserver/handlers"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

// @title Server - GO Expert - Activity
// @version 1.0
// @description API to query USD current value
// @termsOfService http://swagger.io/terms/

// @contact.name Diego Cordeiro
// @contact.url https://github.com/DiegoJCordeiro/client-server-activity-go
// @contact.email diegocordeiro.contatos@gmail.com

// @license.name Diego Cordeiro License
// @license.url  https://github.com/DiegoJCordeiro/client-server-activity-go/blob/main/LICENSE

// @host localhost:8081
// @BasePath /
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

	server := handlers.NewWebServer(fmt.Sprintf(":%s", configuration.WebserverPort))

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

	server.AddHandler("GET /quotation", queryQuotationHandler.QueryHandler)
	server.AddHandler("PUT /quotation", updateQuotationHandler.UpdateHandler)
	server.AddHandler("POST /quotation", createQuotationHandler.CreateHandler)
	server.AddHandler("DELETE /quotation", deleteQuotationHandler.DeleteHandler)
}
