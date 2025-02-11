package quotation_usecase

import "github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"

type IQuotationUseCase interface {
	Execute(input dto.QuotationInputDTO) (dto.QuotationOutputDTO, error)
}
