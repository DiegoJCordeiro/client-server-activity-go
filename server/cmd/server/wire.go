//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"database/sql"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/clients"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/database/repository"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/webserver/handlers/quotation_handler"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/usecases/quotation_usecase"
	"github.com/google/wire"
	"net/http"
)

var setQuotationRepositoryDependency = wire.NewSet(
	repository.NewQuotationRepositoryImpl,
	wire.Bind(new(repository.IQuotationRepository), new(*repository.QuotationRepositoryImpl)),
)

var setQuotationClientDependency = wire.NewSet(
	clients.NewQuotationClient,
	wire.Bind(new(clients.IClient), new(*clients.QuotationClient)),
)

var setQuotationQueryUseCaseDependency = wire.NewSet(
	quotation_usecase.NewQueryQuotationUseCase,
	wire.Bind(new(quotation_usecase.IQuotationUseCase), new(*quotation_usecase.QueryQuotationUseCase)),
)

var setQuotationUpdateUseCaseDependency = wire.NewSet(
	quotation_usecase.NewUpdateQuotationUseCase,
	wire.Bind(new(quotation_usecase.IQuotationUseCase), new(*quotation_usecase.UpdateQuotationUseCase)),
)

var setQuotationDeleteUseCaseDependency = wire.NewSet(
	quotation_usecase.NewDeleteQuotationUseCase,
	wire.Bind(new(quotation_usecase.IQuotationUseCase), new(*quotation_usecase.DeleteQuotationUseCase)),
)

var setCreateQuotationUseCaseDependency = wire.NewSet(
	quotation_usecase.NewCreateQuotationUseCase,
	wire.Bind(new(quotation_usecase.IQuotationUseCase), new(*quotation_usecase.CreateQuotationUseCase)),
)

func NewQueryQuotationUseCase(db *sql.DB, client *http.Client) *quotation_usecase.QueryQuotationUseCase {

	wire.Build(
		setQuotationRepositoryDependency,
		setQuotationClientDependency,
		quotation_usecase.NewQueryQuotationUseCase,
	)

	return &quotation_usecase.QueryQuotationUseCase{}
}

func NewQueryQuotationHandler(db *sql.DB, client *http.Client) *quotation_handler.QueryAllQuotationsHandler {

	wire.Build(
		setQuotationRepositoryDependency,
		setQuotationClientDependency,
		setQuotationQueryUseCaseDependency,
		quotation_handler.NewQueryAllQuotationsHandler,
	)

	return &quotation_handler.QueryAllQuotationsHandler{}
}

func NewCreateQuotationHandler(db *sql.DB) *quotation_handler.CreateQuotationHandler {

	wire.Build(
		setQuotationRepositoryDependency,
		setCreateQuotationUseCaseDependency,
		quotation_handler.NewCreateQuotationHandler,
	)

	return &quotation_handler.CreateQuotationHandler{}
}

func NewDeleteQuotationHandler(db *sql.DB) *quotation_handler.DeleteQuotationHandler {

	wire.Build(
		setQuotationRepositoryDependency,
		setQuotationDeleteUseCaseDependency,
		quotation_handler.NewDeleteQuotationsHandler,
	)

	return &quotation_handler.DeleteQuotationHandler{}
}

func NewUpdateQuotationHandler(db *sql.DB) *quotation_handler.UpdateQuotationHandler {

	wire.Build(
		setQuotationRepositoryDependency,
		setQuotationUpdateUseCaseDependency,
		quotation_handler.NewUpdateQuotationsHandler,
	)

	return &quotation_handler.UpdateQuotationHandler{}
}
