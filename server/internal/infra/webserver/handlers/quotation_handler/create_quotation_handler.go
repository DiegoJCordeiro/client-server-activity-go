package quotation_handler

import (
	"encoding/json"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/formatter"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/database/repository"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/usecases/quotation_usecase"
	"net/http"
)

type CreateQuotationHandler struct {
	Repository repository.IQuotationRepository
	UseCase    quotation_usecase.IQuotationUseCase
	Formatter  formatter.IFormatter
}

func NewCreateQuotationHandler(repository repository.IQuotationRepository, useCase quotation_usecase.IQuotationUseCase, formatter formatter.IFormatter) *CreateQuotationHandler {
	return &CreateQuotationHandler{
		Repository: repository,
		UseCase:    useCase,
		Formatter:  formatter,
	}
}

func (handler *CreateQuotationHandler) Handler(response http.ResponseWriter, request *http.Request) {

	var errDto dto.ErrorDTO
	var input dto.QuotationInputDTO
	var quotationCreated dto.QuotationOutputDTO

	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		errDto = dto.ErrorDTO{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		_ = handler.Formatter.EncodeObjectToJson(errDto, response)
		return
	}

	quotationCreated, err := handler.UseCase.Execute(input)

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
	err = json.NewEncoder(response).Encode(quotationCreated)
}
