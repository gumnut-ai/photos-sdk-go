// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/photos-go/internal/apijson"
	"github.com/stainless-sdks/photos-go/internal/requestconfig"
	"github.com/stainless-sdks/photos-go/option"
	"github.com/stainless-sdks/photos-go/packages/respjson"
)

// UserService contains methods and other services that help with interacting with
// the Gumnut API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// Returns information about the authenticated user making the request.
func (r *UserService) Me(ctx context.Context, opts ...option.RequestOption) (res *UserResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Represents a user account with profile information.
type UserResponse struct {
	// Unique user identifier with 'intuser\_' prefix
	ID string `json:"id" api:"required"`
	// When this user account was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this user account is currently active
	IsActive bool `json:"is_active" api:"required"`
	// Whether this user has superuser/admin privileges
	IsSuperuser bool `json:"is_superuser" api:"required"`
	// Whether this user's email is verified
	IsVerified bool `json:"is_verified" api:"required"`
	// When this user account was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// User's email address
	Email string `json:"email" api:"nullable"`
	// User's first name
	FirstName string `json:"first_name" api:"nullable"`
	// User's last name
	LastName string `json:"last_name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsActive    respjson.Field
		IsSuperuser respjson.Field
		IsVerified  respjson.Field
		UpdatedAt   respjson.Field
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserResponse) RawJSON() string { return r.JSON.raw }
func (r *UserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
