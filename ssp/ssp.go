package ssp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	// Import DSP packages here.
	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/common"
)

func getBidFromDSPs(bidRequest common.BidRequest) (highestBid common.Bid) {
	// Call the GetBidHandler function of dsp1 and dsp2 via HTTP
	dsp1Bid := getBidFromDSP(bidRequest, "http://localhost:8081/get-bid")
	dsp2Bid := getBidFromDSP(bidRequest, "http://localhost:8082/get-bid")

	// Compare the bids and return the highest
	if dsp1Bid.Bid > dsp2Bid.Bid {
		return dsp1Bid
	} else {
		return dsp2Bid
	}
}

func getBidFromDSP(bidRequest common.BidRequest, url string) (bid common.Bid) {
	// Convert bidRequest to json
	jsonData, err := json.Marshal(bidRequest)
	if err != nil {
		log.Println(err)
	}

	// Create a new request using http
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalln(err)
	}

	// We add headers to the request
	req.Header.Add("Content-Type", "application/json")

	// Send the request and get the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	// We read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// We close the response body
	resp.Body.Close()

	// Unmarshal the body into a bid
	json.Unmarshal(body, &bid)

	return bid
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
