package common

type BidRequest struct {
	ID  string       `json:"id"`
	Imp []Impression `json:"imp"`
}

type Impression struct {
	ID     string `json:"id"`
	Banner Banner `json:"banner"`
}

type Banner struct {
	W int `json:"w"`
	H int `json:"h"`
}

type BidResponse struct {
	ID      string    `json:"id"`
	SeatBid []SeatBid `json:"seatbid"`
}

type SeatBid struct {
	Bid []Bid `json:"bid"`
}

type Bid struct {
	ID    string  `json:"id"`
	ImpID string  `json:"impid"`
	Price float64 `json:"price"`
	AdID  string  `json:"adid"`
	AdM   string  `json:"adm"`
	NURL  string  `json:"nurl"`
}
