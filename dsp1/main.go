package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/common"
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
	rand.Seed(time.Now().UnixNano())
	bidAmount := rand.Float64() * 100
	bid := common.Bid{
		ID:     bidRequest.ID,
		Bid:    bidAmount,
		AdHTML: "<h1>This is an ad1</h1>",
	}

	// Send the bid
	json.NewEncoder(w).Encode(bid)
}

func main() {
	http.HandleFunc("/get-bid", GetBidHandler)
	http.ListenAndServe(":8081", nil)
}
