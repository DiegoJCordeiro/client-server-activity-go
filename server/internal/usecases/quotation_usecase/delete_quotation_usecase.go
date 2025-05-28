package quotation_usecase

import (
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/database/repository"
)

type DeleteQuotationUseCase struct {
	Repository repository.IQuotationRepository
}

func NewDeleteQuotationUseCase(repository repository.IQuotationRepository) *DeleteQuotationUseCase {
	return &DeleteQuotationUseCase{
		Repository: repository,
	}
}

func (quotationUseCase *DeleteQuotationUseCase) Execute(inputDto dto.QuotationInputUseCaseDTO) (dto.QuotationOutputUseCaseDTO, error) {

	quotationDeleted, err := quotationUseCase.Repository.Delete(inputDto.ID)

	if err != nil {
		return dto.QuotationOutputUseCaseDTO{}, err
	}

	return dto.QuotationOutputUseCaseDTO{
		ID:        quotationDeleted.ID.String(),
		Bid:       quotationDeleted.Bid,
		Ask:       quotationDeleted.Ask,
		CreatedAt: quotationDeleted.CreatedAt,
		UpdatedAt: quotationDeleted.UpdatedAt,
		DeletedAt: quotationDeleted.DeletedAt,
	}, nil
}
