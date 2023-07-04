package cart

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/url"
)

type HttpOrderClient struct {
	client http.Client
}

func NewOrderClient(client http.Client) *HttpOrderClient {
	return &HttpOrderClient{
		client,
	}
}

type MakeOrderRequest struct {
	Items []Item `json:"items"`
}

type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

func (c *HttpOrderClient) MakeOrder(ctx context.Context, items []Item) (uuid.UUID, error) {
	log.Println("placed and order")

	reqBody, _ := json.Marshal(MakeOrderRequest{items})
	address, _ := url.Parse("http://localhost:8001/v1/orders")
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, address.String(), bytes.NewBuffer(reqBody))

	req.Header.Add("Content-Type", "application/json")

	log.Println("sending request to order service")
	res, err := c.client.Do(req)
	if err != nil {
		log.Println(err)
		return uuid.Nil, err
	}
	log.Println("request to order service succeeded")
	defer res.Body.Close()

	var resObj CreateResponse
	err = json.NewDecoder(res.Body).Decode(&resObj)
	if err != nil {
		log.Println(err)
	}

	return resObj.ID, err
}
