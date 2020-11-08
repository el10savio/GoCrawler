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
	GetLinksTestCase{
		Description:   "http://www.monzo.com/",
		URL:           "http://www.monzo.com/",
		ExpectedList:  []string{"http://www.monzo.com//i/business", "http://www.monzo.com//i/current-account", "http://www.monzo.com//i/monzo-plus", "http://www.monzo.com//i/business", "http://www.monzo.com//features/joint-accounts", "http://www.monzo.com//features/16-plus", "http://www.monzo.com//features/switch", "http://www.monzo.com//i/savingwithmonzo", "http://www.monzo.com//features/savings", "http://www.monzo.com//isa", "http://www.monzo.com//i/overdrafts", "http://www.monzo.com//i/loans", "http://www.monzo.com//blog/2019/11/12/what-are-unsecured-loans", "http://www.monzo.com//features/travel", "http://www.monzo.com//features/energy-switching", "http://www.monzo.com//i/shared-tabs-more", "http://www.monzo.com//community/making-monzo", "http://www.monzo.com//help", "http://www.monzo.com//i/coronavirus-faq", "http://www.monzo.com//i/monzo-plus", "http://www.monzo.com//features/savings", "http://www.monzo.com//features/travel", "http://www.monzo.com//i/loans", "http://www.monzo.com//i/security", "http://www.monzo.com//service-quality-results", "http://www.monzo.com//about", "http://www.monzo.com//usa", "http://www.monzo.com//blog", "http://www.monzo.com//press", "http://www.monzo.com//careers", "http://www.monzo.com//i/socialimpact", "http://www.monzo.com//i/community", "http://www.monzo.com//community/making-monzo", "http://www.monzo.com//transparency", "http://www.monzo.com//i/coronavirus-update", "http://www.monzo.com//i/fraud", "http://www.monzo.com//tone-of-voice", "http://www.monzo.com//i/business", "http://www.monzo.com//static/docs/modern-slavery-statement/modern-slavery-statement-2020.pdf", "http://www.monzo.com//faq", "http://www.monzo.com//legal/terms-and-conditions", "http://www.monzo.com//legal/fscs-information", "http://www.monzo.com//legal/privacy-notice", "http://www.monzo.com//legal/cookie-notice", "http://www.monzo.com//information-about-current-account-services", "http://www.monzo.com//service-information"},
		ExpectedError: nil,
	},
}
