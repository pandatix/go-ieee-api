package goieeeapi_test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	goieeeapi "github.com/pandatix/go-ieee-api"
)

func TestGetKeywords(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client           goieeeapi.HTTPClient
		ID               int
		ExpectedKeywords *goieeeapi.GetKeywordsResponse
		ExpectedErr      error
	}{
		"nil-client": {
			Client:           nil,
			ID:               0,
			ExpectedKeywords: nil,
			ExpectedErr:      goieeeapi.ErrNilClient,
		},
		"failing-client": {
			Client:           newFakeHTTPClient(``, 0, errFake),
			ID:               0,
			ExpectedKeywords: nil,
			ExpectedErr:      errFake,
		},
		"unexpected-statuscode": {
			Client:           newFakeHTTPClient(``, 0, nil),
			ID:               0,
			ExpectedKeywords: nil,
			ExpectedErr: &goieeeapi.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:           newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			ID:               0,
			ExpectedKeywords: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"userInfo":{"institute":false,"member":false,"individual":false,"guest":false,"subscribedContent":false,"fileCabinetContent":false,"fileCabinetUser":false,"institutionalFileCabinetUser":false,"showPatentCitations":true,"showGet802Link":false,"showOpenUrlLink":false,"tracked":false,"delegatedAdmin":false,"desktop":false,"isInstitutionDashboardEnabled":false,"isInstitutionProfileEnabled":false,"isRoamingEnabled":false,"isDelegatedAdmin":false,"isMdl":false,"isCwg":false},"articleNumber":"8900441","getProgramTermsAccepted":false,"formulaStrippedArticleTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","pubLink":"/xpl/conhome/8891871/proceeding","keywords":[{"type":"IEEE Keywords","kwd":["Forestry","Training","Deep learning","Remote sensing","Vegetation mapping","Synthetic aperture radar","Task analysis"]},{"type":"INSPEC: Controlled Indexing","kwd":["computer vision","forestry","geophysical image processing","image classification","learning (artificial intelligence)","object detection","remote sensing"]},{"type":"INSPEC: Non-Controlled Indexing","kwd":["deep learning solutions","tandem-X-based forest classification","computer vision","discriminative tasks","image classification","remote sensing applications","domain-specific peculiarities","DL methods","RS tasks","TanDEM-X data","RS applications","object detection","nonforest classification problem","forest classification problem"]},{"type":"Author Keywords ","kwd":["Deep Learning","Convolutional Neural Network (CNN)","Vegetation Monitoring","Forest Classification","TanDEM-X"]}],"allowComments":false,"issueLink":"/xpl/tocresult.jsp?isnumber=null","isReadingRoomArticle":false,"isGetArticle":false,"isGetAddressInfoCaptured":false,"isMarketingOptIn":false,"publisher":"IEEE","xploreDocumentType":"Conference Publication","isPromo":false,"isNotDynamicOrStatic":true,"htmlAbstractLink":"/document/8900441/","isCustomDenial":false,"isSAE":false,"isDynamicHtml":false,"isFreeDocument":false,"displayDocTitle":"Deep Learning Solutions for Tandem-X-Based Forest Classification","isStandard":false,"isSMPTE":false,"isOUP":false,"isNow":false,"isProduct":false,"isMorganClaypool":false,"isJournal":false,"isBook":false,"isBookWithoutChapters":false,"isOpenAccess":false,"isEphemera":false,"isConference":true,"isChapter":false,"isStaticHtml":false,"isEarlyAccess":false,"persistentLink":"https://ieeexplore.ieee.org/servlet/opac?punumber=8891871","contentTypeDisplay":"Conferences","mlTime":"PT0.017138S","lastupdate":"2021-08-21","mediaPath":"/mediastore_new/IEEE/content/media/8891871/8897702/8900441","title":"Deep Learning Solutions for Tandem-X-Based Forest Classification","contentType":"conferences","publicationNumber":"8891871"}`, http.StatusOK, nil),
			ID:     0,
			ExpectedKeywords: &goieeeapi.GetKeywordsResponse{
				UserInfo: goieeeapi.UserInfo{
					Institute:                     false,
					Member:                        false,
					Individual:                    false,
					Guest:                         false,
					SubscribedContent:             false,
					FileCabinetContent:            false,
					FileCabinetUser:               false,
					InstitutionalFileCabinetUser:  false,
					ShowPatentCitations:           true,
					ShowGet802Link:                false,
					ShowOpenURLLink:               false,
					Tracked:                       false,
					DelegatedAdmin:                false,
					Desktop:                       false,
					IsInstitutionDashboardEnabled: false,
					IsInstitutionProfileEnabled:   false,
					IsRoamingEnabled:              false,
					IsDelegatedAdmin:              false,
					IsMdl:                         false,
					IsCwg:                         false,
				},
				ArticleNumber:               "8900441",
				GetProgramTermsAccepted:     false,
				FormulaStrippedArticleTitle: "Deep Learning Solutions for Tandem-X-Based Forest Classification",
				PubLink:                     "/xpl/conhome/8891871/proceeding",
				Keywords: []goieeeapi.Keyword{
					{
						Type: "IEEE Keywords",
						Kwd: []string{
							"Forestry",
							"Training",
							"Deep learning",
							"Remote sensing",
							"Vegetation mapping",
							"Synthetic aperture radar",
							"Task analysis",
						},
					}, {
						Type: "INSPEC: Controlled Indexing",
						Kwd: []string{
							"computer vision",
							"forestry",
							"geophysical image processing",
							"image classification",
							"learning (artificial intelligence)",
							"object detection",
							"remote sensing",
						},
					}, {
						Type: "INSPEC: Non-Controlled Indexing",
						Kwd: []string{
							"deep learning solutions",
							"tandem-X-based forest classification",
							"computer vision",
							"discriminative tasks",
							"image classification",
							"remote sensing applications",
							"domain-specific peculiarities",
							"DL methods",
							"RS tasks",
							"TanDEM-X data",
							"RS applications",
							"object detection",
							"nonforest classification problem",
							"forest classification problem",
						},
					}, {
						Type: "Author Keywords ",
						Kwd: []string{
							"Deep Learning",
							"Convolutional Neural Network (CNN)",
							"Vegetation Monitoring",
							"Forest Classification",
							"TanDEM-X",
						},
					},
				},
				AllowComments:            false,
				IssueLink:                "/xpl/tocresult.jsp?isnumber=null",
				IsReadingRoomArticle:     false,
				IsGetArticle:             false,
				IsGetAddressInfoCaptured: false,
				IsMarketingOptIn:         false,
				Publisher:                "IEEE",
				XploreDocumentType:       "Conference Publication",
				IsPromo:                  false,
				IsNotDynamicOrStatic:     true,
				HTMLAbstractLink:         "/document/8900441/",
				IsCustomDenial:           false,
				IsSAE:                    false,
				IsDynamicHTML:            false,
				IsFreeDocument:           false,
				DisplayDocTitle:          "Deep Learning Solutions for Tandem-X-Based Forest Classification",
				IsStandard:               false,
				IsSMPTE:                  false,
				IsOUP:                    false,
				IsNow:                    false,
				IsProduct:                false,
				IsMorganClaypool:         false,
				IsJournal:                false,
				IsBook:                   false,
				IsBookWithoutChapters:    false,
				IsOpenAccess:             false,
				IsEphemera:               false,
				IsConference:             true,
				IsChapter:                false,
				IsStaticHTML:             false,
				IsEarlyAccess:            false,
				PersistentLink:           "https://ieeexplore.ieee.org/servlet/opac?punumber=8891871",
				Title:                    "Deep Learning Solutions for Tandem-X-Based Forest Classification",
				ContentTypeDisplay:       "Conferences",
				MlTime:                   "PT0.017138S",
				LastUpdate:               "2021-08-21",
				MediaPath:                "/mediastore_new/IEEE/content/media/8891871/8897702/8900441",
				ContentType:              "conferences",
				PublicationNumber:        "8891871",
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			kwds, err := goieeeapi.GetKeywords(tt.Client, tt.ID)

			if !reflect.DeepEqual(kwds, tt.ExpectedKeywords) {
				t.Errorf("Failed to get expected keywords: got \"%v\" instead of \"%v\".", kwds, tt.ExpectedKeywords)
			}
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
