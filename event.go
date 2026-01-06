// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/photos-go/internal/apijson"
	"github.com/stainless-sdks/photos-go/internal/apiquery"
	"github.com/stainless-sdks/photos-go/internal/requestconfig"
	"github.com/stainless-sdks/photos-go/option"
	"github.com/stainless-sdks/photos-go/packages/param"
	"github.com/stainless-sdks/photos-go/packages/respjson"
)

// EventService contains methods and other services that help with interacting with
// the Gumnut API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.Options = opts
	return
}

// Retrieves a list of entity change events for syncing.
//
// Events are returned in order of entity type priority (assets first, then exif,
// albums, etc.), then by `updated_at` timestamp (oldest first), then by entity ID
// for tie-breaking.
//
// **Pagination:** Use `updated_at_gte` with the timestamp of the last received
// event to fetch the next page. When multiple entities share the same timestamp,
// also provide `starting_after_id` with the last entity's ID to avoid duplicates.
// Use `updated_at_lt` to bound the sync window and prevent infinite loops when new
// events are created during sync.
//
// **Important:** When using `starting_after_id`, you must specify exactly one
// `entity_types` value. This ensures the cursor ID is unambiguous. To sync all
// entity types with cursor support, query each entity type separately.
//
// **Recommended sync pattern (per entity type):**
//
//  1. Capture current time as `sync_started_at`
//  2. For each entity type, fetch events with
//     `entity_types={type}&updated_at_lt=sync_started_at`
//  3. For subsequent pages, use
//     `entity_types={type}&updated_at_gte={last.updated_at}&starting_after_id={last.id}&updated_at_lt=sync_started_at`
//  4. Continue until an empty result set is returned
//  5. Store `sync_started_at` as checkpoint for next sync
//
// **Entity ID field by type:**
//
// - Most entities: use the `id` field from the response
// - Exif: use the `asset_id` field (exif has no separate id)
func (r *EventService) Get(ctx context.Context, query EventGetParams, opts ...option.RequestOption) (res *EventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Event payload for album_asset entities.
type AlbumAssetEventPayload struct {
	// Full album_asset data
	Data AlbumAssetResponse `json:"data,required"`
	// Any of "album_asset".
	EntityType AlbumAssetEventPayloadEntityType `json:"entity_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		EntityType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlbumAssetEventPayload) RawJSON() string { return r.JSON.raw }
func (r *AlbumAssetEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlbumAssetEventPayloadEntityType string

const (
	AlbumAssetEventPayloadEntityTypeAlbumAsset AlbumAssetEventPayloadEntityType = "album_asset"
)

// Represents a link between an album and an asset.
type AlbumAssetResponse struct {
	// Unique album*asset identifier with 'album_asset*' prefix
	ID string `json:"id,required"`
	// ID of the album
	AlbumID string `json:"album_id,required"`
	// ID of the asset
	AssetID string `json:"asset_id,required"`
	// When this link was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// When this link was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
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

// Event payload for album entities.
type AlbumEventPayload struct {
	// Full album data
	Data AlbumResponse `json:"data,required"`
	// Any of "album".
	EntityType AlbumEventPayloadEntityType `json:"entity_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		EntityType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlbumEventPayload) RawJSON() string { return r.JSON.raw }
func (r *AlbumEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlbumEventPayloadEntityType string

const (
	AlbumEventPayloadEntityTypeAlbum AlbumEventPayloadEntityType = "album"
)

// Event payload for asset entities.
type AssetEventPayload struct {
	// Full asset data
	Data AssetResponse `json:"data,required"`
	// Any of "asset".
	EntityType AssetEventPayloadEntityType `json:"entity_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		EntityType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetEventPayload) RawJSON() string { return r.JSON.raw }
func (r *AssetEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetEventPayloadEntityType string

const (
	AssetEventPayloadEntityTypeAsset AssetEventPayloadEntityType = "asset"
)

// Response containing events.
type EventsResponse struct {
	// List of events, ordered by entity type priority, then updated_at, then entity_id
	Data []EventsResponseDataUnion `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventsResponse) RawJSON() string { return r.JSON.raw }
func (r *EventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventsResponseDataUnion contains all possible properties and values from
// [AssetEventPayload], [AlbumEventPayload], [PersonEventPayload],
// [FaceEventPayload], [AlbumAssetEventPayload], [ExifEventPayload].
//
// Use the [EventsResponseDataUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EventsResponseDataUnion struct {
	// This field is a union of [AssetResponse], [AlbumResponse], [PersonResponse],
	// [FaceResponse], [AlbumAssetResponse], [ExifResponse]
	Data EventsResponseDataUnionData `json:"data"`
	// Any of "asset", "album", "person", "face", "album_asset", "exif".
	EntityType string `json:"entity_type"`
	JSON       struct {
		Data       respjson.Field
		EntityType respjson.Field
		raw        string
	} `json:"-"`
}

// anyEventsResponseData is implemented by each variant of
// [EventsResponseDataUnion] to add type safety for the return type of
// [EventsResponseDataUnion.AsAny]
type anyEventsResponseData interface {
	implEventsResponseDataUnion()
}

func (AssetEventPayload) implEventsResponseDataUnion()      {}
func (AlbumEventPayload) implEventsResponseDataUnion()      {}
func (PersonEventPayload) implEventsResponseDataUnion()     {}
func (FaceEventPayload) implEventsResponseDataUnion()       {}
func (AlbumAssetEventPayload) implEventsResponseDataUnion() {}
func (ExifEventPayload) implEventsResponseDataUnion()       {}

// Use the following switch statement to find the correct variant
//
//	switch variant := EventsResponseDataUnion.AsAny().(type) {
//	case photos.AssetEventPayload:
//	case photos.AlbumEventPayload:
//	case photos.PersonEventPayload:
//	case photos.FaceEventPayload:
//	case photos.AlbumAssetEventPayload:
//	case photos.ExifEventPayload:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u EventsResponseDataUnion) AsAny() anyEventsResponseData {
	switch u.EntityType {
	case "asset":
		return u.AsAsset()
	case "album":
		return u.AsAlbum()
	case "person":
		return u.AsPerson()
	case "face":
		return u.AsFace()
	case "album_asset":
		return u.AsAlbumAsset()
	case "exif":
		return u.AsExif()
	}
	return nil
}

func (u EventsResponseDataUnion) AsAsset() (v AssetEventPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsResponseDataUnion) AsAlbum() (v AlbumEventPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsResponseDataUnion) AsPerson() (v PersonEventPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsResponseDataUnion) AsFace() (v FaceEventPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsResponseDataUnion) AsAlbumAsset() (v AlbumAssetEventPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventsResponseDataUnion) AsExif() (v ExifEventPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventsResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *EventsResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventsResponseDataUnionData is an implicit subunion of
// [EventsResponseDataUnion]. EventsResponseDataUnionData provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [EventsResponseDataUnion].
type EventsResponseDataUnionData struct {
	ID string `json:"id"`
	// This field is from variant [AssetResponse].
	Checksum  string    `json:"checksum"`
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [AssetResponse].
	DeviceAssetID string `json:"device_asset_id"`
	// This field is from variant [AssetResponse].
	DeviceID string `json:"device_id"`
	// This field is from variant [AssetResponse].
	FileCreatedAt time.Time `json:"file_created_at"`
	// This field is from variant [AssetResponse].
	FileModifiedAt time.Time `json:"file_modified_at"`
	// This field is from variant [AssetResponse].
	LocalDatetime time.Time `json:"local_datetime"`
	// This field is from variant [AssetResponse].
	MimeType string `json:"mime_type"`
	// This field is from variant [AssetResponse].
	OriginalFileName string    `json:"original_file_name"`
	UpdatedAt        time.Time `json:"updated_at"`
	// This field is from variant [AssetResponse].
	ChecksumSha1 string `json:"checksum_sha1"`
	// This field is from variant [AssetResponse].
	DownloadURL string `json:"download_url"`
	// This field is from variant [AssetResponse].
	Exif ExifResponse `json:"exif"`
	// This field is from variant [AssetResponse].
	Faces []FaceResponse `json:"faces"`
	// This field is from variant [AssetResponse].
	FileSizeBytes int64 `json:"file_size_bytes"`
	// This field is from variant [AssetResponse].
	Height int64 `json:"height"`
	// This field is from variant [AssetResponse].
	Metrics map[string]float64 `json:"metrics"`
	// This field is from variant [AssetResponse].
	People       []PersonResponse `json:"people"`
	ThumbnailURL string           `json:"thumbnail_url"`
	// This field is from variant [AssetResponse].
	Width int64 `json:"width"`
	// This field is from variant [AlbumResponse].
	AssetCount int64  `json:"asset_count"`
	Name       string `json:"name"`
	// This field is from variant [AlbumResponse].
	AlbumCoverAssetID string `json:"album_cover_asset_id"`
	Description       string `json:"description"`
	// This field is from variant [AlbumResponse].
	EndDate time.Time `json:"end_date"`
	// This field is from variant [AlbumResponse].
	StartDate time.Time `json:"start_date"`
	// This field is from variant [PersonResponse].
	IsFavorite bool `json:"is_favorite"`
	// This field is from variant [PersonResponse].
	IsHidden bool `json:"is_hidden"`
	// This field is from variant [PersonResponse].
	BirthDate time.Time `json:"birth_date"`
	// This field is from variant [PersonResponse].
	ThumbnailFaceID string `json:"thumbnail_face_id"`
	// This field is from variant [PersonResponse].
	ThumbnailFaceURL string `json:"thumbnail_face_url"`
	AssetID          string `json:"asset_id"`
	// This field is from variant [FaceResponse].
	BoundingBox map[string]int64 `json:"bounding_box"`
	// This field is from variant [FaceResponse].
	PersonID string `json:"person_id"`
	// This field is from variant [FaceResponse].
	TimestampMs int64 `json:"timestamp_ms"`
	// This field is from variant [AlbumAssetResponse].
	AlbumID string `json:"album_id"`
	// This field is from variant [ExifResponse].
	Altitude float64 `json:"altitude"`
	// This field is from variant [ExifResponse].
	AutoStackID string `json:"auto_stack_id"`
	// This field is from variant [ExifResponse].
	City string `json:"city"`
	// This field is from variant [ExifResponse].
	Country string `json:"country"`
	// This field is from variant [ExifResponse].
	DigitizedDatetime time.Time `json:"digitized_datetime"`
	// This field is from variant [ExifResponse].
	ExposureBias float64 `json:"exposure_bias"`
	// This field is from variant [ExifResponse].
	ExposureTime float64 `json:"exposure_time"`
	// This field is from variant [ExifResponse].
	FNumber float64 `json:"f_number"`
	// This field is from variant [ExifResponse].
	FocalLength float64 `json:"focal_length"`
	// This field is from variant [ExifResponse].
	Fps float64 `json:"fps"`
	// This field is from variant [ExifResponse].
	ISO int64 `json:"iso"`
	// This field is from variant [ExifResponse].
	Latitude float64 `json:"latitude"`
	// This field is from variant [ExifResponse].
	LensModel string `json:"lens_model"`
	// This field is from variant [ExifResponse].
	LivePhotoCid string `json:"live_photo_cid"`
	// This field is from variant [ExifResponse].
	Longitude float64 `json:"longitude"`
	// This field is from variant [ExifResponse].
	Make string `json:"make"`
	// This field is from variant [ExifResponse].
	Model string `json:"model"`
	// This field is from variant [ExifResponse].
	ModifiedDatetime time.Time `json:"modified_datetime"`
	// This field is from variant [ExifResponse].
	Orientation int64 `json:"orientation"`
	// This field is from variant [ExifResponse].
	OriginalDatetime time.Time `json:"original_datetime"`
	// This field is from variant [ExifResponse].
	ProfileDescription string `json:"profile_description"`
	// This field is from variant [ExifResponse].
	ProjectionType string `json:"projection_type"`
	// This field is from variant [ExifResponse].
	Rating int64 `json:"rating"`
	// This field is from variant [ExifResponse].
	State string `json:"state"`
	JSON  struct {
		ID                 respjson.Field
		Checksum           respjson.Field
		CreatedAt          respjson.Field
		DeviceAssetID      respjson.Field
		DeviceID           respjson.Field
		FileCreatedAt      respjson.Field
		FileModifiedAt     respjson.Field
		LocalDatetime      respjson.Field
		MimeType           respjson.Field
		OriginalFileName   respjson.Field
		UpdatedAt          respjson.Field
		ChecksumSha1       respjson.Field
		DownloadURL        respjson.Field
		Exif               respjson.Field
		Faces              respjson.Field
		FileSizeBytes      respjson.Field
		Height             respjson.Field
		Metrics            respjson.Field
		People             respjson.Field
		ThumbnailURL       respjson.Field
		Width              respjson.Field
		AssetCount         respjson.Field
		Name               respjson.Field
		AlbumCoverAssetID  respjson.Field
		Description        respjson.Field
		EndDate            respjson.Field
		StartDate          respjson.Field
		IsFavorite         respjson.Field
		IsHidden           respjson.Field
		BirthDate          respjson.Field
		ThumbnailFaceID    respjson.Field
		ThumbnailFaceURL   respjson.Field
		AssetID            respjson.Field
		BoundingBox        respjson.Field
		PersonID           respjson.Field
		TimestampMs        respjson.Field
		AlbumID            respjson.Field
		Altitude           respjson.Field
		AutoStackID        respjson.Field
		City               respjson.Field
		Country            respjson.Field
		DigitizedDatetime  respjson.Field
		ExposureBias       respjson.Field
		ExposureTime       respjson.Field
		FNumber            respjson.Field
		FocalLength        respjson.Field
		Fps                respjson.Field
		ISO                respjson.Field
		Latitude           respjson.Field
		LensModel          respjson.Field
		LivePhotoCid       respjson.Field
		Longitude          respjson.Field
		Make               respjson.Field
		Model              respjson.Field
		ModifiedDatetime   respjson.Field
		Orientation        respjson.Field
		OriginalDatetime   respjson.Field
		ProfileDescription respjson.Field
		ProjectionType     respjson.Field
		Rating             respjson.Field
		State              respjson.Field
		raw                string
	} `json:"-"`
}

func (r *EventsResponseDataUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event payload for exif entities.
type ExifEventPayload struct {
	// Full exif data
	Data ExifResponse `json:"data,required"`
	// Any of "exif".
	EntityType ExifEventPayloadEntityType `json:"entity_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		EntityType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExifEventPayload) RawJSON() string { return r.JSON.raw }
func (r *ExifEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExifEventPayloadEntityType string

const (
	ExifEventPayloadEntityTypeExif ExifEventPayloadEntityType = "exif"
)

// EXIF metadata extracted from image and video files.
type ExifResponse struct {
	// ID of the asset this EXIF data belongs to
	AssetID string `json:"asset_id,required"`
	// When this EXIF record was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// When this EXIF record was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// GPS altitude in meters
	Altitude float64 `json:"altitude,nullable"`
	// Identifier for automatic photo stacking
	AutoStackID string `json:"auto_stack_id,nullable"`
	// City name from GPS/location data
	City string `json:"city,nullable"`
	// Country name from GPS/location data
	Country string `json:"country,nullable"`
	// Image description or caption
	Description string `json:"description,nullable"`
	// When the photo was digitized, with timezone info
	DigitizedDatetime time.Time `json:"digitized_datetime,nullable" format:"date-time"`
	// Exposure compensation in EV (e.g., -1.0, +0.5)
	ExposureBias float64 `json:"exposure_bias,nullable"`
	// Shutter speed in seconds (e.g., 0.001 for 1/1000s)
	ExposureTime float64 `json:"exposure_time,nullable"`
	// Aperture f-stop value (e.g., 2.8, 5.6)
	FNumber float64 `json:"f_number,nullable"`
	// Focal length in millimeters
	FocalLength float64 `json:"focal_length,nullable"`
	// Frame rate for video files
	Fps float64 `json:"fps,nullable"`
	// ISO sensitivity value (e.g., 100, 800, 3200)
	ISO int64 `json:"iso,nullable"`
	// GPS latitude in decimal degrees
	Latitude float64 `json:"latitude,nullable"`
	// Lens model used (e.g., 'EF 24-70mm f/2.8L II USM')
	LensModel string `json:"lens_model,nullable"`
	// Live photo content identifier
	LivePhotoCid string `json:"live_photo_cid,nullable"`
	// GPS longitude in decimal degrees
	Longitude float64 `json:"longitude,nullable"`
	// Camera manufacturer (e.g., 'Canon', 'Nikon')
	Make string `json:"make,nullable"`
	// Camera model (e.g., 'EOS 5D Mark IV')
	Model string `json:"model,nullable"`
	// When the file was last modified, with timezone info
	ModifiedDatetime time.Time `json:"modified_datetime,nullable" format:"date-time"`
	// Image orientation value (1-8) indicating rotation/flip: 1=normal, 2=mirror
	// horizontal, 3=rotate 180°, 4=mirror vertical, 5=mirror horizontal+rotate 90° CW,
	// 6=rotate 90° CW, 7=mirror horizontal+rotate 90° CCW, 8=rotate 90° CCW
	Orientation int64 `json:"orientation,nullable"`
	// When the photo was originally taken, with timezone info
	OriginalDatetime time.Time `json:"original_datetime,nullable" format:"date-time"`
	// Color profile description
	ProfileDescription string `json:"profile_description,nullable"`
	// Projection type (e.g., for 360° photos)
	ProjectionType string `json:"projection_type,nullable"`
	// User or camera rating (typically 1-5 stars)
	Rating int64 `json:"rating,nullable"`
	// State/province name from GPS/location data
	State string `json:"state,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetID            respjson.Field
		CreatedAt          respjson.Field
		UpdatedAt          respjson.Field
		Altitude           respjson.Field
		AutoStackID        respjson.Field
		City               respjson.Field
		Country            respjson.Field
		Description        respjson.Field
		DigitizedDatetime  respjson.Field
		ExposureBias       respjson.Field
		ExposureTime       respjson.Field
		FNumber            respjson.Field
		FocalLength        respjson.Field
		Fps                respjson.Field
		ISO                respjson.Field
		Latitude           respjson.Field
		LensModel          respjson.Field
		LivePhotoCid       respjson.Field
		Longitude          respjson.Field
		Make               respjson.Field
		Model              respjson.Field
		ModifiedDatetime   respjson.Field
		Orientation        respjson.Field
		OriginalDatetime   respjson.Field
		ProfileDescription respjson.Field
		ProjectionType     respjson.Field
		Rating             respjson.Field
		State              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExifResponse) RawJSON() string { return r.JSON.raw }
func (r *ExifResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event payload for face entities.
type FaceEventPayload struct {
	// Full face data
	Data FaceResponse `json:"data,required"`
	// Any of "face".
	EntityType FaceEventPayloadEntityType `json:"entity_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		EntityType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaceEventPayload) RawJSON() string { return r.JSON.raw }
func (r *FaceEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaceEventPayloadEntityType string

const (
	FaceEventPayloadEntityTypeFace FaceEventPayloadEntityType = "face"
)

// Event payload for person entities.
type PersonEventPayload struct {
	// Full person data
	Data PersonResponse `json:"data,required"`
	// Any of "person".
	EntityType PersonEventPayloadEntityType `json:"entity_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		EntityType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonEventPayload) RawJSON() string { return r.JSON.raw }
func (r *PersonEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonEventPayloadEntityType string

const (
	PersonEventPayloadEntityTypePerson PersonEventPayloadEntityType = "person"
)

type EventGetParams struct {
	// Comma-separated list of entity types to include (e.g., 'asset,album'). Valid
	// types: asset, album, person, face, album_asset, exif. Default: all types.
	EntityTypes param.Opt[string] `query:"entity_types,omitzero" json:"-"`
	// Library to list events from. If not provided, uses the user's default library.
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Entity ID to start after for tie-breaking when paginating. Used with
	// updated_at_gte for composite keyset pagination. Requires exactly one
	// entity_types value. For exif entities, use asset_id.
	StartingAfterID param.Opt[string] `query:"starting_after_id,omitzero" json:"-"`
	// Only return events with updated_at >= this timestamp (ISO 8601 format)
	UpdatedAtGte param.Opt[time.Time] `query:"updated_at_gte,omitzero" format:"date-time" json:"-"`
	// Only return events with updated_at < this timestamp (ISO 8601 format).
	// Recommended for bounding sync operations.
	UpdatedAtLt param.Opt[time.Time] `query:"updated_at_lt,omitzero" format:"date-time" json:"-"`
	// Maximum number of events to return (1-500)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventGetParams]'s query parameters as `url.Values`.
func (r EventGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
