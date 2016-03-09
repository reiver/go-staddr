package staddr


import (
	"strings"
)


var (
	upperCaseEnglishDirections = map[string]struct{}{
		"EAST":struct{}{},
		"E":struct{}{},
		"E.":struct{}{},

		"NORTH":struct{}{},
		"N":struct{}{},
		"N.":struct{}{},

		"NORTHEAST":struct{}{},
		"NE":struct{}{},
		"NE.":struct{}{},

		"NORTHWEST":struct{}{},
		"NW":struct{}{},
		"NW.":struct{}{},

		"SOUTH":struct{}{},
		"S":struct{}{},
		"S.":struct{}{},

		"SOUTHEAST":struct{}{},
		"SE":struct{}{},
		"SE.":struct{}{},

		"SOUTHWEST":struct{}{},
		"SW":struct{}{},
		"SW.":struct{}{},

		"WEST":struct{}{},
		"W":struct{}{},
		"W.":struct{}{},
	}

	upperCaseFrenchDirections = map[string]struct{}{
		"EST":struct{}{},
		"E":struct{}{},
		"E.":struct{}{},

		"NORD":struct{}{},
		"N":struct{}{},
		"N.":struct{}{},

		"NORD-EST":struct{}{},
		"NE":struct{}{},
		"NE.":struct{}{},

		"NORD-OUEST":struct{}{},
		"NO":struct{}{},
		"NO.":struct{}{},

		"SUD":struct{}{},
		"S":struct{}{},
		"S.":struct{}{},

		"SUD-EST":struct{}{},
		"SE":struct{}{},
		"SE.":struct{}{},

		"SUD-OUEST":struct{}{},
		"SO":struct{}{},
		"SO.":struct{}{},

		"OUEST":struct{}{},
		"O":struct{}{},
		"O.":struct{}{},
	}
)


func isCanadianDirection(s string) bool {
	_, isEnglishDirection := upperCaseEnglishDirections[ strings.ToUpper(s) ]
	_, isFrenchDirection  := upperCaseFrenchDirections[  strings.ToUpper(s) ]

	return isEnglishDirection || isFrenchDirection
}
