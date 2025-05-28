package repository

import (
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/entity"
)

type IQuotationRepository interface {
	Create(quotationDTO dto.QuotationInputRepositoryDTO) (entity.Quotation, error)
	Update(quotationDTO dto.QuotationInputRepositoryDTO) (entity.Quotation, error)
	Delete(id string) (entity.Quotation, error)
	QueryAll() ([]entity.Quotation, error)
	QueryByID(id string) (entity.Quotation, error)
}
