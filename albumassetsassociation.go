// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/gumnut-ai/photos-sdk-go/internal/apijson"
	shimjson "github.com/gumnut-ai/photos-sdk-go/internal/encoding/json"
	"github.com/gumnut-ai/photos-sdk-go/internal/requestconfig"
	"github.com/gumnut-ai/photos-sdk-go/option"
	"github.com/gumnut-ai/photos-sdk-go/packages/param"
	"github.com/gumnut-ai/photos-sdk-go/packages/respjson"
)

// AlbumAssetsAssociationService contains methods and other services that help with
// interacting with the Gumnut API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlbumAssetsAssociationService] method instead.
type AlbumAssetsAssociationService struct {
	Options []option.RequestOption
}

// NewAlbumAssetsAssociationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAlbumAssetsAssociationService(opts ...option.RequestOption) (r AlbumAssetsAssociationService) {
	r = AlbumAssetsAssociationService{}
	r.Options = opts
	return
}

// Adds one or more existing assets to the specified album. Assets must already be
// in the same library as the album (this tool does not upload new assets). Assets
// already in the album are silently skipped and returned separately as
// `duplicate_assets`. Idempotent: calling with the same IDs twice leaves the album
// in the same state.
func (r *AlbumAssetsAssociationService) Add(ctx context.Context, albumID string, body AlbumAssetsAssociationAddParams, opts ...option.RequestOption) (res *AlbumAssetsAssociationAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if albumID == "" {
		err = errors.New("missing required album_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/albums/%s/assets", albumID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Detaches one or more assets from the given album. The assets remain in the
// library and in any other albums they belong to. Use `delete_asset` to delete the
// asset entirely. To empty an album completely, call `list_album_assets` to get
// the links and then remove them, or delete the album itself with `delete_album`.
func (r *AlbumAssetsAssociationService) Remove(ctx context.Context, albumID string, body AlbumAssetsAssociationRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if albumID == "" {
		err = errors.New("missing required album_id parameter")
		return err
	}
	path := fmt.Sprintf("api/albums/%s/assets", albumID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// The property AssetIDs is required.
type AlbumAssetAssociationParam struct {
	// Asset IDs (with `asset_` prefix) to associate with the album. Get IDs from
	// `list_assets`, `search_assets`, or `list_album_assets`.
	AssetIDs []string `json:"asset_ids,omitzero" api:"required"`
	paramObj
}

func (r AlbumAssetAssociationParam) MarshalJSON() (data []byte, err error) {
	type shadow AlbumAssetAssociationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlbumAssetAssociationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlbumAssetsAssociationAddResponse struct {
	// Asset IDs newly added to the album by this call.
	AddedAssets []string `json:"added_assets" api:"required"`
	// Asset IDs that were already in the album and were skipped (idempotent no-op, not
	// an error).
	DuplicateAssets []string `json:"duplicate_assets" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddedAssets     respjson.Field
		DuplicateAssets respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlbumAssetsAssociationAddResponse) RawJSON() string { return r.JSON.raw }
func (r *AlbumAssetsAssociationAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlbumAssetsAssociationAddParams struct {
	AlbumAssetAssociation AlbumAssetAssociationParam
	paramObj
}

func (r AlbumAssetsAssociationAddParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AlbumAssetAssociation)
}
func (r *AlbumAssetsAssociationAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlbumAssetsAssociationRemoveParams struct {
	AlbumAssetAssociation AlbumAssetAssociationParam
	paramObj
}

func (r AlbumAssetsAssociationRemoveParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AlbumAssetAssociation)
}
func (r *AlbumAssetsAssociationRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
