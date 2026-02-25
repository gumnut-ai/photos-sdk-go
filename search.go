// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/photos-go/internal/apiform"
	"github.com/stainless-sdks/photos-go/internal/apijson"
	"github.com/stainless-sdks/photos-go/internal/apiquery"
	"github.com/stainless-sdks/photos-go/internal/requestconfig"
	"github.com/stainless-sdks/photos-go/option"
	"github.com/stainless-sdks/photos-go/packages/param"
	"github.com/stainless-sdks/photos-go/packages/respjson"
)

// SearchService contains methods and other services that help with interacting
// with the Gumnut API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	Options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r SearchService) {
	r = SearchService{}
	r.Options = opts
	return
}

// Searches for assets using semantic similarity and/or metadata filters. Results
// include asset metadata, faces, and people. At least one search criterion must be
// provided.
func (r *SearchService) Search(ctx context.Context, query SearchSearchParams, opts ...option.RequestOption) (res *SearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Searches for assets using semantic similarity and/or metadata filters. Results
// include asset metadata, faces, and people. At least one search criterion must be
// provided. Can search by text query, uploaded image, or both combined.
func (r *SearchService) SearchAssets(ctx context.Context, body SearchSearchAssetsParams, opts ...option.RequestOption) (res *SearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SearchResponse struct {
	Data []SearchResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchResponseData struct {
	// Represents a photo or video asset with metadata and access URLs.
	Asset    AssetResponse `json:"asset" api:"required"`
	Distance float64       `json:"distance" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset       respjson.Field
		Distance    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponseData) RawJSON() string { return r.JSON.raw }
func (r *SearchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchSearchParams struct {
	// Filter to only include assets captured after this date (ISO format).
	CapturedAfter param.Opt[time.Time] `query:"captured_after,omitzero" format:"date-time" json:"-"`
	// Filter to only include assets captured before this date (ISO format).
	CapturedBefore param.Opt[time.Time] `query:"captured_before,omitzero" format:"date-time" json:"-"`
	// Library to search assets from (optional)
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// The text query to search for. If you want to search for a specific person or set
	// of people, use the person_ids parameter instead.If you want to search for a
	// photos taken during a specific date range, use the captured_before and
	// captured_after parameters instead.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Number of results per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Similarity threshold (lower means more similar)
	Threshold param.Opt[float64] `query:"threshold,omitzero" json:"-"`
	// Filter to only include assets containing ALL of these person IDs. Can be
	// comma-delimited string (e.g. 'person_123,person_abc') or multiple query
	// parameters.
	PersonIDs []string `query:"person_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SearchSearchParams]'s query parameters as `url.Values`.
func (r SearchSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SearchSearchAssetsParams struct {
	// Filter to only include assets captured after this date (ISO format).
	CapturedAfter param.Opt[time.Time] `json:"captured_after,omitzero" format:"date-time"`
	// Filter to only include assets captured before this date (ISO format).
	CapturedBefore param.Opt[time.Time] `json:"captured_before,omitzero" format:"date-time"`
	// Library to search assets from (optional)
	LibraryID param.Opt[string] `json:"library_id,omitzero"`
	// The text query to search for. If you want to search for a specific person or set
	// of people, use the person_ids parameter instead.If you want to search for a
	// photos taken during a specific date range, use the captured_before and
	// captured_after parameters instead.
	Query param.Opt[string] `json:"query,omitzero"`
	// Number of results per page
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Page number
	Page param.Opt[int64] `json:"page,omitzero"`
	// Similarity threshold (lower means more similar)
	Threshold param.Opt[float64] `json:"threshold,omitzero"`
	// Image file to search for similar assets. Can be combined with text query.
	Image io.Reader `json:"image,omitzero" format:"binary"`
	// Filter to only include assets containing ALL of these person IDs. Can be
	// comma-delimited string (e.g. 'person_123,person_abc') or multiple query
	// parameters.
	PersonIDs []string `json:"person_ids,omitzero"`
	paramObj
}

func (r SearchSearchAssetsParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
