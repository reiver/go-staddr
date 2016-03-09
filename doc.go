/*
Pacakge staddr provides tools for dealing with mailing addresses, such splitting up a street address to building number, street name, and street type.

Likely you would load the fields of this, and then call
the normalize method, perhaps to split up a street address
into a building number, street name, and street type.

For example:

	address := Address{
		CountryName: "Canada",
		Region:      "British Columbia",
		Locality:    "Vancouver",
		
		StreetAddress: "70 Murray St. West",
		UnitNumber:    "123",
		PostalCode:    "W3W 2F2",
		
		BuildingNumber: "", // <- We don't have this value.
		StreetName:     "", // <- We don't have this value.
		StreetType:     "", // <- We don't have this value.
	}
	
	err := address.Normalize()
	if nil != err {
		//@TODO: Deal with error.
		panic(err)
	}
	
	// Now what we have is:
	//
	// address := Address{
	//	CountryName: "Canada",
	//	Region:      "British Columbia",
	//	Locality:    "Vancouver",
	//
	//	StreetAddress: "", // <- Note that this is empty now.
	//	UnitNumber:    "123",
	//	PostalCode:    "W3W 2F2",
	//
	//	BuildingNumber: "70",       // <- Note that this has a value now.
	//	StreetName:     "Murray",   // <- Note that this has a value now.
	//	StreetType:     "St. West", // <- Note that this has a value now.
	//}
*/
package staddr
