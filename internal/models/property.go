package models

type Address struct{
	Street string `json:"addressLine1"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zipCode"`
}

type Property struct {
	ID            string  `json:"id"`
	Address       string  `json:"formattedAddress"`
	Street 				string  `json:"addressLine1"`
	City   				string  `json:"city"`
	State  				string  `json:"state"`
	Zip    				string  `json:"zipCode"`
	County        string  `json:"county"`
	Bedrooms      int     `json:"bedrooms"`
	Bathrooms     int     `json:"bathrooms"`
	SquareFootage int     `json:"squareFootage"`
	YearBuilt     int     `json:"yearBuilt"`
	PropertyType  string  `json:"propertyType"`
	LotSize       int     `json:"lotSize,omitempty"`
	LastSaleDate  string  `json:"lastSaleDate,omitempty"`
	LastSalePrice int     `json:"lastSalePrice,omitempty"`
	HOA           struct {
		Fee int `json:"fee"`
	} `json:"hoa"`
}
