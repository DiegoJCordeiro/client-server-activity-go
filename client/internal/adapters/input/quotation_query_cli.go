package input

import "github.com/DiegoJCordeiro/golang-study/activity/client/internal/application/usecases"

type QuotationQueryCli interface {
	Execute() error
}

type QuotationQueryCliImpl struct {
	useCase usecases.QuotationQueryUseCase
}

func NewQuotationQueryCli(quotationQueryUseCase usecases.QuotationQueryUseCase) QuotationQueryCli {
	return &QuotationQueryCliImpl{
		useCase: quotationQueryUseCase,
	}
}

func (quotationQueryCli *QuotationQueryCliImpl) Execute() error {

	return quotationQueryCli.useCase.ExecuteQuery()
}
