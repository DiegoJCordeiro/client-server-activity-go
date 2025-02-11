package clients

import (
	"encoding/json"
	"errors"
	"github.com/DiegoJCordeiro/golang-study/activity/server/internal/dto"
	"io"
	"net"
	"net/http"
)

type IClient interface {
	Call(httpMethod string, apiUrl string, Body io.Reader) (any, error)
}

type QuotationClient struct {
	Client *http.Client
}

func NewQuotationClient(client *http.Client) *QuotationClient {
	return &QuotationClient{
		Client: client,
	}
}

func (quotationClient *QuotationClient) Call(httpMethod string, apiUrl string, body io.Reader) (any, error) {

	req, err := http.NewRequest(
		httpMethod,
		apiUrl,
		body,
	)

	if err != nil {
		return nil, errors.New("erro ao criar requisição para API de câmbio")
	}

	resp, err := quotationClient.Client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			return nil, errors.New("tempo de requisição excedido.")
		}
		if resp.StatusCode != http.StatusOK {
			return nil, errors.New("API de câmbio retornou erro")
		}
		return nil, errors.New("erro ao obter dados da API de câmbio")
	}

	var quotationResponse dto.QuotationDTO

	if err := json.NewDecoder(resp.Body).Decode(&quotationResponse); err != nil {
		return nil, errors.New("Erro ao decodificar resposta da API")
	}

	return quotationResponse, nil
}
