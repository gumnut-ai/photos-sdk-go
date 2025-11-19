// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/photos-go/internal/apijson"
	"github.com/stainless-sdks/photos-go/internal/requestconfig"
	"github.com/stainless-sdks/photos-go/option"
	"github.com/stainless-sdks/photos-go/packages/param"
	"github.com/stainless-sdks/photos-go/packages/respjson"
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

// Creates a new library for the authenticated user.
func (r *LibraryService) New(ctx context.Context, body LibraryNewParams, opts ...option.RequestOption) (res *LibraryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/libraries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns details of a specific library owned by the authenticated user.
func (r *LibraryService) Get(ctx context.Context, libraryID string, opts ...option.RequestOption) (res *LibraryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if libraryID == "" {
		err = errors.New("missing required library_id parameter")
		return
	}
	path := fmt.Sprintf("api/libraries/%s", libraryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the name and/or description of a library owned by the authenticated
// user.
func (r *LibraryService) Update(ctx context.Context, libraryID string, body LibraryUpdateParams, opts ...option.RequestOption) (res *LibraryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if libraryID == "" {
		err = errors.New("missing required library_id parameter")
		return
	}
	path := fmt.Sprintf("api/libraries/%s", libraryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns all libraries owned by the authenticated user.
func (r *LibraryService) List(ctx context.Context, opts ...option.RequestOption) (res *[]LibraryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/libraries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a library and all its associated data (assets, albums, people, faces).
// Cannot delete the user's only library.
func (r *LibraryService) Delete(ctx context.Context, libraryID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if libraryID == "" {
		err = errors.New("missing required library_id parameter")
		return
	}
	path := fmt.Sprintf("api/libraries/%s", libraryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Represents a user's photo library.
type LibraryResponse struct {
	// Unique library identifier with 'lib\_' prefix
	ID string `json:"id,required"`
	// When this library was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Display name of the library
	Name string `json:"name,required"`
	// When this library was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// ID of the user who owns this library
	UserID string `json:"user_id,required"`
	// Optional description text for the library
	Description string `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
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
	Name        string            `json:"name,required"`
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
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r LibraryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LibraryUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LibraryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
