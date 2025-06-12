package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Aranaris/dwellstats/internal/models"
)

func ParseAddress(address string) (*models.Address, error) {
	segments := strings.Split(address, ",")
	if len(segments) != 4 {
		return nil, fmt.Errorf("Invalid address string")
	}

	for i, v := range segments {
		segments[i] = strings.TrimSpace(v)
	}

	a := models.Address{
		Street: segments[0],
		City: segments[1],
		State: segments[2],
		Zip: segments[3],
	}

	return &a, nil
}

func GetPropertyByAddress(endpoint string, address *models.Address) (*models.Property, error) {

	qp := fmt.Sprintf("address=%s,%s,%s,%s", address.Street, address.City, address.State, address.Zip)
	query := url.QueryEscape(qp)
	req := fmt.Sprintf("%s?%s", endpoint, query)

	resp, err := http.Get(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var p models.Property

	err = json.Unmarshal(body, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

