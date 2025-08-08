package case_insensitive_emails

import (
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tests"
)

const testDataDir = "./test/pb_data/"

func generateToken(collectionNameOrId string, email string) (string, error) {
	app, err := tests.NewTestApp(testDataDir)
	if err != nil {
		return "", err
	}
	defer app.Cleanup()

	record, err := app.FindAuthRecordByEmail(collectionNameOrId, email)
	if err != nil {
		return "", err
	}

	return record.NewAuthToken()
}

func TestPlugin(t *testing.T) {
	setupTestApp := func(t testing.TB) *tests.TestApp {
		testApp, err := tests.NewTestApp(testDataDir)
		if err != nil {
			t.Fatal(err)
		}
		(&Plugin{
			// test config will go here
		}).Init(testApp)
		return testApp
	}

	superuserToken, err := generateToken(core.CollectionNameSuperusers, "admin@example.com")
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []tests.ApiScenario{
		{
			Name:   "create record",
			Method: http.MethodPost,
			URL:    "/api/collections/users/records",
			Headers: map[string]string{
				"Authorization": superuserToken,
			},
			Body: strings.NewReader(`{
				"email": "TEST@EXAMPLE.COM",
				"password": "test12345",
				"passwordConfirm": "test12345",
				"emailVisibility": true
			}`),
			ExpectedStatus: http.StatusOK,
			ExpectedContent: []string{
				`"email":"test@example.com"`,
			},
			TestAppFactory: setupTestApp,
		},
		{
			Name:   "update record",
			Method: http.MethodPatch,
			URL:    "/api/collections/users/records/69n7x705xe51e1k",
			Headers: map[string]string{
				"Authorization": superuserToken,
			},
			Body: strings.NewReader(`{
				"email": "TEST@EXAMPLE.COM"
			}`),
			ExpectedStatus: http.StatusOK,
			ExpectedContent: []string{
				`"email":"test@example.com"`,
			},
			TestAppFactory: setupTestApp,
		},
		// query
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
