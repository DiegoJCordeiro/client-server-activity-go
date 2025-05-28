package quotation_usecase

import (
	"encoding/json"
	"errors"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/clients"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/database/repository"
	"time"
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

func (quotationUseCase *QueryQuotationUseCase) Execute(dto.QuotationInputUseCaseDTO) (dto.QuotationOutputUseCaseDTO, error) {

	var quotationInputUseCaseDTO dto.QuotationInputUseCaseDTO
	var quotationInputRepositoryDTO dto.QuotationInputRepositoryDTO

	resp, err := quotationUseCase.Client.Call("GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		return dto.QuotationOutputUseCaseDTO{}, err
	}

	responseToJson, _ := json.Marshal(resp)

	if err := json.Unmarshal(responseToJson, &quotationInputUseCaseDTO); err == nil {

		quotationInputRepositoryDTO = dto.QuotationInputRepositoryDTO{
			Bid:       quotationInputUseCaseDTO.Bid,
			Ask:       quotationInputUseCaseDTO.Ask,
			Code:      quotationInputUseCaseDTO.Code,
			Timestamp: quotationInputUseCaseDTO.Timestamp,
		}

		quotationCreated, errUseCase := quotationUseCase.Repository.Create(quotationInputRepositoryDTO)

		if errUseCase != nil {
			return dto.QuotationOutputUseCaseDTO{}, errUseCase
		}

		return dto.QuotationOutputUseCaseDTO{
			ID:        quotationCreated.ID.String(),
			Bid:       quotationCreated.Bid,
			Ask:       quotationCreated.Ask,
			CreatedAt: quotationCreated.CreatedAt,
			UpdatedAt: time.Time{},
			DeletedAt: time.Time{},
		}, nil
	} else {
		return dto.QuotationOutputUseCaseDTO{}, errors.New("error to convert response")
	}
}
