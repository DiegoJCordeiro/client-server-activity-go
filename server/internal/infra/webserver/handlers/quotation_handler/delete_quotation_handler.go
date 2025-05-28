package quotation_handler

import (
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

// DeleteHandler Delete a quotation godoc
//
// @Summary     Delete a quotation
// @Description This endpoint is used to Delete a quotation.
// @Tags        Quotation
// @Accept      json
// @Produces    json
// @Param 		id 	query string 	true "id quotation"
// @Success     200 {object} dto.QuotationOutputUseCaseDTO
// @Failure     500         {object}      dto.ErrorDTO
// @Router      /quotation  [delete]
func (handler *DeleteQuotationHandler) DeleteHandler(response http.ResponseWriter, request *http.Request) {

	var errDto dto.ErrorDTO
	var quotationDeleted dto.QuotationOutputUseCaseDTO

	paramId := request.URL.Query().Get("id")

	var input dto.QuotationInputUseCaseDTO = dto.QuotationInputUseCaseDTO{
		ID: paramId,
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

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusOK)
	_ = handler.Formatter.EncodeObjectToJson(quotationDeleted, response)
}
