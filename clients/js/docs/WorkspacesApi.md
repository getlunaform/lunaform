# Lunaform.WorkspacesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createWorkspace**](WorkspacesApi.md#createWorkspace) | **PUT** /tf/workspace/{name} | 
[**describeWorkspace**](WorkspacesApi.md#describeWorkspace) | **GET** /tf/workspace/{name} | 
[**listWorkspaces**](WorkspacesApi.md#listWorkspaces) | **GET** /tf/workspaces | 


<a name="createWorkspace"></a>
# **createWorkspace**
> ResourceTfWorkspace createWorkspace(name, opts)



Create a Terraform workspace

### Example
```javascript
var Lunaform = require('lunaform');
var defaultClient = Lunaform.ApiClient.instance;

// Configure API key authorization: api-key
var api-key = defaultClient.authentications['api-key'];
api-key.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-key.apiKeyPrefix = 'Token';

var apiInstance = new Lunaform.WorkspacesApi();

var name = "name_example"; // String | A terraform workspace

var opts = { 
  'terraformWorkspace': new Lunaform.ResourceTfWorkspace() // ResourceTfWorkspace | A terraform workspace
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.createWorkspace(name, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **String**| A terraform workspace | 
 **terraformWorkspace** | [**ResourceTfWorkspace**](ResourceTfWorkspace.md)| A terraform workspace | [optional] 

### Return type

[**ResourceTfWorkspace**](ResourceTfWorkspace.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

<a name="describeWorkspace"></a>
# **describeWorkspace**
> ResourceTfWorkspace describeWorkspace(name)



Describe a terraform workspace

### Example
```javascript
var Lunaform = require('lunaform');
var defaultClient = Lunaform.ApiClient.instance;

// Configure API key authorization: api-key
var api-key = defaultClient.authentications['api-key'];
api-key.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-key.apiKeyPrefix = 'Token';

var apiInstance = new Lunaform.WorkspacesApi();

var name = "name_example"; // String | A terraform workspace


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.describeWorkspace(name, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **String**| A terraform workspace | 

### Return type

[**ResourceTfWorkspace**](ResourceTfWorkspace.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

<a name="listWorkspaces"></a>
# **listWorkspaces**
> ResponseListTfWorkspaces listWorkspaces()



List approved terraform workspaces

### Example
```javascript
var Lunaform = require('lunaform');
var defaultClient = Lunaform.ApiClient.instance;

// Configure API key authorization: api-key
var api-key = defaultClient.authentications['api-key'];
api-key.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-key.apiKeyPrefix = 'Token';

var apiInstance = new Lunaform.WorkspacesApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.listWorkspaces(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseListTfWorkspaces**](ResponseListTfWorkspaces.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

