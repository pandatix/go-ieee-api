package api

import (
	"path"
	"strconv"
)

// GetDocumentReferences fetch the /document/<id>/references REST endpoint and
// returns the document references information.
func (client *IEEEClient) GetDocumentReferences(id int, params *GetDocumentReferencesParams, opts ...Option) (*GetReferencesResponse, error) {
	resp := &GetReferencesResponse{}
	err := client.get(path.Join("document", strconv.Itoa(id), "references"), params, resp, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetDocumentReferencesParams struct {
	Start *int `schema:"start,omitempty"`
	End   *int `schema:"end,omitempty"`
}

type GetReferencesResponse struct {
	UserInfo                    UserInfo    `json:"userInfo"`
	References                  []Reference `json:"references,omitempty"`
	ArticleNumber               *string     `json:"articleNumber,omitempty"`
	GetProgramTermsAccepted     bool        `json:"getProgramTermsAccepted"`
	FormulaStrippedArticleTitle *string     `json:"formulaStrippedArticleTitle,omitempty"`
	PubLink                     *string     `json:"pubLink,omitempty"`
	AllowComments               bool        `json:"allowComments"`
	IssueLink                   string      `json:"issueLink"`
	IsGetAddressInfoCaptured    bool        `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool        `json:"isMarketingOptIn"`
	Publisher                   *string     `json:"publisher,omitempty"`
	XploreDocumentType          *string     `json:"xploreDocumentType,omitempty"`
	IsPromo                     bool        `json:"isPromo"`
	HTMLAbstractLink            string      `json:"htmlAbstractLink"`
	IsCustomDenial              bool        `json:"isCustomDenial"`
	IsSAE                       bool        `json:"isSAE"`
	IsDynamicHTML               bool        `json:"isDynamicHtml"`
	IsFreeDocument              bool        `json:"isFreeDocument"`
	DisplayDocTitle             *string     `json:"displayDocTitle,omitempty"`
	IsStandard                  bool        `json:"isStandard"`
	IsSMPTE                     bool        `json:"isSMPTE"`
	IsOUP                       bool        `json:"isOUP"`
	IsNow                       bool        `json:"isNow"`
	IsProduct                   bool        `json:"isProduct"`
	IsJournal                   bool        `json:"isJournal"`
	IsBook                      bool        `json:"isBook"`
	IsBookWithoutChapters       bool        `json:"isBookWithoutChapters"`
	IsOpenAccess                bool        `json:"isOpenAccess"`
	IsEphemera                  bool        `json:"isEphemera"`
	IsConference                bool        `json:"isConference"`
	HTMLLink                    *string     `json:"htmlLink,omitempty"`
	IsChapter                   bool        `json:"isChapter"`
	IsEarlyAccess               bool        `json:"isEarlyAccess"`
	PersistentLink              *string     `json:"persistentLink,omitempty"`
	ArticleId                   *string     `json:"articleId,omitempty"`
	OpenAccessFlag              *string     `json:"openAccessFlag,omitempty"`
	EphemeraFlag                *string     `json:"ephemeraFlag,omitempty"`
	Title                       *string     `json:"title,omitempty"`
	ContentTypeDisplay          *string     `json:"contentTypeDisplay,omitempty"`
	HTMLFlag                    *string     `json:"html_flag,omitempty"`
	// HTMLFlagLegacy is only used by document 0. Seems to be legacy and hard-coded.
	HTMLFlagLegacy      *string `json:"htmlFlag,omitempty"`
	MlHTMLFlag          *string `json:"ml_html_flag,omitempty"`
	MlTime              string  `json:"mlTime"`
	LastUpdate          *string `json:"lastupdate,omitempty"`
	MediaPath           *string `json:"mediaPath,omitempty"`
	ContentType         *string `json:"contentType,omitempty"`
	Definitions         *string `json:"definitions,omitempty"`
	PublicationNumber   *string `json:"publicationNumber,omitempty"`
	HasStandardVersions bool    `json:"hasStandardVersions"`
	IsGiveaway          bool    `json:"isGiveaway"`
	IsLatestStandard    bool    `json:"isLatestStandard"`
	IsOnlineOnly        bool    `json:"isOnlineOnly"`
	IsSpringer          bool    `json:"isSpringer"`
	IsTranslation       bool    `json:"isTranslation"`
}

type Reference struct {
	Order             string    `json:"order"`
	Text              *string   `json:"text,omitempty"`
	DisplayText       *string   `json:"displayText,omitempty"`
	Title             *string   `json:"title,omitempty"`
	Context           []Context `json:"context,omitempty"`
	Links             *Links    `json:"links,omitempty"`
	GoogleScholarLink *string   `json:"googleScholarLink,omitempty"`
	RefType           *string   `json:"refType,omitempty"`
	ID                *string   `json:"id,omitempty"`
}

type Context struct {
	Sec  string  `json:"sec"`
	Text string  `json:"text"`
	Type *string `json:"type,omitempty"`
	Part *string `json:"part,omitempty"`
}

type Links struct {
	DocumentLink  *string `json:"documentLink,omitempty"`
	PdfLink       *string `json:"pdfLink,omitempty"`
	Abstract      *string `json:"abstract,omitempty"`
	CrossRefLink  *string `json:"crossRefLink,omitempty"`
	ArticleNumber *string `json:"articleNumber,omitempty"`
	OpenURLImgLoc string  `json:"openUrlImgLoc"`
	PdfSize       *string `json:"pdfSize,omitempty"`
}
