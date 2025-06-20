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
		return nil, fmt.Errorf("invalid address string")
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

func GetPropertyByAddress(endpoint string, address string) (*models.Property, error) {
	
	base, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	q := base.Query()
	q.Set("address", address)

	base.RawQuery = q.Encode()

	fmt.Println(base.String())

	resp, err := http.Get(base.String())
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

