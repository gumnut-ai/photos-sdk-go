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

// EventsV2Service contains methods and other services that help with interacting
// with the Gumnut API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventsV2Service] method instead.
type EventsV2Service struct {
	Options []option.RequestOption
}

// NewEventsV2Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEventsV2Service(opts ...option.RequestOption) (r EventsV2Service) {
	r = EventsV2Service{}
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
func (r *EventsV2Service) Get(ctx context.Context, query EventsV2GetParams, opts ...option.RequestOption) (res *EventsV2Response, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v2/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response containing a page of v2 events.
type EventsV2Response struct {
	// List of events, ordered by event ID (monotonically increasing)
	Data []EventsV2ResponseData `json:"data,required"`
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
func (r EventsV2Response) RawJSON() string { return r.JSON.raw }
func (r *EventsV2Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lightweight event record for v2 sync endpoint.
type EventsV2ResponseData struct {
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
func (r EventsV2ResponseData) RawJSON() string { return r.JSON.raw }
func (r *EventsV2ResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventsV2GetParams struct {
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

// URLQuery serializes [EventsV2GetParams]'s query parameters as `url.Values`.
func (r EventsV2GetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
