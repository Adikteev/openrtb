package openrtb

import (
	"errors"
	"github.com/tinylib/msgp/msgp"
)

//go:generate .deps/msgp

// Validation errors
var (
	ErrInvalidImpNoID        = errors.New("openrtb: impression ID missing")
	ErrInvalidImpMultiAssets = errors.New("openrtb: impression has multiple assets") // at least two out of Banner, Video, Native
)

// Impression or the "imp" object describes the ad position or impression being auctioned. A single bid request
// can include multiple "imp" objects, a use case for which might be an exchange that supports
// selling all ad positions on a given page as a bundle.  Each "imp" object has a required ID so that
// bids can reference them individually.  An exchange can also conduct private auctions by
// restricting involvement to specific subsets of seats within bidders.
// The presence of Banner, Video, and/or Native objects
// subordinate to the Imp object indicates the type of impression being offered.
type Impression struct {
	ID                    string         `json:"id" msgp:"id"` // A unique identifier for this impression
	Banner                *Banner        `json:"banner,omitempty" msgp:"banner,omitempty"`
	Video                 *Video         `json:"video,omitempty" msgp:"video,omitempty"`
	Audio                 *Audio         `json:"audio,omitempty" msgp:"audio,omitempty"`
	Native                *Native        `json:"native,omitempty" msgp:"native,omitempty"`
	PMP                   *PMP           `json:"pmp,omitempty" msgp:"pmp,omitempty"`                             // A reference to the PMP object containing any Deals eligible for the impression object.
	DisplayManager        string         `json:"displaymanager,omitempty" msgp:"displaymanager,omitempty"`       // Name of ad mediation partner, SDK technology, etc
	DisplayManagerVersion string         `json:"displaymanagerver,omitempty" msgp:"displaymanagerver,omitempty"` // Version of the above
	Interstitial          int            `json:"instl,omitempty" msgp:"instl,omitempty"`                         // Interstitial, Default: 0 ("1": Interstitial, "0": Something else)
	TagID                 string         `json:"tagid,omitempty" msgp:"tagid,omitempty"`                         // IDentifier for specific ad placement or ad tag
	BidFloor              float64        `json:"bidfloor,omitempty" msgp:"bidfloor,omitempty"`                   // Bid floor for this impression in CPM
	BidFloorCurrency      string         `json:"bidfloorcur,omitempty" msgp:"bidfloorcur,omitempty"`             // Currency of bid floor
	Secure                NumberOrString `json:"secure,omitempty" msgp:"secure,omitempty"`                       // Flag to indicate whether the impression requires secure HTTPS URL creative assets and markup.
	Quantity              *Quantity      `json:"qty,omitempty" msgp:"qty,omitempty"`                             // Includes the impression multiplier, and describes its source.
	Exp                   int            `json:"exp,omitempty" msgp:"exp,omitempty"`                             // Advisory as to the number of seconds that may elapse between the auction and the actual impression.
	IFrameBusters         []string       `json:"iframebuster,omitempty" msgp:"iframebuster,omitempty"`           // Array of names for supportediframe busters.
	Rewarded              int            `json:"rwdd,omitempty" msgp:"rwdd,omitempty"`                           // Impression is rewarded, Default: 0 ("1": yes, "0": no)
	Metric                []Metric       `json:"metric,omitempty" msgp:"metric,omitempty"`
	Ext                   msgp.Raw       `json:"ext,omitempty" msgp:"ext,omitempty"`
}

func (imp *Impression) assetCount() int {
	n := 0
	if imp.Banner != nil {
		n++
	}
	if imp.Video != nil {
		n++
	}
	if imp.Native != nil {
		n++
	}
	return n
}

// Validate the `imp` object
func (imp *Impression) Validate() error {
	if imp.ID == "" {
		return ErrInvalidImpNoID
	}

	if count := imp.assetCount(); count > 1 {
		return ErrInvalidImpMultiAssets
	}

	if imp.Video != nil {
		if err := imp.Video.Validate(); err != nil {
			return err
		}
	}
	if imp.Quantity != nil {
		if err := imp.Quantity.Validate(); err != nil {
			return err
		}
	}

	return nil
}
