# \SubtitleApi

All URIs are relative to *https://api.musixmatch.com/ws/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MatcherSubtitleGetGet**](SubtitleApi.md#MatcherSubtitleGetGet) | **Get** /matcher.subtitle.get | 
[**TrackSubtitleGetGet**](SubtitleApi.md#TrackSubtitleGetGet) | **Get** /track.subtitle.get | 


# **MatcherSubtitleGetGet**
> InlineResponse2008 MatcherSubtitleGetGet($format, $callback, $qTrack, $qArtist, $fSubtitleLength, $fSubtitleLengthMaxDeviation)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 
 **qTrack** | **string**| The song title | [optional] 
 **qArtist** | **string**|  The song artist | [optional] 
 **fSubtitleLength** | **float32**| Filter by subtitle length in seconds | [optional] 
 **fSubtitleLengthMaxDeviation** | **float32**| Max deviation for a subtitle length in seconds | [optional] 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrackSubtitleGetGet**
> InlineResponse2008 TrackSubtitleGetGet($trackId, $format, $callback)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trackId** | **string**| The musiXmatch track id | 
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

