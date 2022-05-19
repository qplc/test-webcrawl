package rest

import (
	"time"

	"gopkg.in/resty.v1"
)

var request *resty.Client

type RequestService struct {
	URL string
}

func init() {

	request = resty.New()

	request.SetDebug(false)
	request.SetTimeout(5 * time.Second)
}

func (r *RequestService) PostData(url string) ([]byte, error) {

	resp, err := request.R().Post(url)

	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}
