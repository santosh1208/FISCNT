package getcustomer

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
)

const (
	ivConnection = "FISConnection"

	ivcustomerid = "CustomerID"
	ovMessageId  = "Message"
)

type GerCustomerActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &GerCustomerActivity{metadata: metadata}
}

func (a *GerCustomerActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *GerCustomerActivity) Eval(context activity.Context) (done bool, err error) {

	if context.GetInput(ivcustomerid) == nil {
		return false, activity.NewError("Customerid is required ", "FIS-GETCUSTOMER-4002", nil)
	}

	//Read connection details
	connectionInfo, _ := data.CoerceToObject(context.GetInput(ivConnection))

	if connectionInfo == nil {
		return false, activity.NewError("FIS connection is not configured", "FIS-GETCUSTOMER-4001", nil)
	}

	var ConsumerKey string
	var ConsumerSecret string

	connectionSettings, _ := connectionInfo["settings"].([]interface{})
	if connectionSettings != nil {
		for _, v := range connectionSettings {
			setting, _ := data.CoerceToObject(v)
			if setting != nil {
				if setting["name"] == "ConsumerKey" {
					ConsumerKey, _ = data.CoerceToString(setting["value"])
				} else if setting["name"] == "ConsumerSecret" {
					ConsumerSecret, _ = data.CoerceToString(setting["value"])
				}

			}
		}
	}

	var customerid = ivcustomerid

	url := "https://api-gw-uat.fisglobal.com/rest/inv-customer/v1.2/customers/" + customerid

	var mod = (ConsumerKey + ":" + ConsumerSecret)
	sEnc := base64.StdEncoding.EncodeToString([]byte(mod))

	var auth = "Basic " + sEnc

	urls := "https://api-gw-uat.fisglobal.com/token"

	payload := strings.NewReader("grant_type=client_credentials")

	req, _ := http.NewRequest("POST", urls, payload)

	req.Header.Add("Authorization", auth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	a1 := []rune(string(body))

	accesstoken := string(a1[17:53])

	var bearer = "Bearer " + accesstoken

	reqs, _ := http.NewRequest("GET", url, nil)

	reqs.Header.Add("Authorization", bearer)

	resp, _ := http.DefaultClient.Do(reqs)

	defer res.Body.Close()
	bodyre, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(bodyre))

	//Set Message ID in the output

	context.SetOutput(ovMessageId, string(bodyre))
	return true, nil
}
