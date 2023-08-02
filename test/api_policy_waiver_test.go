/*
Sonatype IQ Server

Testing PolicyWaiverAPIService

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

func Test_sonatypeiq_PolicyWaiverAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test PolicyWaiverAPIService AddPolicyWaiver", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyViolationId string
		var ownerType string

		httpRes, err := apiClient.PolicyWaiverAPI.AddPolicyWaiver(context.Background(), policyViolationId, ownerType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}