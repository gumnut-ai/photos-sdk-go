// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/photos-go"
	"github.com/stainless-sdks/photos-go/internal/testutil"
	"github.com/stainless-sdks/photos-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := photos.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	t.Skip("Mock server tests are disabled")
	albumResponse, err := client.Albums.New(context.TODO(), photos.AlbumNewParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", albumResponse.ID)
}
