// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos

import (
	"context"
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
// Events are lightweight records indicating that entities have changed. Each event
// contains the entity type, entity ID, and event type (e.g., "asset_created",
// "album_deleted"). Clients should fetch full entity data from the appropriate
// endpoints if needed.
//
// **Pagination:** Use the `after_cursor` parameter with the `cursor` value from
// the last event to fetch the next page. The `has_more` field indicates if more
// events exist.
//
// **Recommended sync pattern:**
//
// 1. Capture current time as `sync_end`
// 2. Fetch events with `created_at_lt=sync_end`
// 3. For subsequent pages, use `after_cursor={last.cursor}&created_at_lt=sync_end`
// 4. Continue until `has_more=false`
// 5. For each event, fetch the entity data from the appropriate endpoint if needed
// 6. Store `sync_end` as checkpoint for next sync
//
// **Handling deletions:** When `event_type` ends with "\_deleted" or "\_removed",
// the entity no longer exists. Remove it from your local cache/database. Some
// deletion events include a `payload` field with additional context (e.g.,
// `album_asset_removed` includes `album_id` and `asset_id` since the junction
// record is deleted).
//
// **Event types:**
//
// - `asset_created`, `asset_updated`, `asset_deleted`
// - `album_created`, `album_updated`, `album_deleted`
// - `person_created`, `person_updated`, `person_deleted`
// - `face_created`, `face_updated`, `face_deleted`
// - `album_asset_added`, `album_asset_removed`
// - `exif_created`, `exif_updated`
func (r *EventService) Get(ctx context.Context, query EventGetParams, opts ...option.RequestOption) (res *EventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response containing a page of events.
type EventsResponse struct {
	// List of events, ordered by event ID (monotonically increasing)
	Data []EventsResponseData `json:"data,required"`
	// True if there are more events after this page. Use the last event's cursor to
	// fetch the next page.
	HasMore bool `json:"has_more,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventsResponse) RawJSON() string { return r.JSON.raw }
func (r *EventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lightweight event record for sync endpoint.
type EventsResponseData struct {
	// When the event was recorded
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Opaque cursor for pagination. Pass as after_cursor to get the next page.
	Cursor string `json:"cursor,required"`
	// ID of the entity that changed
	EntityID string `json:"entity_id,required"`
	// Type of entity that changed (e.g., 'asset', 'album', 'person')
	EntityType string `json:"entity_type,required"`
	// Semantic event type (e.g., 'asset_created', 'album_deleted')
	EventType string `json:"event_type,required"`
	// Optional extra context for the event (e.g., foreign keys for junction table
	// deletions)
	Payload map[string]any `json:"payload,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Cursor      respjson.Field
		EntityID    respjson.Field
		EntityType  respjson.Field
		EventType   respjson.Field
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventsResponseData) RawJSON() string { return r.JSON.raw }
func (r *EventsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type EventGetParams struct {
	// Cursor from the last event to paginate from. Pass the `cursor` field from the
	// last event to get the next page.
	AfterCursor param.Opt[string] `query:"after_cursor,omitzero" json:"-"`
	// Only return events created at or after this timestamp (ISO 8601 format)
	CreatedAtGte param.Opt[time.Time] `query:"created_at_gte,omitzero" format:"date-time" json:"-"`
	// Only return events created before this timestamp (ISO 8601 format). Recommended
	// for bounding sync operations.
	CreatedAtLt param.Opt[time.Time] `query:"created_at_lt,omitzero" format:"date-time" json:"-"`
	// Comma-separated list of entity types to include (e.g., 'asset,album'). Valid
	// types: asset, album, person, face, album_asset, exif. Default: all types.
	EntityTypes param.Opt[string] `query:"entity_types,omitzero" json:"-"`
	// Library to list events from. If not provided, uses the user's default library.
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Maximum number of events to return (1-500)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventGetParams]'s query parameters as `url.Values`.
func (r EventGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
