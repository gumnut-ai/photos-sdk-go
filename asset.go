// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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
	"github.com/stainless-sdks/photos-go/packages/pagination"
	"github.com/stainless-sdks/photos-go/packages/param"
	"github.com/stainless-sdks/photos-go/packages/respjson"
)

// AssetService contains methods and other services that help with interacting with
// the Gumnut API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetService] method instead.
type AssetService struct {
	Options []option.RequestOption
}

// NewAssetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAssetService(opts ...option.RequestOption) (r AssetService) {
	r = AssetService{}
	r.Options = opts
	return
}

// Uploads a new asset file (image or video) along with its metadata to the
// specified library. If no library_id is provided and the user only has one
// library, uses that library. If the user has multiple libraries, library_id is
// required.
func (r *AssetService) New(ctx context.Context, body AssetNewParams, opts ...option.RequestOption) (res *AssetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/assets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves detailed metadata for a specific asset, including EXIF information,
// asset metrics, faces, and people.
func (r *AssetService) Get(ctx context.Context, assetID string, opts ...option.RequestOption) (res *AssetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if assetID == "" {
		err = errors.New("missing required asset_id parameter")
		return
	}
	path := fmt.Sprintf("api/assets/%s", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves a paginated list of assets from the specified library, optionally
// filtered by album, person, or specific asset IDs. Asset data includes metrics,
// EXIF data, faces, and people. Assets are ordered by local creation time,
// descending.
func (r *AssetService) List(ctx context.Context, query AssetListParams, opts ...option.RequestOption) (res *pagination.CursorPage[AssetResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/assets"
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

// Retrieves a paginated list of assets from the specified library, optionally
// filtered by album, person, or specific asset IDs. Asset data includes metrics,
// EXIF data, faces, and people. Assets are ordered by local creation time,
// descending.
func (r *AssetService) ListAutoPaging(ctx context.Context, query AssetListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[AssetResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a specific asset and its associated data (including the file from
// storage).
func (r *AssetService) Delete(ctx context.Context, assetID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if assetID == "" {
		err = errors.New("missing required asset_id parameter")
		return
	}
	path := fmt.Sprintf("api/assets/%s", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Checks which assets exist in the user's library based on checksums or device
// identifiers. Provide exactly one of: checksums, checksum_sha1s, or (deviceId AND
// deviceAssetIds). List parameters are limited to 5000 items.
func (r *AssetService) CheckExistence(ctx context.Context, params AssetCheckExistenceParams, opts ...option.RequestOption) (res *AssetExistenceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/assets/exist"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Downloads the original file for a specific asset.
func (r *AssetService) Download(ctx context.Context, assetID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/*")}, opts...)
	if assetID == "" {
		err = errors.New("missing required asset_id parameter")
		return
	}
	path := fmt.Sprintf("api/assets/%s/download", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Downloads a thumbnail for a specific asset. The exact thumbnail returned depends
// on availability and the optional `size` parameter.
func (r *AssetService) DownloadThumbnail(ctx context.Context, assetID string, query AssetDownloadThumbnailParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/*")}, opts...)
	if assetID == "" {
		err = errors.New("missing required asset_id parameter")
		return
	}
	path := fmt.Sprintf("api/assets/%s/thumbnail", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response for asset existence check endpoint.
type AssetExistenceResponse struct {
	// List of assets matching the query criteria
	Assets []AssetLiteResponse `json:"assets,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assets      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetExistenceResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetExistenceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lightweight asset response for existence checks.
type AssetLiteResponse struct {
	// Unique asset identifier with 'asset\_' prefix
	ID string `json:"id,required"`
	// Base64-encoded SHA-256 hash of the asset contents for duplicate detection and
	// integrity
	Checksum string `json:"checksum,required"`
	// Original asset identifier from the device that uploaded this asset
	DeviceAssetID string `json:"device_asset_id,required"`
	// Identifier of the device that uploaded this asset
	DeviceID string `json:"device_id,required"`
	// Base64-encoded SHA-1 hash for Immich client compatibility. May be null for older
	// assets.
	ChecksumSha1 string `json:"checksum_sha1,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Checksum      respjson.Field
		DeviceAssetID respjson.Field
		DeviceID      respjson.Field
		ChecksumSha1  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetLiteResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetLiteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a photo or video asset with metadata and access URLs.
type AssetResponse struct {
	// Unique asset identifier with 'asset\_' prefix
	ID string `json:"id,required"`
	// Base64-encoded SHA-256 hash of the asset contents for duplicate detection and
	// integrity
	Checksum string `json:"checksum,required"`
	// When this asset record was created in the database
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Original asset identifier from the device that uploaded this asset
	DeviceAssetID string `json:"device_asset_id,required"`
	// Identifier of the device that uploaded this asset
	DeviceID string `json:"device_id,required"`
	// When the file was created on the uploading device
	FileCreatedAt time.Time `json:"file_created_at,required" format:"date-time"`
	// When the file was last modified on the uploading device
	FileModifiedAt time.Time `json:"file_modified_at,required" format:"date-time"`
	// When the photo/video was taken, in the device's local timezone
	LocalDatetime time.Time `json:"local_datetime,required" format:"date-time"`
	// MIME type of the file (e.g., 'image/jpeg', 'video/mp4')
	MimeType string `json:"mime_type,required"`
	// Original filename when the asset was uploaded
	OriginalFileName string `json:"original_file_name,required"`
	// When this asset record was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Base64-encoded SHA-1 hash for Immich client compatibility. May be null for older
	// assets.
	ChecksumSha1 string `json:"checksum_sha1,nullable"`
	// If you need to download the full asset, use this URL. Otherwise, use the
	// thumbnail_url.
	DownloadURL string `json:"download_url,nullable"`
	// EXIF metadata extracted from image and video files.
	Exif ExifResponse `json:"exif,nullable"`
	// All faces detected in this asset
	Faces []FaceResponse `json:"faces"`
	// File size of the asset in bytes
	FileSizeBytes int64 `json:"file_size_bytes"`
	// Height of the asset in pixels
	Height int64 `json:"height"`
	// ML-generated quality scores and other metrics
	Metrics map[string]float64 `json:"metrics,nullable"`
	// All unique people identified in this asset (deduplicated from faces)
	People []PersonResponse `json:"people"`
	// Use this URL to display the asset. Never download the full asset unless you
	// absolutely have to; prefer the thumbnail instead.
	ThumbnailURL string `json:"thumbnail_url,nullable"`
	// Width of the asset in pixels
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Checksum         respjson.Field
		CreatedAt        respjson.Field
		DeviceAssetID    respjson.Field
		DeviceID         respjson.Field
		FileCreatedAt    respjson.Field
		FileModifiedAt   respjson.Field
		LocalDatetime    respjson.Field
		MimeType         respjson.Field
		OriginalFileName respjson.Field
		UpdatedAt        respjson.Field
		ChecksumSha1     respjson.Field
		DownloadURL      respjson.Field
		Exif             respjson.Field
		Faces            respjson.Field
		FileSizeBytes    respjson.Field
		Height           respjson.Field
		Metrics          respjson.Field
		People           respjson.Field
		ThumbnailURL     respjson.Field
		Width            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetNewParams struct {
	AssetData      io.Reader `json:"asset_data,omitzero,required" format:"binary"`
	DeviceAssetID  string    `json:"device_asset_id,required"`
	DeviceID       string    `json:"device_id,required"`
	FileCreatedAt  time.Time `json:"file_created_at,required" format:"date-time"`
	FileModifiedAt time.Time `json:"file_modified_at,required" format:"date-time"`
	// Library to upload asset to (optional)
	LibraryID param.Opt[string] `json:"library_id,omitzero"`
	paramObj
}

func (r AssetNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type AssetListParams struct {
	// Filter by assets in a specific album
	AlbumID param.Opt[string] `query:"album_id,omitzero" json:"-"`
	// Library to list assets from (optional)
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Filter by assets associated with a specific person ID
	PersonID param.Opt[string] `query:"person_id,omitzero" json:"-"`
	// Asset ID to start listing assets after
	StartingAfterID param.Opt[string] `query:"starting_after_id,omitzero" json:"-"`
	Limit           param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Filter by specific asset IDs (max 100)
	IDs []string `query:"ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssetListParams]'s query parameters as `url.Values`.
func (r AssetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AssetCheckExistenceParams struct {
	// Library to check assets in (optional)
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Device ID to filter assets by (required with deviceAssetIds)
	DeviceID param.Opt[string] `json:"deviceId,omitzero"`
	// List of base64-encoded SHA-1 checksums to check for existence (for Immich
	// compatibility)
	ChecksumSha1s []string `json:"checksum_sha1s,omitzero"`
	// List of base64-encoded SHA-256 checksums to check for existence
	Checksums []string `json:"checksums,omitzero"`
	// List of device asset IDs to check for existence (requires deviceId)
	DeviceAssetIDs []string `json:"deviceAssetIds,omitzero"`
	paramObj
}

func (r AssetCheckExistenceParams) MarshalJSON() (data []byte, err error) {
	type shadow AssetCheckExistenceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetCheckExistenceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [AssetCheckExistenceParams]'s query parameters as
// `url.Values`.
func (r AssetCheckExistenceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AssetDownloadThumbnailParams struct {
	// Desired thumbnail size (e.g., thumbnail, preview)
	Size param.Opt[string] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssetDownloadThumbnailParams]'s query parameters as
// `url.Values`.
func (r AssetDownloadThumbnailParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
