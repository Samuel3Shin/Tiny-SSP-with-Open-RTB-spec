package ssp_test

import (
	"testing"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/ssp"
)

type mockBidGetter struct{}

func (mbg *mockBidGetter) GetBidFromDSP(bidRequest common.BidRequest, url string) common.BidResponse {
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
	}
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

	maxBidResponse := sspInstance.GetBidFromDSPs(bidRequest)
	maxBid := maxBidResponse.SeatBid[0].Bid[0]

	if maxBid.ID != "1234" || maxBid.Price != 50.0 || maxBid.AdID != "ad1" {
		t.Errorf("Unexpected bid: %v", maxBid)
	}
}
