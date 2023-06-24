package dsp1

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
	bidResponse := GenerateBid(bidRequest)

	// Send the bid
	json.NewEncoder(w).Encode(bidResponse)
}

func GenerateBid(bidRequest common.BidRequest) common.BidResponse {
	cfg := common.GetConfig()
	rand.Seed(time.Now().UnixNano())
	bidAmount := rand.Float64() * 100
	return common.BidResponse{
		ID: bidRequest.ID,
		SeatBid: []common.SeatBid{
			{
				Bid: []common.Bid{
					{
						ID:    "12",
						ImpID: bidRequest.Imp[0].ID,
						Price: bidAmount,
						AdID:  "ad1",
						NURL:  cfg.LOGSERVER_URL,
					},
				},
			},
		},
	}
}
