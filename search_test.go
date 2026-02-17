// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/photos-go"
	"github.com/stainless-sdks/photos-go/internal/testutil"
	"github.com/stainless-sdks/photos-go/option"
)

func TestSearchSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Search.Search(context.TODO(), photos.SearchSearchParams{
		CapturedAfter:  photos.Time(time.Now()),
		CapturedBefore: photos.Time(time.Now()),
		LibraryID:      photos.String("library_id"),
		Limit:          photos.Int(1),
		Page:           photos.Int(1),
		PersonIDs:      []string{"string", "string"},
		Query:          photos.String("query"),
		Threshold:      photos.Float(0),
	})
	if err != nil {
		var apierr *photos.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSearchSearchAssetsWithOptionalParams(t *testing.T) {
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
	_, err := client.Search.SearchAssets(context.TODO(), photos.SearchSearchAssetsParams{
		CapturedAfter:  photos.Time(time.Now()),
		CapturedBefore: photos.Time(time.Now()),
		Image:          io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		LibraryID:      photos.String("library_id"),
		Limit:          photos.Int(1),
		Page:           photos.Int(1),
		PersonIDs:      []string{"string"},
		Query:          photos.String("query"),
		Threshold:      photos.Float(0),
	})
	if err != nil {
		var apierr *photos.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
