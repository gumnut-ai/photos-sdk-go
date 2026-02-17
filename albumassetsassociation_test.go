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

func TestAlbumAssetsAssociationList(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Albums.AssetsAssociations.List(context.TODO(), "album_id")
	if err != nil {
		var apierr *photos.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAlbumAssetsAssociationAdd(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Albums.AssetsAssociations.Add(
		context.TODO(),
		"album_id",
		photos.AlbumAssetsAssociationAddParams{
			AlbumAssetAssociation: photos.AlbumAssetAssociationParam{
				AssetIDs: []string{"string"},
			},
		},
	)
	if err != nil {
		var apierr *photos.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAlbumAssetsAssociationRemove(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	err := client.Albums.AssetsAssociations.Remove(
		context.TODO(),
		"album_id",
		photos.AlbumAssetsAssociationRemoveParams{
			AlbumAssetAssociation: photos.AlbumAssetAssociationParam{
				AssetIDs: []string{"string"},
			},
		},
	)
	if err != nil {
		var apierr *photos.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
