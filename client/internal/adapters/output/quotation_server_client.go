package output

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type QuotationServerClient interface {
	Call() (string, error)
}

type QuotationServerClientImpl struct {
}

func NewQuotationServerClient() QuotationServerClient {
	return &QuotationServerClientImpl{}
}

func (quotationServerClient *QuotationServerClientImpl) Call() (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/quotation", nil)

	if err != nil {
		log.Fatal("Erro ao criar requisição:", err)
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal("Erro ao obter cotação do servidor:", err)
		return "", err
	}

	defer resp.Body.Close()

	var result map[string]string

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal("Erro ao decodificar resposta:", err)
		return "", err
	}

	quotationFound, ok := result["bid"]
	if !ok {
		log.Fatal("Campo 'bid' não encontrado na resposta")
		return "", err
	}

	return quotationFound, nil
}
