package request

import (
	"testing"

	"github.com/Aranaris/dwellstats/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatRequest(t *testing.T) {
	MockAddress := "5500 Grand Lake Dr, San Antonio, TX, 78244"

	//Test Parse Address returns correct Address struct
	a, err := ParseAddress(MockAddress)
	require.NoError(t, err)
	require.NotNil(t, a)
	assert.Equal(t, "5500 Grand Lake Dr", a.Street)
	assert.Equal(t, "San Antonio", a.City)
	assert.Equal(t, "TX", a.State)
	assert.Equal(t, "78244", a.Zip)

	//Test Get Property Endpoint returns correct Property struct
	MockProperty := models.Property{
		Bedrooms: 3,
		Bathrooms: 2,
		SquareFootage: 1878,
		YearBuilt: 1973,
		PropertyType: "Single Family",
		LastSaleDate: "2017-10-19T00:00:00.000Z",
		LastSalePrice: 185000,
	}
	MockEndpoint := "http://127.0.0.1:1337"

	p, err := GetPropertyByAddress(MockEndpoint, MockAddress)
	require.NoError(t, err)
	require.NotNil(t, p)
	assert.Equal(t, MockProperty.Bedrooms, p.Bedrooms)
	assert.Equal(t, MockProperty.Bathrooms, p.Bathrooms)
	assert.Equal(t, MockProperty.SquareFootage, p.SquareFootage)
	assert.Equal(t, MockProperty.YearBuilt, p.YearBuilt)
	assert.Equal(t, MockProperty.PropertyType, p.PropertyType)
	assert.Equal(t, MockProperty.LastSaleDate, p.LastSaleDate)
	assert.Equal(t, MockProperty.LastSalePrice, p.LastSalePrice)
}
