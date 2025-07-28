package openrtb

import (
	"encoding/json"
	"errors"

	"github.com/bytedance/sonic"
)

// Validation errors
var (
	ErrInvalidVideoNoMIMEs     = errors.New("openrtb: video has no mimes")
	ErrInvalidVideoNoLinearity = errors.New("openrtb: video linearity missing")
	ErrInvalidVideoNoProtocols = errors.New("openrtb: video protocols missing")
)

//go:generate .deps/msgp

// Video object must be included directly in the impression object if the impression offered
// for auction is an in-stream video ad opportunity.
type Video struct {
	MIMEs           []string            `json:"mimes,omitempty" msgp:"mimes,omitempty"`                   // Content MIME types supported.
	MinDuration     int                 `json:"minduration,omitempty" msgp:"minduration,omitempty"`       // Minimum video ad duration in seconds
	MaxDuration     int                 `json:"maxduration,omitempty" msgp:"maxduration,omitempty"`       // Maximum video ad duration in seconds
	Protocols       []Protocol          `json:"protocols,omitempty" msgp:"protocols,omitempty"`           // Video bid response protocols
	Protocol        Protocol            `json:"protocol,omitempty" msgp:"protocol,omitempty"`             // Video bid response protocols DEPRECATED
	Width           int                 `json:"w" msgp:"w"`                                               // Width of the player in pixels
	Height          int                 `json:"h" msgp:"h"`                                               // Height of the player in pixels
	StartDelay      StartDelay          `json:"startdelay,omitempty" msgp:"startdelay,omitempty"`         // Indicates the start delay in seconds
	Linearity       VideoLinearity      `json:"linearity,omitempty" msgp:"linearity,omitempty"`           // Indicates whether the ad impression is linear or non-linear
	Skip            *int                `json:"skip,omitempty" msgp:"skip,omitempty"`                     // Indicates if the player will allow the video to be skipped, where 0 = no, 1 = yes.
	SkipMin         int                 `json:"skipmin,omitempty" msgp:"skipmin,omitempty"`               // Videos of total duration greater than this number of seconds can be skippable
	SkipAfter       int                 `json:"skipafter,omitempty" msgp:"skipafter,omitempty"`           // Number of seconds a video must play before skipping is enabled
	Sequence        int                 `json:"sequence,omitempty" msgp:"sequence,omitempty"`             // Default: 1
	BlockedAttrs    []CreativeAttribute `json:"battr,omitempty" msgp:"battr,omitempty"`                   // Blocked creative attributes
	MaxExtended     int                 `json:"maxextended,omitempty" msgp:"maxextended,omitempty"`       // Maximum extended video ad duration
	MinBitrate      int                 `json:"minbitrate,omitempty" msgp:"minbitrate,omitempty"`         // Minimum bit rate in Kbps
	MaxBitrate      int                 `json:"maxbitrate,omitempty" msgp:"maxbitrate,omitempty"`         // Maximum bit rate in Kbps
	BoxingAllowed   *int                `json:"boxingallowed,omitempty" msgp:"boxingallowed,omitempty"`   // If exchange publisher has rules preventing letter boxing
	PodID           string              `json:"podid,omitempty" msgp:"podid,omitempty"`                   // Pod id unique identifier for video ad pod
	PodDuration     int                 `json:"poddur,omitempty" msgp:"poddur,omitempty"`                 // Pod Duration total amount of time in seconds that advertisers may fill for a video ad pod
	PodSequence     PodSequence         `json:"podseq,omitempty" msgp:"podseq,omitempty"`                 // Pod Sequence position of the video ad pod
	SlotInPod       SlotPositionInPod   `json:"slotinpod,omitempty" msgp:"slotinpod,omitempty"`           // Slot Position in ad
	RqdDurs         []int64             `json:"rqddurs,omitempty" msgp:"rqddurs,omitempty"`               // Precise acceptable durations for video creatives inseconds.
	MinCPMPerSecond float64             `json:"mincpmpersec,omitempty" msgp:"mincpmpersec,omitempty"`     //   Minimum CPM per second. This is a price floor for the portion of a video ad pod
	PlaybackMethods []VideoPlayback     `json:"playbackmethod,omitempty" msgp:"playbackmethod,omitempty"` // List of allowed playback methods
	Delivery        []ContentDelivery   `json:"delivery,omitempty" msgp:"delivery,omitempty"`             // List of supported delivery methods
	Position        AdPosition          `json:"pos,omitempty" msgp:"pos,omitempty"`                       // Ad Position
	CompanionAds    []Banner            `json:"companionad,omitempty" msgp:"companionad,omitempty"`
	APIs            []APIFramework      `json:"api,omitempty" msgp:"api,omitempty"` // List of supported API frameworks
	CompanionTypes  []CompanionType     `json:"companiontype,omitempty" msgp:"companiontype,omitempty"`
	Placement       VideoPlacement      `json:"placement,omitempty" msgp:"placement,omitempty"` // Video placement type, DEPRECATED
	Plcmt           VideoPlcmt          `json:"plcmt,omitempty" msgp:"plcmt,omitempty"`         // Video Plcmt type ad defined in ADCOM1.0
	Ext             json.RawMessage     `json:"ext,omitempty" msgp:"ext,omitempty"`
}

type jsonVideo Video

// Validate the object
func (v *Video) Validate() error {
	if len(v.MIMEs) == 0 {
		return ErrInvalidVideoNoMIMEs
	} else if v.Linearity == 0 {
		return ErrInvalidVideoNoLinearity
	} else if v.Protocol == 0 && len(v.Protocols) == 0 {
		return ErrInvalidVideoNoProtocols
	}
	return nil
}

// GetBoxingAllowed returns the boxing-allowed indicator
func (v *Video) GetBoxingAllowed() int {
	if v.BoxingAllowed != nil {
		return *v.BoxingAllowed
	}
	return 1
}

// MarshalJSON custom marshalling with normalization
func (v *Video) MarshalJSON() ([]byte, error) {
	v.normalize()
	return sonic.Marshal((*jsonVideo)(v))
}

// UnmarshalJSON custom unmarshalling with normalization
func (v *Video) UnmarshalJSON(data []byte) error {
	var h jsonVideo
	if err := sonic.Unmarshal(data, &h); err != nil {
		return err
	}

	*v = (Video)(h)
	v.normalize()
	return nil
}

func (v *Video) normalize() {
	if v.Sequence == 0 {
		v.Sequence = 1
	}
	if v.Linearity == 0 {
		v.Linearity = VideoLinearityLinear
	}
}
