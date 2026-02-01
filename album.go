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

	"github.com/stainless-sdks/photos-go/internal/apijson"
	"github.com/stainless-sdks/photos-go/internal/apiquery"
	"github.com/stainless-sdks/photos-go/internal/requestconfig"
	"github.com/stainless-sdks/photos-go/option"
	"github.com/stainless-sdks/photos-go/packages/pagination"
	"github.com/stainless-sdks/photos-go/packages/param"
	"github.com/stainless-sdks/photos-go/packages/respjson"
)

// AlbumService contains methods and other services that help with interacting with
// the Gumnut API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlbumService] method instead.
type AlbumService struct {
	Options []option.RequestOption
	Assets  AlbumAssetService
}

// NewAlbumService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAlbumService(opts ...option.RequestOption) (r AlbumService) {
	r = AlbumService{}
	r.Options = opts
	r.Assets = NewAlbumAssetService(opts...)
	return
}

// Creates a new, empty album with optional name and description in the specified
// library.
func (r *AlbumService) New(ctx context.Context, body AlbumNewParams, opts ...option.RequestOption) (res *AlbumResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/albums"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves details for a specific album.
func (r *AlbumService) Get(ctx context.Context, albumID string, opts ...option.RequestOption) (res *AlbumResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if albumID == "" {
		err = errors.New("missing required album_id parameter")
		return
	}
	path := fmt.Sprintf("api/albums/%s", albumID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the name and/or description of a specific album.
func (r *AlbumService) Update(ctx context.Context, albumID string, body AlbumUpdateParams, opts ...option.RequestOption) (res *AlbumResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if albumID == "" {
		err = errors.New("missing required album_id parameter")
		return
	}
	path := fmt.Sprintf("api/albums/%s", albumID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieves a paginated list of albums from the specified library, ordered by
// creation time, descending. Can be filtered by asset_id.
func (r *AlbumService) List(ctx context.Context, query AlbumListParams, opts ...option.RequestOption) (res *pagination.CursorPage[AlbumResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/albums"
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

// Retrieves a paginated list of albums from the specified library, ordered by
// creation time, descending. Can be filtered by asset_id.
func (r *AlbumService) ListAutoPaging(ctx context.Context, query AlbumListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[AlbumResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a specific album. Note: This does not delete the assets within the
// album.
func (r *AlbumService) Delete(ctx context.Context, albumID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if albumID == "" {
		err = errors.New("missing required album_id parameter")
		return
	}
	path := fmt.Sprintf("api/albums/%s", albumID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Represents a collection of assets organized by the user.
type AlbumResponse struct {
	// Unique album identifier with 'album\_' prefix
	ID string `json:"id,required"`
	// Total number of assets in this album
	AssetCount int64 `json:"asset_count,required"`
	// When this album was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Display name of the album
	Name string `json:"name,required"`
	// When this album was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// ID of the asset used as the album cover
	AlbumCoverAssetID string `json:"album_cover_asset_id,nullable"`
	// URL to get the album cover thumbnail image
	AlbumCoverThumbnailURL string `json:"album_cover_thumbnail_url,nullable"`
	// Optional description text for the album
	Description string `json:"description,nullable"`
	// The newest asset date (local_datetime) in the album, or null if empty
	EndDate time.Time `json:"end_date,nullable" format:"date-time"`
	// The oldest asset date (local_datetime) in the album, or null if empty
	StartDate time.Time `json:"start_date,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AssetCount             respjson.Field
		CreatedAt              respjson.Field
		Name                   respjson.Field
		UpdatedAt              respjson.Field
		AlbumCoverAssetID      respjson.Field
		AlbumCoverThumbnailURL respjson.Field
		Description            respjson.Field
		EndDate                respjson.Field
		StartDate              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlbumResponse) RawJSON() string { return r.JSON.raw }
func (r *AlbumResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlbumNewParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	LibraryID   param.Opt[string] `json:"library_id,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AlbumNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AlbumNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlbumNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlbumUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AlbumUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AlbumUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlbumUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlbumListParams struct {
	// Filter albums containing this asset ID (optional)
	AssetID param.Opt[string] `query:"asset_id,omitzero" json:"-"`
	// Library to list albums from (optional)
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Album ID to start listing albums after
	StartingAfterID param.Opt[string] `query:"starting_after_id,omitzero" json:"-"`
	// Max number of albums to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AlbumListParams]'s query parameters as `url.Values`.
func (r AlbumListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
