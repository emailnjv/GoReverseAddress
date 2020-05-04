package client

import (
	"fmt"
	"os"

	"GoReverseAddress/resources"

	"github.com/joho/godotenv"
)

const (
	BaseURL = "https://api.recordsapi.com/v2/json.php"
	AppID = "App_ID"
	AppKey = "App_Key"
	Timestamp = "Timestamp"
	Catalogue = "catalogue"
	Name = "Name"
	State = "State"
	County = "County"
	City = "City"
	ZipCode = "ZipCode"
	ExactMatch = "ExactMatch"
)

type RecordsAPIModule struct {
	Caller    Caller
	apiKey    string
	apiID     string
	ipAddress string
}

func NewRecordsAPIModule(caller Caller) (RecordsAPIModule, error) {
	var result RecordsAPIModule

	err := godotenv.Load("../.env")
	if err != nil {
		return result, fmt.Errorf("Error loading .env file")
	}

	apiID := os.Getenv("RECORDS_API_APP_ID")
	apiKey := os.Getenv("RECORDS_API_APP_KEY")
	ipAddress := os.Getenv("PUBLIC_IP")
	if apiID == "" || apiKey == "" || ipAddress == "" {
		return result, fmt.Errorf("Error loading .env variables for Records API")
	}

	caller, err = NewCallOut()
	if err != nil {
		return result, err
	}

	return result, err
}

func (r RecordsAPIModule) RequestRALU(name resources.AddressAndName, err error) (string, error) {

}

func (r RecordsAPIModule) ParseRALU(s string, err error) (resources.PersonalInfo, error) {
	panic("implement me")
}

func (r RecordsAPIModule) builldURL(appID string, appKey string, timestamp string, ip string, catalogue string, name string, state string, county string, city string, zipCode string, exactMatch string) (string, error) {
// https://api.recordsapi.com/v2/json.php?App_ID= [YOUR APP ID]&App_Key= [YOUR APPKEY]&Timestamp=[CURRENT TIMESTAMP]&IP=[YOUR IPADDRESS]&catalogue=USPEOPLE&Name=%20JOSE%20LUIS%20ZUNIGA&%20State=CA&County=&City=&ZipCode=&ExactMatch=Yes
// 	var result string
	
	// result
	panic("implement me")

}
