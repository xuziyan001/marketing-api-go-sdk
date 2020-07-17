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
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
)

type AsyncReportFilesGetExample struct {
	TAds                    *ads.SDKClient
	AccessToken             string
	AccountId               int64
	TaskId                  int64
	FileId                  int64
	AsyncReportFilesGetOpts *api.AsyncReportFilesGetOpts
}

func (e *AsyncReportFilesGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = 789
	e.TaskId = 789
	e.FileId = 789
	e.AsyncReportFilesGetOpts = &api.AsyncReportFilesGetOpts{}
}

func (e *AsyncReportFilesGetExample) RunExample() (interface{}, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.AsyncReportFiles().Get(ctx, e.AccountId, e.TaskId, e.FileId, e.AsyncReportFilesGetOpts)
}

func main() {
	e := &AsyncReportFilesGetExample{}
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
