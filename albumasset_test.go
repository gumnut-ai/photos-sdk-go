// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/photos-go"
	"github.com/stainless-sdks/photos-go/internal/testutil"
	"github.com/stainless-sdks/photos-go/option"
)

func TestAlbumAssetListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.AlbumAssets.List(context.TODO(), photos.AlbumAssetListParams{
		AlbumID:         photos.String("album_id"),
		AssetID:         photos.String("asset_id"),
		IDs:             []string{"string", "string"},
		LibraryID:       photos.String("library_id"),
		Limit:           photos.Int(1),
		StartingAfterID: photos.String("starting_after_id"),
	})
	if err != nil {
		var apierr *photos.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAlbumAssetGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.AlbumAssets.Get(context.TODO(), "album_asset_id")
	if err != nil {
		var apierr *photos.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
