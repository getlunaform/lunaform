# Lunaform.ResourcesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listResourceGroups**](ResourcesApi.md#listResourceGroups) | **GET** / | 
[**listResources**](ResourcesApi.md#listResources) | **GET** /{group} | 


<a name="listResourceGroups"></a>
# **listResourceGroups**
> ResponseListResources listResourceGroups()



List the root resource groups

### Example
```javascript
var Lunaform = require('lunaform');

var apiInstance = new Lunaform.ResourcesApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.listResourceGroups(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**ResponseListResources**](ResponseListResources.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

<a name="listResources"></a>
# **listResources**
> ResponseListResources listResources(group)



### Example
```javascript
var Lunaform = require('lunaform');

var apiInstance = new Lunaform.ResourcesApi();

var group = "group_example"; // String | Root level resources


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.listResources(group, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **String**| Root level resources | 

### Return type

[**ResponseListResources**](ResponseListResources.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.lunaform.v1+json
 - **Accept**: application/vnd.lunaform.v1+json

