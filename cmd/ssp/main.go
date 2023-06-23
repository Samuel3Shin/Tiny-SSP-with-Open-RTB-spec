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

func (bg *bidGetter) GetBidFromDSP(bidRequest common.BidRequest, url string) common.Bid {
	jsonReq, err := json.Marshal(bidRequest)
	if err != nil {
		log.Printf("Failed to marshal bid request: %v", err)
		return common.Bid{}
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Printf("Failed to get bid: %v", err)
		return common.Bid{}
	}
	defer resp.Body.Close()

	var bid common.Bid
	json.NewDecoder(resp.Body).Decode(&bid)
	return bid
}

func main() {
	s := ssp.NewSSP(&bidGetter{})
	s.StartServer()
}
