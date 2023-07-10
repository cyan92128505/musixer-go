# \AlbumApi

All URIs are relative to *https://api.musixmatch.com/ws/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlbumGetGet**](AlbumApi.md#AlbumGetGet) | **Get** /album.get | 
[**ArtistAlbumsGetGet**](AlbumApi.md#ArtistAlbumsGetGet) | **Get** /artist.albums.get | 


# **AlbumGetGet**
> InlineResponse200 AlbumGetGet($albumId, $format, $callback)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **albumId** | **string**| The musiXmatch album id | 
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArtistAlbumsGetGet**
> InlineResponse2002 ArtistAlbumsGetGet($artistId, $format, $callback, $sReleaseDate, $gAlbumName, $pageSize, $page)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistId** | **string**| The musiXmatch artist id | 
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 
 **sReleaseDate** | **string**| Sort by release date (asc|desc) | [optional] 
 **gAlbumName** | **string**| Group by Album Name | [optional] 
 **pageSize** | **float32**| Define the page size for paginated results.Range is 1 to 100. | [optional] 
 **page** | **float32**| Define the page number for paginated results | [optional] 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

