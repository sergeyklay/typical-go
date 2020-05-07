package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestCreateAnnotation(t *testing.T) {
	testcases := []struct {
		testName      string
		decl          *typast.Decl
		raw           string
		expected      *typast.Annot
		expectedError string
	}{
		{
			testName: "tag only",
			raw:      `@autowire`,
			expected: &typast.Annot{
				TagName: "autowire",
			},
		},
		{
			testName: "tag only with space",
			raw:      `@  autowire  `,
			expected: &typast.Annot{
				TagName: "autowire",
			},
		},
		{
			testName: "with attribute",
			raw:      `@mock{"pkg":"mock2"}`,
			expected: &typast.Annot{
				TagName:  "mock",
				TagAttrs: []byte(`{"pkg":"mock2"}`),
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			annotation, err := typast.CreateAnnot(tt.decl, tt.raw)
			if tt.expectedError != "" {
				require.EqualError(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.expected, annotation)
		})
	}
}

func TestUnmarshall(t *testing.T) {
	testcases := []struct {
		testName string
		*typast.Annot
		expected    map[string]string
		expectedErr string
	}{
		{
			testName: "",
			Annot: &typast.Annot{
				TagName:  "mock",
				TagAttrs: []byte(`{"key1":"value1"}`),
			},
			expected: map[string]string{
				"key1": "value1",
			},
		},
		{
			testName: "",
			Annot: &typast.Annot{
				TagName: "mock",
			},
		},
		{
			testName: "",
			Annot: &typast.Annot{
				TagName:  "mock",
				TagAttrs: []byte(`{"key1":"value1"`),
			},
			expectedErr: "unexpected end of JSON input",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			var m map[string]string
			err := tt.Unmarshal(&m)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.expected, m)

		})
	}
}
