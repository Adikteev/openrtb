package openrtb

import (
	"github.com/tinylib/msgp/msgp"
)

// Banner object must be included directly in the impression object if the impression offered
// for auction is display or rich media, or it may be optionally embedded in the video object to
// describe the companion banners available for the linear or non-linear video ad.  The banner
// object may include a unique identifier; this can be useful if these IDs can be leveraged in the
// VAST response to dictate placement of the companion creatives when multiple companion ad
// opportunities of the same size are available on a page.

//go:generate .deps/msgp

type Banner struct {
	Width        int                 `json:"w,omitempty" msgp:"w,omitempty"`               // Width
	Height       int                 `json:"h,omitempty" msgp:"h,omitempty"`               // Height
	Formats      []Format            `json:"format,omitempty" msgp:"format,omitempty"`     // Array of format objects representing the banner sizes permitted.
	WidthMax     int                 `json:"wmax,omitempty" msgp:"wmax,omitempty"`         // Width maximum DEPRECATED
	HeightMax    int                 `json:"hmax,omitempty" msgp:"hmax,omitempty"`         // Height maximum DEPRECATED
	WidthMin     int                 `json:"wmin,omitempty" msgp:"wmin,omitempty"`         // Width minimum DEPRECATED
	HeightMin    int                 `json:"hmin,omitempty" msgp:"hmin,omitempty"`         // Height minimum DEPRECATED
	ID           string              `json:"id,omitempty" msgp:"id,omitempty"`             // A unique identifier
	BlockedTypes []BannerType        `json:"btype,omitempty" msgp:"btype,omitempty"`       // Blocked banner types
	BlockedAttrs []CreativeAttribute `json:"battr,omitempty" msgp:"battr,omitempty"`       // Blocked creative attributes
	Position     AdPosition          `json:"pos,omitempty" msgp:"pos,omitempty"`           // Ad Position
	MIMEs        []string            `json:"mimes,omitempty" msgp:"mimes,omitempty"`       // Whitelist of content MIME types supported
	TopFrame     int                 `json:"topframe,omitempty" msgp:"topframe,omitempty"` // Default: 0 ("1": Delivered in top frame, "0": Elsewhere)
	ExpDirs      []ExpDir            `json:"expdir,omitempty" msgp:"expdir,omitempty"`     // Specify properties for an expandable ad
	APIs         []APIFramework      `json:"api,omitempty" msgp:"api,omitempty"`           // List of supported API frameworks
	VCM          int                 `json:"vcm,omitempty" msgp:"vcm,omitempty"`           // Represents the relationship with video. 0 = concurrent, 1 = end-card
	Ext          msgp.Raw            `json:"ext,omitempty" msgp:"ext,omitempty"`
}
