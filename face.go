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

// FaceService contains methods and other services that help with interacting with
// the Gumnut AI API.
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

// Fetches one face's details (bounding box, assigned person, timestamps,
// thumbnail). Use when you already have a `face_id`.
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

// Assigns a face to a specific person, or detaches it (set `person_id` to null).
// This is the right tool for 'this face is Alice' or 'this face isn't Bob after
// all'.
//
// Currently only the `person_id` field is mutable. To create a brand-new identity
// first, call `create_person`; to delete the face detection entirely, use
// `delete_face`.
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

// Returns a paginated list of individual face detections (with bounding boxes),
// ordered by creation time (newest first). Each row is a single face in a single
// asset — a person with many photos will have many face rows.
//
// **Use `list_people` instead** when the user wants the grouped identities ('list
// everyone in my library') rather than individual face detections. This tool is
// useful for curating clustering results, finding unassigned faces, or picking a
// thumbnail face for a person via `update_person.thumbnail_face_id`.
//
// **Pagination** is cursor-based: when `has_more` is true, pass the `id` of the
// last face in `data` as `starting_after_id` to fetch the next page.
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

// Returns a paginated list of individual face detections (with bounding boxes),
// ordered by creation time (newest first). Each row is a single face in a single
// asset — a person with many photos will have many face rows.
//
// **Use `list_people` instead** when the user wants the grouped identities ('list
// everyone in my library') rather than individual face detections. This tool is
// useful for curating clustering results, finding unassigned faces, or picking a
// thumbnail face for a person via `update_person.thumbnail_face_id`.
//
// **Pagination** is cursor-based: when `has_more` is true, pass the `id` of the
// last face in `data` as `starting_after_id` to fetch the next page.
func (r *FaceService) ListAutoPaging(ctx context.Context, query FaceListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[FaceResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Removes one face detection row. The underlying asset and the person this face
// was assigned to are both preserved.
//
// **Use `update_face` with `person_id=null` instead** when the user wants to
// disassociate the face from a person without discarding the detection (so
// re-clustering can try again). Use `delete_person` to remove a person; use
// `trash_assets` (or `permanently_delete_assets` for irreversible removal) to
// remove the photo entirely.
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

// Per-face cluster-assignment diagnostics: how well the face fits its
// currently-assigned Person, and which other Persons are nearby in embedding
// space. Surfaced via `include=cluster_assignment` on the faces endpoints — used
// by the operator-facing face cleanup dashboard to triage mis-clustered faces.
type ClusterAssignmentResponse struct {
	// Persons in the same library that pass the same gate shape as production face
	// assignment, surfaced with deliberately relaxed thresholds so the list is a
	// superset of what the automated path would admit. Sorted ascending by distance.
	// Excludes the face's currently-assigned Person (its distance is in
	// `distance_to_person`). Empty when no eligible Persons pass the gate.
	Candidates []ClusterAssignmentResponseCandidate `json:"candidates"`
	// Cosine distance from the face's embedding to its currently-assigned Person's
	// centroid. Lower = better fit. Null when the face is unassigned or when the
	// assigned Person has no centroid.
	DistanceToPerson float64 `json:"distance_to_person" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Candidates       respjson.Field
		DistanceToPerson respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAssignmentResponse) RawJSON() string { return r.JSON.raw }
func (r *ClusterAssignmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Person whose centroid is close enough to a given face's embedding that it
// would be considered for assignment — surfaced under
// `ClusterAssignmentResponse.candidates`.
type ClusterAssignmentResponseCandidate struct {
	// Cosine distance from the face's embedding to this Person's centroid (lower =
	// closer).
	Distance float64 `json:"distance" api:"required"`
	// Person ID (with 'person\_' prefix) of the candidate.
	PersonID string `json:"person_id" api:"required"`
	// Display name of the candidate Person, or null for unnamed clusters. Candidates
	// surface the same Persons production assignment considers, which includes unnamed
	// clusters.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance    respjson.Field
		PersonID    respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClusterAssignmentResponseCandidate) RawJSON() string { return r.JSON.raw }
func (r *ClusterAssignmentResponseCandidate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	AssetURLs map[string]shared.AssetVariant `json:"asset_urls" api:"nullable"`
	// Per-face cluster-assignment diagnostics: how well the face fits its
	// currently-assigned Person, and which other Persons are nearby in embedding
	// space. Surfaced via `include=cluster_assignment` on the faces endpoints — used
	// by the operator-facing face cleanup dashboard to triage mis-clustered faces.
	ClusterAssignment ClusterAssignmentResponse `json:"cluster_assignment" api:"nullable"`
	// ID of the person this face belongs to (if identified)
	PersonID string `json:"person_id" api:"nullable"`
	// For video files, timestamp in milliseconds when face appears
	TimestampMs int64 `json:"timestamp_ms" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AssetID           respjson.Field
		BoundingBox       respjson.Field
		CreatedAt         respjson.Field
		UpdatedAt         respjson.Field
		AssetURLs         respjson.Field
		ClusterAssignment respjson.Field
		PersonID          respjson.Field
		TimestampMs       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaceResponse) RawJSON() string { return r.JSON.raw }
func (r *FaceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaceGetParams struct {
	// Comma-separated list of opt-in expansion fields. See `list_faces` for supported
	// values.
	Include param.Opt[string] `query:"include,omitzero" json:"-"`
	// Library the face belongs to. Optional if the user has a single library; required
	// when they have multiple.
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
	// Library the face belongs to. Optional if the user has a single library; required
	// when they have multiple.
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Target person ID (with `person_` prefix) to assign this face to. Pass `null` to
	// detach the face from its current person without deleting either. Get IDs from
	// `list_people`; use `create_person` first if the target identity doesn't exist
	// yet.
	PersonID param.Opt[string] `json:"person_id,omitzero"`
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
	// Return only faces detected in this asset. Useful for 'show me all the faces in
	// this photo'.
	AssetID param.Opt[string] `query:"asset_id,omitzero" json:"-"`
	// Comma-separated list of opt-in expansion fields. Supported values:
	// `cluster_assignment` (adds the nested `cluster_assignment` object —
	// `distance_to_person` and a top-K `candidates` list of nearby Persons).
	Include param.Opt[string] `query:"include,omitzero" json:"-"`
	// Library to list from. Optional if the user has a single library; required when
	// they have multiple.
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Return only faces currently assigned to this person. Useful for reviewing or
	// curating a person's face cluster.
	PersonID param.Opt[string] `query:"person_id,omitzero" json:"-"`
	// Cursor for pagination. Pass the `id` of the last face in the previous response's
	// `data` to fetch the next page. Omit for the first page.
	StartingAfterID param.Opt[string] `query:"starting_after_id,omitzero" json:"-"`
	// Maximum number of faces per page (1–200). Defaults to 20.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Look up specific faces by ID (max 100). IDs use the `face_` prefix.
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
	// Library the face belongs to. Optional if the user has a single library; required
	// when they have multiple.
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
