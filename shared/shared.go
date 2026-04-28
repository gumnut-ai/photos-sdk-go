// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/gumnut-ai/photos-sdk-go/internal/apijson"
	"github.com/gumnut-ai/photos-sdk-go/packages/param"
	"github.com/gumnut-ai/photos-sdk-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// A single image variant with its URL, MIME type, and target width.
type AssetVariant struct {
	// MIME type of the served image
	Mimetype string `json:"mimetype" api:"required"`
	// URL to fetch this image variant
	URL string `json:"url" api:"required"`
	// Target width in pixels (null if unknown)
	Width int64 `json:"width" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mimetype    respjson.Field
		URL         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetVariant) RawJSON() string { return r.JSON.raw }
func (r *AssetVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
