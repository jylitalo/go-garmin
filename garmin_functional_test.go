//go:build functional

package garmin

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/harrybrwn/go-garmin/internal/rt"
)

var (
	cacher   = NewFileTokenCacher(filepath.Join(os.TempDir(), "go-garmin-tests", "test"))
	debugger = &rt.Debugger{SkipBody: true}
)

func testapi(t *testing.T) *API {
	t.Helper()
	client := NewClient(WithCacher(cacher))
	client.prependTransport(debugger)
	email := os.Getenv("GOGARMIN_TEST_EMAIL")
	pw := os.Getenv("GOGARMIN_TEST_PASSWORD")
	if email == "" {
		t.Fatal("no test email found: set GARMIN_TEST_EMAIL")
		return nil
	}
	if pw == "" {
		t.Fatal("no test password found: set GARMIN_TEST_PASSWORD")
		return nil
	}
	err := client.Login(email, pw)
	if err != nil {
		t.Fatal(err)
		return nil
	}
	return NewAPI(client)
}

func testToken(t *testing.T) *AccessToken {
	t.Helper()
	client := NewClient(WithCacher(cacher))
	client.prependTransport(debugger)
	email := os.Getenv("GOGARMIN_TEST_EMAIL")
	pw := os.Getenv("GOGARMIN_TEST_PASSWORD")
	if email == "" {
		t.Fatal("no test email found: set GARMIN_TEST_EMAIL")
		return nil
	}
	if pw == "" {
		t.Fatal("no test password found: set GARMIN_TEST_PASSWORD")
		return nil
	}
	_, accessToken, err := login(client, email, pw)
	if err != nil {
		t.Fatal(err)
		return nil
	}
	return accessToken
}

func TestFunctional_WeightLatest(t *testing.T) {
	t.Skip()
	weight, err := testapi(t).Weight.Latest(time.Now())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", weight)
	fmt.Println(weight.Weight / 453.592)
	fmt.Println(GramsToPounds(weight.Weight))
}

func TestFunctional_Activities(t *testing.T) {
	t.Skip()
	activities, err := testapi(t).ActivityList.Activities(new(ActivitySearch).
		WithActivityType("running").
		WithStart(0))
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range activities {
		fmt.Printf("%+v\n", a)
	}
}
