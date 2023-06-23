package ssp

import (
	"encoding/json"
	"net/http"

	// Import DSP packages here.
	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/common"
	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/dsp1"
	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/dsp2"
)

func getBidFromDSPs(bidRequest common.BidRequest) (highestBid common.Bid) {
	// Call the bid function of dsp1 and dsp2
	// For example purposes, we assume that dsp1 and dsp2 have a function called "getBid"
	dsp1Bid := dsp1.GetBid(bidRequest)
	dsp2Bid := dsp2.GetBid(bidRequest)

	// Compare the bids and return the highest
	if dsp1Bid.Bid > dsp2Bid.Bid {
		return dsp1Bid
	} else {
		return dsp2Bid
	}
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

	// Create and send the BidResponse
	json.NewEncoder(w).Encode(common.BidResponse{
		ID:  bidRequest.ID,
		Bid: highestBid,
	})
}

func StartServer() {
	http.HandleFunc("/bid", BidRequestHandler)
	http.ListenAndServe(":8080", nil)
}
