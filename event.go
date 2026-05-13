// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/gumnut-ai/photos-sdk-go/internal/apijson"
	"github.com/gumnut-ai/photos-sdk-go/internal/apiquery"
	"github.com/gumnut-ai/photos-sdk-go/internal/requestconfig"
	"github.com/gumnut-ai/photos-sdk-go/option"
	"github.com/gumnut-ai/photos-sdk-go/packages/param"
	"github.com/gumnut-ai/photos-sdk-go/packages/respjson"
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

// Returns a paginated stream of change events (create/update/delete) for entities
// in the library. Each event is a lightweight record — `entity_type`, `entity_id`,
// `event_type`, and timestamps — pointing at a concrete entity that has changed.
// Follow up with `get_asset`, `get_album`, `get_person`, or `get_face` to fetch
// full entity data when needed.
//
// **Use this tool** when the user wants to synchronise a local copy of their
// library, audit recent activity, or detect deletions. **Don't use it** for
// content queries — use `search_assets` or `list_assets` instead. Events cannot be
// filtered by content or asset metadata.
//
// **Pagination:** cursor-based via `after_cursor`. When `has_more` is true, pass
// the last event's `cursor` value into `after_cursor` to fetch the next page.
//
// **Recommended sync pattern:**
//
//  1. Capture current time as `sync_end`.
//  2. Fetch events with `created_at_lt=sync_end`.
//  3. For subsequent pages, use
//     `after_cursor={last.cursor}&created_at_lt=sync_end`.
//  4. Continue until `has_more=false`.
//  5. For each event, fetch the entity data from the appropriate endpoint if
//     needed.
//  6. Store `sync_end` as checkpoint for next sync.
//
// **Handling deletions:** when `event_type` ends with `_deleted` or `_removed`,
// the entity no longer exists — remove it from the local cache. Some deletion
// events include a `payload` field with context (e.g., `album_asset_removed`
// carries `album_id` and `asset_id` since the junction row is gone).
//
// **Event types:**
//
// - `asset_created`, `asset_updated`, `asset_deleted`
// - `album_created`, `album_updated`, `album_deleted`
// - `person_created`, `person_updated`, `person_deleted`
// - `face_created`, `face_updated`, `face_deleted`
// - `album_asset_added`, `album_asset_removed`
// - `metadata_updated`
func (r *EventService) Get(ctx context.Context, query EventGetParams, opts ...option.RequestOption) (res *EventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Response containing a page of events.
type EventsResponse struct {
	// List of events, ordered by event ID (monotonically increasing)
	Data []EventsResponseData `json:"data" api:"required"`
	// True if there are more events after this page. Pass the last event's `cursor`
	// value as `after_cursor` to fetch the next page.
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
func (r EventsResponse) RawJSON() string { return r.JSON.raw }
func (r *EventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lightweight event record for sync endpoint.
type EventsResponseData struct {
	// When the event was recorded
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Opaque cursor for pagination. Pass as after_cursor to get the next page.
	Cursor string `json:"cursor" api:"required"`
	// ID of the entity that changed
	EntityID string `json:"entity_id" api:"required"`
	// Type of entity that changed (e.g., 'asset', 'album', 'person')
	EntityType string `json:"entity_type" api:"required"`
	// Semantic event type (e.g., 'asset_created', 'album_deleted')
	EventType string `json:"event_type" api:"required"`
	// Optional extra context for the event (e.g., foreign keys for junction table
	// deletions)
	Payload map[string]any `json:"payload" api:"nullable"`
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

type EventGetParams struct {
	// Opaque cursor from the last event of the previous page. Pass the `cursor` field
	// from the last event to fetch the next page. Omit for the first page.
	AfterCursor param.Opt[string] `query:"after_cursor,omitzero" json:"-"`
	// Only return events created at or after this timestamp (ISO 8601). Set this to
	// the previous sync's checkpoint when doing incremental sync.
	CreatedAtGte param.Opt[time.Time] `query:"created_at_gte,omitzero" format:"date-time" json:"-"`
	// Only return events created strictly before this timestamp (ISO 8601).
	// Recommended for bounding a sync operation — capture `now` once and reuse it as
	// `created_at_lt` across all pages so newly arriving events don't shift the
	// window.
	CreatedAtLt param.Opt[time.Time] `query:"created_at_lt,omitzero" format:"date-time" json:"-"`
	// Library to stream events from. Optional if the user has a single library;
	// required when they have multiple. Use `list_libraries` to enumerate.
	LibraryID param.Opt[string] `query:"library_id,omitzero" json:"-"`
	// Maximum number of events to return per page (1–200). Defaults to 20.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Entity types to include (e.g., `asset`, `album`). Valid values: `asset`,
	// `album`, `person`, `face`, `album_asset`, `metadata`. Accepts multiple
	// `entity_types=` query params or a single comma-delimited value (e.g.,
	// `entity_types=asset,album`). Omit to receive events for all types.
	EntityTypes []string `query:"entity_types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventGetParams]'s query parameters as `url.Values`.
func (r EventGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
