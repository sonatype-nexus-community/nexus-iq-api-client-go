/*
Sonatype Lifecycle Public REST API

Testing UsersAPIService

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

func Test_sonatypeiq_UsersAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test UsersAPIService Add", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.UsersAPI.Add(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService Delete1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		httpRes, err := apiClient.UsersAPI.Delete1(context.Background(), username).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService Get1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.UsersAPI.Get1(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetAll2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.GetAll2(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService Update", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.UsersAPI.Update(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
