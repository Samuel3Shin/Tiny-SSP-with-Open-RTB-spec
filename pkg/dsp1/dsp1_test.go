package dsp1_test

import (
	"testing"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/dsp1"
)

func TestGenerateBid(t *testing.T) {
	bidRequest := common.BidRequest{ID: "abcd", Imp: "ad", Site: "test.com"}
	bid := dsp1.GenerateBid(bidRequest)

	if bid.ID != "abcd" || bid.Bid < 0 || bid.Bid > 100 || bid.AdHTML != "<h1>This is an ad1</h1>" {
		t.Errorf("Unexpected bid: %v", bid)
	}
}
