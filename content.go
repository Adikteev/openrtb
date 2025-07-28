package openrtb

import (
	"github.com/tinylib/msgp/msgp"
)

//go:generate .deps/msgp

// Content object describes the content in which the impression will appear, which may be syndicated or nonsyndicated
// content. This object may be useful when syndicated content contains impressions and does
// not necessarily match the publisher's general content. The exchange might or might not have
// knowledge of the page where the content is running, as a result of the syndication method. For
// example might be a video impression embedded in an iframe on an unknown web property or device.
type Content struct {
	ID                 string            `json:"id,omitempty" msgp:"id,omitempty"`                                 // ID uniquely identifying the content.
	Episode            int               `json:"episode,omitempty" msgp:"episode,omitempty"`                       // Episode number (typically applies to video content).
	Title              string            `json:"title,omitempty" msgp:"title,omitempty"`                           // Content title.
	Series             string            `json:"series,omitempty" msgp:"series,omitempty"`                         // Content series.
	Season             string            `json:"season,omitempty" msgp:"season,omitempty"`                         // Content season.
	Artist             string            `json:"artist,omitempty" msgp:"artist,omitempty"`                         // Artist credited with the content.
	Genre              string            `json:"genre,omitempty" msgp:"genre,omitempty"`                           // Genre that best describes the content
	Album              string            `json:"album,omitempty" msgp:"album,omitempty"`                           // Album to which the content belongs; typically for audio.
	ISRC               string            `json:"isrc,omitempty" msgp:"isrc,omitempty"`                             // International Standard Recording Code conforming to ISO - 3901.
	Producer           *Producer         `json:"producer,omitempty" msgp:"producer,omitempty"`                     // The producer.
	URL                string            `json:"url,omitempty" msgp:"url,omitempty"`                               // URL of the content, for buy-side contextualization or review.
	CategoryTaxonomy   CategoryTaxonomy  `json:"cattax,omitempty" msgp:"cattax,omitempty"`                         // Defines the taxonomy in use.
	Categories         []ContentCategory `json:"cat,omitempty" msgp:"cat,omitempty"`                               // Array of IAB content categories that describe the content.
	ProductionQuality  ProductionQuality `json:"prodq,omitempty" msgp:"prodq,omitempty"`                           // Production quality per IAB's classification.
	VideoQuality       ProductionQuality `json:"videoquality,omitempty" msgp:"videoquality,omitempty"`             // DEPRECATED. Video quality per IAB's classification.
	Context            ContentContext    `json:"context,omitempty" msgp:"context,omitempty"`                       // Type of content (game, video, text, etc.).
	ContentRating      string            `json:"contentrating,omitempty" msgp:"contentrating,omitempty"`           // Content rating (e.g., MPAA).
	UserRating         string            `json:"userrating,omitempty" msgp:"userrating,omitempty"`                 // User rating of the content (e.g., number of stars, likes, etc.).
	MediaRating        IQGRating         `json:"qagmediarating,omitempty" msgp:"qagmediarating,omitempty"`         // Media rating per QAG guidelines.
	Keywords           string            `json:"keywords,omitempty" msgp:"keywords,omitempty"`                     // Comma separated list of keywords describing the content.
	LiveStream         int               `json:"livestream,omitempty" msgp:"livestream,omitempty"`                 // 0 = not live, 1 = content is live (e.g., stream, live blog).
	SourceRelationship int               `json:"sourcerelationship,omitempty" msgp:"sourcerelationship,omitempty"` // 0 = indirect, 1 = direct.
	Length             int               `json:"len,omitempty" msgp:"len,omitempty"`                               // Length of content in seconds; appropriate for video or audio.
	Language           string            `json:"language,omitempty" msgp:"language,omitempty"`                     // Content language using ISO-639-1-alpha-2.
	LanguageB          string            `json:"langb,omitempty" msgp:"langb,omitempty"`                           // Content language using IETF BCP 47. Only one of language or langb should be present.
	Embeddable         int               `json:"embeddable,omitempty" msgp:"embeddable,omitempty"`                 // Indicator of whether or not the content is embeddable (e.g., an embeddable video player), where 0 = no, 1 = yes.
	Data               []Data            `json:"data,omitempty" msgp:"data,omitempty"`                             // Additional content data.
	Network            *ChannelEntity    `json:"network,omitempty" msgp:"network,omitempty"`                       // Details about the network the content is on.
	Channel            *ChannelEntity    `json:"channel,omitempty" msgp:"channel,omitempty"`                       // Details about the channel the content is on.
	KwArray            []string          `json:"kwarray,omitempty" msgp:"kwarray,omitempty"`                       // Array of keywords about the site. Only one of ‘keywords’ or‘kwarray’ may be present.
	Ext                msgp.Raw          `json:"ext,omitempty" msgp:"ext,omitempty"`
}
