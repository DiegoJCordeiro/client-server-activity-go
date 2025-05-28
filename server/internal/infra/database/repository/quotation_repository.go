package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/entity"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"time"
)

type QuotationRepositoryImpl struct {
	DB      *sql.DB
	Queries sqlc.Queries
}

func NewQuotationRepositoryImpl(db *sql.DB) *QuotationRepositoryImpl {

	var quotationRepository = QuotationRepositoryImpl{
		DB:      db,
		Queries: *sqlc.New(db),
	}

	return &quotationRepository
}

func (quotation *QuotationRepositoryImpl) Create(quotationDTO dto.QuotationInputRepositoryDTO) (entity.Quotation, error) {

	id := uuid.New()
	ctx := context.Background()
	creationDate := time.Now()

	insertQuotationParams := sqlc.InsertQuotationParams{
		ID:        id.String(),
		Bid:       sql.NullString{String: quotationDTO.Bid, Valid: true},
		Ask:       sql.NullString{String: quotationDTO.Ask, Valid: true},
		Timestamp: sql.NullString{String: quotationDTO.Timestamp, Valid: true},
		CreatedAt: sql.NullTime{Time: creationDate, Valid: true},
	}

	err := quotation.Queries.InsertQuotation(ctx, insertQuotationParams)

	if err != nil {
		return entity.Quotation{}, errors.New("error to persist data")
	}

	return entity.Quotation{
		ID:        id,
		Bid:       quotationDTO.Bid,
		Ask:       quotationDTO.Ask,
		Timestamp: quotationDTO.Timestamp,
		CreatedAt: creationDate,
	}, nil
}

func (quotation *QuotationRepositoryImpl) Update(quotationDTO dto.QuotationInputRepositoryDTO) (entity.Quotation, error) {

	ctx := context.Background()
	updateDate := time.Now()

	updateQuotationParams := sqlc.UpdateQuotationParams{
		ID:        quotationDTO.ID,
		Bid:       sql.NullString{String: quotationDTO.Bid, Valid: true},
		Ask:       sql.NullString{String: quotationDTO.Ask, Valid: true},
		UpdatedAt: sql.NullTime{Time: updateDate, Valid: true},
	}

	err := quotation.Queries.UpdateQuotation(ctx, updateQuotationParams)

	if err != nil {
		return entity.Quotation{}, errors.New("error to update data")
	}

	idParsed, err := uuid.Parse(quotationDTO.ID)

	return entity.Quotation{
		ID:        idParsed,
		Bid:       quotationDTO.Bid,
		Ask:       quotationDTO.Ask,
		Timestamp: quotationDTO.Timestamp,
		UpdatedAt: updateDate,
	}, nil
}

func (quotation *QuotationRepositoryImpl) Delete(id string) (entity.Quotation, error) {

	ctx := context.Background()
	deletedDate := time.Now()

	deleteQuotationParams := sqlc.DeleteQuotationParams{
		DeletedAt: sql.NullTime{Time: deletedDate, Valid: true},
		ID:        id,
	}

	err := quotation.Queries.DeleteQuotation(ctx, deleteQuotationParams)

	if err != nil {
		return entity.Quotation{}, errors.New("error to delete data")
	}

	quotationFound, _ := quotation.QueryByID(id)

	return quotationFound, nil
}

func (quotation *QuotationRepositoryImpl) QueryAll() ([]entity.Quotation, error) {

	ctx := context.Background()

	quotations, err := quotation.Queries.QueryAllQuotation(ctx)

	if err != nil {
		return []entity.Quotation{}, errors.New("error to query all data")
	}

	quotationsSlice := make([]entity.Quotation, 0, cap(quotations))

	for _, v := range quotations {

		uuidAux, _ := uuid.Parse(v.ID)
		quotationsAux := entity.Quotation{
			ID:        uuidAux,
			Bid:       v.Bid.String,
			Ask:       v.Ask.String,
			Timestamp: v.Timestamp.String,
			CreatedAt: v.CreatedAt.Time,
			DeletedAt: v.DeletedAt.Time,
			UpdatedAt: v.UpdatedAt.Time,
		}

		quotationsSlice = append(quotationsSlice, quotationsAux)
	}

	return quotationsSlice, nil
}

func (quotation *QuotationRepositoryImpl) QueryByID(id string) (entity.Quotation, error) {

	ctx := context.Background()

	quotationFound, err := quotation.Queries.QueryQuotationById(ctx, id)

	if err != nil {
		return entity.Quotation{}, errors.New("error to query data")
	}

	idParsed, err := uuid.Parse(quotationFound.ID)

	if err != nil {
		return entity.Quotation{}, errors.New("error to parse data")
	}

	return entity.Quotation{
		ID:        idParsed,
		Ask:       quotationFound.Ask.String,
		Bid:       quotationFound.Bid.String,
		Timestamp: quotationFound.Timestamp.String,
		CreatedAt: quotationFound.CreatedAt.Time,
		DeletedAt: quotationFound.DeletedAt.Time,
		UpdatedAt: quotationFound.UpdatedAt.Time,
	}, nil
}
