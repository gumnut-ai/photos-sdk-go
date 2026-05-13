// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package photos_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/gumnut-ai/photos-sdk-go"
	"github.com/gumnut-ai/photos-sdk-go/internal/testutil"
	"github.com/gumnut-ai/photos-sdk-go/option"
)

func TestEventGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Events.Get(context.TODO(), photos.EventGetParams{
		AfterCursor:  photos.String("after_cursor"),
		CreatedAtGte: photos.Time(time.Now()),
		CreatedAtLt:  photos.Time(time.Now()),
		EntityTypes:  []string{"string", "string"},
		LibraryID:    photos.String("library_id"),
		Limit:        photos.Int(1),
	})
	if err != nil {
		var apierr *photos.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
