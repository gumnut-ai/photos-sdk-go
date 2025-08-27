// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/photos-go/internal/apijson"
	shimjson "github.com/stainless-sdks/photos-go/internal/encoding/json"
	"github.com/stainless-sdks/photos-go/internal/requestconfig"
	"github.com/stainless-sdks/photos-go/option"
	"github.com/stainless-sdks/photos-go/packages/param"
	"github.com/stainless-sdks/photos-go/packages/respjson"
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

// Retrieves a list of all assets contained within a specific album, along with
// their associated metrics, EXIF data, faces, and people.
func (r *AlbumAssetService) List(ctx context.Context, albumID string, opts ...option.RequestOption) (res *[]AssetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if albumID == "" {
		err = errors.New("missing required album_id parameter")
		return
	}
	path := fmt.Sprintf("api/albums/%s/assets", albumID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Adds one or more existing assets to a specific album. Assets must be in the same
// library as the album. Duplicate assets are ignored.
func (r *AlbumAssetService) Add(ctx context.Context, albumID string, body AlbumAssetAddParams, opts ...option.RequestOption) (res *AlbumAssetAddResponse, err error) {
	opts = append(r.Options[:], opts...)
	if albumID == "" {
		err = errors.New("missing required album_id parameter")
		return
	}
	path := fmt.Sprintf("api/albums/%s/assets", albumID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes one or more assets from a specific album. Note: This does not delete the
// assets themselves.
func (r *AlbumAssetService) Remove(ctx context.Context, albumID string, body AlbumAssetRemoveParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if albumID == "" {
		err = errors.New("missing required album_id parameter")
		return
	}
	path := fmt.Sprintf("api/albums/%s/assets", albumID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// The property AssetIDs is required.
type AlbumAssetAssociationParam struct {
	AssetIDs []string `json:"asset_ids,omitzero,required"`
	paramObj
}

func (r AlbumAssetAssociationParam) MarshalJSON() (data []byte, err error) {
	type shadow AlbumAssetAssociationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlbumAssetAssociationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlbumAssetAddResponse struct {
	AddedAssets     []string `json:"added_assets,required"`
	DuplicateAssets []string `json:"duplicate_assets,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddedAssets     respjson.Field
		DuplicateAssets respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlbumAssetAddResponse) RawJSON() string { return r.JSON.raw }
func (r *AlbumAssetAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlbumAssetAddParams struct {
	AlbumAssetAssociation AlbumAssetAssociationParam
	paramObj
}

func (r AlbumAssetAddParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AlbumAssetAssociation)
}
func (r *AlbumAssetAddParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AlbumAssetAssociation)
}

type AlbumAssetRemoveParams struct {
	AlbumAssetAssociation AlbumAssetAssociationParam
	paramObj
}

func (r AlbumAssetRemoveParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AlbumAssetAssociation)
}
func (r *AlbumAssetRemoveParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AlbumAssetAssociation)
}
