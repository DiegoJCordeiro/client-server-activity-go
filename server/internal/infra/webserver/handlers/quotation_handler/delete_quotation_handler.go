package quotation_handler

import (
	"encoding/json"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/formatter"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/database/repository"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/usecases/quotation_usecase"
	"net/http"
)

type DeleteQuotationHandler struct {
	Repository repository.IQuotationRepository
	UseCase    quotation_usecase.IQuotationUseCase
	Formatter  formatter.IFormatter
}

func NewDeleteQuotationsHandler(repository repository.IQuotationRepository, useCase quotation_usecase.IQuotationUseCase, formatter formatter.IFormatter) *DeleteQuotationHandler {
	return &DeleteQuotationHandler{
		Repository: repository,
		UseCase:    useCase,
		Formatter:  formatter,
	}
}

func (handler *DeleteQuotationHandler) Handler(response http.ResponseWriter, request *http.Request) {

	var errDto dto.ErrorDTO
	var input dto.QuotationInputDTO
	var quotationDeleted dto.QuotationOutputDTO

	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		errDto = dto.ErrorDTO{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		_ = handler.Formatter.EncodeObjectToJson(errDto, response)
		return
	}

	quotationDeleted, err := handler.UseCase.Execute(input)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errDto = dto.ErrorDTO{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		_ = handler.Formatter.EncodeObjectToJson(errDto, response)
		return
	}

	response.WriteHeader(http.StatusOK)
	_ = handler.Formatter.EncodeObjectToJson(quotationDeleted, response)
}
