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

func (quotationUseCase *UpdateQuotationUseCase) Execute(inputDto dto.QuotationInputDTO) (dto.QuotationOutputDTO, error) {

	quotationUpdated, err := quotationUseCase.Repository.Update(inputDto)

	if err != nil {
		return dto.QuotationOutputDTO{}, err
	}

	return dto.QuotationOutputDTO{
		ID:        quotationUpdated.ID.String(),
		Bid:       quotationUpdated.Bid,
		Ask:       quotationUpdated.Ask,
		CreatedAt: quotationUpdated.CreatedAt,
		UpdatedAt: quotationUpdated.UpdatedAt,
		DeletedAt: quotationUpdated.DeletedAt,
	}, nil
}
