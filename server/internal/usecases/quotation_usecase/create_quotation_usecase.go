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

func (quotationUseCase *CreateQuotationUseCase) Execute(inputUseCase dto.QuotationInputUseCaseDTO) (dto.QuotationOutputUseCaseDTO, error) {

	var inputRepository = dto.QuotationInputRepositoryDTO{
		ID:        inputUseCase.ID,
		Bid:       inputUseCase.Bid,
		Ask:       inputUseCase.Ask,
		Code:      inputUseCase.Code,
		Timestamp: inputUseCase.Timestamp,
	}

	quotationCreated, err := quotationUseCase.Repository.Create(inputRepository)

	if err != nil {
		return dto.QuotationOutputUseCaseDTO{}, err
	}

	return dto.QuotationOutputUseCaseDTO{
		ID:        quotationCreated.ID.String(),
		Bid:       quotationCreated.Bid,
		Ask:       quotationCreated.Ask,
		CreatedAt: quotationCreated.CreatedAt,
		UpdatedAt: quotationCreated.UpdatedAt,
		DeletedAt: quotationCreated.DeletedAt,
	}, nil
}
