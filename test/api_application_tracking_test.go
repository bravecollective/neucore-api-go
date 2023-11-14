/*
Neucore API

Testing ApplicationTrackingApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package neucoreapi

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_neucoreapi_ApplicationTrackingApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ApplicationTrackingApiService MemberTrackingV1", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id int32

        resp, httpRes, err := apiClient.ApplicationTrackingApi.MemberTrackingV1(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}