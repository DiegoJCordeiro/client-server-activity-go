package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Quotation struct {
	ID        uuid.UUID
	Bid       string
	Ask       string
	Timestamp string
	CreatedAt time.Time
	DeletedAt time.Time
	UpdatedAt time.Time
}

func NewQuotation(id uuid.UUID, bid, ask, timestamp string, createdAt, deletedAt, updatedAt time.Time) (*Quotation, error) {
	return &Quotation{
		ID:        id,
		Bid:       bid,
		Ask:       ask,
		Timestamp: timestamp,
		CreatedAt: createdAt,
		DeletedAt: deletedAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (quotation *Quotation) IsValid() error {

	if quotation.Bid == "" {
		return errors.New("quotation_usecase bid cannot be empty")
	}

	if quotation.Ask == "" {
		return errors.New("quotation_usecase Ask cannot be empty")
	}

	if quotation.Timestamp == "" {
		return errors.New("quotation_usecase Timestamp cannot be empty")
	}

	return nil
}
