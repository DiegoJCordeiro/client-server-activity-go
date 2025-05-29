package usecases

import (
	"github.com/DiegoJCordeiro/golang-study/activity/client/internal/adapters/output"
	"github.com/DiegoJCordeiro/golang-study/activity/client/internal/application/services"
	"log"
)

type QuotationQueryUseCase interface {
	ExecuteQuery() error
}

type QuotationQueryUseCaseImpl struct {
	textEngineService     services.TextEngineService
	quotationServerClient output.QuotationServerClient
}

func NewQuotationQueryUseCase(textEngineService services.TextEngineService, quotationServerClient output.QuotationServerClient) QuotationQueryUseCase {
	return &QuotationQueryUseCaseImpl{
		textEngineService:     textEngineService,
		quotationServerClient: quotationServerClient,
	}
}

func (quotationQueryUseCase *QuotationQueryUseCaseImpl) ExecuteQuery() error {

	quotationFound, err := quotationQueryUseCase.quotationServerClient.Call()

	if err != nil {
		return err
	}

	err = quotationQueryUseCase.textEngineService.WriteText("./", quotationFound)

	if err != nil {
		log.Fatal("Erro ao escrever no arquivo:", err)
		return err
	}

	log.Println("Cotação salva com sucesso:", quotationFound)

	return nil
}
