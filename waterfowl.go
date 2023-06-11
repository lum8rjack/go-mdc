package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Area struct {
	Key  string
	Name string
}

var Areas = [15]Area{
	{Name: "B.K. Leach", Key: "8514"},
	{Name: "Ten Mile Pond", Key: "8241"},
	{Name: "Otter Slough", Key: "5004"},
	{Name: "Schell-Osage", Key: "5701"},
	{Name: "Four Rivers", Key: "8238"},
	{Name: "Montrose", Key: "5604"},
	{Name: "Duck Creek", Key: "5001"},
	{Name: "Columbia Bottoms", Key: "9736"},
	{Name: "Marais Temps Clair", Key: "7902"},
	{Name: "Ted Shanks", Key: "7011"},
	{Name: "Eagle Bluffs", Key: "8931"},
	{Name: "Grand Pass", Key: "8010"},
	{Name: "Fountain Grove", Key: "4601"},
	{Name: "Bob Brown", Key: "8142"},
	{Name: "Nodaway Valley", Key: "9134"},
}

// Get waterfowl area by ney
func GetAreaByKey(key string) (Area, error) {
	a := Area{}
	for _, area := range Areas {
		if area.Key == key {
			return area, nil
		}
	}
	return a, errors.New("area not found")
}

// Get waterfowl area by name
func GetAreaByName(name string) (Area, error) {
	a := Area{}
	for _, area := range Areas {
		if strings.Contains(strings.ToLower(area.Name), strings.ToLower(name)) {
			return area, nil
		}
	}
	return a, errors.New("area not found")
}

const (
	HarvestDataMinYear int    = 2016
	HarvestDataUrl     string = "https://extra.mdc.mo.gov/widgets/wtrfwl_harvest/dataJSONservice.php"
	HarvestDataOrigin  string = "https://extra.mdc.mo.gov"
	HarvestDataReferer string = "https://extra.mdc.mo.gov/widgets/wtrfwl_harvest/"
)

var (
	HarvestDataUserAgent string = "go-mdc"
)

type HarvestData struct {
	WfHrvUpdatesRecs []HarvestDay `json:"wfHrv_Updates_Recs"`
}

type HarvestDay struct {
	RecID        int     `json:"RecId"`
	Area         string  `json:"Area"`
	Season       string  `json:"Season"`
	Week         int     `json:"Week"`
	ReportDate   string  `json:"Report_Date"`
	DuckPop      int     `json:"Duck_Pop"`
	GoosePop     int     `json:"Goose_Pop"`
	DuckHarv     int     `json:"Duck_Harv"`
	CGHarv       int     `json:"CG_Harv"`
	SBGHarv      int     `json:"SBG_Harv"`
	WFGHarv      int     `json:"WFG_Harv"`
	WatfowlHarv  int     `json:"Watfowl_Harv"`
	NumHunters   int     `json:"Num_Hunters"`
	BirdsPerHunt float64 `json:"Birds_Per_Hunt"`
	DucksPerHunt float64 `json:"Ducks_Per_Hunt"`
	BlindVac     int     `json:"Blind_Vac"`
	WadeVac      int     `json:"Wade_Vac"`
	TotalPos     int     `json:"Total_Pos"`
	PLAvail      int     `json:"PL_Avail"`
	PLDraw       int     `json:"PL_Draw"`
	Comments     string  `json:"Comments"`
	AreaID       int     `json:"Area_ID"`
}

// Get current hunting season
func GetCurrentSeason() string {
	// Get current year/month
	currenttime := time.Now()
	currentyear := currenttime.Year()
	currentmonth := int(currenttime.Month())

	if currentmonth < 11 {
		currentyear--
	}

	return fmt.Sprintf("%d-%d", currentyear, currentyear+1)
}

// Check if the year is valid and returns the season
// that is needed with the web request.
// e.g. 2022 -> "2022-2023"
func CheckYear(year int) (string, error) {
	// Make sure its a valid season
	currentyear := time.Now().Year()
	if year < HarvestDataMinYear || year > currentyear {
		return "", errors.New("invalid year")
	}

	currentmonth := int(time.Now().Month())
	if currentmonth < 11 && currentyear == year {
		return "", errors.New("invalid year - season hasn't started")
	}

	// Set the season
	return fmt.Sprintf("%d-%d", year, year+1), nil
}

// Send web reques to get harvest data for the provided year
// If you want the 2022-2023 season, the year should be 2022
func GetData(year int, area Area) (HarvestData, error) {
	// Setup empty harvest data to start
	hd := HarvestData{}

	// Validate the year
	season, err := CheckYear(year)
	if err != nil {
		return hd, err
	}

	// Setup the HTTP client
	client := &http.Client{
		Timeout: time.Second * 10,
		// Accept any cert - needed for the scratch docker instance
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	// Setup the params
	// paravalue=?areas=7011&season=2021-2022
	paramdata := fmt.Sprintf("?areas=%s&season=%s", area.Key, season)
	params := url.Values{}
	params.Add("paravalue", paramdata)

	// Setup the POST request
	req, err := http.NewRequest("POST", HarvestDataUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return hd, err
	}

	// Setup the headers
	req.Header.Set("User-Agent", HarvestDataUserAgent)
	req.Header.Add("Origin", HarvestDataOrigin)
	req.Header.Add("Referer", HarvestDataReferer)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	response, err := client.Do(req)
	if err != nil {
		return hd, err
	}
	defer response.Body.Close()

	// Parse the response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return hd, err
	}

	// Convert body to struct
	err = json.Unmarshal(body, &hd)
	if err != nil {
		return hd, err
	}

	return hd, nil
}
