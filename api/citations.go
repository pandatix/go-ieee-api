package api

import (
	"path"
	"strconv"
)

// GetDocumentCitations fetch the /document/<id>/citations REST endpoint and
// returns the document citations information.
func (client *IEEEClient) GetDocumentCitations(id int, opts ...Option) (*GetDocumentCitationsResponse, error) {
	resp := &GetDocumentCitationsResponse{}
	err := client.get(path.Join("document", strconv.Itoa(id), "citations"), nil, resp, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetDocumentCitationsResponse struct {
	GetProgramTermsAccepted     bool             `json:"getProgramTermsAccepted"`
	FormulaStrippedArticleTitle *string          `json:"formulaStrippedArticleTitle,omitempty"`
	PaperCitations              *PaperCitations  `json:"paperCitations,omitempty"`
	IsGetAddressInfoCaptured    bool             `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn            bool             `json:"isMarketingOptIn"`
	IsProduct                   bool             `json:"isProduct"`
	IsChapter                   bool             `json:"isChapter"`
	IsEarlyAccess               bool             `json:"isEarlyAccess"`
	NonIeeeCitationCount        *string          `json:"nonIeeeCitationCount,omitempty"`
	ContentTypeDisplay          *string          `json:"contentTypeDisplay,omitempty"`
	PatentCitationCount         string           `json:"patentCitationCount"`
	PatentCitations             []PatentCitation `json:"patentCitations,omitempty"`
	MlTime                      *string          `json:"mlTime,omitempty"`
	LastUpdate                  *string          `json:"lastupdate,omitempty"`
	MediaPath                   *string          `json:"mediaPath,omitempty"`
	Title                       *string          `json:"title,omitempty"`
	ContentType                 *string          `json:"contentType,omitempty"`
	IEEECitationCount           *string          `json:"ieeeCitationCount,omitempty"`
	PublicationNumber           *string          `json:"publicationNumber,omitempty"`
	HTMLFlagLegacy              *string          `json:"htmlFlag,omitempty"`
}

type PaperCitations struct {
	IEEE    []Reference `json:"ieee,omitempty"`
	NonIEEE []Reference `json:"nonIeee,omitempty"`
}

type PatentCitation struct {
	AppNum            string   `json:"appNum"`
	Assignees         []string `json:"assignees"`
	FilingDate        string   `json:"filingDate"`
	GoogleScholarLink string   `json:"googleScholarLink"`
	GrantDate         string   `json:"grantDate"`
	ID                string   `json:"id"`
	Inventors         string   `json:"inventors"`
	IPCClassList      string   `json:"ipcClassList"`
	IPCClasses        []string `json:"ipcClasses"`
	Order             string   `json:"order"`
	PatentAbstract    string   `json:"patentAbstract"`
	PatentLink        string   `json:"patentLink"`
	PatentNumber      string   `json:"patentNumber"`
	PdfLink           *string  `json:"pdfLink,omitempty"`
	PocClassList      *string  `json:"pocClassList,omitempty"`
	PocClasses        []string `json:"pocClasses,omitempty"`
	Title             string   `json:"title"`
}
