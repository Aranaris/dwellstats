package models

type Address struct{
	Street string
	City string
	State string
	Zip string
}

type Property struct{
	ID string
	Address Address
	County string
	Bedrooms int
	Bathrooms int
	SquareFootage int
	YearBuilt int
	PropertyType string
}

