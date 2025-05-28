package dto

import (
	"encoding/json"
	"fmt"
	"time"
)

type QuotationDTO struct {
	CurrencyDTO CurrencyDTO `json:"USDBRL"`
}

type CurrencyDTO struct {
	Code       string    `json:"code"`
	CodeIn     string    `json:"codein"`
	Name       string    `json:"name"`
	High       string    `json:"high"`
	Low        string    `json:"low"`
	VarBid     string    `json:"varBid"`
	PctChange  string    `json:"pctChange"`
	Bid        string    `json:"bid"`
	Ask        string    `json:"ask"`
	Timestamp  string    `json:"timestamp"`
	CreateDate time.Time `json:"create_date"`
}

type QuotationClientOutput struct {
	Currency CurrencyClientOutput `json:"USDBRL"`
}

type CurrencyClientOutput struct {
	Code       string `json:"code"`
	CodeIn     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type QuotationInputHandlerCreateDTO struct {
	ID  string `json:"id"`
	Bid string `json:"bid"`
	Ask string `json:"ask"`
}

type QuotationInputHandlerUpdateDTO struct {
	*QuotationInputHandlerCreateDTO
}

type QuotationInputUseCaseDTO struct {
	ID        string `json:"id"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	Code      string `json:"code"`
	CodeIn    string `json:"codein"`
	Name      string `json:"name"`
	High      string `json:"high"`
	Low       string `json:"low"`
	VarBid    string `json:"varBid"`
	PctChange string `json:"pctChange"`
	Timestamp string `json:"timestamp"`
}

type QuotationInputRepositoryDTO struct {
	ID        string    `json:"id"`
	Bid       string    `json:"bid"`
	Ask       string    `json:"ask"`
	Code      string    `json:"code"`
	Timestamp string    `json:"timestamp"`
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt time.Time `json:"deletedAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type QuotationOutputUseCaseDTO struct {
	ID        string    `json:"id"`
	Bid       string    `json:"bid"`
	Ask       string    `json:"ask"`
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt time.Time `json:"deletedAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (quotation *QuotationDTO) UnmarshalJSON(data []byte) error {

	const layout = "2006-01-02 15:04:05"

	var quotationClientResponse = new(QuotationClientOutput)

	if err := json.Unmarshal(data, quotationClientResponse); err != nil {
		return err
	}

	var err error

	quotation.CurrencyDTO.Code = quotationClientResponse.Currency.Code
	quotation.CurrencyDTO.CodeIn = quotationClientResponse.Currency.CodeIn
	quotation.CurrencyDTO.Name = quotationClientResponse.Currency.Name
	quotation.CurrencyDTO.High = quotationClientResponse.Currency.High
	quotation.CurrencyDTO.Low = quotationClientResponse.Currency.Low
	quotation.CurrencyDTO.VarBid = quotationClientResponse.Currency.VarBid
	quotation.CurrencyDTO.PctChange = quotationClientResponse.Currency.PctChange
	quotation.CurrencyDTO.Bid = quotationClientResponse.Currency.Bid
	quotation.CurrencyDTO.Ask = quotationClientResponse.Currency.Ask
	quotation.CurrencyDTO.Timestamp = quotationClientResponse.Currency.Timestamp
	quotation.CurrencyDTO.CreateDate, err = time.Parse(layout, quotationClientResponse.Currency.CreateDate)

	if err != nil {
		return fmt.Errorf("falha ao parsear CreateDate: %v", err)
	}

	return nil
}

func (quotation *QuotationInputUseCaseDTO) UnmarshalJSON(data []byte) error {

	var quotationClientResponse = new(QuotationClientOutput)

	if err := json.Unmarshal(data, quotationClientResponse); err != nil {
		return err
	}

	if quotationClientResponse.Currency.Code != "" {

		quotation.Code = quotationClientResponse.Currency.Code
		quotation.CodeIn = quotationClientResponse.Currency.CodeIn
		quotation.Name = quotationClientResponse.Currency.Name
		quotation.High = quotationClientResponse.Currency.High
		quotation.Low = quotationClientResponse.Currency.Low
		quotation.VarBid = quotationClientResponse.Currency.VarBid
		quotation.PctChange = quotationClientResponse.Currency.PctChange
		quotation.Bid = quotationClientResponse.Currency.Bid
		quotation.Ask = quotationClientResponse.Currency.Ask
		quotation.Timestamp = quotationClientResponse.Currency.Timestamp

		return nil
	}

	return nil
}
