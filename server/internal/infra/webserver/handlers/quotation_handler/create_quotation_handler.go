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

// CreateHandler Create a quotation godoc
//
// @Summary     Create a quotation
// @Description This endpoint is used to Query a quotation.
// @Tags        Quotation
// @Accept      json
// @Produces    json
// @Param       request      body   	dto.QuotationInputHandlerCreateDTO      true      "Quotation Request"
// @Success     200 {object} dto.QuotationOutputUseCaseDTO
// @Failure     500         {object}      dto.ErrorDTO
// @Router      /quotation  [post]
func (handler *CreateQuotationHandler) CreateHandler(response http.ResponseWriter, request *http.Request) {

	var errDto dto.ErrorDTO
	var input dto.QuotationInputHandlerCreateDTO
	var inputUseCase dto.QuotationInputUseCaseDTO
	var quotationCreated dto.QuotationOutputUseCaseDTO

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

	quotationCreated, err := handler.UseCase.Execute(inputUseCase)

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
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(quotationCreated)
}
