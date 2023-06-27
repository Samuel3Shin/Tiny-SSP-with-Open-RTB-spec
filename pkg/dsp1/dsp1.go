package dsp1

import (
	"encoding/json"
	"fmt"
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
	seatBids := []common.SeatBid{}
	sampleAdms := [][]string{{
		"https://patron.naver.com/ntv/c/intro/globalcontents",
		"https://dsp-ad-objects.s3.amazonaws.com/samchongsa.jpeg"},
		{"https://happybean.naver.com/fundings/detail/F944",
			"https://dsp-ad-objects.s3.amazonaws.com/happybean.jpeg"},
		{"https://wooltariusa.com/collections/snacks?utm_source=naver&utm_medium=display&utm_campaign=PC_%EB%A9%94%EC%9D%B8_%ED%83%80%EC%9E%84%EB%B3%B4%EB%93%9C&utm_content=%EC%9A%B8%ED%83%80%EB%A6%AC%EB%AA%B0_230619&utm_term=k%EA%B0%84%EC%8B%9D",
			"https://dsp-ad-objects.s3.amazonaws.com/wooltari.jpeg"}}

	for i := 0; i < 3; i++ {
		bids := []common.Bid{}
		for j := 0; j < 3; j++ {
			bidAmount := rand.Float64() * 100
			bids = append(bids, common.Bid{
				ID:    fmt.Sprintf("bid%d", j+1),
				ImpID: bidRequest.Imp[0].ID,
				Price: bidAmount,
				AdID:  fmt.Sprintf("ad%d", j+1),
				AdM:   fmt.Sprintf("<a href='%s' target='_blank'><img src='%s' /><img src='%s' style='height:1px; width:1px;' /></a>", sampleAdms[j][0], sampleAdms[j][1], fmt.Sprintf("%s?adID=%s", cfg.LOGSERVER_URL, fmt.Sprintf("ad%d", j+1))),
				NURL:  fmt.Sprintf("%s?adID=%s", cfg.LOGSERVER_URL, fmt.Sprintf("ad%d", j+1)),
			})
		}

		seatBids = append(seatBids, common.SeatBid{
			Bid: bids,
		})
	}

	return common.BidResponse{
		ID:      bidRequest.ID,
		SeatBid: seatBids,
	}
}
