// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/gumnut-ai/photos-sdk-go/internal/apijson"
	"github.com/gumnut-ai/photos-sdk-go/internal/apiquery"
	"github.com/gumnut-ai/photos-sdk-go/internal/requestconfig"
	"github.com/gumnut-ai/photos-sdk-go/option"
	"github.com/gumnut-ai/photos-sdk-go/packages/param"
	"github.com/gumnut-ai/photos-sdk-go/packages/respjson"
)

// OAuthService contains methods and other services that help with interacting with
// the Gumnut API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOAuthService] method instead.
type OAuthService struct {
	Options []option.RequestOption
}

// NewOAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOAuthService(opts ...option.RequestOption) (r OAuthService) {
	r = OAuthService{}
	r.Options = opts
	return
}

// Generate OAuth authorization URL with state and nonce for CSRF and replay attack
// protection. State is stored with TTL for validation.
func (r *OAuthService) AuthURL(ctx context.Context, query OAuthAuthURLParams, opts ...option.RequestOption) (res *AuthURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/oauth/auth-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Exchange OAuth authorization code for application JWT after validating state,
// nonce, and ID token signature. User is retrieved from or created in the database
// and details added to the JWT.
func (r *OAuthService) Exchange(ctx context.Context, body OAuthExchangeParams, opts ...option.RequestOption) (res *ExchangeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/oauth/exchange"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the OAuth provider's logout endpoint URL from OIDC discovery. This can
// be used to redirect users to logout from the OAuth provider after logging out
// locally.
func (r *OAuthService) LogoutEndpoint(ctx context.Context, opts ...option.RequestOption) (res *LogoutEndpointResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/oauth/logout-endpoint"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Response containing OAuth authorization URL
type AuthURLResponse struct {
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthURLResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing JWT and user info
type ExchangeResponse struct {
	AccessToken string `json:"access_token" api:"required"`
	// User information in token exchange response
	User ExchangeResponseUser `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExchangeResponse) RawJSON() string { return r.JSON.raw }
func (r *ExchangeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User information in token exchange response
type ExchangeResponseUser struct {
	ID          string `json:"id" api:"required"`
	ClerkUserID string `json:"clerk_user_id" api:"required"`
	Email       string `json:"email" api:"required"`
	FirstName   string `json:"first_name" api:"required"`
	IsActive    bool   `json:"is_active" api:"required"`
	IsVerified  bool   `json:"is_verified" api:"required"`
	LastName    string `json:"last_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ClerkUserID respjson.Field
		Email       respjson.Field
		FirstName   respjson.Field
		IsActive    respjson.Field
		IsVerified  respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExchangeResponseUser) RawJSON() string { return r.JSON.raw }
func (r *ExchangeResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing OAuth provider logout endpoint
type LogoutEndpointResponse struct {
	LogoutEndpoint string `json:"logout_endpoint" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LogoutEndpoint respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogoutEndpointResponse) RawJSON() string { return r.JSON.raw }
func (r *LogoutEndpointResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthAuthURLParams struct {
	// The URI to redirect to after OAuth consent. Must match the registered redirect
	// URI in OAuth client configuration.
	RedirectUri string `query:"redirect_uri" api:"required" json:"-"`
	// PKCE code challenge derived from code_verifier. Required for public clients to
	// prevent authorization code interception attacks.
	CodeChallenge param.Opt[string] `query:"code_challenge,omitzero" json:"-"`
	// PKCE code challenge method, typically 'S256' (SHA-256 hash). Must be provided if
	// code_challenge is specified.
	CodeChallengeMethod param.Opt[string] `query:"code_challenge_method,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OAuthAuthURLParams]'s query parameters as `url.Values`.
func (r OAuthAuthURLParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OAuthExchangeParams struct {
	// Authorization code returned by the OAuth provider after user consent
	Code param.Opt[string] `json:"code,omitzero"`
	// PKCE code verifier that corresponds to the code_challenge sent in the
	// authorization request
	CodeVerifier param.Opt[string] `json:"code_verifier,omitzero"`
	// Error code if OAuth provider returned an error instead of authorization code
	Error param.Opt[string] `json:"error,omitzero"`
	// State token from the initial auth request, used for CSRF protection
	State param.Opt[string] `json:"state,omitzero"`
	paramObj
}

func (r OAuthExchangeParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthExchangeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthExchangeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
