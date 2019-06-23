# Lunaform.ModulesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createModule**](ModulesApi.md#createModule) | **POST** /tf/modules | 
[**getModule**](ModulesApi.md#getModule) | **GET** /tf/module/{id} | 
[**listModules**](ModulesApi.md#listModules) | **GET** /tf/modules | 


<a name="createModule"></a>
# **createModule**
> ResourceTfModule createModule(opts)



Upload a Terraform module

### Example
```javascript
var Lunaform = require('lunaform');
var defaultClient = Lunaform.ApiClient.instance;

// Configure API key authorization: api-key
var api-key = defaultClient.authentications['api-key'];
api-key.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-key.apiKeyPrefix = 'Token';

var apiInstance = new Lunaform.ModulesApi();

var opts = { 
  'terraformModule': new Lunaform.ResourceTfModule() // ResourceTfModule | A terraform module
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.createModule(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terraformModule** | [**ResourceTfModule**](ResourceTfModule.md)| A terraform module | [optional] 

### Return type

[**ResourceTfModule**](ResourceTfModule.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

<a name="getModule"></a>
# **getModule**
> ResourceTfModule getModule(id)



Get a Terraform module

### Example
```javascript
var Lunaform = require('lunaform');
var defaultClient = Lunaform.ApiClient.instance;

// Configure API key authorization: api-key
var api-key = defaultClient.authentications['api-key'];
api-key.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-key.apiKeyPrefix = 'Token';

var apiInstance = new Lunaform.ModulesApi();

var id = "id_example"; // String | Unique identifier for this module


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getModule(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Unique identifier for this module | 

### Return type

[**ResourceTfModule**](ResourceTfModule.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

<a name="listModules"></a>
# **listModules**
> ResponseListTfModules listModules()



List TF modules

### Example
```javascript
var Lunaform = require('lunaform');
var defaultClient = Lunaform.ApiClient.instance;

// Configure API key authorization: api-key
var api-key = defaultClient.authentications['api-key'];
api-key.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-key.apiKeyPrefix = 'Token';

var apiInstance = new Lunaform.ModulesApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.listModules(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseListTfModules**](ResponseListTfModules.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

