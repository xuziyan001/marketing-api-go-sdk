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

	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type UserActionSetReportsGetExample struct {
	TAds                        *ads.SDKClient
	AccessToken                 string
	AccountId                   int64
	UserActionSetId             int64
	DateRange                   model.DateRange
	TimeGranularity             string
	UserActionSetReportsGetOpts *api.UserActionSetReportsGetOpts
}

func (e *UserActionSetReportsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.UserActionSetId = int64(0)
	e.DateRange = model.DateRange{
		StartDate: "REPORT START DATE",
		EndDate:   "REPORT END DATE",
	}
	e.TimeGranularity = "DAILY"
	e.UserActionSetReportsGetOpts = &api.UserActionSetReportsGetOpts{

		Fields: optional.NewInterface([]string{"date", "action_type", "custom_action", "raw_action_count", "identified_action_count", "identified_user_count"}),
	}
}

func (e *UserActionSetReportsGetExample) RunExample() (model.UserActionSetReportsGetResponseData, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.UserActionSetReports().Get(ctx, e.AccountId, e.UserActionSetId, e.DateRange, e.TimeGranularity, e.UserActionSetReportsGetOpts)
}

func main() {
	e := &UserActionSetReportsGetExample{}
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
