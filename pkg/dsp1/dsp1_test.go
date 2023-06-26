package dsp1_test

import (
	"fmt"
	"testing"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/dsp1"
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
	bidResponse := dsp1.GenerateBid(bidRequest)

	if bidResponse.ID != "abcd" {
		t.Errorf("Unexpected response ID: %v", bidResponse.ID)
	}

	for i, seatBid := range bidResponse.SeatBid {
		for j, bid := range seatBid.Bid {
			if bid.Price < 0 || bid.Price > 100 || bid.ID != fmt.Sprintf("bid%d", j+1) || bid.AdID != fmt.Sprintf("ad%d", j+1) {
				t.Errorf("Unexpected bid in SeatBid %d: %v", i, bid)
			}
		}
	}
}
