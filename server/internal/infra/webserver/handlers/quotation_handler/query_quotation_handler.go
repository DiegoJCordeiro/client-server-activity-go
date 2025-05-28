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
	Formatter  formatter.IFormatter
}

func NewQueryAllQuotationsHandler(repository repository.IQuotationRepository, useCase quotation_usecase.IQuotationUseCase, formatter formatter.IFormatter) *QueryAllQuotationsHandler {
	return &QueryAllQuotationsHandler{
		Repository: repository,
		UseCase:    useCase,
		Formatter:  formatter,
	}
}

// QueryHandler Query a quotation godoc
//
// @Summary     Query a quotation
// @Description This endpoint is used to Query a quotation.
// @Tags        Quotation
// @Accept      json
// @Produces    json
// @Success     200 {object} dto.QuotationOutputUseCaseDTO
// @Failure     500         {object}      dto.ErrorDTO
// @Router      /quotation  [get]
func (handler *QueryAllQuotationsHandler) QueryHandler(response http.ResponseWriter, request *http.Request) {

	var errorDto dto.ErrorDTO
	quotationDTO, err := handler.UseCase.Execute(dto.QuotationInputUseCaseDTO{})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errorDto = dto.ErrorDTO{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		_ = handler.Formatter.EncodeObjectToJson(errorDto, response)
		return
	}

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.WriteHeader(http.StatusOK)
	_ = handler.Formatter.EncodeObjectToJson(quotationDTO, response)
}
