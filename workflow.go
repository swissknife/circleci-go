package sdk

import (
	"context"
	"fmt"
	"github.com/swissknife/circle-go/pkg/models/operations"
	"github.com/swissknife/circle-go/pkg/utils"
	"net/http"
)

type Workflow struct {
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

func NewWorkflow(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *Workflow {
	return &Workflow{
		_defaultClient:  defaultClient,
		_securityClient: securityClient,
		_serverURL:      serverURL,
		_language:       language,
		_sdkVersion:     sdkVersion,
		_genVersion:     genVersion,
	}
}

// ApprovePendingApprovalJobByID - Approve a job
// Approves a pending approval job in a workflow.
func (s *Workflow) ApprovePendingApprovalJobByID(ctx context.Context, request operations.ApprovePendingApprovalJobByIDRequest) (*operations.ApprovePendingApprovalJobByIDResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/workflow/{id}/approve/{approval_request_id}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
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

	res := &operations.ApprovePendingApprovalJobByIDResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 202:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ApprovePendingApprovalJobByIDMessageResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.MessageResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ApprovePendingApprovalJobByIDDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ApprovePendingApprovalJobByIDDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// CancelWorkflow - Cancel a workflow
// Cancels a running workflow.
func (s *Workflow) CancelWorkflow(ctx context.Context, request operations.CancelWorkflowRequest) (*operations.CancelWorkflowResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/workflow/{id}/cancel", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
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

	res := &operations.CancelWorkflowResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 202:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CancelWorkflowMessageResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.MessageResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CancelWorkflowDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CancelWorkflowDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetWorkflowByID - Get a workflow
// Returns summary fields of a workflow by ID.
func (s *Workflow) GetWorkflowByID(ctx context.Context, request operations.GetWorkflowByIDRequest) (*operations.GetWorkflowByIDResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/workflow/{id}", request.PathParams)

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

	res := &operations.GetWorkflowByIDResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetWorkflowByIDWorkflow
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Workflow = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetWorkflowByIDDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetWorkflowByIDDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListWorkflowJobs - Get a workflow's jobs
// Returns a sequence of jobs for a workflow.
func (s *Workflow) ListWorkflowJobs(ctx context.Context, request operations.ListWorkflowJobsRequest) (*operations.ListWorkflowJobsResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/workflow/{id}/job", request.PathParams)

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

	res := &operations.ListWorkflowJobsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListWorkflowJobsWorkflowJobListResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.WorkflowJobListResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListWorkflowJobsDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListWorkflowJobsDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// RerunWorkflow - Rerun a workflow
// Reruns a workflow.
func (s *Workflow) RerunWorkflow(ctx context.Context, request operations.RerunWorkflowRequest) (*operations.RerunWorkflowResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/workflow/{id}/rerun", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.RerunWorkflowResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 202:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.RerunWorkflow202ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.RerunWorkflow202ApplicationJSONObject = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.RerunWorkflowDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.RerunWorkflowDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}
