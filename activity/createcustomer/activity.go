/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package cratecustomer

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivConnection = "FISConnection"
	ivinput      = "Customer"
	ovMessage    = "Message"
)

var activityLog = logger.GetLogger("Logger")

type CreateCustomerActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &CreateCustomerActivity{metadata: metadata}
}

func (a *CreateCustomerActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *CreateCustomerActivity) Eval(context activity.Context) (done bool, err error) {
	//	activityLog.Info("Executing SQS Receive Message activity")

	//if context.GetInput(ivQueueUrl) == nil {
	//		return false, activity.NewError("SQS Queue URL is not configured", "SQS-RECEIVEMESSAGE-4002", nil)
	//	}

	//Read connection details
	connectionInfo, _ := data.CoerceToObject(context.GetInput(ivConnection))

	if connectionInfo == nil {
		return false, activity.NewError("SQS connection is not configured", "SQS-RECEIVEMESSAGE-4001", nil)
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

	context.SetOutput(ivinput, "{\n\"birthDate\": \"1990-05-02\",\n\"branch\": \"ICICIEPIP\",\n\"email\": \"ram@gmail.com\",\n\"idType\": \"Individual\",\n\"phone\": \"979012321\",\n\"primaryOfficer\": \"Primary\",\n\"profileImage\": \"profileImage\",\n\"residentAlien\": \"residentAlien\",\n\"ipStatus\": \"Open\",\n\"taxId\": \"Tax001001\",\n\"vipCode\": \"Vip001\",\n\"prefix\": \"Mr\",\n\"firstName\": \"Ram\",\n\"middleInitial\": \"Lakshman\",\n\"lastName\": \"Bharath\",\n\"suffix\": \"Jr\",\n\"gender\": \"Male\",\n\"ssn\": \"211\",\n\"mothersMaidenName\": \"Kaushalya\",\n\"addresses\": [{\n\"line1\": \"whitefield\",\n\"line2\": \"EpipZone\",\n\"city\": \"Bangalore\",\n\"state\": \"KA\",\n\"zip\": \"521229\"\n}],\n\"ipPoc\": {\n\"ipEmail\": [{\n\"addr\": \"jsmm@gamil.com\",\n\"emailTyp\": \"Primary\",\n\"endDte\": \"2017-04-28\",\n\"prfrdInd\": \"true\",\n\"seqNbr\": 1,\n\"spclInst\": \"Email Special Instructions\",\n\"strtDte\": \"2017-04-28\"\n}],\n\"ipPhone\": [{\n\"phTmeframe\": [{\n\"tmeRng\": [{\n\"type\": \"Primary\",\n\"strtTme\": \"02:02:03\",\n\"endTme\": \"18:42:50\",\n\"tzNme\": \"Canada/Central\",\n\"utcOfst\": \"+05:00\"\n}],\n\"startDay\": \"Sunday\",\n\"endDay\": \"Sunday\"\n}],\n\"strtDte\": \"2017-04-28\",\n\"endDate\": \"2017-04-28\",\n\"nbr\": \"1231312\",\n\"intlPrfx\": 1,\n\"ctryCallngCde\": 1,\n\"trnkCde\": 1,\n\"areaCde\": 1,\n\"teleNbr\": 10,\n\"extnNbr\": 10,\n\"txtEnabl\": \"true\",\n\"phTyp\": \"Mobile\",\n\"seqNbr\": 1,\n\"spclInst\": \"Phone Special Instructions\",\n\"tzCode\": \"Canada/Central\",\n\"utcOfSet\": \"+05:00\",\n\"prfrdPhInd\": \"true\",\n\"prfrdTxtInd\": \"true\"\n}],\n\"ipPostalAddress\": [{\n\"ipPostalAddressDemo\": {\n\"occpyDte\": \"2017-04-28\",\n\"prchDte\": \"2017-04-28\",\n\"mlngCde\": \"All Mail Allowed\",\n\"occpyEndDte\": \"2017-04-28\",\n\"spclInst\": \"Special Instructions\"\n},\n\"ipPostalAddressLine\": [{\n\"pstlAddrLneTxt\": \"House No 26\",\n\"pstlAddrLneTyp\": \"Destination Line 1\"\n}],\n\"ipRcpntLne\": [{\n\"ipNmeTyp\": \"Primary\",\n\"rcpntLneTxt\": \"Mr John\",\n\"rcpntLneTyp\": \"Recipient Line 1\"\n}],\n\"pstlAddrTyp\": \"Home\",\n\"endDte\": \"2017-04-28\",\n\"prfrdInd\": \"true\",\n\"pstlAddrCat\": \"Master\",\n\"pstlAddrFmtCntryCde\": \"USA\",\n\"pstlAddrValDte\": \"2017-04-28\",\n\"pstlAddrValSrc\": \"Utility Bills\",\n\"seqNbr\": 1,\n\"strtDte\": \"2017-04-28\",\n\"addrLne1\": \"string\",\n\"addrLne2\": \"string\",\n\"addrLne3\": \"string\",\n\"str\": \"string\",\n\"cty\": \"string\",\n\"st\": \"State\",\n\"pstlCde\": \"5210067\"\n}],\n\"ipSocialMedia\": [{\n\"endDte\": \"2017-04-28\",\n\"prfrdInd\": \"true\",\n\"seqNbr\": 1,\n\"socMediaAccId\": \"Jsss@gmail.com\",\n\"socMediaTyp\": \"Facebook\",\n\"spclInst\": \"Only Facebook\",\n\"strtDte\": \"2017-04-28\"\n}]\n},\n\"nme\": {\n\"strtDte\": \"2017-06-16\",\n\"endDte\": \"2017-06-16\",\n\"type\": \"Primary\",\n\"frstNme\": \"Madhu\",\n\"midNme\": \"Moahn\",\n\"lstNme\": \"Rao\",\n\"nmeSfx\": \"PhD\",\n\"nmePrfx\": \"Mr.\",\n\"nmeLne\": \"string\",\n\"nmeTtl\": \"Honorable\",\n\"scndLstNme\": \"Simha\"\n},\n\"ipFeAssgn\": {\n\"cstCntrId\": \"CostCID123\",\n\"cstCntrDesc\": \"Cost center\",\n\"brnchAssgn\": [{\n\"brnchId\": \"8990\",\n\"brnchDesc\": \"Branch Desc\"\n}],\n\"rltMgrAssgn\": [{\n\"rltMgrCde\": \"RC0012\",\n\"rltMgrNme\": \"Prasanna\",\n\"rltMgrCntc\": \"6767617177\"\n}]\n},\n\"ipRole\": {\n\"ipId\": 0,\n\"strtDte\": \"2017-06-16\",\n\"endDte\": \"2017-06-16\",\n\"percentage\": 100,\n\"code\": \"CXDTYY\",\n\"role\": \"string\",\n\"roleType\": \"Primary\"\n},\n\"issId\": {\n\n\"issDte\": \"2017-06-16\",\n\"expDte\": \"2017-06-16\",\n\"issBy\": \"Govt Of India\",\n\"issByDesc\": \"Government of India\",\n\"issIdTyp\": \"Aadhar\",\n\"val\": \"1213177800\",\n\"issDesc\": \"AaadharId\",\n\"seqNbr\": 1\n}\n}")

	var mod = (ConsumerKey + ":" + ConsumerSecret)
	sEnc := base64.StdEncoding.EncodeToString([]byte(mod))
	fmt.Println(sEnc)

	var auth = "Basic " + sEnc

	url := "https://api-gw-uat.fisglobal.com/token"

	payload := strings.NewReader("grant_type=client_credentials")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", auth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	a1 := []rune(string(body))

	accesstoken := string(a1[17:53])

	fmt.Println(accesstoken)

	//urlto := "https://api-gw-uat.fisglobal.com/rest/inv-customer/v1.2/customers"

	//jsonSchema, _ := data.CoerceToComplexObject(context.GetInput(ivinput))

	//bytes := json.Marshal(jsonSchema)
	//payload := strings.NewReader(bytes)

	//	reqx, _ := http.NewRequest("POST", url, payload)

	//var author = "Bearer " + accesstoken
	//req.Header.Add("Content-Type", "application/json")
	//	req.Header.Add("Authorization", author)
	//	req.Header.Add("feid", "5100")

	//Set Message details in the output
	//msgs := make([]map[string]interface{}, len(reqx.body))

	//read message attributes

	//output := &data.ComplexObject{Metadata: "", Value: msgs}
	//	context.SetOutput(ovMessage, output)
	return true, nil
}
