/*
Sonatype Lifecycle Public REST API

Testing ClaimComponentsAPIService

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

func Test_sonatypeiq_ClaimComponentsAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test ClaimComponentsAPIService Delete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hash string

		httpRes, err := apiClient.ClaimComponentsAPI.Delete(context.Background(), hash).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClaimComponentsAPIService Get", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var hash string

		resp, httpRes, err := apiClient.ClaimComponentsAPI.Get(context.Background(), hash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClaimComponentsAPIService GetAll", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ClaimComponentsAPI.GetAll(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClaimComponentsAPIService Set", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ClaimComponentsAPI.Set(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}