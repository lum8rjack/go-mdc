package mdc

const (
	HarvestDataMinYear   int    = 2016
	HarvestDataUrl       string = "https://extra.mdc.mo.gov/widgets/wtrfwl_harvest/dataJSONservice.php"
	HarvestDataOrigin    string = "https://extra.mdc.mo.gov"
	HarvestDataReferer   string = "https://extra.mdc.mo.gov/widgets/wtrfwl_harvest/"
	HarvestDataUserAgent string = "go-mdc http client"
)

var Areas = [15]Area{
	{
		Name:      "B.K. Leach",
		Key:       "8514",
		Zipcode:   63347,
		City:      "Folley",
		Longitude: -90.759,
		Latitude:  39.07,
		URL:       "https://mdc.mo.gov/discover-nature/places/b-k-leach-memorial-conservation-area",
	},
	{
		Name:      "Ten Mile Pond",
		Key:       "8241",
		Zipcode:   63845,
		City:      "East Prairie",
		Longitude: -89.361,
		Latitude:  36.744,
		URL:       "https://mdc.mo.gov/discover-nature/places/ten-mile-pond-conservation-area",
	},
	{
		Name:      "Otter Slough",
		Key:       "5004",
		Zipcode:   63841,
		City:      "Liberty Township",
		Longitude: -89.974,
		Latitude:  36.784,
		URL:       "https://mdc.mo.gov/discover-nature/places/otter-slough-conservation-area",
	},
	{
		Name:      "Schell-Osage",
		Key:       "5701",
		Zipcode:   64783,
		City:      "Schell City",
		Longitude: -94.1,
		Latitude:  38.006,
		URL:       "https://mdc.mo.gov/discover-nature/places/schell-osage-conservation-area",
	},
	{
		Name:      "Four Rivers",
		Key:       "8238",
		Zipcode:   64779,
		City:      "Rich Hill",
		Longitude: -94.382,
		Latitude:  38.086,
		URL:       "https://mdc.mo.gov/discover-nature/places/august-busch-jr-memorial-wetlands-four-rivers-conservation-area",
	},
	{
		Name:      "Montrose",
		Key:       "5604",
		Zipcode:   64770,
		City:      "Montrose",
		Longitude: -93.99,
		Latitude:  38.278,
		URL:       "https://mdc.mo.gov/discover-nature/places/montrose-conservation-area",
	},
	{
		Name:      "Duck Creek",
		Key:       "5001",
		Zipcode:   63960,
		City:      "Puxico",
		Longitude: -90.131,
		Latitude:  36.957,
		URL:       "https://mdc.mo.gov/discover-nature/places/duck-creek-conservation-area",
	},
	{
		Name:      "Columbia Bottoms",
		Key:       "9736",
		Zipcode:   63138,
		City:      "St. Louis",
		Longitude: -90.21,
		Latitude:  38.788,
		URL:       "https://mdc.mo.gov/discover-nature/places/columbia-bottom-conservation-area",
	},
	{
		Name:      "Marais Temps Clair",
		Key:       "7902",
		Zipcode:   63301,
		City:      "St Charles",
		Longitude: -90.507,
		Latitude:  38.809,
		URL:       "https://mdc.mo.gov/discover-nature/places/marais-temps-clair-conservation-area",
	},
	{
		Name:      "Ted Shanks",
		Key:       "7011",
		Zipcode:   63433,
		City:      "Ashburn",
		Longitude: -91.182,
		Latitude:  39.55,
		URL:       "https://mdc.mo.gov/discover-nature/places/ted-shanks-conservation-area",
	},
	{
		Name:      "Eagle Bluffs",
		Key:       "8931",
		Zipcode:   65203,
		City:      "Columbia",
		Longitude: -92.369,
		Latitude:  38.925,
		URL:       "https://mdc.mo.gov/discover-nature/places/eagle-bluffs-conservation-area",
	},
	{
		Name:      "Grand Pass",
		Key:       "8010",
		Zipcode:   65344,
		City:      "Miami",
		Longitude: -93.226,
		Latitude:  39.311,
		URL:       "https://mdc.mo.gov/discover-nature/places/grand-pass-conservation-area",
	},
	{
		Name:      "Fountain Grove",
		Key:       "4601",
		Zipcode:   64659,
		City:      "Meadville",
		Longitude: -93.3,
		Latitude:  39.786,
		URL:       "https://mdc.mo.gov/discover-nature/places/fountain-grove-conservation-area",
	},
	{
		Name:      "Bob Brown",
		Key:       "8142",
		Zipcode:   64451,
		City:      "Forest City",
		Longitude: -95.188,
		Latitude:  39.983,
		URL:       "https://mdc.mo.gov/discover-nature/places/bob-brown-conservation-area",
	},
	{
		Name:      "Nodaway Valley",
		Key:       "9134",
		Zipcode:   64466,
		City:      "Maitland",
		Longitude: -95.088,
		Latitude:  40.176,
		URL:       "https://mdc.mo.gov/discover-nature/places/nodaway-valley-conservation-area",
	},
}
