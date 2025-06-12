package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatRequest(t *testing.T) {
	MockAddress := "5500 Grand Lake Dr, San Antonio, TX, 78244"

	a, err := ParseAddress(MockAddress)
	require.NoError(t, err)
	require.NotNil(t, a)
	assert.Equal(t, "5500 Grand Lake Dr", a.Street)
	assert.Equal(t, "San Antonio", a.City)
	assert.Equal(t, "TX", a.State)
	assert.Equal(t, "78244", a.Zip)

}
