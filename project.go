package sdk

import (
	"context"
	"fmt"
	"github.com/swissknife/circle-go/pkg/models/operations"
	"github.com/swissknife/circle-go/pkg/utils"
	"net/http"
)

type Project struct {
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

func NewProject(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *Project {
	return &Project{
		_defaultClient:  defaultClient,
		_securityClient: securityClient,
		_serverURL:      serverURL,
		_language:       language,
		_sdkVersion:     sdkVersion,
		_genVersion:     genVersion,
	}
}

// CreateCheckoutKey - Create a new checkout key
// Creates a new checkout key. This API request is only usable with a user API token.
func (s *Project) CreateCheckoutKey(ctx context.Context, request operations.CreateCheckoutKeyRequest) (*operations.CreateCheckoutKeyResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/checkout-key", request.PathParams)

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

	res := &operations.CreateCheckoutKeyResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateCheckoutKeyCheckoutKey
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CheckoutKey = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateCheckoutKeyDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateCheckoutKeyDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// CreateEnvVar - Create an environment variable
// Creates a new environment variable.
func (s *Project) CreateEnvVar(ctx context.Context, request operations.CreateEnvVarRequest) (*operations.CreateEnvVarResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/envvar", request.PathParams)

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

	res := &operations.CreateEnvVarResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateEnvVarEnvironmentVariablePair
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.EnvironmentVariablePair = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateEnvVarDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateEnvVarDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// DeleteCheckoutKey - Delete a checkout key
// Deletes the checkout key.
func (s *Project) DeleteCheckoutKey(ctx context.Context, request operations.DeleteCheckoutKeyRequest) (*operations.DeleteCheckoutKeyResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/checkout-key/{fingerprint}", request.PathParams)

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

	res := &operations.DeleteCheckoutKeyResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteCheckoutKeyMessageResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.MessageResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteCheckoutKeyDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.DeleteCheckoutKeyDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// DeleteEnvVar - Delete an environment variable
// Deletes the environment variable named :name.
func (s *Project) DeleteEnvVar(ctx context.Context, request operations.DeleteEnvVarRequest) (*operations.DeleteEnvVarResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/envvar/{name}", request.PathParams)

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

	res := &operations.DeleteEnvVarResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteEnvVarMessageResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.MessageResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DeleteEnvVarDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.DeleteEnvVarDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetCheckoutKey - Get a checkout key
// Returns an individual checkout key.
func (s *Project) GetCheckoutKey(ctx context.Context, request operations.GetCheckoutKeyRequest) (*operations.GetCheckoutKeyResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/checkout-key/{fingerprint}", request.PathParams)

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

	res := &operations.GetCheckoutKeyResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetCheckoutKeyCheckoutKey
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CheckoutKey = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetCheckoutKeyDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetCheckoutKeyDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetEnvVar - Get a masked environment variable
// Returns the masked value of environment variable :name.
func (s *Project) GetEnvVar(ctx context.Context, request operations.GetEnvVarRequest) (*operations.GetEnvVarResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/envvar/{name}", request.PathParams)

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

	res := &operations.GetEnvVarResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetEnvVarEnvironmentVariablePair
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.EnvironmentVariablePair = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetEnvVarDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetEnvVarDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// GetProjectBySlug - Get a project
// Retrieves a project by project slug.
func (s *Project) GetProjectBySlug(ctx context.Context, request operations.GetProjectBySlugRequest) (*operations.GetProjectBySlugResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}", request.PathParams)

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

	res := &operations.GetProjectBySlugResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetProjectBySlugProject
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Project = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetProjectBySlugDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetProjectBySlugDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListCheckoutKeys - Get all checkout keys
// Returns a sequence of checkout keys for `:project`.
func (s *Project) ListCheckoutKeys(ctx context.Context, request operations.ListCheckoutKeysRequest) (*operations.ListCheckoutKeysResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/checkout-key", request.PathParams)

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

	res := &operations.ListCheckoutKeysResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListCheckoutKeysCheckoutKeyListResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CheckoutKeyListResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListCheckoutKeysDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListCheckoutKeysDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListEnvVars - List all environment variables
// Returns four 'x' characters, in addition to the last four ASCII characters of the value, consistent with the display of environment variable values on the CircleCI website.
func (s *Project) ListEnvVars(ctx context.Context, request operations.ListEnvVarsRequest) (*operations.ListEnvVarsResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/project/{project-slug}/envvar", request.PathParams)

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

	res := &operations.ListEnvVarsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListEnvVarsEnvironmentVariableListResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.EnvironmentVariableListResponse = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListEnvVarsDefaultApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListEnvVarsDefaultApplicationJSONObject = out
		}
	}

	return res, nil
}
