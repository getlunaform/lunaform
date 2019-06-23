# Lunaform.StacksApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**deployStack**](StacksApi.md#deployStack) | **POST** /tf/stacks | 
[**getStack**](StacksApi.md#getStack) | **GET** /tf/stack/{id} | 
[**listDeployments**](StacksApi.md#listDeployments) | **GET** /tf/stack/{id}/deployments | 
[**listStacks**](StacksApi.md#listStacks) | **GET** /tf/stacks | 


<a name="deployStack"></a>
# **deployStack**
> ResourceTfStack deployStack(opts)



Deploy a stack from a module

### Example
```javascript
var Lunaform = require('lunaform');
var defaultClient = Lunaform.ApiClient.instance;

// Configure API key authorization: api-key
var api-key = defaultClient.authentications['api-key'];
api-key.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-key.apiKeyPrefix = 'Token';

var apiInstance = new Lunaform.StacksApi();

var opts = { 
  'terraformStack': new Lunaform.ResourceTfStack() // ResourceTfStack | A deployed terraform module
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.deployStack(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terraformStack** | [**ResourceTfStack**](ResourceTfStack.md)| A deployed terraform module | [optional] 

### Return type

[**ResourceTfStack**](ResourceTfStack.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

<a name="getStack"></a>
# **getStack**
> ResourceTfStack getStack(id)



Get a stack

### Example
```javascript
var Lunaform = require('lunaform');
var defaultClient = Lunaform.ApiClient.instance;

// Configure API key authorization: api-key
var api-key = defaultClient.authentications['api-key'];
api-key.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-key.apiKeyPrefix = 'Token';

var apiInstance = new Lunaform.StacksApi();

var id = "id_example"; // String | Unique identifier for this stack


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getStack(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Unique identifier for this stack | 

### Return type

[**ResourceTfStack**](ResourceTfStack.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

<a name="listDeployments"></a>
# **listDeployments**
> ResponseListTfDeployments listDeployments(id)



List the deployments for this stack

### Example
```javascript
var Lunaform = require('lunaform');
var defaultClient = Lunaform.ApiClient.instance;

// Configure API key authorization: api-key
var api-key = defaultClient.authentications['api-key'];
api-key.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-key.apiKeyPrefix = 'Token';

var apiInstance = new Lunaform.StacksApi();

var id = "id_example"; // String | Unique identifier for this stack


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.listDeployments(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Unique identifier for this stack | 

### Return type

[**ResponseListTfDeployments**](ResponseListTfDeployments.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

<a name="listStacks"></a>
# **listStacks**
> ResponseListTfStacks listStacks()



List deployed TF modules

### Example
```javascript
var Lunaform = require('lunaform');
var defaultClient = Lunaform.ApiClient.instance;

// Configure API key authorization: api-key
var api-key = defaultClient.authentications['api-key'];
api-key.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-key.apiKeyPrefix = 'Token';

var apiInstance = new Lunaform.StacksApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.listStacks(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseListTfStacks**](ResponseListTfStacks.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

