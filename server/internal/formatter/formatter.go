package formatter

import (
	"encoding/json"
	"encoding/xml"
	"github.com/google/uuid"
	"io"
)

type IFormatter interface {
	EncodeObjectToJson(object any, writer io.Writer) error
	EncodeObjectToXml(object any, writer io.Writer) error
	DecodeObjectToJson(object any, reader io.Reader) error
	DecodeObjectToXml(object any, reader io.Reader) error
}

type Formatter struct {
	code uuid.UUID
}

func NewFormatter() *Formatter {
	return &Formatter{code: uuid.New()}
}

func (formatter *Formatter) EncodeObjectToJson(object any, writer io.Writer) error {

	if err := json.NewEncoder(writer).Encode(&object); err != nil {
		return err
	}

	return nil
}

func (formatter *Formatter) EncodeObjectToXml(object any, writer io.Writer) error {

	if err := xml.NewEncoder(writer).Encode(&object); err != nil {
		return err
	}

	return nil
}

func (formatter *Formatter) DecodeObjectToJson(object any, reader io.Reader) error {

	if err := json.NewDecoder(reader).Decode(&object); err != nil {
		return err
	}

	return nil
}

func (formatter *Formatter) DecodeObjectToXml(object any, reader io.Reader) error {

	if err := xml.NewDecoder(reader).Decode(&object); err != nil {
		return err
	}

	return nil
}
