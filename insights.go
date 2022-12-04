package sdk

import (
	"context"
	"fmt"
	"github.com/swissknife/circle-go/pkg/models/operations"
	"github.com/swissknife/circle-go/pkg/utils"
	"net/http"
)

type Insights struct {
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

func NewInsights(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *Insights {
	return &Insights{
		_defaultClient:  defaultClient,
		_securityClient: securityClient,
		_serverURL:      serverURL,
		_language:       language,
		_sdkVersion:     sdkVersion,
		_genVersion:     genVersion,
	}
}

// GetAllInsightsBranches - Get all branches for a project
// Get a list of all branches for a specified project. The list will only contain branches currently available within Insights. The maximum number of branches returned by this endpoint is 5,000.
func (s *Insights) GetAllInsightsBranches(ctx context.Context, request operations.GetAllInsightsBranchesRequest) (*operations.GetAllInsightsBranchesResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/insights/{project-slug}/branches", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAllInsightsBranchesResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetAllInsightsBranches200ApplicationJSONAny = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetAllInsightsBranchesDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetAllInsightsBranchesDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetFlakyTests - Get flaky tests for a project
// Get a list of flaky tests for a given project. Flaky tests are branch agnostic.
//
//	A flaky test is a test that passed and failed in the same commit.
func (s *Insights) GetFlakyTests(ctx context.Context, request operations.GetFlakyTestsRequest) (*operations.GetFlakyTestsResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/insights/{project-slug}/flaky-tests", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetFlakyTestsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetFlakyTests200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetFlakyTests200ApplicationJSONObject = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetFlakyTestsDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetFlakyTestsDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetJobTimeseries - Job timeseries data
// Get timeseries data for all jobs within a workflow.
func (s *Insights) GetJobTimeseries(ctx context.Context, request operations.GetJobTimeseriesRequest) (*operations.GetJobTimeseriesResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/insights/time-series/{project-slug}/workflows/{workflow-name}/jobs", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetJobTimeseriesResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetJobTimeseries200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetJobTimeseries200ApplicationJSONObject = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetJobTimeseriesDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetJobTimeseriesDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetOrgSummaryData - Get summary metrics with trends for the entire org, and for each project.
// Gets aggregated summary metrics with trends for the entire org.
//
//	Also gets aggregated metrics and trends for each project belonging to the org.
func (s *Insights) GetOrgSummaryData(ctx context.Context, request operations.GetOrgSummaryDataRequest) (*operations.GetOrgSummaryDataResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/insights/{org-slug}/summary", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetOrgSummaryDataResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetOrgSummaryData200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetOrgSummaryData200ApplicationJSONObject = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetOrgSummaryDataDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetOrgSummaryDataDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetProjectWorkflowJobMetrics - Get summary metrics for a project workflow's jobs.
// Get summary metrics for a project workflow's jobs. Job runs going back at most 90 days are included in the aggregation window. Metrics are refreshed daily, and thus may not include executions from the last 24 hours. Please note that Insights is not a real time financial reporting tool and should not be used for credit reporting. The most up to date credit information can be found in Plan Overview in the CircleCI UI.
func (s *Insights) GetProjectWorkflowJobMetrics(ctx context.Context, request operations.GetProjectWorkflowJobMetricsRequest) (*operations.GetProjectWorkflowJobMetricsResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/insights/{project-slug}/workflows/{workflow-name}/jobs", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetProjectWorkflowJobMetricsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetProjectWorkflowJobMetrics200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetProjectWorkflowJobMetrics200ApplicationJSONObject = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetProjectWorkflowJobMetricsDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetProjectWorkflowJobMetricsDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetProjectWorkflowMetrics - Get summary metrics for a project's workflows
// Get summary metrics for a project's workflows. Workflow runs going back at most 90 days are included in the aggregation window. Metrics are refreshed daily, and thus may not include executions from the last 24 hours. Please note that Insights is not a real time financial reporting tool and should not be used for credit reporting. The most up to date credit information can be found in Plan Overview in the CircleCI UI.
func (s *Insights) GetProjectWorkflowMetrics(ctx context.Context, request operations.GetProjectWorkflowMetricsRequest) (*operations.GetProjectWorkflowMetricsResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/insights/{project-slug}/workflows", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetProjectWorkflowMetricsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetProjectWorkflowMetrics200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetProjectWorkflowMetrics200ApplicationJSONObject = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetProjectWorkflowMetricsDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetProjectWorkflowMetricsDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetProjectWorkflowRuns - Get recent runs of a workflow
// Get recent runs of a workflow. Runs going back at most 90 days are returned. Please note that Insights is not a real time financial reporting tool and should not be used for credit reporting. The most up to date credit information can be found in Plan Overview in the CircleCI UI.
func (s *Insights) GetProjectWorkflowRuns(ctx context.Context, request operations.GetProjectWorkflowRunsRequest) (*operations.GetProjectWorkflowRunsResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/insights/{project-slug}/workflows/{workflow-name}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetProjectWorkflowRunsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetProjectWorkflowRuns200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetProjectWorkflowRuns200ApplicationJSONObject = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetProjectWorkflowRunsDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetProjectWorkflowRunsDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetProjectWorkflowTestMetrics - Get test metrics for a project's workflows
// Get test metrics for a project's workflows. Currently tests metrics are calculated based on 10 most recent workflow runs.
func (s *Insights) GetProjectWorkflowTestMetrics(ctx context.Context, request operations.GetProjectWorkflowTestMetricsRequest) (*operations.GetProjectWorkflowTestMetricsResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/insights/{project-slug}/workflows/{workflow-name}/test-metrics", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetProjectWorkflowTestMetricsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetProjectWorkflowTestMetrics200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetProjectWorkflowTestMetrics200ApplicationJSONObject = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetProjectWorkflowTestMetricsDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetProjectWorkflowTestMetricsDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetProjectWorkflowsPageData - Get summary metrics and trends for a project across it's workflows and branches
// Get summary metrics and trends for a project at workflow and branch level.
//
//	Workflow runs going back at most 90 days are included in the aggregation window.
//	Trends are only supported upto last 30 days.
//	Metrics are refreshed daily, and thus may not include executions from the last 24 hours.
//	Please note that Insights is not a real time financial reporting tool and should not be used for credit reporting.
//	The most up to date credit information can be found in Plan Overview in the CircleCI UI.
func (s *Insights) GetProjectWorkflowsPageData(ctx context.Context, request operations.GetProjectWorkflowsPageDataRequest) (*operations.GetProjectWorkflowsPageDataResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/insights/pages/{project-slug}/summary", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetProjectWorkflowsPageDataResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetProjectWorkflowsPageData200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetProjectWorkflowsPageData200ApplicationJSONObject = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetProjectWorkflowsPageDataDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetProjectWorkflowsPageDataDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetWorkflowSummary - Get metrics and trends for workflows
// Get the metrics and trends for a particular workflow on a single branch or all branches
func (s *Insights) GetWorkflowSummary(ctx context.Context, request operations.GetWorkflowSummaryRequest) (*operations.GetWorkflowSummaryResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/insights/{project-slug}/workflows/{workflow-name}/summary", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetWorkflowSummaryResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetWorkflowSummary200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetWorkflowSummary200ApplicationJSONObject = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetWorkflowSummaryDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetWorkflowSummaryDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}
