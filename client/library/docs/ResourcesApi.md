# \ResourcesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListResourceGroups**](ResourcesApi.md#ListResourceGroups) | **Get** / | 
[**ListResources**](ResourcesApi.md#ListResources) | **Get** /{group} | 


# **ListResourceGroups**
> ResponseListResources ListResourceGroups(ctx, )


List the root resource groups

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseListResources**](response-list-resources.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.terraform.server.v1+json
 - **Accept**: application/vnd.terraform.server.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListResources**
> ResponseListResources ListResources(ctx, group)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **group** | **string**| Root level resources | 

### Return type

[**ResponseListResources**](response-list-resources.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.terraform.server.v1+json
 - **Accept**: application/vnd.terraform.server.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

