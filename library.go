// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/gumnut-ai/photos-sdk-go/internal/apijson"
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

// Fetches one library's metadata by ID (name, description, asset count). Use when
// you already have a specific `library_id`; for enumerating a user's libraries
// prefer `list_libraries`.
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

// Returns every library owned by the authenticated user (no pagination — users
// typically have one or a handful). Call this when another tool's `library_id`
// parameter is required but you don't yet know which libraries exist. A
// single-library user can usually omit `library_id` on other tools entirely.
func (r *LibraryService) List(ctx context.Context, opts ...option.RequestOption) (res *[]LibraryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/libraries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deletes the library and all its contents — assets (including their stored
// files), albums, people, and faces. **Destructive and irreversible** — should be
// used only when the user explicitly confirms they want to destroy an entire
// library.
func (r *LibraryService) Delete(ctx context.Context, libraryID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if libraryID == "" {
		err = errors.New("missing required library_id parameter")
		return err
	}
	path := fmt.Sprintf("api/libraries/%s", libraryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
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
