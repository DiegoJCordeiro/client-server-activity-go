package main

import (
	"github.com/DiegoJCordeiro/golang-study/activity/client/internal/adapters/input"
	"github.com/DiegoJCordeiro/golang-study/activity/client/internal/adapters/output"
	"github.com/DiegoJCordeiro/golang-study/activity/client/internal/application/services"
	"github.com/DiegoJCordeiro/golang-study/activity/client/internal/application/usecases"
	"github.com/DiegoJCordeiro/golang-study/activity/client/internal/infrastructure/textengine"
)

func main() {

	var textEngineInfra = textengine.NewTextEngine("./")
	err := textEngineInfra.AddTextFile("./", "cotacao", "txt")

	if err != nil {
		panic(err)
	}

	var textEngineService = services.NewTextEngineServiceImpl(textEngineInfra)
	var quotationServerClient = output.NewQuotationServerClient()
	var quotationQueryUseCase = usecases.NewQuotationQueryUseCase(textEngineService, quotationServerClient)
	var quotationQueryCli = input.NewQuotationQueryCli(quotationQueryUseCase)

	err = quotationQueryCli.Execute()

	if err != nil {
		panic(err)
	}
}
