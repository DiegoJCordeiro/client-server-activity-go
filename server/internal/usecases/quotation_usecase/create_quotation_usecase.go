package quotation_usecase

import (
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/database/repository"
)

type CreateQuotationUseCase struct {
	Repository repository.IQuotationRepository
}

func NewCreateQuotationUseCase(repository repository.IQuotationRepository) *CreateQuotationUseCase {
	return &CreateQuotationUseCase{
		Repository: repository,
	}
}

func (quotationUseCase *CreateQuotationUseCase) Execute(input dto.QuotationInputDTO) (dto.QuotationOutputDTO, error) {

	quotationCreated, err := quotationUseCase.Repository.Create(input)

	if err != nil {
		return dto.QuotationOutputDTO{}, err
	}

	return dto.QuotationOutputDTO{
		ID:        quotationCreated.ID.String(),
		Bid:       quotationCreated.Bid,
		Ask:       quotationCreated.Ask,
		CreatedAt: quotationCreated.CreatedAt,
		UpdatedAt: quotationCreated.UpdatedAt,
		DeletedAt: quotationCreated.DeletedAt,
	}, nil
}
