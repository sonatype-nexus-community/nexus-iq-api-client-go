/*
Sonatype IQ Server

Testing ScanAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	iqclient "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func Test_openapi_ScanAPIService(t *testing.T) {

	configuration := iqclient.NewConfiguration()
	apiClient := iqclient.NewAPIClient(configuration)

	t.Run("Test ScanAPIService GetIdeUsersOverview", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ScanAPI.GetIdeUsersOverview(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScanAPIService GetScanStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var scanRequestId string

		resp, httpRes, err := apiClient.ScanAPI.GetScanStatus(context.Background(), applicationId, scanRequestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScanAPIService ScanComponents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var source string

		httpRes, err := apiClient.ScanAPI.ScanComponents(context.Background(), applicationId, source).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
