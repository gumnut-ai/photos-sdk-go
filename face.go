// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/gumnut-ai/photos-sdk-go/internal/apijson"
	"github.com/gumnut-ai/photos-sdk-go/internal/apiquery"
	"github.com/gumnut-ai/photos-sdk-go/internal/requestconfig"
	"github.com/gumnut-ai/photos-sdk-go/option"
	"github.com/gumnut-ai/photos-sdk-go/packages/pagination"
	"github.com/gumnut-ai/photos-sdk-go/packages/param"
	"github.com/gumnut-ai/photos-sdk-go/packages/respjson"
)

// FaceService contains methods and other services that help with interacting with
// the Gumnut API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFaceService] method instead.
type FaceService struct {
	Options []option.RequestOption
}

// NewFaceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFaceService(opts ...option.RequestOption) (r FaceService) {
	r = FaceService{}
	r.Options = opts
	return
}

// Retrieves details for a specific face.
func (r *FaceService) Get(ctx context.Context, faceID string, query FaceGetParams, opts ...option.RequestOption) (res *FaceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if faceID == "" {
		err = errors.New("missing required face_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/faces/%s", faceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Updates the details of a specific face, currently only supporting
// associating/disassociating with a person.
func (r *FaceService) Update(ctx context.Context, faceID string, params FaceUpdateParams, opts ...option.RequestOption) (res *FaceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if faceID == "" {
		err = errors.New("missing required face_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/faces/%s", faceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieves a paginated list of faces, optionally filtered by asset, person, or
// specific face IDs, ordered by creation time, descending.
//
// **Pagination:** When `has_more` is true, pass the `id` of the last face in
// `data` as `starting_after_id` to fetch the next page.
func (r *FaceService) List(ctx context.Context, query FaceListParams, opts ...option.RequestOption) (res *pagination.CursorPage[FaceResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/faces"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieves a paginated list of faces, optionally filtered by asset, person, or
// specific face IDs, ordered by creation time, descending.
//
// **Pagination:** When `has_more` is true, pass the `id` of the last face in
// `data` as `starting_after_id` to fetch the next page.
func (r *FaceService) ListAutoPaging(ctx context.Context, query FaceListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[FaceResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a specific face entry. This does not delete the associated asset or
// person.
func (r *FaceService) Delete(ctx context.Context, faceID string, body FaceDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if faceID == "" {
		err = errors.New("missing required face_id parameter")
		return err
	}
	path := fmt.Sprintf("api/faces/%s", faceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Represents a detected face in an asset with facial recognition data.
type FaceResponse struct {
	// Unique face identifier with 'face\_' prefix
	ID string `json:"id" api:"required"`
	// ID of the asset containing this face
	AssetID string `json:"asset_id" api:"required"`
	// Face location as {x, y, w, h} coordinates in pixels
	BoundingBox map[string]int64 `json:"bounding_box" api:"required"`
	// When this face was detected and recorded
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When this face record was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Asset variants for this face: 'thumbnail' with face crop
	AssetURLs map[string]FaceResponseAssetURL `json:"asset_urls" api:"nullable"`
	// ID of the person this face belongs to (if identified)
	PersonID string `json:"person_id" api:"nullable"`
	// For video files, timestamp in milliseconds when face appears
	TimestampMs int64 `json:"timestamp_ms" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AssetID     respjson.Field
		BoundingBox respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		AssetURLs   respjson.Field
		PersonID    respjson.Field
		TimestampMs respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaceResponse) RawJSON() string { return r.JSON.raw }
func (r *FaceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single image variant with its URL, MIME type, and target width.
type FaceResponseAssetURL struct {
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
func (r FaceResponseAssetURL) RawJSON() string { return r.JSON.raw }
func (r *FaceResponseAssetURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaceGetParams struct {
	// Library ID (required if user has multiple libraries)
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FaceGetParams]'s query parameters as `url.Values`.
func (r FaceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FaceUpdateParams struct {
	// Library ID (required if user has multiple libraries)
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	PersonID  param.Opt[string] `json:"person_id,omitzero"`
	paramObj
}

func (r FaceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FaceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FaceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [FaceUpdateParams]'s query parameters as `url.Values`.
func (r FaceUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FaceListParams struct {
	// Filter by faces in a specific asset
	AssetID param.Opt[string] `query:"asset_id,omitzero" json:"-"`
	// Library ID (required if user has multiple libraries)
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Filter by faces associated with a specific person
	PersonID param.Opt[string] `query:"person_id,omitzero" json:"-"`
	// Cursor for pagination. Pass the `id` of the last face from the previous page to
	// get the next page.
	StartingAfterID param.Opt[string] `query:"starting_after_id,omitzero" json:"-"`
	// Max number of faces to return (1-200)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by specific face IDs (max 100)
	IDs []string `query:"ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FaceListParams]'s query parameters as `url.Values`.
func (r FaceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FaceDeleteParams struct {
	// Library ID (required if user has multiple libraries)
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FaceDeleteParams]'s query parameters as `url.Values`.
func (r FaceDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
