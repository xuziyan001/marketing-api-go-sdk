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
	"os"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type CustomAudienceFilesAddExample struct {
	TAds                       *ads.SDKClient
	AccessToken                string
	AccountId                  int64
	AudienceId                 int64
	UserIdType                 string
	File                       *os.File
	CustomAudienceFilesAddOpts *api.CustomAudienceFilesAddOpts
}

func (e *CustomAudienceFilesAddExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.AudienceId = int64(0)
	e.UserIdType = "QQ"
	file, err := os.Open("YOUR FILE PATH")
	if err != nil {
		e.File = file
	}
	e.CustomAudienceFilesAddOpts = &api.CustomAudienceFilesAddOpts{}
}

func (e *CustomAudienceFilesAddExample) RunExample() (model.CustomAudienceFilesAddResponseData, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.CustomAudienceFiles().Add(ctx, e.AccountId, e.AudienceId, e.UserIdType, e.File, e.CustomAudienceFilesAddOpts)
}

func main() {
	e := &CustomAudienceFilesAddExample{}
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
