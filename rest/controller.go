package rest

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/test-webcrawl/errorhandler"
	"bitbucket.org/test-webcrawl/models"
	"bitbucket.org/test-webcrawl/utils"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

//HealthCheckAPI to check health of the application
func HealthCheckAPI(w http.ResponseWriter, r *http.Request) {
	utils.LogD.Info("EntryPoint", zap.Any("inside ", "HealthCheckAPI"))

	defer r.Body.Close()

	respObj := models.ResponseObject{}
	respObj.Success = true

	render.Status(r, 200)
	render.JSON(w, r, respObj)
}

//WebCrawlAPI to get weburl response
func WebCrawlAPI(w http.ResponseWriter, r *http.Request) {
	utils.LogD.Info("EntryPoint", zap.Any("inside ", "WebCrawlAPI"))

	errorhandler.Block{
		Try: func() {
			defer r.Body.Close()

			var requestData models.IncomingRequest
			if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
				utils.LogD.Error("Validation of JSON parsing failed", zap.Error(err))
				render.Status(r, 400)
				return
			}

			utils.LogD.Info("Incoming Request Data", zap.Any("", requestData))

			urls := requestData.WebUrl

			urllen := len(urls)

			WebCrawlResponse := models.WebCrawlResponse{}

			var WebCrawlResults = make([]models.WebCrawlResult, urllen)

			for i, url := range urls {
				reqObj := RequestService{}

				rawData, err := reqObj.PostData("http://google.com")
				if err != nil {
					utils.LogD.Error("WebCrawl Error", zap.Error(err))
				}
				utils.LogD.Info("rawData", zap.Any(url, rawData))

				WebCrawlResult := models.WebCrawlResult{}
				WebCrawlResult.WebUrl = url
				WebCrawlResult.Data = string(rawData)

				WebCrawlResults[i] = WebCrawlResult
			}

			WebCrawlResponse.WebCrawlResult = WebCrawlResults

			render.Status(r, 200)
			render.JSON(w, r, WebCrawlResponse)
		},
		Catch: func(e errorhandler.Exception) {
			errorhandler.LogError(e)
		},
	}.Do()
}
