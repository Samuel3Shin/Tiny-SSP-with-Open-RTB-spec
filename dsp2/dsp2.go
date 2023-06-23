package dsp2

import (
	"math/rand"
	"time"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/common"
)

func GetBid(bidRequest common.BidRequest) (bid common.Bid) {
	// For this MVP, let's create a dummy bid
	// Bid is a random float
	rand.Seed(time.Now().UnixNano())
	bidAmount := rand.Float64() * 100

	return common.Bid{
		ID:     bidRequest.ID,
		Bid:    bidAmount,
		AdHTML: "<h1>This is an ad</h1>",
	}
}
