package api

import (
	"path"
	"strconv"
)

// GetDocumentAbstract fetch the /document/<id>/abstract REST endpoint and
// returns the document abstract information.
func (client *IEEEClient) GetDocumentAbstract(id int, opts ...Option) (*GetDocumentAbstractResponse, error) {
	resp := &GetDocumentAbstractResponse{}
	err := client.get(path.Join("document", strconv.Itoa(id), "abstract"), nil, resp, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetDocumentAbstractResponse struct {
	UserInfo                    UserInfo           `json:"userInfo"`
	Authors                     []Author           `json:"authors,omitempty"`
	ISBN                        []ISBN             `json:"isbn,omitempty"`
	ISSN                        []ISSN             `json:"issn,omitempty"`
	ArticleNumber               string             `json:"articleNumber"`
	AuthorNames                 *string            `json:"authorNames,omitempty"`
	GetProgramTermsAccepted     bool               `json:"getProgramTermsAccepted"`
	GraphicalAbstract           *GraphicalAbstract `json:"graphicalAbstract,omitempty"`
	Sections                    *Sections          `json:"sections,omitempty"`
	PdfURL                      *string            `json:"pdfUrl,omitempty"`
	Keywords                    []Keyword          `json:"keywords,omitempty"`
	PubLink                     *string            `json:"pubLink,omitempty"`
	AllowComments               bool               `json:"allowComments"`
	ArticleCopyright            *string            `json:"articleCopyRight,omitempty"`
	Abstract                    *string            `json:"abstract,omitempty"`
	DOI                         *string            `json:"doi,omitempty"`
	DOILink                     *string            `json:"doiLink,omitempty"`
	RightsLink                  *string            `json:"rightsLink,omitempty"`
	StartPage                   *string            `json:"startPage,omitempty"`
	EndPage                     *string            `json:"endPage,omitempty"`
	PublicationTitle            *string            `json:"publicationTitle,omitempty"`
	DisplayPublicationDate      *string            `json:"displayPublicationDate,omitempty"`
	PublicationYear             *string            `json:"publicationYear,omitempty"`
	DisplayPublicationTitle     *string            `json:"displayPublicationTitle,omitempty"`
	PdfPath                     *string            `json:"pdfPath,omitempty"`
	Issue                       *string            `json:"issue,omitempty"`
	IssueLink                   string             `json:"issueLink"`
	FormulaStrippedArticleTitle *string            `json:"formulaStrippedArticleTitle,omitempty"`
	FundingAgencies             *FundingAgencies   `json:"fundingAgencies,omitempty"`
	IsGiveaway                  bool               `json:"isGiveaway"`
	IsGetAddressInfoCaptured    bool               `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool               `json:"isMarketingOptIn"`
	PubTopics                   []PubTopic         `json:"pubTopics,omitempty"`
	ReferenceCount              *int               `json:"referenceCount,omitempty"`
	Publisher                   *string            `json:"publisher,omitempty"`
	IsDynamicHTML               bool               `json:"isDynamicHtml"`
	IsFreeDocument              bool               `json:"isFreeDocument"`
	HasStandardVersions         *bool              `json:"hasStandardVersions,omitempty"`
	HTMLAbstractLink            string             `json:"htmlAbstractLink"`
	IsCustomDenial              bool               `json:"isCustomDenial"`
	IsSMPTE                     bool               `json:"isSMPTE"`
	IsSpringer                  bool               `json:"isSpringer"`
	IsTranslation               bool               `json:"isTranslation"`
	IsOUP                       bool               `json:"isOUP"`
	IsOnlineOnly                bool               `json:"isOnlineOnly"`
	ConferenceDate              *string            `json:"conferenceDate,omitempty"`
	ChronOrPublicationDate      *string            `json:"chronOrPublicationDate,omitempty"`
	CitationCount               *string            `json:"citationCount,omitempty"`
	IsNow                       bool               `json:"isNow"`
	IsSAE                       bool               `json:"isSAE"`
	IsPromo                     bool               `json:"isPromo"`
	DateOfInsertion             *string            `json:"dateOfInsertion,omitempty"`
	DisplayDocTitle             *string            `json:"displayDocTitle,omitempty"`
	IsStandard                  bool               `json:"isStandard"`
	IsConference                bool               `json:"isConference"`
	IsProduct                   bool               `json:"isProduct"`
	IsJournal                   bool               `json:"isJournal"`
	IsLatestStandard            bool               `json:"isLatestStandard"`
	InsertDate                  *string            `json:"insertDate,omitempty"`
	IsBook                      bool               `json:"isBook"`
	IsBookWithoutChapters       bool               `json:"isBookWithoutChapters"`
	IsEarlyAccess               bool               `json:"isEarlyAccess"`
	AccessionNumber             *string            `json:"accessionNumber,omitempty"`
	IsOpenAccess                bool               `json:"isOpenAccess"`
	PublicationDate             *string            `json:"publicationDate,omitempty"`
	IsEphemera                  bool               `json:"isEphemera"`
	HTMLLink                    *string            `json:"htmlLink,omitempty"`
	PersistentLink              *string            `json:"persistentLink,omitempty"`
	IsChapter                   bool               `json:"isChapter"`
	XploreDocumentType          *string            `json:"xploreDocumentType,omitempty"`
	OpenAccessFlag              *string            `json:"openAccessFlag,omitempty"`
	EphemeraFlag                *string            `json:"ephemeraFlag,omitempty"`
	Title                       *string            `json:"title,omitempty"`
	ConfLoc                     *string            `json:"confLoc,omitempty"`
	HTMLFlag                    *string            `json:"html_flag,omitempty"`
	MlHTMLFlag                  *string            `json:"ml_html_flag,omitempty"`
	SourcePdf                   *string            `json:"sourcePdf,omitempty"`
	MlTime                      string             `json:"mlTime"`
	XplorePubID                 *string            `json:"xplore-pub-id,omitempty"`
	IsNumber                    *string            `json:"isNumber,omitempty"`
	RightsLinkFlag              *string            `json:"rightsLinkFlag,omitempty"`
	ContentType                 *string            `json:"contentType,omitempty"`
	PublicationNumber           *string            `json:"publicationNumber,omitempty"`
	Volume                      *string            `json:"volume,omitempty"`
	XploreIssue                 *string            `json:"xplore-issue,omitempty"`
	ArticleID                   *string            `json:"articleId,omitempty"`
	ContentTypeDisplay          *string            `json:"contentTypeDisplay,omitempty"`
	SubType                     *string            `json:"subType,omitempty"`
	Value                       *string            `json:"_value,omitempty"`
	LastUpdate                  *string            `json:"lastupdate,omitempty"`
	MediaPath                   *string            `json:"mediaPath,omitempty"`
	HTMLFlagLegacy              *string            `json:"htmlFlag,omitempty"`
	Sponsors                    []Sponsor          `json:"sponsors,omitempty"`
	SubjectCategories           []string           `json:"subjectCategories,omitempty"`
}

type ISBN struct {
	Format   string `json:"format"`
	Value    string `json:"value"`
	ISBNType string `json:"isbnType"`
}

type ISSN struct {
	Format string `json:"format"`
	Value  string `json:"value"`
}

type Sections struct {
	Abstract       string `json:"abstract"`
	Authors        string `json:"authors"`
	Figures        string `json:"figures"`
	Multimedia     string `json:"multimedia"`
	References     string `json:"references"`
	CitedBy        string `json:"citedby"`
	Keywords       string `json:"keywords"`
	Definitions    string `json:"definitions"`
	Algorithm      string `json:"algorithm"`
	Dataset        string `json:"dataset"`
	Cadmore        string `json:"cadmore"`
	Footnotes      string `json:"footnotes"`
	Disclaimer     string `json:"disclaimer"`
	RelatedContent string `json:"relatedContent"`
}

type PubTopic struct {
	Name string `json:"name"`
}

type FundingAgencies struct {
	FundingAgency []FundingAgency `json:"fundingAgency"`
}

type FundingAgency struct {
	FundingID   *string `json:"fundingId,omitempty"`
	FundingName string  `json:"fundingName"`
	GrantNumber *string `json:"grantNumber,omitempty"`
}

type GraphicalAbstract struct {
	CoverImage *string `json:"coverImage,omitempty"`
	File       *string `json:"file,omitempty"`
	FileSize   *string `json:"fileSize,omitempty"`
	Summary    *string `json:"summary,omitempty"`
	Type       *string `json:"type,omitempty"`
}

type Sponsor struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
