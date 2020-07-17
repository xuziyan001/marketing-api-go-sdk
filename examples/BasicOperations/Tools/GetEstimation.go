/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type EstimationGetExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.EstimationGetRequest
}

func (e *EstimationGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.EstimationGetRequest{
		Targeting: model.EstimationReadTargetingSetting{
			Age: []model.AgeStruct{model.AgeStruct{
				Max: int64(0),
				Min: int64(0),
			}},
			Gender: []string{"YOUR TARGETING GENDER"},
		},
		AccountId: int64(0),
		CampaignSpec: model.CampaignTargeting{
			CampaignId: int64(0),
		},
	}
}

func (e *EstimationGetExample) RunExample() (model.EstimationGetResponseData, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.Estimation().Get(ctx, e.Data)
}

func main() {
	e := &EstimationGetExample{}
	e.Init()
	response, httpResponse, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Http response:", httpResponse)
}
