package testing

import (
	"testing"

	"github.com/jerakia/go-jerakia"
	th "github.com/jerakia/go-jerakia/testhelper"
	fake "github.com/jerakia/go-jerakia/testhelper/client"

	"github.com/stretchr/testify/assert"
)

func TestLookupBasic(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLookupBasic(t)

	lookupOpts := &jerakia.LookupOpts{
		Namespace: "test",
	}

	actual, err := jerakia.Lookup(fake.FakeClient(), "cities", lookupOpts)
	if err != nil {
		t.Fatal(err)
	}

	expected := LookupBasicResult
	assert.Equal(t, expected, *actual)
}

func TestLookupSingleBool(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLookupSingleBool(t)

	lookupOpts := &jerakia.LookupOpts{
		Namespace: "test",
	}

	actual, err := jerakia.Lookup(fake.FakeClient(), "booltrue", lookupOpts)
	if err != nil {
		t.Fatal(err)
	}

	expected := LookupSingleBoolResult
	assert.Equal(t, expected, *actual)
}

func TestLookupMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleLookupMetadata(t)

	lookupOpts := &jerakia.LookupOpts{
		Namespace: "test",
		Metadata: map[string]string{
			"hostname": "example",
		},
	}

	actual, err := jerakia.Lookup(fake.FakeClient(), "users", lookupOpts)
	if err != nil {
		t.Fatal(err)
	}

	expected := LookupMetadataResult
	assert.Equal(t, expected, *actual)
}
