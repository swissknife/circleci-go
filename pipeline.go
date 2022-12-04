package sdk

import (
	"context"
	"fmt"
	"github.com/swissknife/circle-go/pkg/models/operations"
	"github.com/swissknife/circle-go/pkg/utils"
	"net/http"
	"strings"
)

type Pipeline struct {
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

func NewPipeline(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *Pipeline {
	return &Pipeline{
		_defaultClient:  defaultClient,
		_securityClient: securityClient,
		_serverURL:      serverURL,
		_language:       language,
		_sdkVersion:     sdkVersion,
		_genVersion:     genVersion,
	}
}

// ContinuePipeline - Continue a pipeline
// Continue a pipeline from the setup phase.
func (s *Pipeline) ContinuePipeline(ctx context.Context, request operations.ContinuePipelineRequest) (*operations.ContinuePipelineResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/pipeline/continue"

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

	res := &operations.ContinuePipelineResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ContinuePipelineMessageResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.MessageResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ContinuePipelineDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ContinuePipelineDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetPipelineByID - Get a pipeline by ID
// Returns a pipeline by the pipeline ID.
func (s *Pipeline) GetPipelineByID(ctx context.Context, request operations.GetPipelineByIDRequest) (*operations.GetPipelineByIDResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/pipeline/{pipeline-id}", request.PathParams)

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

	res := &operations.GetPipelineByIDResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetPipelineByIDPipeline
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Pipeline = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetPipelineByIDDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetPipelineByIDDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetPipelineByNumber - Get a pipeline by pipeline number
// Returns a pipeline by the pipeline number.
func (s *Pipeline) GetPipelineByNumber(ctx context.Context, request operations.GetPipelineByNumberRequest) (*operations.GetPipelineByNumberResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/pipeline/{pipeline-number}", request.PathParams)

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

	res := &operations.GetPipelineByNumberResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetPipelineByNumberPipeline
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Pipeline = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetPipelineByNumberDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetPipelineByNumberDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetPipelineConfigByID - Get a pipeline's configuration
// Returns a pipeline's configuration by ID.
func (s *Pipeline) GetPipelineConfigByID(ctx context.Context, request operations.GetPipelineConfigByIDRequest) (*operations.GetPipelineConfigByIDResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/pipeline/{pipeline-id}/config", request.PathParams)

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

	res := &operations.GetPipelineConfigByIDResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetPipelineConfigByIDPipelineConfig
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PipelineConfig = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetPipelineConfigByIDDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetPipelineConfigByIDDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListMyPipelines - Get your pipelines
// Returns a sequence of all pipelines for this project triggered by the user.
func (s *Pipeline) ListMyPipelines(ctx context.Context, request operations.ListMyPipelinesRequest) (*operations.ListMyPipelinesResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/pipeline/mine", request.PathParams)

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

	res := &operations.ListMyPipelinesResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListMyPipelinesPipelineListResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PipelineListResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListMyPipelinesDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListMyPipelinesDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListPipelines - Get a list of pipelines
// Returns all pipelines for the most recently built projects (max 250) you follow in an organization.
func (s *Pipeline) ListPipelines(ctx context.Context, request operations.ListPipelinesRequest) (*operations.ListPipelinesResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/pipeline"

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

	res := &operations.ListPipelinesResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListPipelinesPipelineListResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PipelineListResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListPipelinesDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListPipelinesDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListPipelinesForProject - Get all pipelines
// Returns all pipelines for this project.
func (s *Pipeline) ListPipelinesForProject(ctx context.Context, request operations.ListPipelinesForProjectRequest) (*operations.ListPipelinesForProjectResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/pipeline", request.PathParams)

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

	res := &operations.ListPipelinesForProjectResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListPipelinesForProjectPipelineListResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PipelineListResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListPipelinesForProjectDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListPipelinesForProjectDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListWorkflowsByPipelineID - Get a pipeline's workflows
// Returns a paginated list of workflows by pipeline ID.
func (s *Pipeline) ListWorkflowsByPipelineID(ctx context.Context, request operations.ListWorkflowsByPipelineIDRequest) (*operations.ListWorkflowsByPipelineIDResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/pipeline/{pipeline-id}/workflow", request.PathParams)

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

	res := &operations.ListWorkflowsByPipelineIDResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListWorkflowsByPipelineIDWorkflowListResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.WorkflowListResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListWorkflowsByPipelineIDDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListWorkflowsByPipelineIDDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// TriggerPipeline - Trigger a new pipeline
// Triggers a new pipeline on the project.
func (s *Pipeline) TriggerPipeline(ctx context.Context, request operations.TriggerPipelineRequest) (*operations.TriggerPipelineResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/pipeline", request.PathParams)

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

	res := &operations.TriggerPipelineResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.TriggerPipelinePipelineCreation
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PipelineCreation = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.TriggerPipelineDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.TriggerPipelineDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}
