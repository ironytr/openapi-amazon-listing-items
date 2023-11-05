/*
Selling Partner API for Listings Items

Testing ListingsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ListingsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ListingsAPIService DeleteListingsItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sellerId string
		var sku string

		resp, httpRes, err := apiClient.ListingsAPI.DeleteListingsItem(context.Background(), sellerId, sku).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListingsAPIService GetListingsItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sellerId string
		var sku string

		resp, httpRes, err := apiClient.ListingsAPI.GetListingsItem(context.Background(), sellerId, sku).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListingsAPIService PatchListingsItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sellerId string
		var sku string

		resp, httpRes, err := apiClient.ListingsAPI.PatchListingsItem(context.Background(), sellerId, sku).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListingsAPIService PutListingsItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sellerId string
		var sku string

		resp, httpRes, err := apiClient.ListingsAPI.PutListingsItem(context.Background(), sellerId, sku).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
