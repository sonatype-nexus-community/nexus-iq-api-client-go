/*
Sonatype IQ Server

Testing VulnerabilitiesAPIService

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

func Test_openapi_VulnerabilitiesAPIService(t *testing.T) {

	configuration := iqclient.NewConfiguration()
	apiClient := iqclient.NewAPIClient(configuration)

	t.Run("Test VulnerabilitiesAPIService GetSecurityVulnerabilityDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var refId string

		resp, httpRes, err := apiClient.VulnerabilitiesAPI.GetSecurityVulnerabilityDetails(context.Background(), refId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
