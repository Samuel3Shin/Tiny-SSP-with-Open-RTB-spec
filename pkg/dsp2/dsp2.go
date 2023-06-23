package dsp2

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
)

func GetBidHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body
	var bidRequest common.BidRequest
	err := json.NewDecoder(r.Body).Decode(&bidRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a bid
	bid := GenerateBid(bidRequest)

	// Send the bid
	json.NewEncoder(w).Encode(bid)
}

func GenerateBid(bidRequest common.BidRequest) common.Bid {
	rand.Seed(time.Now().UnixNano())
	bidAmount := rand.Float64() * 100
	return common.Bid{
		ID:     bidRequest.ID,
		Bid:    bidAmount,
		AdHTML: "<h1>This is an ad2</h1>",
	}
}
