package ssp_test

import (
	"testing"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/ssp"
)

type mockBidGetter struct{}

func (mbg *mockBidGetter) GetBidFromDSP(bidRequest common.BidRequest, url string) common.Bid {
	// This mock will return the same bid for every URL
	return common.Bid{
		ID:     "1234",
		Bid:    50.0,
		AdHTML: "<h1>This is a mock ad</h1>",
	}
}

func TestGetBidFromDSPs(t *testing.T) {
	s := ssp.NewSSP(&mockBidGetter{})
	bidRequest := common.BidRequest{ID: "abcd", Imp: "ad", Site: "test.com"}
	bid := s.GetBidFromDSPs(bidRequest)

	if bid.ID != "1234" || bid.Bid != 50.0 || bid.AdHTML != "<h1>This is a mock ad</h1>" {
		t.Errorf("Unexpected bid: %v", bid)
	}
}
