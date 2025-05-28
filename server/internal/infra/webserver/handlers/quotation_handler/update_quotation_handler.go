package quotation_handler

import (
	"encoding/json"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/formatter"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/database/repository"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/usecases/quotation_usecase"
	"net/http"
)

type UpdateQuotationHandler struct {
	Repository repository.IQuotationRepository
	UseCase    quotation_usecase.IQuotationUseCase
	Formatter  formatter.IFormatter
}

func NewUpdateQuotationsHandler(repository repository.IQuotationRepository, useCase quotation_usecase.IQuotationUseCase, formatter formatter.IFormatter) *UpdateQuotationHandler {

	return &UpdateQuotationHandler{
		Repository: repository,
		UseCase:    useCase,
		Formatter:  formatter,
	}
}

// UpdateHandler Update a quotation godoc
//
// @Summary     Update a quotation
// @Description This endpoint is used to Update a quotation.
// @Tags        Quotation
// @Accept      json
// @Produces    json
// @Param       request      body   	dto.QuotationInputHandlerUpdateDTO      true      "Quotation Request"
// @Success     204 {object} dto.QuotationOutputUseCaseDTO
// @Failure     500         {object}      dto.ErrorDTO
// @Router      /quotation  [put]
func (handler *UpdateQuotationHandler) UpdateHandler(response http.ResponseWriter, request *http.Request) {

	var errDto dto.ErrorDTO
	var input dto.QuotationInputHandlerUpdateDTO
	var inputUseCase dto.QuotationInputUseCaseDTO
	var quotationUpdated dto.QuotationOutputUseCaseDTO

	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		errDto = dto.ErrorDTO{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		_ = handler.Formatter.EncodeObjectToJson(errDto, response)
		return
	}

	inputUseCase = dto.QuotationInputUseCaseDTO{
		ID:  input.ID,
		Bid: input.Bid,
		Ask: input.Ask,
	}

	quotationUpdated, err := handler.UseCase.Execute(inputUseCase)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errDto = dto.ErrorDTO{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		_ = handler.Formatter.EncodeObjectToJson(errDto, response)
		return
	}

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusNoContent)
	_ = handler.Formatter.EncodeObjectToJson(quotationUpdated, response)
}
