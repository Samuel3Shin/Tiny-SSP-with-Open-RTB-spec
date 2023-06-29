package ssp_test

import (
	"reflect"
	"testing"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/ssp"
)

type mockBidGetter struct{}

func (mbg *mockBidGetter) GetBidFromDSP(bidRequest common.BidRequest, url string) (common.BidResponse, error) {
	cfg := common.GetConfig()
	return common.BidResponse{
		ID: "abcd",
		SeatBid: []common.SeatBid{
			{
				Bid: []common.Bid{
					{
						ID:    "1234",
						ImpID: "imp1",
						Price: 50.0,
						AdID:  "ad1",
						NURL:  cfg.LOGSERVER_URL,
					},
				},
			},
		},
	}, nil
}

func TestGetBidFromDSPs(t *testing.T) {
	sspInstance := ssp.NewSSP(&mockBidGetter{})
	bidRequest := common.BidRequest{
		ID: "abcd",
		Imp: []common.Impression{
			{
				ID:     "imp1",
				Banner: common.Banner{W: 300, H: 250},
			},
		},
	}

	// Test that the result is not in the cache
	maxBidResponse, err := sspInstance.GetBidFromDSPs(bidRequest)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	maxBid := maxBidResponse.SeatBid[0].Bid[0]

	if maxBid.ID != "1234" || maxBid.Price != 50.0 || maxBid.AdID != "ad1" {
		t.Errorf("Unexpected bid: %v", maxBid)
	}

	// Test that the result is in the cache
	maxBidResponseCached, err := sspInstance.GetBidFromDSPs(bidRequest)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(maxBidResponse, maxBidResponseCached) {
		t.Errorf("Unexpected difference between initial response and cached response")
	}
}
