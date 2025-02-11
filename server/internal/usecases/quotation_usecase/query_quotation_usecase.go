package quotation_usecase

import (
	"errors"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/clients"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/database/repository"
)

type QueryQuotationUseCase struct {
	Repository repository.IQuotationRepository
	Client     clients.IClient
}

func NewQueryQuotationUseCase(repository repository.IQuotationRepository, client clients.IClient) *QueryQuotationUseCase {
	return &QueryQuotationUseCase{
		Repository: repository,
		Client:     client,
	}
}

func (quotationUseCase *QueryQuotationUseCase) Execute(dto.QuotationInputDTO) (dto.QuotationOutputDTO, error) {

	resp, err := quotationUseCase.Client.Call("GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		return dto.QuotationOutputDTO{}, err
	}

	if quotationDTO, ok := resp.(dto.QuotationInputDTO); ok {

		quotationCreated, errUseCase := quotationUseCase.Repository.Create(quotationDTO)

		if errUseCase != nil {
			return dto.QuotationOutputDTO{}, errUseCase
		}

		return dto.QuotationOutputDTO{
			ID:        quotationCreated.ID.String(),
			Bid:       quotationCreated.Bid,
			Ask:       quotationCreated.Ask,
			CreatedAt: quotationCreated.CreatedAt,
			UpdatedAt: quotationCreated.UpdatedAt,
			DeletedAt: quotationCreated.DeletedAt,
		}, nil
	} else {
		return dto.QuotationOutputDTO{}, errors.New("error to convert response")
	}
}
