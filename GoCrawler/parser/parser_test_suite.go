package parser

// ValidateURLTestCase is the test case
// struct definition for ValidateURL()
type ValidateURLTestCase struct {
	Description   string
	URL           string
	ExpectedError error
}

// GetLinksTestCase is the test case
// struct definition for GetLinks()
type GetLinksTestCase struct {
	Description   string
	URL           string
	ExpectedList  []string
	ExpectedError error
}

// ValidateURLTestSuite is the test suite cases for ValidateURL()
var ValidateURLTestSuite = []ValidateURLTestCase{
	ValidateURLTestCase{
		Description:   "Validate Valid URL",
		URL:           "http://www.google.com",
		ExpectedError: nil,
	},
	ValidateURLTestCase{
		Description:   "Validate URL With Port",
		URL:           "https://www.google.com:443",
		ExpectedError: nil,
	},
	ValidateURLTestCase{
		Description:   "Validate Empty URL",
		URL:           "",
		ExpectedError: ErrEmptyURL,
	},
	ValidateURLTestCase{
		Description:   "Validate URL With Only Path",
		URL:           "/testing-path",
		ExpectedError: ErrEmptyURLHost,
	},
	ValidateURLTestCase{
		Description:   "Validate Empty Host URL",
		URL:           "https://",
		ExpectedError: ErrEmptyURLHost,
	},
	ValidateURLTestCase{
		Description:   "Invalid Case",
		URL:           "alskjff#?asf//dfas",
		ExpectedError: ErrInvalidURL,
	},
	ValidateURLTestCase{
		Description:   "Invalid Case",
		URL:           "https",
		ExpectedError: ErrInvalidURL,
	},
	ValidateURLTestCase{
		Description:   "Invalid Case",
		URL:           "google",
		ExpectedError: ErrInvalidURL,
	},
	ValidateURLTestCase{
		Description:   "Invalid Case",
		URL:           "google.com",
		ExpectedError: ErrInvalidURL,
	},
	ValidateURLTestCase{
		Description:   "Invalid Case",
		URL:           "testing-path",
		ExpectedError: ErrInvalidURL,
	},
}

// GetLinksTestSuite is the test suite cases for GetLinks()
var GetLinksTestSuite = []GetLinksTestCase{
	GetLinksTestCase{
		Description:   "http://blank.org/",
		URL:           "http://blank.org/",
		ExpectedList:  []string{"http://blank.org/blank.html"},
		ExpectedError: nil,
	},
	GetLinksTestCase{
		Description:   "http://www.blankwebsite.com/",
		URL:           "http://www.blankwebsite.com/",
		ExpectedList:  []string{},
		ExpectedError: nil,
	},
}
