/*
Sonatype Lifecycle Public REST API

Testing FeatureConfigurationAPIService

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

func Test_sonatypeiq_FeatureConfigurationAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test FeatureConfigurationAPIService DisableFeature", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var feature string

		httpRes, err := apiClient.FeatureConfigurationAPI.DisableFeature(context.Background(), feature).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeatureConfigurationAPIService EnabledFeature", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var feature string

		httpRes, err := apiClient.FeatureConfigurationAPI.EnabledFeature(context.Background(), feature).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
