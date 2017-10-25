package webservices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// GroupClient is the these APIs allow end users to operate on Azure Machine Learning Web Services resources. They
// support the following operations:<ul><li>Create or update a web service</li><li>Get a web service</li><li>Patch a
// web service</li><li>Delete a web service</li><li>Get All Web Services in a Resource Group </li><li>Get All Web
// Services in a Subscription</li><li>Get Web Services Keys</li></ul>
type GroupClient struct {
	ManagementClient
}

// NewGroupClient creates an instance of the GroupClient client.
func NewGroupClient(subscriptionID string) GroupClient {
	return NewGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGroupClientWithBaseURI creates an instance of the GroupClient client.
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return GroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a web service. This call will overwrite an existing web service. Note that there is
// no warning or confirmation. This is a nonrecoverable operation. If your intent is to create a new web service, call
// the Get operation first to verify that it does not exist. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is name of the resource group in which the web service is located. webServiceName is the name of
// the web service. createOrUpdatePayload is the payload that is used to create or update the web service.
func (client GroupClient) CreateOrUpdate(resourceGroupName string, webServiceName string, createOrUpdatePayload WebService, cancel <-chan struct{}) (<-chan WebService, <-chan error) {
	resultChan := make(chan WebService, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: createOrUpdatePayload,
			Constraints: []validation.Constraint{{Target: "createOrUpdatePayload.Properties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.RealtimeConfiguration", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.RealtimeConfiguration.MaxConcurrentCalls", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.RealtimeConfiguration.MaxConcurrentCalls", Name: validation.InclusiveMaximum, Rule: 200, Chain: nil},
							{Target: "createOrUpdatePayload.Properties.RealtimeConfiguration.MaxConcurrentCalls", Name: validation.InclusiveMinimum, Rule: 4, Chain: nil},
						}},
					}},
					{Target: "createOrUpdatePayload.Properties.MachineLearningWorkspace", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.MachineLearningWorkspace.ID", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "createOrUpdatePayload.Properties.CommitmentPlan", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.CommitmentPlan.ID", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "createOrUpdatePayload.Properties.Input", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.Input.Type", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "createOrUpdatePayload.Properties.Input.Properties", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "createOrUpdatePayload.Properties.Output", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.Output.Type", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "createOrUpdatePayload.Properties.Output.Properties", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "createOrUpdatePayload.Properties.PayloadsLocation", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.PayloadsLocation.URI", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "webservices.GroupClient", "CreateOrUpdate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result WebService
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, webServiceName, createOrUpdatePayload, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "webservices.GroupClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "webservices.GroupClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "webservices.GroupClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client GroupClient) CreateOrUpdatePreparer(resourceGroupName string, webServiceName string, createOrUpdatePayload WebService, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webServiceName":    autorest.Encode("path", webServiceName),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", pathParameters),
		autorest.WithJSON(createOrUpdatePayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client GroupClient) CreateOrUpdateResponder(resp *http.Response) (result WebService, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateRegionalProperties creates an encrypted credentials parameter blob for the specified region. To get the web
// service from a region other than the region in which it has been created, you must first call Create Regional Web
// Services Properties to create a copy of the encrypted credential parameter blob in that region. You only need to do
// this before the first time that you get the web service in the new region. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is name of the resource group in which the web service is located. webServiceName is the name of
// the web service. region is the region for which encrypted credential parameters are created.
func (client GroupClient) CreateRegionalProperties(resourceGroupName string, webServiceName string, region string, cancel <-chan struct{}) (<-chan AsyncOperationStatus, <-chan error) {
	resultChan := make(chan AsyncOperationStatus, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result AsyncOperationStatus
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateRegionalPropertiesPreparer(resourceGroupName, webServiceName, region, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "webservices.GroupClient", "CreateRegionalProperties", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateRegionalPropertiesSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "webservices.GroupClient", "CreateRegionalProperties", resp, "Failure sending request")
			return
		}

		result, err = client.CreateRegionalPropertiesResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "webservices.GroupClient", "CreateRegionalProperties", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateRegionalPropertiesPreparer prepares the CreateRegionalProperties request.
func (client GroupClient) CreateRegionalPropertiesPreparer(resourceGroupName string, webServiceName string, region string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webServiceName":    autorest.Encode("path", webServiceName),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"region":      autorest.Encode("query", region),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}/CreateRegionalBlob", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateRegionalPropertiesSender sends the CreateRegionalProperties request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) CreateRegionalPropertiesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateRegionalPropertiesResponder handles the response to the CreateRegionalProperties request. The method always
// closes the http.Response Body.
func (client GroupClient) CreateRegionalPropertiesResponder(resp *http.Response) (result AsyncOperationStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the Web Service Definition as specified by a subscription, resource group, and name. Note that the storage
// credentials and web service keys are not returned by this call. To get the web service access keys, call List Keys.
//
// resourceGroupName is name of the resource group in which the web service is located. webServiceName is the name of
// the web service. region is the region for which encrypted credential parameters are valid.
func (client GroupClient) Get(resourceGroupName string, webServiceName string, region string) (result WebService, err error) {
	req, err := client.GetPreparer(resourceGroupName, webServiceName, region)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.GroupClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "webservices.GroupClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.GroupClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client GroupClient) GetPreparer(resourceGroupName string, webServiceName string, region string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webServiceName":    autorest.Encode("path", webServiceName),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(region) > 0 {
		queryParameters["region"] = autorest.Encode("query", region)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GroupClient) GetResponder(resp *http.Response) (result WebService, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets the web services in the specified resource group.
//
// resourceGroupName is name of the resource group in which the web service is located. skiptoken is continuation token
// for pagination.
func (client GroupClient) ListByResourceGroup(resourceGroupName string, skiptoken string) (result PaginatedWebServicesList, err error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroupName, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.GroupClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "webservices.GroupClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.GroupClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client GroupClient) ListByResourceGroupPreparer(resourceGroupName string, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client GroupClient) ListByResourceGroupResponder(resp *http.Response) (result PaginatedWebServicesList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client GroupClient) ListByResourceGroupNextResults(lastResults PaginatedWebServicesList) (result PaginatedWebServicesList, err error) {
	req, err := lastResults.PaginatedWebServicesListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "webservices.GroupClient", "ListByResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "webservices.GroupClient", "ListByResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.GroupClient", "ListByResourceGroup", resp, "Failure responding to next results request")
	}

	return
}

// ListByResourceGroupComplete gets all elements from the list without paging.
func (client GroupClient) ListByResourceGroupComplete(resourceGroupName string, skiptoken string, cancel <-chan struct{}) (<-chan WebService, <-chan error) {
	resultChan := make(chan WebService)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByResourceGroup(resourceGroupName, skiptoken)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByResourceGroupNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// ListBySubscriptionID gets the web services in the specified subscription.
//
// skiptoken is continuation token for pagination.
func (client GroupClient) ListBySubscriptionID(skiptoken string) (result PaginatedWebServicesList, err error) {
	req, err := client.ListBySubscriptionIDPreparer(skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.GroupClient", "ListBySubscriptionID", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "webservices.GroupClient", "ListBySubscriptionID", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.GroupClient", "ListBySubscriptionID", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionIDPreparer prepares the ListBySubscriptionID request.
func (client GroupClient) ListBySubscriptionIDPreparer(skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearning/webServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListBySubscriptionIDSender sends the ListBySubscriptionID request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) ListBySubscriptionIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListBySubscriptionIDResponder handles the response to the ListBySubscriptionID request. The method always
// closes the http.Response Body.
func (client GroupClient) ListBySubscriptionIDResponder(resp *http.Response) (result PaginatedWebServicesList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscriptionIDNextResults retrieves the next set of results, if any.
func (client GroupClient) ListBySubscriptionIDNextResults(lastResults PaginatedWebServicesList) (result PaginatedWebServicesList, err error) {
	req, err := lastResults.PaginatedWebServicesListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "webservices.GroupClient", "ListBySubscriptionID", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListBySubscriptionIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "webservices.GroupClient", "ListBySubscriptionID", resp, "Failure sending next results request")
	}

	result, err = client.ListBySubscriptionIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.GroupClient", "ListBySubscriptionID", resp, "Failure responding to next results request")
	}

	return
}

// ListBySubscriptionIDComplete gets all elements from the list without paging.
func (client GroupClient) ListBySubscriptionIDComplete(skiptoken string, cancel <-chan struct{}) (<-chan WebService, <-chan error) {
	resultChan := make(chan WebService)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListBySubscriptionID(skiptoken)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListBySubscriptionIDNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// ListKeys gets the access keys for the specified web service.
//
// resourceGroupName is name of the resource group in which the web service is located. webServiceName is the name of
// the web service.
func (client GroupClient) ListKeys(resourceGroupName string, webServiceName string) (result Keys, err error) {
	req, err := client.ListKeysPreparer(resourceGroupName, webServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.GroupClient", "ListKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "webservices.GroupClient", "ListKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.GroupClient", "ListKeys", resp, "Failure responding to request")
	}

	return
}

// ListKeysPreparer prepares the ListKeys request.
func (client GroupClient) ListKeysPreparer(resourceGroupName string, webServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webServiceName":    autorest.Encode("path", webServiceName),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}/listKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListKeysSender sends the ListKeys request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) ListKeysSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListKeysResponder handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (client GroupClient) ListKeysResponder(resp *http.Response) (result Keys, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Patch modifies an existing web service resource. The PATCH API call is an asynchronous operation. To determine
// whether it has completed successfully, you must perform a Get operation. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is name of the resource group in which the web service is located. webServiceName is the name of
// the web service. patchPayload is the payload to use to patch the web service.
func (client GroupClient) Patch(resourceGroupName string, webServiceName string, patchPayload WebService, cancel <-chan struct{}) (<-chan WebService, <-chan error) {
	resultChan := make(chan WebService, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result WebService
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.PatchPreparer(resourceGroupName, webServiceName, patchPayload, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "webservices.GroupClient", "Patch", nil, "Failure preparing request")
			return
		}

		resp, err := client.PatchSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "webservices.GroupClient", "Patch", resp, "Failure sending request")
			return
		}

		result, err = client.PatchResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "webservices.GroupClient", "Patch", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// PatchPreparer prepares the Patch request.
func (client GroupClient) PatchPreparer(resourceGroupName string, webServiceName string, patchPayload WebService, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webServiceName":    autorest.Encode("path", webServiceName),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", pathParameters),
		autorest.WithJSON(patchPayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) PatchSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client GroupClient) PatchResponder(resp *http.Response) (result WebService, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Remove deletes the specified web service. This method may poll for completion. Polling can be canceled by passing
// the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is name of the resource group in which the web service is located. webServiceName is the name of
// the web service.
func (client GroupClient) Remove(resourceGroupName string, webServiceName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.RemovePreparer(resourceGroupName, webServiceName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "webservices.GroupClient", "Remove", nil, "Failure preparing request")
			return
		}

		resp, err := client.RemoveSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "webservices.GroupClient", "Remove", resp, "Failure sending request")
			return
		}

		result, err = client.RemoveResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "webservices.GroupClient", "Remove", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// RemovePreparer prepares the Remove request.
func (client GroupClient) RemovePreparer(resourceGroupName string, webServiceName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webServiceName":    autorest.Encode("path", webServiceName),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) RemoveSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client GroupClient) RemoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
