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

// Creates a new person entry.
func (r *PersonService) New(ctx context.Context, body PersonNewParams, opts ...option.RequestOption) (res *PersonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/people"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves details for a specific person.
func (r *PersonService) Get(ctx context.Context, personID string, opts ...option.RequestOption) (res *PersonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/people/%s", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the details of a specific person.
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

// Retrieves a paginated list of people, ordered by creation time, descending. Can
// be filtered by specific person IDs.
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

// Retrieves a paginated list of people, ordered by creation time, descending. Can
// be filtered by specific person IDs.
func (r *PersonService) ListAutoPaging(ctx context.Context, query PersonListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[PersonResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a specific person. Associated faces will have their person_id set to the
// closest matching person, or null if no one matches.
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
	// Optional birth date of this person
	BirthDate time.Time `json:"birth_date" api:"nullable" format:"date"`
	// Optional name assigned to this person
	Name string `json:"name" api:"nullable"`
	// ID of the face resource used as this person's thumbnail
	ThumbnailFaceID string `json:"thumbnail_face_id" api:"nullable"`
	// URL for this person's profile thumbnail image
	ThumbnailFaceURL string `json:"thumbnail_face_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		IsFavorite       respjson.Field
		IsHidden         respjson.Field
		UpdatedAt        respjson.Field
		AssetCount       respjson.Field
		BirthDate        respjson.Field
		Name             respjson.Field
		ThumbnailFaceID  respjson.Field
		ThumbnailFaceURL respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonNewParams struct {
	BirthDate       param.Opt[time.Time] `json:"birth_date,omitzero" format:"date"`
	IsFavorite      param.Opt[bool]      `json:"is_favorite,omitzero"`
	IsHidden        param.Opt[bool]      `json:"is_hidden,omitzero"`
	LibraryID       param.Opt[string]    `json:"library_id,omitzero"`
	Name            param.Opt[string]    `json:"name,omitzero"`
	ThumbnailFaceID param.Opt[string]    `json:"thumbnail_face_id,omitzero"`
	paramObj
}

func (r PersonNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonUpdateParams struct {
	BirthDate       param.Opt[time.Time] `json:"birth_date,omitzero" format:"date"`
	IsFavorite      param.Opt[bool]      `json:"is_favorite,omitzero"`
	IsHidden        param.Opt[bool]      `json:"is_hidden,omitzero"`
	Name            param.Opt[string]    `json:"name,omitzero"`
	ThumbnailFaceID param.Opt[string]    `json:"thumbnail_face_id,omitzero"`
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
	// Include only people associated with this album ID
	AlbumID param.Opt[string] `query:"album_id,omitzero" json:"-"`
	// Include only people associated with this asset ID
	AssetID param.Opt[string] `query:"asset_id,omitzero" json:"-"`
	// Library ID (required if user has multiple libraries)
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Person ID to start listing people after
	StartingAfterID param.Opt[string] `query:"starting_after_id,omitzero" json:"-"`
	// Max number of people to return (1-200)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by specific person IDs (max 100)
	IDs []string `query:"ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PersonListParams]'s query parameters as `url.Values`.
func (r PersonListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
