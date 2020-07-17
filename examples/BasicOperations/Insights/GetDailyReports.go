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

type DailyReportsGetExample struct {
	TAds                *ads.SDKClient
	AccessToken         string
	AccountId           int64
	Level               string
	DateRange           model.ReportDateRange
	DailyReportsGetOpts *api.DailyReportsGetOpts
}

func (e *DailyReportsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.Level = "REPORT_LEVEL_ADGROUP"
	e.DateRange = model.ReportDateRange{
		StartDate: "REPORT START DATE",
		EndDate:   "REPORT END DATE",
	}
	e.DailyReportsGetOpts = &api.DailyReportsGetOpts{

		Fields: optional.NewInterface([]string{"date", "view_count", "valid_click_count", "ctr", "cpc", "cost"}),
	}
}

func (e *DailyReportsGetExample) RunExample() (model.DailyReportsGetResponseData, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.DailyReports().Get(ctx, e.AccountId, e.Level, e.DateRange, e.DailyReportsGetOpts)
}

func main() {
	e := &DailyReportsGetExample{}
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
