package quotation_usecase

import "github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"

type IQuotationUseCase interface {
	Execute(input dto.QuotationInputUseCaseDTO) (dto.QuotationOutputUseCaseDTO, error)
}
