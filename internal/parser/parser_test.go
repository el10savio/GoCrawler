package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	for _, ValidateURLTest := range ValidateURLTestSuite {
		ValidateURLTest := ValidateURLTest
		t.Run(ValidateURLTest.Description, func(st *testing.T) {
			assert.Equal(
				st,
				ValidateURLTest.ExpectedError,
				ValidateURL(ValidateURLTest.URL),
			)
			st.Parallel()
		})
	}
}

func TestGetLinks(t *testing.T) {
	for _, GetLinksTest := range GetLinksTestSuite {
		GetLinksTest := GetLinksTest
		t.Run(GetLinksTest.Description, func(st *testing.T) {
			actualList, actualError := GetLinks(GetLinksTest.URL)
			assert.Equal(st, GetLinksTest.ExpectedList, actualList)
			assert.Equal(st, GetLinksTest.ExpectedError, actualError)
			st.Parallel()
		})
	}
}
