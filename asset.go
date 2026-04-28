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
	"github.com/gumnut-ai/photos-sdk-go/shared"
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

// Fetches one asset and its associated metadata. Use this when you already have a
// specific asset ID (e.g., from `list_assets`, `search_assets`, or
// `list_album_assets`) and need its full details. For bulk fetch of multiple known
// IDs, prefer `list_assets` with the `ids` parameter to avoid N round trips.
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

// Returns a paginated list of assets ordered by local capture time (newest first).
// Use this tool for structured browsing and filtering — when the request can be
// expressed as exact filters on album membership, people, date range, or specific
// asset IDs.
//
// **Use `search_assets` instead** when the request involves natural-language image
// content ('photos of sunsets', 'pictures with my dog'), location or place
// ('photos from Japan'), or any concept requiring semantic understanding of what's
// in the image. `list_assets` does not filter by image content, location, or
// caption text.
//
// **To present a curated set of specific assets to the user** (e.g., a hand-picked
// subset of `search_assets` results), call this tool with `ids=[...]` rather than
// building a custom gallery — the asset IDs you already have are enough to
// re-render them through the interactive widget.
//
// **Pagination** is cursor-based: when `has_more` is true, pass the `id` of the
// last asset in `data` as `starting_after_id` to fetch the next page.
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

// Returns a paginated list of assets ordered by local capture time (newest first).
// Use this tool for structured browsing and filtering — when the request can be
// expressed as exact filters on album membership, people, date range, or specific
// asset IDs.
//
// **Use `search_assets` instead** when the request involves natural-language image
// content ('photos of sunsets', 'pictures with my dog'), location or place
// ('photos from Japan'), or any concept requiring semantic understanding of what's
// in the image. `list_assets` does not filter by image content, location, or
// caption text.
//
// **To present a curated set of specific assets to the user** (e.g., a hand-picked
// subset of `search_assets` results), call this tool with `ids=[...]` rather than
// building a custom gallery — the asset IDs you already have are enough to
// re-render them through the interactive widget.
//
// **Pagination** is cursor-based: when `has_more` is true, pass the `id` of the
// last asset in `data` as `starting_after_id` to fetch the next page.
func (r *AssetService) ListAutoPaging(ctx context.Context, query AssetListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[AssetResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes the asset entirely — the database record, the stored file, and all
// associated data (faces, album links, etc.). This is irreversible.
//
// **Use `remove_assets_from_album` instead** when the user only wants to remove an
// asset from a specific album but keep the file in their library. Use
// `delete_album` to remove an album without deleting its assets.
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
//
// **Pagination:** When `has_more` is true, pass the last `time_bucket` value from
// `data` as `local_datetime_before` to fetch the next page.
func (r *AssetService) Counts(ctx context.Context, query AssetCountsParams, opts ...option.RequestOption) (res *AssetCountResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/assets/counts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AssetCountResponse struct {
	// Time bucket and count pairs, ordered by time bucket descending
	Data []AssetCountResponseData `json:"data" api:"required"`
	// True if there are more time buckets. To fetch the next page, pass the last
	// `time_bucket` value as `local_datetime_before` (exclusive — buckets starting
	// before that value are returned).
	HasMore bool `json:"has_more" api:"required"`
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
	AssetURLs map[string]shared.AssetVariant `json:"asset_urls" api:"nullable"`
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
	// Metadata for an asset — camera/EXIF fields, GPS, and location names.
	Metadata MetadataResponse `json:"metadata" api:"nullable"`
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
		Metadata         respjson.Field
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

// Metadata for an asset — camera/EXIF fields, GPS, and location names.
type MetadataResponse struct {
	// ID of the asset this metadata belongs to
	AssetID string `json:"asset_id" api:"required"`
	// When this metadata record was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When this metadata record was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// GPS altitude in meters
	Altitude float64 `json:"altitude" api:"nullable"`
	// Identifier for automatic photo stacking
	AutoStackID string `json:"auto_stack_id" api:"nullable"`
	// City name
	City string `json:"city" api:"nullable"`
	// Country name
	Country string `json:"country" api:"nullable"`
	// ISO 3166-1 alpha-2 country code (e.g., 'US', 'JP')
	CountryCode string `json:"country_code" api:"nullable"`
	// Image description or caption
	Description string `json:"description" api:"nullable"`
	// When the photo was digitized, with timezone offset if available
	DigitizedDatetime time.Time `json:"digitized_datetime" api:"nullable" format:"date-time"`
	// Exposure compensation in EV (e.g., -1.0, +0.5)
	ExposureBias float64 `json:"exposure_bias" api:"nullable"`
	// Shutter speed in seconds (e.g., 0.001 for 1/1000s)
	ExposureTime float64 `json:"exposure_time" api:"nullable"`
	// Aperture f-stop value (e.g., 2.8, 5.6)
	FNumber float64 `json:"f_number" api:"nullable"`
	// Focal length in millimeters
	FocalLength float64 `json:"focal_length" api:"nullable"`
	// Frame rate for video files
	Fps float64 `json:"fps" api:"nullable"`
	// ISO sensitivity value (e.g., 100, 800, 3200)
	ISO int64 `json:"iso" api:"nullable"`
	// GPS latitude in decimal degrees
	Latitude float64 `json:"latitude" api:"nullable"`
	// Lens model used (e.g., 'EF 24-70mm f/2.8L II USM')
	LensModel string `json:"lens_model" api:"nullable"`
	// Live photo content identifier
	LivePhotoCid string `json:"live_photo_cid" api:"nullable"`
	// GPS longitude in decimal degrees
	Longitude float64 `json:"longitude" api:"nullable"`
	// Camera manufacturer (e.g., 'Canon', 'Nikon')
	Make string `json:"make" api:"nullable"`
	// Camera model (e.g., 'EOS 5D Mark IV')
	Model string `json:"model" api:"nullable"`
	// When the file was last modified, with timezone offset if available
	ModifiedDatetime time.Time `json:"modified_datetime" api:"nullable" format:"date-time"`
	// Image orientation value (1-8) indicating rotation/flip: 1=normal, 2=mirror
	// horizontal, 3=rotate 180°, 4=mirror vertical, 5=mirror horizontal+rotate 90° CW,
	// 6=rotate 90° CW, 7=mirror horizontal+rotate 90° CCW, 8=rotate 90° CCW
	Orientation int64 `json:"orientation" api:"nullable"`
	// When the photo was originally taken, with timezone offset if available
	OriginalDatetime time.Time `json:"original_datetime" api:"nullable" format:"date-time"`
	// Landmark or point-of-interest name
	PlaceName string `json:"place_name" api:"nullable"`
	// Projection type (e.g., for 360° photos)
	ProjectionType string `json:"projection_type" api:"nullable"`
	// User or camera rating (typically 1-5 stars)
	Rating int64 `json:"rating" api:"nullable"`
	// State/province name
	State string `json:"state" api:"nullable"`
	// Neighborhood or district
	Sublocation string `json:"sublocation" api:"nullable"`
	// IANA timezone identifier (e.g., 'America/Los_Angeles')
	Timezone string `json:"timezone" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetID           respjson.Field
		CreatedAt         respjson.Field
		UpdatedAt         respjson.Field
		Altitude          respjson.Field
		AutoStackID       respjson.Field
		City              respjson.Field
		Country           respjson.Field
		CountryCode       respjson.Field
		Description       respjson.Field
		DigitizedDatetime respjson.Field
		ExposureBias      respjson.Field
		ExposureTime      respjson.Field
		FNumber           respjson.Field
		FocalLength       respjson.Field
		Fps               respjson.Field
		ISO               respjson.Field
		Latitude          respjson.Field
		LensModel         respjson.Field
		LivePhotoCid      respjson.Field
		Longitude         respjson.Field
		Make              respjson.Field
		Model             respjson.Field
		ModifiedDatetime  respjson.Field
		Orientation       respjson.Field
		OriginalDatetime  respjson.Field
		PlaceName         respjson.Field
		ProjectionType    respjson.Field
		Rating            respjson.Field
		State             respjson.Field
		Sublocation       respjson.Field
		Timezone          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetadataResponse) RawJSON() string { return r.JSON.raw }
func (r *MetadataResponse) UnmarshalJSON(data []byte) error {
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
	// Return only assets that are in the album with this ID. Equivalent to calling
	// `list_album_assets` with `album_id` and then fetching each asset — prefer this
	// param when you need the full asset metadata in one call.
	AlbumID param.Opt[string] `query:"album_id,omitzero" json:"-"`
	// Library to list assets from. Optional if the user has a single library; required
	// when they have multiple. Use `list_libraries` to enumerate available libraries.
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Only include assets captured strictly after this instant (ISO 8601; exclusive).
	// `local_datetime` is the photo's wall-clock time in the device's own timezone.
	// Naive values compare directly against `local_datetime`. Timezone-aware values:
	// assets with a known offset are compared in UTC (`local_datetime - offset`);
	// assets without an offset fall back to wall-clock comparison against
	// `local_datetime`. Equivalent in purpose to `captured_after` on `search_assets`
	// (naming inconsistency is tracked as a follow-up).
	LocalDatetimeAfter param.Opt[time.Time] `query:"local_datetime_after,omitzero" format:"date-time" json:"-"`
	// Only include assets captured strictly before this instant (ISO 8601; exclusive).
	// Same awareness/offset semantics as `local_datetime_after`. Equivalent in purpose
	// to `captured_before` on `search_assets` (naming inconsistency is tracked as a
	// follow-up).
	LocalDatetimeBefore param.Opt[time.Time] `query:"local_datetime_before,omitzero" format:"date-time" json:"-"`
	// Return only assets containing a face belonging to this person. Singular on this
	// tool; the sibling `search_assets` uses `person_ids` (plural, ALL-of).
	PersonID param.Opt[string] `query:"person_id,omitzero" json:"-"`
	// Cursor for pagination. Pass the `id` of the last asset in the previous
	// response's `data` to fetch the next page. Omit for the first page. `list_assets`
	// uses cursor pagination; the sibling `search_assets` uses 1-indexed `page`
	// numbers (naming inconsistency is tracked as a follow-up).
	StartingAfterID param.Opt[string] `query:"starting_after_id,omitzero" json:"-"`
	// Maximum number of assets to return per page (1–200). Defaults to 20.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Look up specific assets by ID (max 100; each ID has the `asset_` prefix). Use
	// this for bulk fetch when you already have asset IDs. Combines with other filters
	// (album_id, person_id, datetime range) using AND logic — the result is the
	// intersection.
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
	// values compare directly against local_datetime. Timezone-aware values: assets
	// with a known offset are compared in UTC (local_datetime - offset); assets
	// without an offset fall back to wall-clock comparison against local_datetime.
	LocalDatetimeAfter param.Opt[time.Time] `query:"local_datetime_after,omitzero" format:"date-time" json:"-"`
	// Only include assets with local_datetime before this value (ISO 8601). Naive
	// values compare directly against local_datetime. Timezone-aware values: assets
	// with a known offset are compared in UTC (local_datetime - offset); assets
	// without an offset fall back to wall-clock comparison against local_datetime. Use
	// the last time_bucket from a previous response to paginate.
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
