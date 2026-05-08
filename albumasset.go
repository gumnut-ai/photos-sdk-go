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

// AlbumAssetService contains methods and other services that help with interacting
// with the Gumnut API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlbumAssetService] method instead.
type AlbumAssetService struct {
	Options []option.RequestOption
}

// NewAlbumAssetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlbumAssetService(opts ...option.RequestOption) (r AlbumAssetService) {
	r = AlbumAssetService{}
	r.Options = opts
	return
}

// Returns paginated _link_ records (lightweight join rows between albums and
// assets) describing which assets are in which albums — each row contains
// `album_id` + `asset_id` + link timestamps, not the full asset or album metadata.
// Use this when you specifically need the junction records (for sync or change
// tracking).
//
// **For most use cases you want a different tool:** use `list_assets` with
// `album_id` to get the full asset metadata for a specific album; use
// `list_albums` with `asset_id` to find which albums contain an asset.
//
// **Pagination** is cursor-based: when `has_more` is true, pass the `id` of the
// last album-asset in `data` as `starting_after_id` to fetch the next page.
func (r *AlbumAssetService) List(ctx context.Context, query AlbumAssetListParams, opts ...option.RequestOption) (res *pagination.CursorPage[AlbumAssetResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/album-assets"
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

// Returns paginated _link_ records (lightweight join rows between albums and
// assets) describing which assets are in which albums — each row contains
// `album_id` + `asset_id` + link timestamps, not the full asset or album metadata.
// Use this when you specifically need the junction records (for sync or change
// tracking).
//
// **For most use cases you want a different tool:** use `list_assets` with
// `album_id` to get the full asset metadata for a specific album; use
// `list_albums` with `asset_id` to find which albums contain an asset.
//
// **Pagination** is cursor-based: when `has_more` is true, pass the `id` of the
// last album-asset in `data` as `starting_after_id` to fetch the next page.
func (r *AlbumAssetService) ListAutoPaging(ctx context.Context, query AlbumAssetListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[AlbumAssetResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Fetches one album-asset link record by ID (the junction row between an album and
// an asset). Rarely needed directly; most callers want `get_asset` or `get_album`
// instead.
func (r *AlbumAssetService) Get(ctx context.Context, albumAssetID string, opts ...option.RequestOption) (res *AlbumAssetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if albumAssetID == "" {
		err = errors.New("missing required album_asset_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/album-assets/%s", albumAssetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Represents a link between an album and an asset.
type AlbumAssetResponse struct {
	// Unique album*asset identifier with 'album_asset*' prefix
	ID string `json:"id" api:"required"`
	// ID of the album
	AlbumID string `json:"album_id" api:"required"`
	// ID of the asset
	AssetID string `json:"asset_id" api:"required"`
	// When this link was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When this link was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AlbumID     respjson.Field
		AssetID     respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlbumAssetResponse) RawJSON() string { return r.JSON.raw }
func (r *AlbumAssetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlbumAssetListParams struct {
	// Return only link records for this album ID. Equivalent to 'list the assets in
	// this album' — in most cases prefer `list_assets` with `album_id` to get the
	// asset metadata directly instead of the lightweight link records.
	AlbumID param.Opt[string] `query:"album_id,omitzero" json:"-"`
	// Return only link records for this asset ID. Equivalent to 'which albums contain
	// this asset' — in most cases prefer `list_albums` with `asset_id` to get the
	// album metadata directly.
	AssetID param.Opt[string] `query:"asset_id,omitzero" json:"-"`
	// Library to list from. Optional if the user has a single library; required when
	// they have multiple.
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Cursor for pagination. Pass the `id` of the last album-asset in the previous
	// response's `data` to fetch the next page. Omit for the first page.
	StartingAfterID param.Opt[string] `query:"starting_after_id,omitzero" json:"-"`
	// Maximum number of link records per page (1–200). Defaults to 20.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Look up specific album-asset link records by ID (max 100). The ID has the
	// `album_asset_` prefix.
	IDs []string `query:"ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AlbumAssetListParams]'s query parameters as `url.Values`.
func (r AlbumAssetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
