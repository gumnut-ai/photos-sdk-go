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
	"github.com/gumnut-ai/photos-sdk-go/packages/param"
	"github.com/gumnut-ai/photos-sdk-go/packages/respjson"
)

// LibraryService contains methods and other services that help with interacting
// with the Gumnut API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLibraryService] method instead.
type LibraryService struct {
	Options []option.RequestOption
}

// NewLibraryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLibraryService(opts ...option.RequestOption) (r LibraryService) {
	r = LibraryService{}
	r.Options = opts
	return
}

// Creates a new, empty photo library for the authenticated user. A library is the
// top-level container for assets, albums, people, and faces — most users have
// exactly one. Only create a new library when the user explicitly asks for a
// separate container.
func (r *LibraryService) New(ctx context.Context, body LibraryNewParams, opts ...option.RequestOption) (res *LibraryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/libraries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetches one library's metadata by ID. Returns the library regardless of trash
// state.
func (r *LibraryService) Get(ctx context.Context, libraryID string, opts ...option.RequestOption) (res *LibraryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if libraryID == "" {
		err = errors.New("missing required library_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/libraries/%s", libraryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Renames a library or changes its description. Only the fields included in the
// request body are changed. Library contents (assets, albums, people, faces) are
// not affected.
func (r *LibraryService) Update(ctx context.Context, libraryID string, body LibraryUpdateParams, opts ...option.RequestOption) (res *LibraryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if libraryID == "" {
		err = errors.New("missing required library_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/libraries/%s", libraryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns libraries owned by the authenticated user (no pagination — users
// typically have one or a handful). Call this when another tool's `library_id`
// parameter is required but you don't yet know which libraries exist. A
// single-library user can usually omit `library_id` on other tools entirely.
//
// By default trashed libraries are excluded. Pass `state=trashed` to list the
// trash drawer (ordered by most recently trashed) or `state=all` for both.
func (r *LibraryService) List(ctx context.Context, query LibraryListParams, opts ...option.RequestOption) (res *[]LibraryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/libraries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Expedites the background purge on a **trashed** library: the 90-day undo window
// is waived and the drain begins claiming this library on the next scheduled tick.
// Returns immediately; the drain proceeds asynchronously in bounded batches and
// does not block on completion. `restore_library` still works until the drain
// finishes purging all assets, but past this point it will recover only the assets
// the drain hasn't gotten to yet. Returns 409 if the library has not been trashed
// yet — call `trash_library` first.
func (r *LibraryService) Delete(ctx context.Context, libraryID string, opts ...option.RequestOption) (res *LibraryDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if libraryID == "" {
		err = errors.New("missing required library_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/libraries/%s", libraryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Restores a previously-trashed library so it reappears in default list/search
// results. Works as long as the library row still exists — once `get_library`
// returns 404 the row is gone and restore is no longer possible. If the background
// drain has already started purging assets, restore succeeds but recovers only the
// assets the drain hasn't gotten to yet.
//
// Pairs with `trash_library`. To restore individual trashed assets within an
// untrashed library, use `restore_assets` instead.
func (r *LibraryService) Restore(ctx context.Context, libraryID string, opts ...option.RequestOption) (res *LibraryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if libraryID == "" {
		err = errors.New("missing required library_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/libraries/%s/restore", libraryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Moves the library and all its contents into the trash. The library becomes
// inaccessible by default and can be fully restored within 90 days by calling
// `restore_library`. After 90 days the library's assets are gradually purged in
// the background; until the library row itself is removed, restore still works but
// recovers only the assets not yet purged.
//
// Idempotent — a second call on an already-trashed library no-ops. To trash
// individual assets without trashing the whole library, use `trash_assets`
// instead.
func (r *LibraryService) Trash(ctx context.Context, libraryID string, opts ...option.RequestOption) (res *LibraryTrashResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if libraryID == "" {
		err = errors.New("missing required library_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/libraries/%s/trash", libraryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Represents a user's photo library.
type LibraryResponse struct {
	// Unique library identifier with 'lib\_' prefix
	ID string `json:"id" api:"required"`
	// Total number of assets in this library
	AssetCount int64 `json:"asset_count" api:"required"`
	// When this library was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the library
	Name string `json:"name" api:"required"`
	// When this library was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// ID of the user who owns this library
	UserID string `json:"user_id" api:"required"`
	// Optional description text for the library
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AssetCount  respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		UserID      respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LibraryResponse) RawJSON() string { return r.JSON.raw }
func (r *LibraryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LibraryDeleteResponse = any

type LibraryTrashResponse = any

type LibraryNewParams struct {
	// Display name for the new library. Required.
	Name string `json:"name" api:"required"`
	// Optional free-form description shown alongside the library name.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r LibraryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LibraryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LibraryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LibraryUpdateParams struct {
	// New free-form description for the library. Omit to leave unchanged.
	Description param.Opt[string] `json:"description,omitzero"`
	// New display name for the library. Omit to leave unchanged.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r LibraryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LibraryUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LibraryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LibraryListParams struct {
	// Which set of libraries to return: `live` (default — excludes trashed), `trashed`
	// (only trashed, ordered by most recently trashed), or `all` (both).
	//
	// Any of "live", "trashed", "all".
	State LibraryListParamsState `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LibraryListParams]'s query parameters as `url.Values`.
func (r LibraryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Which set of libraries to return: `live` (default — excludes trashed), `trashed`
// (only trashed, ordered by most recently trashed), or `all` (both).
type LibraryListParamsState string

const (
	LibraryListParamsStateLive    LibraryListParamsState = "live"
	LibraryListParamsStateTrashed LibraryListParamsState = "trashed"
	LibraryListParamsStateAll     LibraryListParamsState = "all"
)
