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

	"github.com/gumnut-ai/photos-sdk-go/internal/apiform"
	"github.com/gumnut-ai/photos-sdk-go/internal/apijson"
	"github.com/gumnut-ai/photos-sdk-go/internal/apiquery"
	"github.com/gumnut-ai/photos-sdk-go/internal/requestconfig"
	"github.com/gumnut-ai/photos-sdk-go/option"
	"github.com/gumnut-ai/photos-sdk-go/packages/param"
	"github.com/gumnut-ai/photos-sdk-go/packages/respjson"
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

// Searches for assets using semantic (CLIP-based) image-content matching and/or
// structured filters on people and date range. Use this tool when the user
// describes _what's in_ the photos they want — subjects, scenes, places,
// activities, moods, objects — as opposed to browsing by album membership or exact
// ID.
//
// A natural-language `query` can be combined with structured filters
// (`person_ids`, `captured_before`, `captured_after`) for precision. For example,
// 'photos of my kids at the beach last summer' becomes
// `query='kids at the beach'` + `captured_after=2025-06-01` +
// `captured_before=2025-09-01`.
//
// **Use `list_assets` instead** when the request can be answered with exact
// filters alone (album, person, date range, ID) — it's cheaper and more
// deterministic than semantic search.
//
// Does not filter by location/place today; pass place names as part of `query` and
// rely on semantic matching until a structured location filter lands.
//
// At least one of `query`, `person_ids`, `captured_before`, or `captured_after`
// must be provided.
func (r *SearchService) Search(ctx context.Context, query SearchSearchParams, opts ...option.RequestOption) (res *SearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Searches for assets using semantic similarity and/or metadata filters. Results
// include asset metadata, faces, and people. At least one search criterion must be
// provided. Can search by text query, uploaded image, or both combined.
func (r *SearchService) SearchAssets(ctx context.Context, body SearchSearchAssetsParams, opts ...option.RequestOption) (res *SearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SearchResponse struct {
	// Matching assets ordered by semantic distance (closest first) when `query` is
	// set.
	Data []SearchResultItem `json:"data" api:"required"`
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

type SearchResultItem struct {
	// The matching asset.
	Asset AssetResponse `json:"asset" api:"required"`
	// Semantic distance from `query` (0.0 = identical, 1.0 = unrelated); lower is more
	// similar — inverted from the usual 'similarity score' convention. Null when no
	// semantic `query` was provided (structured-filter-only search).
	Distance float64 `json:"distance" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset       respjson.Field
		Distance    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResultItem) RawJSON() string { return r.JSON.raw }
func (r *SearchResultItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchSearchParams struct {
	// Only include assets captured strictly after this instant (ISO 8601; exclusive).
	// Equivalent in purpose to `local_datetime_after` on `list_assets` (naming
	// inconsistency is tracked as a follow-up).
	CapturedAfter param.Opt[time.Time] `query:"captured_after,omitzero" format:"date-time" json:"-"`
	// Only include assets captured strictly before this instant (ISO 8601; exclusive).
	// Equivalent in purpose to `local_datetime_before` on `list_assets` (naming
	// inconsistency is tracked as a follow-up).
	CapturedBefore param.Opt[time.Time] `query:"captured_before,omitzero" format:"date-time" json:"-"`
	// Library to search. Optional if the user has a single library; required when they
	// have multiple. Use `list_libraries` to enumerate available libraries.
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Natural-language description of the image content to search for. Matched against
	// CLIP image embeddings, so it works best with concrete visual concepts: subjects,
	// scenes, objects, settings ('beach sunset', 'birthday cake', 'mountain hike').
	//
	// Prefer structured params when available: use `person_ids` for people (not names
	// in `query`) and `captured_before`/`captured_after` for dates (not phrases like
	// 'in 2023' in `query`).
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Maximum number of results per page (1–200). Defaults to 20.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// 1-indexed page number. `search_assets` uses page-number pagination; the sibling
	// `list_assets` uses cursor pagination via `starting_after_id`. Increment `page`
	// to fetch subsequent pages.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Maximum semantic distance for a result to be included (0.0 = identical, 1.0 =
	// unrelated). Lower values return fewer, more confident matches; higher values
	// return more results with looser matching. Default 0.8 is moderate — try 0.6 for
	// high-precision queries, 0.9 for exploratory searches. **Note:** this is inverted
	// from the usual 'similarity score' convention where higher means more similar.
	Threshold param.Opt[float64] `query:"threshold,omitzero" json:"-"`
	// Filter to assets containing ALL of these person IDs (intersection, not union).
	// Accepts multiple `person_ids=` query params or a single comma-delimited value
	// (e.g., `person_123,person_abc`). Get person IDs from `list_people`. Plural on
	// this tool; the sibling `list_assets` uses `person_id` (singular).
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
	// Number of results per page (1-200)
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Page number
	Page param.Opt[int64] `json:"page,omitzero"`
	// Similarity threshold (lower means more similar)
	Threshold param.Opt[float64] `json:"threshold,omitzero"`
	// Image file to search for similar assets. Can be combined with text query.
	Image io.Reader `json:"image,omitzero" format:"binary"`
	// Filter to assets containing ALL of these person IDs (intersection, not union).
	// Accepts multiple `person_ids=` form fields or a single comma-delimited value
	// (e.g., `person_123,person_abc`). Get person IDs from `list_people`.
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
