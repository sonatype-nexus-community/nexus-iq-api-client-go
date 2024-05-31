/*
Sonatype Lifecycle Public REST API

Testing PolicyEvaluationAPIService

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

func Test_sonatypeiq_PolicyEvaluationAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test PolicyEvaluationAPIService EvaluateComponents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.PolicyEvaluationAPI.EvaluateComponents(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyEvaluationAPIService EvaluateSourceControl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.PolicyEvaluationAPI.EvaluateSourceControl(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyEvaluationAPIService GetApplicationEvaluationStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var statusId string

		resp, httpRes, err := apiClient.PolicyEvaluationAPI.GetApplicationEvaluationStatus(context.Background(), applicationId, statusId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyEvaluationAPIService GetComponentEvaluation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var resultId string

		resp, httpRes, err := apiClient.PolicyEvaluationAPI.GetComponentEvaluation(context.Background(), applicationId, resultId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyEvaluationAPIService PromoteScan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.PolicyEvaluationAPI.PromoteScan(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}