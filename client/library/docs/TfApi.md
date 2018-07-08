# \TfApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateModule**](TfApi.md#CreateModule) | **Post** /tf/modules | 
[**ListModules**](TfApi.md#ListModules) | **Get** /tf/modules | 


# **CreateModule**
> ResponseTfModule CreateModule(ctx, optional)


Upload a Terraform module

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terraformModule** | [**Terraformmodule**](Terraformmodule.md)| A terraform module | 

### Return type

[**ResponseTfModule**](response-tf-module.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.terraform.server.v1+json
 - **Accept**: application/vnd.terraform.server.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListModules**
> ResponseListTfModules ListModules(ctx, )


List TF modules

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseListTfModules**](response-list-tf-modules.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.terraform.server.v1+json
 - **Accept**: application/vnd.terraform.server.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

