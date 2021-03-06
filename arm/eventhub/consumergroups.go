package eventhub

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
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ConsumerGroupsClient is the azure Event Hubs client
type ConsumerGroupsClient struct {
	ManagementClient
}

// NewConsumerGroupsClient creates an instance of the ConsumerGroupsClient
// client.
func NewConsumerGroupsClient(subscriptionID string) ConsumerGroupsClient {
	return NewConsumerGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConsumerGroupsClientWithBaseURI creates an instance of the
// ConsumerGroupsClient client.
func NewConsumerGroupsClientWithBaseURI(baseURI string, subscriptionID string) ConsumerGroupsClient {
	return ConsumerGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an Event Hubs consumer group as a nested
// resource within a namespace.
//
// resourceGroupName is name of the Resource group within the Azure
// subscription. namespaceName is the namespace name eventHubName is the
// eventhub name consumerGroupName is the consumer group name parameters is
// parameters supplied to create a consumer group resource.
func (client ConsumerGroupsClient) CreateOrUpdate(resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, parameters ConsumerGroup) (result ConsumerGroup, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: eventHubName,
			Constraints: []validation.Constraint{{Target: "eventHubName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "eventHubName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: consumerGroupName,
			Constraints: []validation.Constraint{{Target: "consumerGroupName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "consumerGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "eventhub.ConsumerGroupsClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, namespaceName, eventHubName, consumerGroupName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ConsumerGroupsClient) CreateOrUpdatePreparer(resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, parameters ConsumerGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"consumerGroupName": autorest.Encode("path", consumerGroupName),
		"eventHubName":      autorest.Encode("path", eventHubName),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ConsumerGroupsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ConsumerGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result ConsumerGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a consumer group from the specified Event Hub and resource
// group.
//
// resourceGroupName is name of the Resource group within the Azure
// subscription. namespaceName is the namespace name eventHubName is the
// eventhub name consumerGroupName is the consumer group name
func (client ConsumerGroupsClient) Delete(resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: eventHubName,
			Constraints: []validation.Constraint{{Target: "eventHubName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "eventHubName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: consumerGroupName,
			Constraints: []validation.Constraint{{Target: "consumerGroupName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "consumerGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "eventhub.ConsumerGroupsClient", "Delete")
	}

	req, err := client.DeletePreparer(resourceGroupName, namespaceName, eventHubName, consumerGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ConsumerGroupsClient) DeletePreparer(resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"consumerGroupName": autorest.Encode("path", consumerGroupName),
		"eventHubName":      autorest.Encode("path", eventHubName),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ConsumerGroupsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ConsumerGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a description for the specified consumer group.
//
// resourceGroupName is name of the Resource group within the Azure
// subscription. namespaceName is the namespace name eventHubName is the
// eventhub name consumerGroupName is the consumer group name
func (client ConsumerGroupsClient) Get(resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (result ConsumerGroup, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: eventHubName,
			Constraints: []validation.Constraint{{Target: "eventHubName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "eventHubName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: consumerGroupName,
			Constraints: []validation.Constraint{{Target: "consumerGroupName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "consumerGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "eventhub.ConsumerGroupsClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, namespaceName, eventHubName, consumerGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConsumerGroupsClient) GetPreparer(resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"consumerGroupName": autorest.Encode("path", consumerGroupName),
		"eventHubName":      autorest.Encode("path", eventHubName),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConsumerGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConsumerGroupsClient) GetResponder(resp *http.Response) (result ConsumerGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEventHub gets all the consumer groups in a eventhub. An empty feed is
// returned if no consumer group exists in the eventhub.
//
// resourceGroupName is name of the Resource group within the Azure
// subscription. namespaceName is the namespace name eventHubName is the
// eventhub name
func (client ConsumerGroupsClient) ListByEventHub(resourceGroupName string, namespaceName string, eventHubName string) (result ConsumerGroupListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: eventHubName,
			Constraints: []validation.Constraint{{Target: "eventHubName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "eventHubName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "eventhub.ConsumerGroupsClient", "ListByEventHub")
	}

	req, err := client.ListByEventHubPreparer(resourceGroupName, namespaceName, eventHubName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "ListByEventHub", nil, "Failure preparing request")
	}

	resp, err := client.ListByEventHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "ListByEventHub", resp, "Failure sending request")
	}

	result, err = client.ListByEventHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "ListByEventHub", resp, "Failure responding to request")
	}

	return
}

// ListByEventHubPreparer prepares the ListByEventHub request.
func (client ConsumerGroupsClient) ListByEventHubPreparer(resourceGroupName string, namespaceName string, eventHubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventHubName":      autorest.Encode("path", eventHubName),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByEventHubSender sends the ListByEventHub request. The method will close the
// http.Response Body if it receives an error.
func (client ConsumerGroupsClient) ListByEventHubSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByEventHubResponder handles the response to the ListByEventHub request. The method always
// closes the http.Response Body.
func (client ConsumerGroupsClient) ListByEventHubResponder(resp *http.Response) (result ConsumerGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEventHubNextResults retrieves the next set of results, if any.
func (client ConsumerGroupsClient) ListByEventHubNextResults(lastResults ConsumerGroupListResult) (result ConsumerGroupListResult, err error) {
	req, err := lastResults.ConsumerGroupListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "ListByEventHub", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByEventHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "ListByEventHub", resp, "Failure sending next results request")
	}

	result, err = client.ListByEventHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventhub.ConsumerGroupsClient", "ListByEventHub", resp, "Failure responding to next results request")
	}

	return
}
