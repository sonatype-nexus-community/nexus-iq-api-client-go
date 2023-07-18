/*
Sonatype IQ Server

Testing PoliciesAPIService

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

func Test_openapi_PoliciesAPIService(t *testing.T) {

	configuration := iqclient.NewConfiguration()
	apiClient := iqclient.NewAPIClient(configuration)

	t.Run("Test PoliciesAPIService GetPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PoliciesAPI.GetPolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
