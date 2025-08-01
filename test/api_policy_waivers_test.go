/*
Sonatype Lifecycle Public REST API

Testing PolicyWaiversAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sonatypeiq

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	sonatypeiq "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func Test_sonatypeiq_PolicyWaiversAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test PolicyWaiversAPIService AddPolicyWaiverByPolicyViolationId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string
		var policyViolationId string

		httpRes, err := apiClient.PolicyWaiversAPI.AddPolicyWaiverByPolicyViolationId(context.Background(), ownerType, ownerId, policyViolationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyWaiversAPIService AddWaiverToTransitivePolicyViolationsByAppScanComponent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string
		var scanId string

		httpRes, err := apiClient.PolicyWaiversAPI.AddWaiverToTransitivePolicyViolationsByAppScanComponent(context.Background(), ownerType, ownerId, scanId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyWaiversAPIService AddWaiverToTransitivePolicyViolationsByOwnerStageComponent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string
		var stageId string

		httpRes, err := apiClient.PolicyWaiversAPI.AddWaiverToTransitivePolicyViolationsByOwnerStageComponent(context.Background(), ownerType, ownerId, stageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyWaiversAPIService DeletePolicyWaiver", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string
		var policyWaiverId string

		httpRes, err := apiClient.PolicyWaiversAPI.DeletePolicyWaiver(context.Background(), ownerType, ownerId, policyWaiverId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyWaiversAPIService GetPolicyWaiver", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string
		var policyWaiverId string

		resp, httpRes, err := apiClient.PolicyWaiversAPI.GetPolicyWaiver(context.Background(), ownerType, ownerId, policyWaiverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyWaiversAPIService GetPolicyWaivers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string

		resp, httpRes, err := apiClient.PolicyWaiversAPI.GetPolicyWaivers(context.Background(), ownerType, ownerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyWaiversAPIService GetTransitivePolicyWaiversByAppScanComponent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string
		var scanId string

		resp, httpRes, err := apiClient.PolicyWaiversAPI.GetTransitivePolicyWaiversByAppScanComponent(context.Background(), ownerType, ownerId, scanId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyWaiversAPIService RequestPolicyWaiver", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyViolationId string

		httpRes, err := apiClient.PolicyWaiversAPI.RequestPolicyWaiver(context.Background(), policyViolationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyWaiversAPIService UpdatePolicyWaiver", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string
		var policyWaiverId string

		httpRes, err := apiClient.PolicyWaiversAPI.UpdatePolicyWaiver(context.Background(), ownerType, ownerId, policyWaiverId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
