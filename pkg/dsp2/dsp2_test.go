package dsp2_test

import (
	"fmt"
	"strings"
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

	if bidResponse.ID != "abcd" {
		t.Errorf("Unexpected response ID: %v", bidResponse.ID)
	}

	for i, seatBid := range bidResponse.SeatBid {
		for j, bid := range seatBid.Bid {
			if bid.Price < 0 || bid.Price > 100 || bid.ID != fmt.Sprintf("bid%d", j+4) || bid.AdID != fmt.Sprintf("ad%d", j+4) {
				t.Errorf("Unexpected bid in SeatBid %d: %v", i, bid)
			}

			if bid.AdM == "" {
				t.Errorf("AdM field should not be empty")
			} else if !strings.Contains(bid.AdM, "<a href=") {
				t.Errorf("AdM field should contain valid ad markup")
			}
		}
	}
}
