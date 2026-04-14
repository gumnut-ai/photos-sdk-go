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

	"github.com/gumnut-ai/photos-sdk-go/internal/apiform"
	"github.com/gumnut-ai/photos-sdk-go/internal/apijson"
	"github.com/gumnut-ai/photos-sdk-go/internal/apiquery"
	"github.com/gumnut-ai/photos-sdk-go/internal/requestconfig"
	"github.com/gumnut-ai/photos-sdk-go/option"
	"github.com/gumnut-ai/photos-sdk-go/packages/pagination"
	"github.com/gumnut-ai/photos-sdk-go/packages/param"
	"github.com/gumnut-ai/photos-sdk-go/packages/respjson"
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
	return res, err
}

// Retrieves detailed metadata for a specific asset, including EXIF information,
// asset metrics, faces, and people.
func (r *AssetService) Get(ctx context.Context, assetID string, opts ...option.RequestOption) (res *AssetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if assetID == "" {
		err = errors.New("missing required asset_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/assets/%s", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
		return err
	}
	path := fmt.Sprintf("api/assets/%s", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Checks which assets exist in the user's library based on checksums or device
// identifiers. Provide exactly one of: checksums, checksum_sha1s, or (deviceId AND
// deviceAssetIds). List parameters are limited to 5000 items.
func (r *AssetService) CheckExistence(ctx context.Context, params AssetCheckExistenceParams, opts ...option.RequestOption) (res *AssetExistenceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/assets/exist"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns asset counts grouped by time period. Supports optional filtering by
// album, person, or date range. Results are ordered by time bucket descending.
func (r *AssetService) Counts(ctx context.Context, query AssetCountsParams, opts ...option.RequestOption) (res *AssetCountResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/assets/counts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AssetCountResponse struct {
	Data    []AssetCountResponseData `json:"data" api:"required"`
	HasMore bool                     `json:"has_more" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetCountResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetCountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetCountResponseData struct {
	// Number of assets in this time period
	Count int64 `json:"count" api:"required"`
	// Start of the time period
	TimeBucket time.Time `json:"time_bucket" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		TimeBucket  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetCountResponseData) RawJSON() string { return r.JSON.raw }
func (r *AssetCountResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for asset existence check endpoint.
type AssetExistenceResponse struct {
	// List of assets matching the query criteria
	Assets []AssetLiteResponse `json:"assets" api:"required"`
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
	ID string `json:"id" api:"required"`
	// Base64-encoded SHA-256 hash of the asset contents for duplicate detection and
	// integrity
	Checksum string `json:"checksum" api:"required"`
	// Original asset identifier from the device that uploaded this asset
	DeviceAssetID string `json:"device_asset_id" api:"required"`
	// Identifier of the device that uploaded this asset
	DeviceID string `json:"device_id" api:"required"`
	// Base64-encoded SHA-1 hash for Immich client compatibility. May be null for older
	// assets.
	ChecksumSha1 string `json:"checksum_sha1" api:"nullable"`
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
	ID string `json:"id" api:"required"`
	// Base64-encoded SHA-256 hash of the asset contents for duplicate detection and
	// integrity
	Checksum string `json:"checksum" api:"required"`
	// When this asset record was created in the database
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Original asset identifier from the device that uploaded this asset
	DeviceAssetID string `json:"device_asset_id" api:"required"`
	// Identifier of the device that uploaded this asset
	DeviceID string `json:"device_id" api:"required"`
	// When the file was created on the uploading device
	FileCreatedAt time.Time `json:"file_created_at" api:"required" format:"date-time"`
	// When the file was last modified on the uploading device
	FileModifiedAt time.Time `json:"file_modified_at" api:"required" format:"date-time"`
	// When the photo/video was taken, in the device's local timezone
	LocalDatetime time.Time `json:"local_datetime" api:"required" format:"date-time"`
	// MIME type of the file (e.g., 'image/jpeg', 'video/mp4')
	MimeType string `json:"mime_type" api:"required"`
	// Original filename when the asset was uploaded
	OriginalFileName string `json:"original_file_name" api:"required"`
	// When this asset record was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Named asset variants: 'original', 'thumbnail', 'preview', 'fullsize' for images;
	// 'original' only for videos
	AssetURLs map[string]AssetResponseAssetURL `json:"asset_urls" api:"nullable"`
	// Base64-encoded SHA-1 hash for Immich client compatibility. May be null for older
	// assets.
	ChecksumSha1 string `json:"checksum_sha1" api:"nullable"`
	// AI-generated description of the asset's content, quality, and composition. null
	// means description generation has not yet run; empty string means the model
	// refused to describe the asset. Distinct from exif.description (camera-embedded
	// EXIF metadata).
	Description string `json:"description" api:"nullable"`
	// EXIF metadata extracted from image and video files.
	Exif ExifResponse `json:"exif" api:"nullable"`
	// All faces detected in this asset
	Faces []FaceResponse `json:"faces"`
	// File size of the asset in bytes
	FileSizeBytes int64 `json:"file_size_bytes"`
	// Height of the asset in pixels
	Height int64 `json:"height"`
	// ML-generated quality scores and other metrics
	Metrics map[string]float64 `json:"metrics" api:"nullable"`
	// All unique people identified in this asset (deduplicated from faces)
	People []PersonResponse `json:"people"`
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
		AssetURLs        respjson.Field
		ChecksumSha1     respjson.Field
		Description      respjson.Field
		Exif             respjson.Field
		Faces            respjson.Field
		FileSizeBytes    respjson.Field
		Height           respjson.Field
		Metrics          respjson.Field
		People           respjson.Field
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

// A single image variant with its URL, MIME type, and target width.
type AssetResponseAssetURL struct {
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
func (r AssetResponseAssetURL) RawJSON() string { return r.JSON.raw }
func (r *AssetResponseAssetURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetNewParams struct {
	// The asset file to upload
	AssetData      io.Reader `json:"asset_data,omitzero" api:"required" format:"binary"`
	DeviceAssetID  string    `json:"device_asset_id" api:"required"`
	DeviceID       string    `json:"device_id" api:"required"`
	FileCreatedAt  time.Time `json:"file_created_at" api:"required" format:"date-time"`
	FileModifiedAt time.Time `json:"file_modified_at" api:"required" format:"date-time"`
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
	// Only include assets with local_datetime after this value (ISO 8601). Naive
	// values compare directly against local_datetime; timezone-aware values are
	// converted to UTC and compared against local_datetime adjusted by its stored
	// offset.
	LocalDatetimeAfter param.Opt[time.Time] `query:"local_datetime_after,omitzero" format:"date-time" json:"-"`
	// Only include assets with local_datetime before this value (ISO 8601). Naive
	// values compare directly against local_datetime; timezone-aware values are
	// converted to UTC and compared against local_datetime adjusted by its stored
	// offset.
	LocalDatetimeBefore param.Opt[time.Time] `query:"local_datetime_before,omitzero" format:"date-time" json:"-"`
	// Filter by assets associated with a specific person ID
	PersonID param.Opt[string] `query:"person_id,omitzero" json:"-"`
	// Asset ID to start listing assets after
	StartingAfterID param.Opt[string] `query:"starting_after_id,omitzero" json:"-"`
	// Max number of assets to return (1-200)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
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

type AssetCountsParams struct {
	// Filter by assets in a specific album
	AlbumID param.Opt[string] `query:"album_id,omitzero" json:"-"`
	// Library to count assets in (optional)
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Only include assets with local_datetime after this value (ISO 8601). Naive
	// values compare directly against local_datetime; timezone-aware values are
	// converted to UTC and compared against local_datetime adjusted by its stored
	// offset.
	LocalDatetimeAfter param.Opt[time.Time] `query:"local_datetime_after,omitzero" format:"date-time" json:"-"`
	// Only include assets with local_datetime before this value (ISO 8601). Naive
	// values compare directly against local_datetime; timezone-aware values are
	// converted to UTC and compared against local_datetime adjusted by its stored
	// offset. Use the last time_bucket from a previous response to paginate.
	LocalDatetimeBefore param.Opt[time.Time] `query:"local_datetime_before,omitzero" format:"date-time" json:"-"`
	// Filter by assets associated with a specific person ID
	PersonID param.Opt[string] `query:"person_id,omitzero" json:"-"`
	// Time period to group counts by. Currently only 'month' is supported.
	GroupBy param.Opt[string] `query:"group_by,omitzero" json:"-"`
	// Maximum number of time buckets to return (1-200)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssetCountsParams]'s query parameters as `url.Values`.
func (r AssetCountsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
