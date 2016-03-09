package staddr


// Address represents a mailing address.
type Address struct {
	CountryName string
	Region      string
	Locality    string

	StreetAddress  string
	UnitNumber     string
	PostalCode     string

	BuildingNumber string
	StreetName     string
	StreetType     string
}
