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

// APIKeyService contains methods and other services that help with interacting
// with the Gumnut API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyService] method instead.
type APIKeyService struct {
	Options []option.RequestOption
}

// NewAPIKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIKeyService(opts ...option.RequestOption) (r APIKeyService) {
	r = APIKeyService{}
	r.Options = opts
	return
}

// Creates a new API key for the current user
func (r *APIKeyService) New(ctx context.Context, body APIKeyNewParams, opts ...option.RequestOption) (res *APIKeyNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api-keys/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates the name of a specific API key
func (r *APIKeyService) Update(ctx context.Context, keyID string, body APIKeyUpdateParams, opts ...option.RequestOption) (res *APIKeyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("api-keys/%s", keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieves a list of all API keys for the current user
func (r *APIKeyService) List(ctx context.Context, opts ...option.RequestOption) (res *[]APIKeyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api-keys/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a specific API key
func (r *APIKeyService) Delete(ctx context.Context, keyID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("api-keys/%s", keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Represents an API key for authentication (without exposing the actual key).
type APIKeyResponse struct {
	// Unique API key identifier with 'apikey\_' prefix
	ID string `json:"id" api:"required"`
	// When this API key was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this API key is currently valid and can be used
	IsActive bool `json:"is_active" api:"required"`
	// When this API key was last used for authentication
	LastUsedAt time.Time `json:"last_used_at" api:"nullable" format:"date-time"`
	// Optional descriptive name for this API key
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsActive    respjson.Field
		LastUsedAt  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response when creating a new API key - includes the actual key value.
//
// This is the only time the raw API key is exposed. After creation, only the
// hashed version is stored and the raw key cannot be retrieved.
type APIKeyNewResponse struct {
	// Unique API key identifier with 'apikey\_' prefix
	ID string `json:"id" api:"required"`
	// The actual API key value - store this securely as it cannot be retrieved later
	APIKey string `json:"api_key" api:"required"`
	// When this API key was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this API key is currently valid and can be used
	IsActive bool `json:"is_active" api:"required"`
	// When this API key was last used for authentication
	LastUsedAt time.Time `json:"last_used_at" api:"nullable" format:"date-time"`
	// Optional descriptive name for this API key
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		APIKey      respjson.Field
		CreatedAt   respjson.Field
		IsActive    respjson.Field
		LastUsedAt  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyNewParams struct {
	Name string `json:"name" api:"required"`
	paramObj
}

func (r APIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyUpdateParams struct {
	Name string `json:"name" api:"required"`
	paramObj
}

func (r APIKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
