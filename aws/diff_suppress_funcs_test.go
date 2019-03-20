package aws

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

func TestSuppressEquivalentJsonDiffsWhitespaceAndNoWhitespace(t *testing.T) {
	d := new(schema.ResourceData)

	noWhitespace := `{"test":"test"}`
	whitespace := `
{
  "test": "test"
}`

	if !suppressEquivalentJsonDiffs("", noWhitespace, whitespace, d) {
		t.Errorf("Expected suppressEquivalentJsonDiffs to return true for %s == %s", noWhitespace, whitespace)
	}

	noWhitespaceDiff := `{"test":"test"}`
	whitespaceDiff := `
{
  "test": "tested"
}`

	if suppressEquivalentJsonDiffs("", noWhitespaceDiff, whitespaceDiff, d) {
		t.Errorf("Expected suppressEquivalentJsonDiffs to return false for %s == %s", noWhitespaceDiff, whitespaceDiff)
	}
}

func TestSuppressEquivalentTypeStringBoolean(t *testing.T) {
	testCases := []struct {
		old        string
		new        string
		equivalent bool
	}{
		{
			old:        "false",
			new:        "0",
			equivalent: true,
		},
		{
			old:        "true",
			new:        "1",
			equivalent: true,
		},
		{
			old:        "",
			new:        "0",
			equivalent: false,
		},
		{
			old:        "",
			new:        "1",
			equivalent: false,
		},
	}

	for i, tc := range testCases {
		value := suppressEquivalentTypeStringBoolean("test_property", tc.old, tc.new, nil)

		if tc.equivalent && !value {
			t.Fatalf("expected test case %d to be equivalent", i)
		}

		if !tc.equivalent && value {
			t.Fatalf("expected test case %d to not be equivalent", i)
		}
	}
}

func TestSuppressRoute53ZoneNameWithTrailingDot(t *testing.T) {
	testCases := []struct {
		old        string
		new        string
		equivalent bool
	}{
		{
			old:        "example.com",
			new:        "example.com",
			equivalent: true,
		},
		{
			old:        "example.com.",
			new:        "example.com.",
			equivalent: true,
		},
		{
			old:        "example.com.",
			new:        "example.com",
			equivalent: true,
		},
		{
			old:        "example.com",
			new:        "example.com.",
			equivalent: true,
		},
		{
			old:        ".",
			new:        "",
			equivalent: false,
		},
		{
			old:        "",
			new:        ".",
			equivalent: false,
		},
		{
			old:        ".",
			new:        ".",
			equivalent: true,
		},
	}

	for i, tc := range testCases {
		value := suppressRoute53ZoneNameWithTrailingDot("test_property", tc.old, tc.new, nil)

		if tc.equivalent && !value {
			t.Fatalf("expected test case %d to be equivalent", i)
		}

		if !tc.equivalent && value {
			t.Fatalf("expected test case %d to not be equivalent", i)
		}
	}
}

func TestSuppressStringsEqualFold(t *testing.T) {
	testCases := []struct {
		old        string
		new        string
		equivalent bool
	}{
		{
			old:        "ABC",
			new:        "ABC",
			equivalent: true,
		},
		{
			old:        "ABC",
			new:        "abC",
			equivalent: true,
		},
		{
			old:        "ABC",
			new:        "ab",
			equivalent: false,
		},
		{
			old:        "",
			new:        "ABC",
			equivalent: false,
		},
		{
			old:        "ABC",
			new:        "",
			equivalent: false,
		},
	}

	for i, tc := range testCases {
		value := suppressStringsEqualFold("test_property", tc.old, tc.new, nil)

		if tc.equivalent && !value {
			t.Fatalf("expected test case %d to be equivalent", i)
		}

		if !tc.equivalent && value {
			t.Fatalf("expected test case %d to not be equivalent", i)
		}
	}
}
