/*
Sonatype Lifecycle Public REST API

Testing LicenseLegalMetadataTemplatesAPIService

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

func Test_sonatypeiq_LicenseLegalMetadataTemplatesAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test LicenseLegalMetadataTemplatesAPIService DeleteAttributionReportTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.LicenseLegalMetadataTemplatesAPI.DeleteAttributionReportTemplate(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LicenseLegalMetadataTemplatesAPIService GetAllAttributionReportTemplates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LicenseLegalMetadataTemplatesAPI.GetAllAttributionReportTemplates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LicenseLegalMetadataTemplatesAPIService GetAttributionReportTemplateById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.LicenseLegalMetadataTemplatesAPI.GetAttributionReportTemplateById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LicenseLegalMetadataTemplatesAPIService SaveAttributionReportTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LicenseLegalMetadataTemplatesAPI.SaveAttributionReportTemplate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
