package staddr


import (
	"fmt"
)


func (address Address) debugString() string {

	return fmt.Sprintf(
`
Address{
	CountryName: %q,
	Region:      %q,
	Locality:    %q

	StreetAddress:  %q,
	UnitNumber:     %q,
	PostalCode:     %q,

	BuildingNumber: %q,
	StreetName:     %q,
	StreetType:     %q,
}
`,
		address.CountryName,
		address.Region,
		address.Locality,

		address.StreetAddress,
		address.UnitNumber,
		address.PostalCode,

		address.BuildingNumber,
		address.StreetName,
		address.StreetType)

}
