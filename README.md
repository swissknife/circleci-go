# CircleCI API Go SDK

The CircleCI API Go SDK is a Golang package for accessing the resources that make up the [CircleCI API V2](https://circleci.com/docs/api/v2/index.html).

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/swissknife/circle-go
```
<!-- End SDK Installation -->

<!-- Start SDK Example Usage -->
## SDK Example Usage

```go
package main

import (
    "github.com/swissknife/circle-go"
    "github.com/swissknife/circle-go/pkg/models/shared"
    "github.com/swissknife/circle-go/pkg/models/operations"
)

func main() {
    opts := []sdk.SDKOption{
        sdk.WithSecurity(
            shared.Security{
                APIKeyHeader: &shared.SchemeAPIKeyHeader{
                    APIKey: "YOUR_API_KEY_HERE",
                },
            }
        ),
    }

    s := sdk.New(opts...)
    
    req := operations.AddEnvironmentVariableToContextRequest{
        PathParams: operations.AddEnvironmentVariableToContextPathParams{
            ContextID: "sit",
            EnvVarName: "voluptas",
        },
        Request: &operations.AddEnvironmentVariableToContextRequestBody{
            Value: "culpa",
        },
    }
    
    res, err := s.Context.AddEnvironmentVariableToContext(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AddEnvironmentVariableToContext200ApplicationJSONAnyOf != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations

### Context

* `AddEnvironmentVariableToContext` - Add or update an environment variable
* `CreateContext` - Create a new context
* `DeleteContext` - Delete a context
* `DeleteEnvironmentVariableFromContext` - Remove an environment variable
* `GetContext` - Get a context
* `ListContexts` - List contexts
* `ListEnvironmentVariablesFromContext` - List environment variables

### Insights

* `GetAllInsightsBranches` - Get all branches for a project
* `GetFlakyTests` - Get flaky tests for a project
* `GetJobTimeseries` - Job timeseries data
* `GetOrgSummaryData` - Get summary metrics with trends for the entire org, and for each project.
* `GetProjectWorkflowJobMetrics` - Get summary metrics for a project workflow's jobs.
* `GetProjectWorkflowMetrics` - Get summary metrics for a project's workflows
* `GetProjectWorkflowRuns` - Get recent runs of a workflow
* `GetProjectWorkflowTestMetrics` - Get test metrics for a project's workflows
* `GetProjectWorkflowsPageData` - Get summary metrics and trends for a project across it's workflows and branches
* `GetWorkflowSummary` - Get metrics and trends for workflows

### Job

* `CancelJob` - Cancel job
* `GetJobArtifacts` - Get a job's artifacts
* `GetJobDetails` - Get job details
* `GetTests` - Get test metadata

### Pipeline

* `ContinuePipeline` - Continue a pipeline
* `GetPipelineByID` - Get a pipeline by ID
* `GetPipelineByNumber` - Get a pipeline by pipeline number
* `GetPipelineConfigByID` - Get a pipeline's configuration
* `ListMyPipelines` - Get your pipelines
* `ListPipelines` - Get a list of pipelines
* `ListPipelinesForProject` - Get all pipelines
* `ListWorkflowsByPipelineID` - Get a pipeline's workflows
* `TriggerPipeline` - Trigger a new pipeline

### Project

* `CreateCheckoutKey` - Create a new checkout key
* `CreateEnvVar` - Create an environment variable
* `DeleteCheckoutKey` - Delete a checkout key
* `DeleteEnvVar` - Delete an environment variable
* `GetCheckoutKey` - Get a checkout key
* `GetEnvVar` - Get a masked environment variable
* `GetProjectBySlug` - Get a project
* `ListCheckoutKeys` - Get all checkout keys
* `ListEnvVars` - List all environment variables

### Schedule

* `CreateSchedule` - Create a schedule
* `DeleteScheduleByID` - Delete a schedule
* `GetScheduleByID` - Get a schedule
* `ListSchedulesForProject` - Get all schedules
* `UpdateSchedule` - Update a schedule

### User

* `GetCollaborations` - Collaborations
* `GetCurrentUser` - User Information
* `GetUser` - User Information

### Webhook

* `CreateWebhook` - Create a webhook
* `DeleteWebhook` - Delete a webhook
* `GetWebhookByID` - Get a webhook
* `GetWebhooks` - List webhooks
* `UpdateWebhook` - Update a webhook

### Workflow

* `ApprovePendingApprovalJobByID` - Approve a job
* `CancelWorkflow` - Cancel a workflow
* `GetWorkflowByID` - Get a workflow
* `ListWorkflowJobs` - Get a workflow's jobs
* `RerunWorkflow` - Rerun a workflow

<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
