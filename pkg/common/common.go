package common

// Assuming this structure for a bid
type Bid struct {
	ID     string  `json:"id"`
	Bid    float64 `json:"bid"`
	AdHTML string  `json:"adhtml"`
}

// BidRequest structure
type BidRequest struct {
	ID   string `json:"id"`
	Imp  string `json:"imp"` // for example, this could be the HTML of the ad to be shown
	Site string `json:"site"`
}

// BidResponse structure
type BidResponse struct {
	ID  string `json:"id"`
	Bid Bid    `json:"bid"`
}

// TODO: apply OpenRTB 2.5 spec
// type BidRequest struct {
// 	ID          string       `json:"id"`
// 	Impressions []Impression `json:"imp,omitempty"`
// 	Device      Device       `json:"device"`
// 	User        User         `json:"user"`
// 	Test        int          `json:"test,omitempty"`
// 	At          int          `json:"at,omitempty"`
// 	TMax        int          `json:"tmax,omitempty"`
// 	// Add more fields according to the 2.5 spec as required
// }

// type Impression struct {
// 	ID           string  `json:"id,omitempty"`
// 	Banner       *Banner `json:"banner,omitempty"`
// 	Video        *Video  `json:"video,omitempty"`
// 	Bidfloor     float64 `json:"bidfloor,omitempty"`
// 	BidfloorCur  string  `json:"bidfloorcur,omitempty"`
// 	Secure       int     `json:"secure,omitempty"`
// 	Instl        int     `json:"instl,omitempty"`
// 	ClickBrowser int     `json:"clickbrowser,omitempty"`
// 	// Add more fields according to the 2.5 spec as required
// }

// type Banner struct {
// 	Width  int `json:"w,omitempty"`
// 	Height int `json:"h,omitempty"`
// 	// Add more fields according to the 2.5 spec as required
// }

// type Video struct {
// 	Width  int      `json:"w,omitempty"`
// 	Height int      `json:"h,omitempty"`
// 	Mimes  []string `json:"mimes,omitempty"`
// 	// Add more fields according to the 2.5 spec as required
// }

// type Device struct {
// 	UserAgent string `json:"ua,omitempty"`
// 	IP        string `json:"ip,omitempty"`
// 	// Add more fields according to the 2.5 spec as required
// }

// type User struct {
// 	ID string `json:"id,omitempty"`
// 	// Add more fields according to the 2.5 spec as required
// }

// type BidResponse struct {
// 	ID       string    `json:"id,omitempty"`
// 	Seatbids []Seatbid `json:"seatbid,omitempty"`
// 	// Add more fields according to the 2.5 spec as required
// }

// type Seatbid struct {
// 	Bids []Bid `json:"bid,omitempty"`
// 	// Add more fields according to the 2.5 spec as required
// }

// type Bid struct {
// 	ImpressionID    string  `json:"impid,omitempty"`
// 	Price           float64 `json:"price,omitempty"`
// 	AdMarkup        string  `json:"adm,omitempty"`
// 	NotificationURL string  `json:"nurl,omitempty"`
// 	// Add more fields according to the 2.5 spec as required
// }
