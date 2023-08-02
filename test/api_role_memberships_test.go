/*
Sonatype IQ Server

Testing RoleMembershipsAPIService

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

func Test_sonatypeiq_RoleMembershipsAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test RoleMembershipsAPIService GetRoleMembershipsApplicationOrOrganization", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var internalOwnerId string

		resp, httpRes, err := apiClient.RoleMembershipsAPI.GetRoleMembershipsApplicationOrOrganization(context.Background(), ownerType, internalOwnerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleMembershipsAPIService GetRoleMembershipsGlobalOrRepositoryContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string

		resp, httpRes, err := apiClient.RoleMembershipsAPI.GetRoleMembershipsGlobalOrRepositoryContainer(context.Background(), ownerType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleMembershipsAPIService GrantRoleMembershipApplicationOrOrganization", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var internalOwnerId string
		var roleId string
		var memberType string
		var memberName string

		httpRes, err := apiClient.RoleMembershipsAPI.GrantRoleMembershipApplicationOrOrganization(context.Background(), ownerType, internalOwnerId, roleId, memberType, memberName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleMembershipsAPIService GrantRoleMembershipGlobalOrRepositoryContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var roleId string
		var memberType string
		var memberName string

		httpRes, err := apiClient.RoleMembershipsAPI.GrantRoleMembershipGlobalOrRepositoryContainer(context.Background(), ownerType, roleId, memberType, memberName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleMembershipsAPIService RevokeRoleMembershipApplicationOrOrganization", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var internalOwnerId string
		var roleId string
		var memberType string
		var memberName string

		httpRes, err := apiClient.RoleMembershipsAPI.RevokeRoleMembershipApplicationOrOrganization(context.Background(), ownerType, internalOwnerId, roleId, memberType, memberName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleMembershipsAPIService RevokeRoleMembershipGlobalOrRepositoryContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var roleId string
		var memberType string
		var memberName string

		httpRes, err := apiClient.RoleMembershipsAPI.RevokeRoleMembershipGlobalOrRepositoryContainer(context.Background(), ownerType, roleId, memberType, memberName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}