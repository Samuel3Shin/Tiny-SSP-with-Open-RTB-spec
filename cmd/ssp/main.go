package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/ssp"
)

type bidGetter struct{}

func (bg *bidGetter) GetBidFromDSP(bidRequest common.BidRequest, url string) common.BidResponse {
	requestJSON, err := json.Marshal(bidRequest)
	if err != nil {
		log.Printf("Failed to marshal bid request: %v", err)
		return common.BidResponse{}
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		log.Printf("Failed to get bid: %v", err)
		return common.BidResponse{}
	}
	defer response.Body.Close()

	var bidResponse common.BidResponse
	if err := json.NewDecoder(response.Body).Decode(&bidResponse); err != nil {
		log.Printf("Failed to decode bid response: %v", err)
	}

	return bidResponse
}

func main() {
	sspInstance := ssp.NewSSP(&bidGetter{})
	sspInstance.StartServer()
}
