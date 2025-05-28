package quotation_usecase

import (
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/database/repository"
)

type UpdateQuotationUseCase struct {
	Repository repository.IQuotationRepository
}

func NewUpdateQuotationUseCase(repository repository.IQuotationRepository) *UpdateQuotationUseCase {
	return &UpdateQuotationUseCase{
		Repository: repository,
	}
}

func (quotationUseCase *UpdateQuotationUseCase) Execute(inputUseCase dto.QuotationInputUseCaseDTO) (dto.QuotationOutputUseCaseDTO, error) {

	var inputRepository = dto.QuotationInputRepositoryDTO{
		ID:        inputUseCase.ID,
		Bid:       inputUseCase.Bid,
		Ask:       inputUseCase.Ask,
		Code:      inputUseCase.Code,
		Timestamp: inputUseCase.Timestamp,
	}

	quotationUpdated, err := quotationUseCase.Repository.Update(inputRepository)

	if err != nil {
		return dto.QuotationOutputUseCaseDTO{}, err
	}

	return dto.QuotationOutputUseCaseDTO{
		ID:        quotationUpdated.ID.String(),
		Bid:       quotationUpdated.Bid,
		Ask:       quotationUpdated.Ask,
		CreatedAt: quotationUpdated.CreatedAt,
		UpdatedAt: quotationUpdated.UpdatedAt,
		DeletedAt: quotationUpdated.DeletedAt,
	}, nil
}
