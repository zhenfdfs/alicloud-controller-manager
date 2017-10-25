package commitmentplans

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
	"net/http"
)

// CommitmentAssociationsClient is the these APIs allow end users to operate on Azure Machine Learning Commitment Plans
// resources and their child Commitment Association resources. They support CRUD operations for commitment plans, get
// and list operations for commitment associations, moving commitment associations between commitment plans, and
// retrieving commitment plan usage history.
type CommitmentAssociationsClient struct {
	ManagementClient
}

// NewCommitmentAssociationsClient creates an instance of the CommitmentAssociationsClient client.
func NewCommitmentAssociationsClient(subscriptionID string) CommitmentAssociationsClient {
	return NewCommitmentAssociationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCommitmentAssociationsClientWithBaseURI creates an instance of the CommitmentAssociationsClient client.
func NewCommitmentAssociationsClientWithBaseURI(baseURI string, subscriptionID string) CommitmentAssociationsClient {
	return CommitmentAssociationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a commitment association.
//
// resourceGroupName is the resource group name. commitmentPlanName is the Azure ML commitment plan name.
// commitmentAssociationName is the commitment association name.
func (client CommitmentAssociationsClient) Get(resourceGroupName string, commitmentPlanName string, commitmentAssociationName string) (result CommitmentAssociation, err error) {
	req, err := client.GetPreparer(resourceGroupName, commitmentPlanName, commitmentAssociationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentAssociationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentAssociationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentAssociationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client CommitmentAssociationsClient) GetPreparer(resourceGroupName string, commitmentPlanName string, commitmentAssociationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"commitmentAssociationName": autorest.Encode("path", commitmentAssociationName),
		"commitmentPlanName":        autorest.Encode("path", commitmentPlanName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/commitmentPlans/{commitmentPlanName}/commitmentAssociations/{commitmentAssociationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CommitmentAssociationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CommitmentAssociationsClient) GetResponder(resp *http.Response) (result CommitmentAssociation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get all commitment associations for a parent commitment plan.
//
// resourceGroupName is the resource group name. commitmentPlanName is the Azure ML commitment plan name. skipToken is
// continuation token for pagination.
func (client CommitmentAssociationsClient) List(resourceGroupName string, commitmentPlanName string, skipToken string) (result CommitmentAssociationListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName, commitmentPlanName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentAssociationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentAssociationsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentAssociationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client CommitmentAssociationsClient) ListPreparer(resourceGroupName string, commitmentPlanName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"commitmentPlanName": autorest.Encode("path", commitmentPlanName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/commitmentPlans/{commitmentPlanName}/commitmentAssociations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CommitmentAssociationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CommitmentAssociationsClient) ListResponder(resp *http.Response) (result CommitmentAssociationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client CommitmentAssociationsClient) ListNextResults(lastResults CommitmentAssociationListResult) (result CommitmentAssociationListResult, err error) {
	req, err := lastResults.CommitmentAssociationListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "commitmentplans.CommitmentAssociationsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "commitmentplans.CommitmentAssociationsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentAssociationsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client CommitmentAssociationsClient) ListComplete(resourceGroupName string, commitmentPlanName string, skipToken string, cancel <-chan struct{}) (<-chan CommitmentAssociation, <-chan error) {
	resultChan := make(chan CommitmentAssociation)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(resourceGroupName, commitmentPlanName, skipToken)
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
			list, err = client.ListNextResults(list)
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

// Move re-parent a commitment association from one commitment plan to another.
//
// resourceGroupName is the resource group name. commitmentPlanName is the Azure ML commitment plan name.
// commitmentAssociationName is the commitment association name. movePayload is the move request payload.
func (client CommitmentAssociationsClient) Move(resourceGroupName string, commitmentPlanName string, commitmentAssociationName string, movePayload MoveCommitmentAssociationRequest) (result CommitmentAssociation, err error) {
	req, err := client.MovePreparer(resourceGroupName, commitmentPlanName, commitmentAssociationName, movePayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentAssociationsClient", "Move", nil, "Failure preparing request")
		return
	}

	resp, err := client.MoveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentAssociationsClient", "Move", resp, "Failure sending request")
		return
	}

	result, err = client.MoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentAssociationsClient", "Move", resp, "Failure responding to request")
	}

	return
}

// MovePreparer prepares the Move request.
func (client CommitmentAssociationsClient) MovePreparer(resourceGroupName string, commitmentPlanName string, commitmentAssociationName string, movePayload MoveCommitmentAssociationRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"commitmentAssociationName": autorest.Encode("path", commitmentAssociationName),
		"commitmentPlanName":        autorest.Encode("path", commitmentPlanName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/commitmentPlans/{commitmentPlanName}/commitmentAssociations/{commitmentAssociationName}/move", pathParameters),
		autorest.WithJSON(movePayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// MoveSender sends the Move request. The method will close the
// http.Response Body if it receives an error.
func (client CommitmentAssociationsClient) MoveSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// MoveResponder handles the response to the Move request. The method always
// closes the http.Response Body.
func (client CommitmentAssociationsClient) MoveResponder(resp *http.Response) (result CommitmentAssociation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
