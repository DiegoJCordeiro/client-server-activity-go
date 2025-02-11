package quotation_handler

import (
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/formatter"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/database/repository"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/usecases/quotation_usecase"
	"net/http"
)

type QueryAllQuotationsHandler struct {
	Repository repository.IQuotationRepository
	UseCase    quotation_usecase.IQuotationUseCase
}

func NewQueryAllQuotationsHandler(repository repository.IQuotationRepository, useCase quotation_usecase.IQuotationUseCase) *QueryAllQuotationsHandler {
	return &QueryAllQuotationsHandler{
		Repository: repository,
		UseCase:    useCase,
	}
}

func (handler *QueryAllQuotationsHandler) Handler(response http.ResponseWriter, request *http.Request) {

	var errorDto dto.ErrorDTO
	quotationDTO, err := handler.UseCase.Execute(dto.QuotationInputDTO{})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errorDto = dto.ErrorDTO{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		_ = formatter.EncodeObjectToJson(errorDto, response)
		return
	}

	response.WriteHeader(http.StatusOK)
	_ = formatter.EncodeObjectToJson(quotationDTO, response)
}
