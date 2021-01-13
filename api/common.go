package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiResponse struct {
	IsSuccessful bool                   `json:"is_successful"`
	Error        string                 `json:"error"`
	Response     map[string]interface{} `json:"response"` // interface{} has been used because we are unable to expect the return type of response
}

type ApiRequest struct {
	Region         string `json:"region"`
	ClientId       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	TenantId       string `json:"tenant_id"`
	SubscriptionId string `json:"subscription_id"`
}

func ListBucket(w http.ResponseWriter, r *http.Request) {

	var response ApiResponse
	w.Header().Set("Access-Control-Allow-Origin", "*")
	apiRequest := &ApiRequest{}
	err := json.NewDecoder(r.Body).Decode(apiRequest)
	if err != nil {
		fmt.Print(err)
	} else {
		res, awsErr := AwsBucketList(apiRequest.ClientId, apiRequest.ClientSecret, apiRequest.Region)
		if awsErr != nil {
			fmt.Println(awsErr)
		}
		response.Response = map[string]interface{}{"bucklet_list": res}
		response.IsSuccessful = true
		fmt.Println(res)
		Respond(w, response)
	}
}

func CreateBucket(w http.ResponseWriter, r *http.Request) {

	var response ApiResponse
	w.Header().Set("Access-Control-Allow-Origin", "*")
	apiRequest := &ApiRequest{}
	err := json.NewDecoder(r.Body).Decode(apiRequest)
	if err != nil {
		fmt.Print(err)
	} else {
		res, awsErr := AwsCreateBucket(apiRequest.ClientId, apiRequest.ClientSecret, apiRequest.Region)
		if awsErr != nil {
			fmt.Println(awsErr)
		}
		response.Response = map[string]interface{}{"bucklet_list": res}
		response.IsSuccessful = true
		fmt.Println(res)
		Respond(w, response)
	}
}

func DeleteBucket(w http.ResponseWriter, r *http.Request) {

	var response ApiResponse
	w.Header().Set("Access-Control-Allow-Origin", "*")
	apiRequest := &ApiRequest{}
	err := json.NewDecoder(r.Body).Decode(apiRequest)
	if err != nil {
		fmt.Print(err)
	} else {
		res, awsErr := AwsDeleteBucket(apiRequest.ClientId, apiRequest.ClientSecret, apiRequest.Region)
		if awsErr != nil {
			fmt.Println(awsErr)
		}
		response.Response = map[string]interface{}{"bucklet_list": res}
		response.IsSuccessful = true
		fmt.Println(res)
		Respond(w, response)
	}
}
