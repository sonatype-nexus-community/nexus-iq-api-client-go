/*
Sonatype Lifecycle Public REST API

Testing DeveloperPrioritiesAPIService

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

func Test_sonatypeiq_DeveloperPrioritiesAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test DeveloperPrioritiesAPIService GetPriorities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var scanId string

		resp, httpRes, err := apiClient.DeveloperPrioritiesAPI.GetPriorities(context.Background(), applicationId, scanId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeveloperPrioritiesAPIService GetPrioritiesExport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var scanId string

		httpRes, err := apiClient.DeveloperPrioritiesAPI.GetPrioritiesExport(context.Background(), applicationId, scanId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}