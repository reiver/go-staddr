package staddr


import (
	"strings"
)


var (
	upperCaseCanadianStreetTypesAndAbbreviations = map[string]struct{}{
		"ABBEY":struct{}{},
		"ACRES":struct{}{},
		"ALLÉE":struct{}{},
		"ALLEY":struct{}{},
		"AUT":struct{}{},
		"AUT.":struct{}{},
		"AUTOROUTE":struct{}{},
		"AV":struct{}{},
		"AV.":struct{}{},
		"AVE":struct{}{},
		"AVE.":struct{}{},
		"AVENUE":struct{}{},

		"BAY":struct{}{},
		"BEACH":struct{}{},
		"BEND":struct{}{},
		"BLVD":struct{}{},
		"BLVD.":struct{}{},
		"BOUL":struct{}{},
		"BOUL.":struct{}{},
		"BOULEVARD":struct{}{},
		"BY-PASS":struct{}{},
		"BYPASS":struct{}{},
		"BYWAY":struct{}{},

		"C":struct{}{},
		"C.":struct{}{},
		"CAMPUS":struct{}{},
		"CAPE":struct{}{},
		"CAR":struct{}{},
		"CAR.":struct{}{},
		"CARRÉ":struct{}{},
		"CARREF":struct{}{},
		"CARREF.":struct{}{},
		"CARREFOUR":struct{}{},
		"CDS":struct{}{},
		"CDS.":struct{}{},
		"CENTRE":struct{}{},
		"CERCLE":struct{}{},
		"CH":struct{}{},
		"CH.":struct{}{},
		"CHASE":struct{}{},
		"CHEMIN":struct{}{},
		"CIR":struct{}{},
		"CIR.":struct{}{},
		"CIRCLE":struct{}{},
		"CIRCT":struct{}{},
		"CIRCT.":struct{}{},
		"CIRCUIT":struct{}{},
		"CLOSE":struct{}{},
		"COMMON":struct{}{},
		"CONC":struct{}{},
		"CONC.":struct{}{},
		"CONCESSION":struct{}{},
		"CORNERS":struct{}{},
		"CÔTE":struct{}{},
		"COUR":struct{}{},
		"COURT":struct{}{},
		"COURS":struct{}{},
		"COVE":struct{}{},
		"CRES":struct{}{},
		"CRES.":struct{}{},
		"CRESCENT":struct{}{},
		"CRNRS":struct{}{},
		"CRNRS.":struct{}{},
		"CROIS":struct{}{},
		"CROIS.":struct{}{},
		"CROISSANT":struct{}{},
		"CROSS":struct{}{},
		"CROSS.":struct{}{},
		"CROSSING":struct{}{},
		"CRT":struct{}{},
		"CRT.":struct{}{},
		"CTR":struct{}{},
		"CTR.":struct{}{},
		"CUL-DE-SAC":struct{}{},

		"DALE":struct{}{},
		"DELL":struct{}{},
		"DOWNS":struct{}{},
		"DIVERS":struct{}{},
		"DIVERS.":struct{}{},
		"DIVERSION":struct{}{},
		"DR":struct{}{},
		"DR.":struct{}{},
		"DRIVE":struct{}{},

		"ÉCH":struct{}{},
		"ÉCH.":struct{}{},
		"ÉCHANGEUR":struct{}{},
		"END":struct{}{},
		"ESPLANADE":struct{}{},
		"ESPL":struct{}{},
		"ESPL.":struct{}{},
		"ESTATE":struct{}{},
		"ESTATE.":struct{}{},
		"ESTATES":struct{}{},
		"EXPRESSWAY":struct{}{},
		"EXPY":struct{}{},
		"EXPY.":struct{}{},
		"EXTEN":struct{}{},
		"EXTEN.":struct{}{},
		"EXTENSION":struct{}{},

		"FARM":struct{}{},
		"FIELD":struct{}{},
		"FOREST":struct{}{},
		"FREEWAY":struct{}{},
		"FRONT":struct{}{},
		"FWY":struct{}{},
		"FWY.":struct{}{},

		"GARDENS":struct{}{},
		"GATE":struct{}{},
		"GDNS":struct{}{},
		"GDNS.":struct{}{},
		"GLADE":struct{}{},
		"GLEN":struct{}{},
		"GREEN":struct{}{},
		"GROUNDS":struct{}{},
		"GRNDS":struct{}{},
		"GRNDS.":struct{}{},
		"GROVE":struct{}{},

		"HARBOUR":struct{}{},
		"HARBR":struct{}{},
		"HARBR.":struct{}{},
		"HEATH":struct{}{},
		"HEIGHTS":struct{}{},
		"HGHLDS":struct{}{},
		"HGHLDS.":struct{}{},
		"HIGHLANDS":struct{}{},
		"HIGHWAY":struct{}{},
		"HILL":struct{}{},
		"HOLLOW":struct{}{},
		"HTS":struct{}{},
		"HTS.":struct{}{},
		"HWY":struct{}{},
		"HWY.":struct{}{},

		"ÎLE":struct{}{},
		"IMP":struct{}{},
		"IMP.":struct{}{},
		"IMPASSE":struct{}{},
		"INLET":struct{}{},
		"ISLAND":struct{}{},

		"KEY":struct{}{},
		"KNOLL":struct{}{},

		"LANDING":struct{}{},
		"LANDNG":struct{}{},
		"LANE":struct{}{},
		"LIMITS":struct{}{},
		"LINE":struct{}{},
		"LINK":struct{}{},
		"LKOUT":struct{}{},
		"LKOUT.":struct{}{},
		"LMTS":struct{}{},
		"LMTS.":struct{}{},
		"LOOKOUT":struct{}{},
		"LOOP":struct{}{},

		"MALL":struct{}{},
		"MANOR":struct{}{},
		"MAZE":struct{}{},
		"MEADOW":struct{}{},
		"MEWS":struct{}{},
		"MONTÉE":struct{}{},
		"MOOR":struct{}{},
		"MOUNT":struct{}{},
		"MOUNTAIN":struct{}{},
		"MTN":struct{}{},
		"MTN.":struct{}{},

		"ORCH":struct{}{},
		"ORCH.":struct{}{},
		"ORCHARD":struct{}{},

		"PARADE":struct{}{},
		"PARC":struct{}{},
		"PARK":struct{}{},
		"PARKWAY":struct{}{},
		"PASS":struct{}{},
		"PASSAGE":struct{}{},
		"PATH":struct{}{},
		"PATHWAY":struct{}{},
		"PINES":struct{}{},
		"PK":struct{}{},
		"PK.":struct{}{},
		"PKY":struct{}{},
		"PKY.":struct{}{},
		"PL":struct{}{},
		"PL.":struct{}{},
		"PLACE":struct{}{},
		"PLATEAU":struct{}{},
		"PLAT":struct{}{},
		"PLAT.":struct{}{},
		"PLAZA":struct{}{},
		"POINT":struct{}{},
		"POINTE":struct{}{},
		"PORT":struct{}{},
		"PRIVATE":struct{}{},
		"PROM":struct{}{},
		"PROM.":struct{}{},
		"PROMENADE":struct{}{},
		"PT":struct{}{},
		"PT.":struct{}{},
		"PTWAY":struct{}{},
		"PTWAY.":struct{}{},
		"PVT":struct{}{},
		"PVT.":struct{}{},

		"QUAI":struct{}{},
		"QUAY":struct{}{},

		"RAMP":struct{}{},
		"RANG":struct{}{},
		"RANGE":struct{}{},
		"RD":struct{}{},
		"RD.":struct{}{},
		"RDPT":struct{}{},
		"RDPT.":struct{}{},
		"RG":struct{}{},
		"RG.":struct{}{},
		"RIDGE":struct{}{},
		"RISE":struct{}{},
		"RLE":struct{}{},
		"RLE.":struct{}{},
		"ROAD":struct{}{},
		"ROND-POINT":struct{}{},
		"ROUTE":struct{}{},
		"ROW":struct{}{},
		"RTE":struct{}{},
		"RTE.":struct{}{},
		"RUE":struct{}{},
		"RUELLE":struct{}{},
		"RUN":struct{}{},

		"SENT":struct{}{},
		"SENT.":struct{}{},
		"SENTIER":struct{}{},
		"SQ":struct{}{},
		"SQ.":struct{}{},
		"SQUARE":struct{}{},
		"ST":struct{}{},
		"ST.":struct{}{},
		"STREET":struct{}{},
		"SUBDIV":struct{}{},
		"SUBDIV.":struct{}{},
		"SUBDIVISION":struct{}{},

		"TERR":struct{}{},
		"TERR.":struct{}{},
		"TERRACE":struct{}{},
		"TERRASSE":struct{}{},
		"THICK":struct{}{},
		"THICK.":struct{}{},
		"THICKET":struct{}{},
		"TLINE":struct{}{},
		"TLINE.":struct{}{},
		"TOWERS":struct{}{},
		"TOWNLINE":struct{}{},
		"TRAIL":struct{}{},
		"TRNABT":struct{}{},
		"TRNABT.":struct{}{},
		"TSSE":struct{}{},
		"TSSE.":struct{}{},
		"TURNABOUT":struct{}{},

		"VALE":struct{}{},
		"VIA":struct{}{},
		"VIEW":struct{}{},
		"VILLAS":struct{}{},
		"VILLGE":struct{}{},
		"VILLGE.":struct{}{},
		"VISTA":struct{}{},
		"VOIE":struct{}{},

		"WALK":struct{}{},
		"WAY":struct{}{},
		"WHARF":struct{}{},
		"WOOD":struct{}{},
		"WYND":struct{}{},
	}
)


func isCanadianStreetType(s string) bool {
	_, isStreetType := upperCaseCanadianStreetTypesAndAbbreviations[ strings.ToUpper(s) ]

	return isStreetType
}
