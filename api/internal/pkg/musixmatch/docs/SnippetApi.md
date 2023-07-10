# \SnippetApi

All URIs are relative to *https://api.musixmatch.com/ws/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrackSnippetGetGet**](SnippetApi.md#TrackSnippetGetGet) | **Get** /track.snippet.get | 


# **TrackSnippetGetGet**
> InlineResponse20010 TrackSnippetGetGet($trackId, $format, $callback)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trackId** | **string**| The musiXmatch track id | 
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

