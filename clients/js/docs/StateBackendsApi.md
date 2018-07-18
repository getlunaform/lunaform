# Lunaform.StateBackendsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createStateBackend**](StateBackendsApi.md#createStateBackend) | **POST** /tf/state-backends | 
[**listStateBackends**](StateBackendsApi.md#listStateBackends) | **GET** /tf/state-backends | 
[**updateStateBackend**](StateBackendsApi.md#updateStateBackend) | **PUT** /tf/state-backend/{id} | 


<a name="createStateBackend"></a>
# **createStateBackend**
> ResourceTfStateBackend createStateBackend(opts)



Define a Terraform state backend

### Example
```javascript
var Lunaform = require('lunaform');
var defaultClient = Lunaform.ApiClient.instance;

// Configure API key authorization: api-key
var api-key = defaultClient.authentications['api-key'];
api-key.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-key.apiKeyPrefix = 'Token';

var apiInstance = new Lunaform.StateBackendsApi();

var opts = { 
  'terraformStateBackend': new Lunaform.ResourceTfStateBackend() // ResourceTfStateBackend | A terraform state backend
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.createStateBackend(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terraformStateBackend** | [**ResourceTfStateBackend**](ResourceTfStateBackend.md)| A terraform state backend | [optional] 

### Return type

[**ResourceTfStateBackend**](ResourceTfStateBackend.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

<a name="listStateBackends"></a>
# **listStateBackends**
> ResponseListTfStateBackends listStateBackends()



List TF State Backends

### Example
```javascript
var Lunaform = require('lunaform');
var defaultClient = Lunaform.ApiClient.instance;

// Configure API key authorization: api-key
var api-key = defaultClient.authentications['api-key'];
api-key.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-key.apiKeyPrefix = 'Token';

var apiInstance = new Lunaform.StateBackendsApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.listStateBackends(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseListTfStateBackends**](ResponseListTfStateBackends.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

<a name="updateStateBackend"></a>
# **updateStateBackend**
> ResourceTfStateBackend updateStateBackend(id, opts)



Define a Terraform state backend

### Example
```javascript
var Lunaform = require('lunaform');
var defaultClient = Lunaform.ApiClient.instance;

// Configure API key authorization: api-key
var api-key = defaultClient.authentications['api-key'];
api-key.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-key.apiKeyPrefix = 'Token';

var apiInstance = new Lunaform.StateBackendsApi();

var id = "id_example"; // String | ID of a terraform state backend

var opts = { 
  'terraformStateBackend': new Lunaform.ResourceTfStateBackend() // ResourceTfStateBackend | A terraform state backend
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.updateStateBackend(id, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| ID of a terraform state backend | 
 **terraformStateBackend** | [**ResourceTfStateBackend**](ResourceTfStateBackend.md)| A terraform state backend | [optional] 

### Return type

[**ResourceTfStateBackend**](ResourceTfStateBackend.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

