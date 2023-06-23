package ssp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	// Import DSP packages here.
	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
)

func getBidFromDSPs(bidRequest common.BidRequest) (highestBid common.Bid) {
	// Create a channel to receive bids
	bids := make(chan common.Bid, 2)

	// Call the bid function of dsp1 and dsp2 concurrently
	// Higher QPS
	go func() {
		bids <- getBidFromDSP(bidRequest, "http://localhost:8081/get-bid")
	}()
	go func() {
		bids <- getBidFromDSP(bidRequest, "http://localhost:8082/get-bid")
	}()

	// Receive the bids
	bid1 := <-bids
	bid2 := <-bids

	// Compare the bids and return the highest
	if bid1.Bid > bid2.Bid {
		return bid1
	} else {
		return bid2
	}
}
func getBidFromDSP(bidRequest common.BidRequest, url string) (bid common.Bid) {
	// Convert bidRequest to JSON
	jsonReq, err := json.Marshal(bidRequest)
	if err != nil {
		log.Printf("Failed to marshal bid request: %v", err)
		return
	}

	// Send the request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Printf("Failed to get bid: %v", err)
		return
	}
	defer resp.Body.Close()

	// Read and return the bid
	json.NewDecoder(resp.Body).Decode(&bid)
	return
}

func BidRequestHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body
	var bidRequest common.BidRequest
	err := json.NewDecoder(r.Body).Decode(&bidRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the highest bid
	highestBid := getBidFromDSPs(bidRequest)

	// Fire the impression pixel
	fireImpressionPixel(highestBid)

	// Create and send the BidResponse
	json.NewEncoder(w).Encode(common.BidResponse{
		ID:  bidRequest.ID,
		Bid: highestBid,
	})
}

func fireImpressionPixel(bid common.Bid) {
	// Create the string to log
	logString := fmt.Sprintf("ID: %s, Bid: %f, AdHTML: %s", bid.ID, bid.Bid, bid.AdHTML)

	// Make a POST request to the log server
	_, err := http.Post("http://localhost:8083", "text/plain", strings.NewReader(logString))
	if err != nil {
		log.Printf("Failed to fire impression pixel: %v", err)
	}
}

func StartServer() {
	http.HandleFunc("/bid", BidRequestHandler)
	http.ListenAndServe(":8080", nil)
}
