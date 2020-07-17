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

type AdcreativePreviewsGetExample struct {
	TAds                      *ads.SDKClient
	AccessToken               string
	AccountId                 int64
	Filtering                 []model.FilteringStruct
	AdcreativePreviewsGetOpts *api.AdcreativePreviewsGetOpts
}

func (e *AdcreativePreviewsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.Filtering = []model.FilteringStruct{model.FilteringStruct{
		Field:    "adgroup_id",
		Operator: "EQUALS",
		Values:   []string{"YOUR ADGROUP ID"},
	}}
	e.AdcreativePreviewsGetOpts = &api.AdcreativePreviewsGetOpts{

		Fields: optional.NewInterface([]string{"user_id", "user_id_type"}),
	}
}

func (e *AdcreativePreviewsGetExample) RunExample() (model.AdcreativePreviewsGetResponseData, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.AdcreativePreviews().Get(ctx, e.AccountId, e.Filtering, e.AdcreativePreviewsGetOpts)
}

func main() {
	e := &AdcreativePreviewsGetExample{}
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
