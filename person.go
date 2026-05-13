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
	"github.com/gumnut-ai/photos-sdk-go/shared"
)

// PersonService contains methods and other services that help with interacting
// with the Gumnut API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPersonService] method instead.
type PersonService struct {
	Options []option.RequestOption
}

// NewPersonService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPersonService(opts ...option.RequestOption) (r PersonService) {
	r = PersonService{}
	r.Options = opts
	return
}

// Creates a new person record (a named identity for grouping faces). Most people
// are auto-created by face clustering, so this tool is typically used only when
// the user explicitly wants to introduce a new identity before any faces are
// attached.
//
// To assign an existing face to an existing person, use `update_face` with the
// target `person_id`.
func (r *PersonService) New(ctx context.Context, body PersonNewParams, opts ...option.RequestOption) (res *PersonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/people"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetches one person's metadata by ID (name, asset count, thumbnail, etc.). Use
// this when you already have a `person_id`. To find photos that contain this
// person, use `search_assets` with `person_ids` or `list_assets` with `person_id`.
func (r *PersonService) Get(ctx context.Context, personID string, query PersonGetParams, opts ...option.RequestOption) (res *PersonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/people/%s", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Updates a person's name, birth date, visibility, or thumbnail. Only the fields
// included in the request body are changed. Typical use: assigning a name ('name
// this face cluster "Alice"') or choosing a better thumbnail.
//
// This tool does not move faces between people — use `update_face` with a new
// `person_id` for that.
func (r *PersonService) Update(ctx context.Context, personID string, body PersonUpdateParams, opts ...option.RequestOption) (res *PersonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/people/%s", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of people (named identities that group one or more
// faces), ordered by creation time (newest first), optionally filtered by asset,
// album, name, or ID. Use this to enumerate who appears in the library, to resolve
// a user-typed name to a `person_id`, or to find who appears in a specific asset
// or album.
//
// By default only **named** people are returned; pass `name_filter=all` or
// `name_filter=unnamed` to include clusters that haven't been named yet.
//
// To list the underlying faces for a specific person, use `list_faces` with
// `person_id`.
//
// **Pagination** is cursor-based: when `has_more` is true, pass the `id` of the
// last person in `data` as `starting_after_id` to fetch the next page.
func (r *PersonService) List(ctx context.Context, query PersonListParams, opts ...option.RequestOption) (res *pagination.CursorPage[PersonResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/people"
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

// Returns a paginated list of people (named identities that group one or more
// faces), ordered by creation time (newest first), optionally filtered by asset,
// album, name, or ID. Use this to enumerate who appears in the library, to resolve
// a user-typed name to a `person_id`, or to find who appears in a specific asset
// or album.
//
// By default only **named** people are returned; pass `name_filter=all` or
// `name_filter=unnamed` to include clusters that haven't been named yet.
//
// To list the underlying faces for a specific person, use `list_faces` with
// `person_id`.
//
// **Pagination** is cursor-based: when `has_more` is true, pass the `id` of the
// last person in `data` as `starting_after_id` to fetch the next page.
func (r *PersonService) ListAutoPaging(ctx context.Context, query PersonListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[PersonResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes the person record; the faces that were attached to this person are not
// deleted — they become unassigned and will be re-clustered on the next clustering
// pass.
//
// Use `update_face` with `person_id=null` to detach a specific face without
// deleting the whole person. Use `delete_face` to remove a face detection
// entirely.
func (r *PersonService) Delete(ctx context.Context, personID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return err
	}
	path := fmt.Sprintf("api/people/%s", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Merges one or more source people into the primary person identified by the URL.
// All faces from source people are reassigned to the primary person. Source people
// are permanently deleted (this cannot be undone). The primary person's centroid
// embedding is recalculated.
//
// In the degenerate case where the primary and all sources are unnamed and have
// zero faces, the primary is auto-deleted by the post-merge centroid recompute
// (GUM-681) and the response is `204 No Content`.
func (r *PersonService) Merge(ctx context.Context, personID string, body PersonMergeParams, opts ...option.RequestOption) (res *PersonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/people/%s/merge", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Cohesion metrics for a Person's face cluster — surfaced via
// `include=cluster_metrics` on the people endpoints. These describe how tight the
// cluster is in embedding space (lower = more cohesive) and drive both the
// production face-assignment cohesion gate and the operator-facing face cleanup
// dashboard.
type ClusterMetricsResponse struct {
	// Number of faces that fed into the centroid and pairwise metrics. This is the
	// cluster-membership count, **not** the same as `asset_count` — `face_count`
	// counts every face row, while `asset_count` counts distinct assets (one asset can
	// contribute multiple faces of the same person).
	FaceCount int64 `json:"face_count" api:"required"`
	// Mean pairwise cosine distance between faces in this person's cluster.
	PairwiseMean float64 `json:"pairwise_mean" api:"required"`
	// 90th-percentile pairwise cosine distance between faces in this person's cluster.
	// Lower = more cohesive cluster; loose clusters (higher pairwise_p90) are gated
	// out of the face-assignment path to prevent further drift.
	PairwiseP90 float64 `json:"pairwise_p90" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FaceCount    respjson.Field
		PairwiseMean respjson.Field
		PairwiseP90  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterMetricsResponse) RawJSON() string { return r.JSON.raw }
func (r *ClusterMetricsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a person identified through face clustering and recognition.
type PersonResponse struct {
	// Unique person identifier with 'person\_' prefix
	ID string `json:"id" api:"required"`
	// When this person record was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this person is marked as a favorite
	IsFavorite bool `json:"is_favorite" api:"required"`
	// Whether this person should be hidden from the UI
	IsHidden bool `json:"is_hidden" api:"required"`
	// When this person record was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Number of unique photos this person appears in, or null if not computed
	AssetCount int64 `json:"asset_count" api:"nullable"`
	// Asset variants from this person's thumbnail face. May be null when embedded in
	// an AssetResponse; use /api/people endpoints for full person data.
	AssetURLs map[string]shared.AssetVariant `json:"asset_urls" api:"nullable"`
	// Optional birth date of this person
	BirthDate time.Time `json:"birth_date" api:"nullable" format:"date"`
	// Cohesion metrics for a Person's face cluster — surfaced via
	// `include=cluster_metrics` on the people endpoints. These describe how tight the
	// cluster is in embedding space (lower = more cohesive) and drive both the
	// production face-assignment cohesion gate and the operator-facing face cleanup
	// dashboard.
	ClusterMetrics ClusterMetricsResponse `json:"cluster_metrics" api:"nullable"`
	// Optional name assigned to this person
	Name string `json:"name" api:"nullable"`
	// ID of the face resource used as this person's thumbnail
	ThumbnailFaceID string `json:"thumbnail_face_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		IsFavorite      respjson.Field
		IsHidden        respjson.Field
		UpdatedAt       respjson.Field
		AssetCount      respjson.Field
		AssetURLs       respjson.Field
		BirthDate       respjson.Field
		ClusterMetrics  respjson.Field
		Name            respjson.Field
		ThumbnailFaceID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonNewParams struct {
	// Optional birth date (ISO 8601 date, YYYY-MM-DD) for this person.
	BirthDate param.Opt[time.Time] `json:"birth_date,omitzero" format:"date"`
	// If true, the person is marked as a favorite. Defaults to false.
	IsFavorite param.Opt[bool] `json:"is_favorite,omitzero"`
	// If true, the person is hidden from default listings. Defaults to false.
	IsHidden param.Opt[bool] `json:"is_hidden,omitzero"`
	// Library to create the person in. Optional if the user has a single library;
	// required when they have multiple.
	LibraryID param.Opt[string] `json:"library_id,omitzero"`
	// Display name for the new person (e.g., 'Alice'). Optional — unnamed people can
	// be named later via `update_person`.
	Name param.Opt[string] `json:"name,omitzero"`
	// ID of the face to use as this person's thumbnail (with `face_` prefix).
	// Typically set after the person has at least one associated face — get face IDs
	// from `list_faces`.
	ThumbnailFaceID param.Opt[string] `json:"thumbnail_face_id,omitzero"`
	paramObj
}

func (r PersonNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonGetParams struct {
	// Opt-in expansion fields. See `list_people` for supported values. Accepts
	// multiple `include=` query params or a single comma-delimited value.
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PersonGetParams]'s query parameters as `url.Values`.
func (r PersonGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PersonUpdateParams struct {
	// New birth date (ISO 8601 date). Omit to leave unchanged.
	BirthDate param.Opt[time.Time] `json:"birth_date,omitzero" format:"date"`
	// Mark or unmark this person as a favorite. Omit to leave unchanged.
	IsFavorite param.Opt[bool] `json:"is_favorite,omitzero"`
	// Hide or unhide this person. Omit to leave unchanged.
	IsHidden param.Opt[bool] `json:"is_hidden,omitzero"`
	// New display name. Omit to leave unchanged.
	Name param.Opt[string] `json:"name,omitzero"`
	// New thumbnail face ID for this person. Omit to leave unchanged. Get face IDs
	// from `list_faces`.
	ThumbnailFaceID param.Opt[string] `json:"thumbnail_face_id,omitzero"`
	paramObj
}

func (r PersonUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonListParams struct {
	// Return only people who appear in at least one asset of this album. Useful for
	// 'who is in this album?'.
	AlbumID param.Opt[string] `query:"album_id,omitzero" json:"-"`
	// Return only people who have at least one face in this asset. Useful for 'who is
	// in this photo?'.
	AssetID param.Opt[string] `query:"asset_id,omitzero" json:"-"`
	// Library to list from. Optional if the user has a single library; required when
	// they have multiple.
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Filter by name using case-insensitive substring matching. Use this to resolve a
	// user-supplied name like 'Alice' into a `person_id`, then pass that ID into
	// `search_assets.person_ids` or `list_assets.person_id`.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Cursor for pagination. Pass the `id` of the last person in the previous
	// response's `data` to fetch the next page. Omit for the first page.
	StartingAfterID param.Opt[string] `query:"starting_after_id,omitzero" json:"-"`
	// Maximum number of people to return per page (1–200). Defaults to 20.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Look up specific people by ID (max 100; each ID has the `person_` prefix).
	// Accepts multiple `ids=` query params or a single comma-delimited value (e.g.,
	// `ids=person_1,person_2`). When set, `name_filter` defaults to `all` so unnamed
	// clusters are included in the lookup.
	IDs []string `query:"ids,omitzero" json:"-"`
	// Opt-in expansion fields. Supported values: `cluster_metrics` (adds the nested
	// `cluster_metrics` object — `pairwise_p90`, `pairwise_mean`, `face_count` — for
	// each Person with a populated centroid). Accepts multiple `include=` query params
	// or a single comma-delimited value. Unknown values return 422.
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by name status: `named` returns only people with a name; `unnamed`
	// returns only nameless face clusters awaiting a name; `all` returns both.
	// Defaults to `named` (or `all` when `ids` is provided).
	//
	// Any of "named", "unnamed", "all".
	NameFilter PersonListParamsNameFilter `query:"name_filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PersonListParams]'s query parameters as `url.Values`.
func (r PersonListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by name status: `named` returns only people with a name; `unnamed`
// returns only nameless face clusters awaiting a name; `all` returns both.
// Defaults to `named` (or `all` when `ids` is provided).
type PersonListParamsNameFilter string

const (
	PersonListParamsNameFilterNamed   PersonListParamsNameFilter = "named"
	PersonListParamsNameFilterUnnamed PersonListParamsNameFilter = "unnamed"
	PersonListParamsNameFilterAll     PersonListParamsNameFilter = "all"
)

type PersonMergeParams struct {
	// IDs of the people to merge into the primary person. These people will be deleted
	// after their faces are moved.
	SourcePersonIDs []string `json:"source_person_ids,omitzero" api:"required"`
	paramObj
}

func (r PersonMergeParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonMergeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonMergeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
