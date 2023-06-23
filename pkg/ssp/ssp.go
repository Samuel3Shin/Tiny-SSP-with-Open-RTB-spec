package ssp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Samuel3Shin/Tiny-SSP-with-Open-RTB-spec/pkg/common"
)

type BidGetter interface {
	GetBidFromDSP(bidRequest common.BidRequest, url string) common.Bid
}

type SSP struct {
	BidGetter
}

func NewSSP(bg BidGetter) *SSP {
	return &SSP{BidGetter: bg}
}

func (s *SSP) GetBidFromDSPs(bidRequest common.BidRequest) (highestBid common.Bid) {
	bids := make(chan common.Bid, 2)
	go func() {
		bids <- s.GetBidFromDSP(bidRequest, "http://localhost:8081/get-bid")
	}()
	go func() {
		bids <- s.GetBidFromDSP(bidRequest, "http://localhost:8082/get-bid")
	}()

	bid1 := <-bids
	bid2 := <-bids

	if bid1.Bid > bid2.Bid {
		return bid1
	} else {
		return bid2
	}
}

func (s *SSP) BidRequestHandler(w http.ResponseWriter, r *http.Request) {
	var bidRequest common.BidRequest
	err := json.NewDecoder(r.Body).Decode(&bidRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	highestBid := s.GetBidFromDSPs(bidRequest)

	s.fireImpressionPixel(highestBid)

	json.NewEncoder(w).Encode(common.BidResponse{
		ID:  bidRequest.ID,
		Bid: highestBid,
	})
}

func (s *SSP) fireImpressionPixel(bid common.Bid) {
	logString := fmt.Sprintf("ID: %s, Bid: %f, AdHTML: %s", bid.ID, bid.Bid, bid.AdHTML)

	_, err := http.Post("http://localhost:8083", "text/plain", strings.NewReader(logString))
	if err != nil {
		log.Printf("Failed to fire impression pixel: %v", err)
	}
}

func (s *SSP) StartServer() {
	http.HandleFunc("/bid", s.BidRequestHandler)
	http.ListenAndServe(":8080", nil)
}
