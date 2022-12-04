package sdk

import (
	"context"
	"fmt"
	"github.com/swissknife/circle-go/pkg/models/operations"
	"github.com/swissknife/circle-go/pkg/utils"
	"net/http"
)

type Job struct {
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

func NewJob(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *Job {
	return &Job{
		_defaultClient:  defaultClient,
		_securityClient: securityClient,
		_serverURL:      serverURL,
		_language:       language,
		_sdkVersion:     sdkVersion,
		_genVersion:     genVersion,
	}
}

// CancelJob - Cancel job
// Cancel job with a given job number.
func (s *Job) CancelJob(ctx context.Context, request operations.CancelJobRequest) (*operations.CancelJobResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/job/{job-number}/cancel", request.PathParams)

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

	res := &operations.CancelJobResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CancelJobMessageResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.MessageResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CancelJobDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CancelJobDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetJobArtifacts - Get a job's artifacts
// Returns a job's artifacts.
func (s *Job) GetJobArtifacts(ctx context.Context, request operations.GetJobArtifactsRequest) (*operations.GetJobArtifactsResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/{job-number}/artifacts", request.PathParams)

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

	res := &operations.GetJobArtifactsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetJobArtifactsArtifactListResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ArtifactListResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetJobArtifactsDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetJobArtifactsDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetJobDetails - Get job details
// Returns job details.
func (s *Job) GetJobDetails(ctx context.Context, request operations.GetJobDetailsRequest) (*operations.GetJobDetailsResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/job/{job-number}", request.PathParams)

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

	res := &operations.GetJobDetailsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetJobDetailsJobDetails
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.JobDetails = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetJobDetailsDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetJobDetailsDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetTests - Get test metadata
// Get test metadata for a build. In the rare case where there is more than 250MB of test data on the job, no results will be returned.
func (s *Job) GetTests(ctx context.Context, request operations.GetTestsRequest) (*operations.GetTestsResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/{job-number}/tests", request.PathParams)

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

	res := &operations.GetTestsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetTestsTestsResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.TestsResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetTestsDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetTestsDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}
