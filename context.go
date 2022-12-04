package sdk

import (
	"context"
	"fmt"
	"github.com/swissknife/circle-go/pkg/models/operations"
	"github.com/swissknife/circle-go/pkg/utils"
	"net/http"
	"strings"
)

type Context struct {
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

func NewContext(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *Context {
	return &Context{
		_defaultClient:  defaultClient,
		_securityClient: securityClient,
		_serverURL:      serverURL,
		_language:       language,
		_sdkVersion:     sdkVersion,
		_genVersion:     genVersion,
	}
}

// AddEnvironmentVariableToContext - Add or update an environment variable
// Create or update an environment variable within a context. Returns information about the environment variable, not including its value.
func (s *Context) AddEnvironmentVariableToContext(ctx context.Context, request operations.AddEnvironmentVariableToContextRequest) (*operations.AddEnvironmentVariableToContextResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/context/{context-id}/environment-variable/{env-var-name}", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
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

	res := &operations.AddEnvironmentVariableToContextResponse{
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

			res.AddEnvironmentVariableToContext200ApplicationJSONAnyOf = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.AddEnvironmentVariableToContextDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.AddEnvironmentVariableToContextDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// CreateContext - Create a new context
func (s *Context) CreateContext(ctx context.Context, request operations.CreateContextRequest) (*operations.CreateContextResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/context"

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

	res := &operations.CreateContextResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateContextContext
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Context = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateContextDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateContextDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// DeleteContext - Delete a context
func (s *Context) DeleteContext(ctx context.Context, request operations.DeleteContextRequest) (*operations.DeleteContextResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/context/{context-id}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
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

	res := &operations.DeleteContextResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteContextMessageResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.MessageResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteContextDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.DeleteContextDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// DeleteEnvironmentVariableFromContext - Remove an environment variable
// Delete an environment variable from a context.
func (s *Context) DeleteEnvironmentVariableFromContext(ctx context.Context, request operations.DeleteEnvironmentVariableFromContextRequest) (*operations.DeleteEnvironmentVariableFromContextResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/context/{context-id}/environment-variable/{env-var-name}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
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

	res := &operations.DeleteEnvironmentVariableFromContextResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteEnvironmentVariableFromContextMessageResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.MessageResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteEnvironmentVariableFromContextDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.DeleteEnvironmentVariableFromContextDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetContext - Get a context
// Returns basic information about a context.
func (s *Context) GetContext(ctx context.Context, request operations.GetContextRequest) (*operations.GetContextResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/context/{context-id}", request.PathParams)

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

	res := &operations.GetContextResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetContextContext
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Context = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetContextDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetContextDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListContexts - List contexts
// List all contexts for an owner.
func (s *Context) ListContexts(ctx context.Context, request operations.ListContextsRequest) (*operations.ListContextsResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/context"

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

	res := &operations.ListContextsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListContexts200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListContexts200ApplicationJSONObject = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListContextsDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListContextsDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListEnvironmentVariablesFromContext - List environment variables
// List information about environment variables in a context, not including their values.
func (s *Context) ListEnvironmentVariablesFromContext(ctx context.Context, request operations.ListEnvironmentVariablesFromContextRequest) (*operations.ListEnvironmentVariablesFromContextResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/context/{context-id}/environment-variable", request.PathParams)

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

	res := &operations.ListEnvironmentVariablesFromContextResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListEnvironmentVariablesFromContext200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListEnvironmentVariablesFromContext200ApplicationJSONObject = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListEnvironmentVariablesFromContextDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListEnvironmentVariablesFromContextDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}
