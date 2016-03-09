package staddr


import (
	"fmt"
	"strings"
)


func countryIsCanada(countryName string) bool {

	switch strings.ToUpper(countryName) {
	case "CANADA", "CA":
		return true
	default:
		return false
	}
}


func normalizeCanada(address *Address) error {

	contextAddress := *address

	if countryName := contextAddress.CountryName; !countryIsCanada(countryName) {
//@TODO: Make a better error.
		return fmt.Errorf("Country %q not supported.", countryName)
	}



	if "" == contextAddress.StreetAddress {
		return nil
	}

	fields := strings.Fields(contextAddress.StreetAddress)
	lenFields := len(fields)
	if 3 > lenFields {
//@TODO: Make a better error.
		return fmt.Errorf("Could not normalize street address: %q.", contextAddress.StreetAddress)
	}
	switch {
	default:
//@TODO: Make a better error.
		return fmt.Errorf("Could not normalize street address: %q.", contextAddress.StreetAddress)
	case 3 == lenFields:
		address.BuildingNumber = fields[0]
		address.StreetName     = fields[1]
		address.StreetType     = fields[2]

		address.StreetAddress = ""

	case 3 <= lenFields && isCanadianStreetType(fields[lenFields-1]):
		address.BuildingNumber = fields[0]
		address.StreetName     = strings.Join(fields[1:lenFields-1], " ")
		address.StreetType     = fields[ lenFields-1 ]

		address.StreetAddress = ""

	case 4 <= lenFields && isCanadianDirection(fields[lenFields-3]) && isCanadianStreetType(fields[lenFields-1]):
		address.BuildingNumber = fields[0]
		address.StreetName     = strings.Join(fields[1:lenFields-1], " ")
		address.StreetType     = fields[ lenFields-1 ]

		address.StreetAddress = ""

	case 4 <= lenFields && isCanadianStreetType(fields[lenFields-2]) && isCanadianDirection(fields[lenFields-1]):
		address.BuildingNumber = fields[0]
		address.StreetName     = strings.Join(fields[1:lenFields-2], " ")
		address.StreetType     = strings.Join(fields[lenFields-2:], " ")

		address.StreetAddress = ""
	}


        return nil
}
