/*
Sonatype Lifecycle Public REST API

Testing LicenseOverridesAPIService

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

func Test_sonatypeiq_LicenseOverridesAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test LicenseOverridesAPIService AddLicenseOverride", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string

		httpRes, err := apiClient.LicenseOverridesAPI.AddLicenseOverride(context.Background(), ownerType, ownerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LicenseOverridesAPIService DeleteLicenseOverride", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string
		var licenseOverrideId string

		httpRes, err := apiClient.LicenseOverridesAPI.DeleteLicenseOverride(context.Background(), ownerType, ownerId, licenseOverrideId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LicenseOverridesAPIService GetAppliedLicenseOverrides", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string

		httpRes, err := apiClient.LicenseOverridesAPI.GetAppliedLicenseOverrides(context.Background(), ownerType, ownerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LicenseOverridesAPIService GetAppliedLicenseOverridesForLegalReviewer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string

		resp, httpRes, err := apiClient.LicenseOverridesAPI.GetAppliedLicenseOverridesForLegalReviewer(context.Background(), ownerType, ownerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
