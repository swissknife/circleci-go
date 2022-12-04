package operations

type GetTestsPathParams struct {
	JobNumber   interface{} `pathParam:"style=simple,explode=false,name=job-number"`
	ProjectSlug string      `pathParam:"style=simple,explode=false,name=project-slug"`
}

type GetTestsTestsResponseItems struct {
	Classname string  `json:"classname"`
	File      string  `json:"file"`
	Message   string  `json:"message"`
	Name      string  `json:"name"`
	Result    string  `json:"result"`
	RunTime   float64 `json:"run_time"`
	Source    string  `json:"source"`
}

type GetTestsTestsResponse struct {
	Items         []GetTestsTestsResponseItems `json:"items"`
	NextPageToken string                       `json:"next_page_token"`
}

type GetTestsDefaultApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetTestsRequest struct {
	PathParams GetTestsPathParams
}

type GetTestsResponse struct {
	ContentType                          string
	StatusCode                           int64
	TestsResponse                        *GetTestsTestsResponse
	GetTestsDefaultApplicationJSONObject *GetTestsDefaultApplicationJSON
}
