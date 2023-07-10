# \ArtistApi

All URIs are relative to *https://api.musixmatch.com/ws/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArtistGetGet**](ArtistApi.md#ArtistGetGet) | **Get** /artist.get | 
[**ArtistRelatedGetGet**](ArtistApi.md#ArtistRelatedGetGet) | **Get** /artist.related.get | 
[**ArtistSearchGet**](ArtistApi.md#ArtistSearchGet) | **Get** /artist.search | 
[**ChartArtistsGetGet**](ArtistApi.md#ChartArtistsGetGet) | **Get** /chart.artists.get | 


# **ArtistGetGet**
> InlineResponse2003 ArtistGetGet($artistId, $format, $callback)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistId** | **string**|  The musiXmatch artist id | 
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArtistRelatedGetGet**
> InlineResponse2004 ArtistRelatedGetGet($artistId, $format, $callback, $pageSize, $page)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistId** | **string**| The musiXmatch artist id | 
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 
 **pageSize** | **float32**| Define the page size for paginated results.Range is 1 to 100. | [optional] 
 **page** | **float32**| Define the page number for paginated results | [optional] 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArtistSearchGet**
> InlineResponse2004 ArtistSearchGet($format, $callback, $qArtist, $fArtistId, $page, $pageSize)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 
 **qArtist** | **string**| The song artist | [optional] 
 **fArtistId** | **float32**| When set, filter by this artist id | [optional] 
 **page** | **float32**| Define the page number for paginated results | [optional] 
 **pageSize** | **float32**| Define the page size for paginated results.Range is 1 to 100. | [optional] 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartArtistsGetGet**
> InlineResponse2005 ChartArtistsGetGet($format, $callback, $page, $pageSize, $country)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 
 **page** | **float32**| Define the page number for paginated results | [optional] 
 **pageSize** | **float32**| Define the page size for paginated results.Range is 1 to 100. | [optional] 
 **country** | **string**| A valid ISO 3166 country code | [optional] [default to us]

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

