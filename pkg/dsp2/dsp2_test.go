package dsp2_test

import (
	"testing"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/dsp2"
)

func TestGenerateBid(t *testing.T) {
	bidRequest := common.BidRequest{
		ID: "abcd",
		Imp: []common.Impression{
			{
				ID:     "imp1",
				Banner: common.Banner{W: 300, H: 250},
			},
		},
	}
	bidResponse := dsp2.GenerateBid(bidRequest)
	bid := bidResponse.SeatBid[0].Bid[0]

	if bidResponse.ID != "abcd" || bid.Price < 0 || bid.Price > 100 || bid.ID != "1234" || bid.AdID != "ad1" {
		t.Errorf("Unexpected bid: %v", bid)
	}
}
