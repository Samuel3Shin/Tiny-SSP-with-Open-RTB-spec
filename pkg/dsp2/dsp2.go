package dsp2

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
		"https://www.hyundaiusa.com/us/ko/vehicles/ioniq-6?cmpid=tnad_ko_hotplace_nv_830x130&utm_source=tnad_nv&utm_medium=tnad_ko_banner",
		"https://dsp-ad-objects.s3.amazonaws.com/hyundai_ioniq6.jpeg"},
		{"https://tv.naver.com/ktstudiogenie", "https://dsp-ad-objects.s3.amazonaws.com/madang.jpeg"},
		{"https://myprofile.naver.com/main", "https://dsp-ad-objects.s3.amazonaws.com/people.png"}}

	for i := 0; i < 3; i++ {
		bids := []common.Bid{}
		for j := 0; j < 3; j++ {
			bidAmount := rand.Float64() * 100
			bids = append(bids, common.Bid{
				ID:    fmt.Sprintf("bid%d", j+4),
				ImpID: bidRequest.Imp[0].ID,
				Price: bidAmount,
				AdID:  fmt.Sprintf("ad%d", j+4),
				AdM:   fmt.Sprintf("<a href='%s' target='_blank'><img src='%s' /></a>", sampleAdms[j][0], sampleAdms[j][1]),
				NURL:  cfg.LOGSERVER_URL,
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
