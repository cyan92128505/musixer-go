# \LyricsApi

All URIs are relative to *https://api.musixmatch.com/ws/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MatcherLyricsGetGet**](LyricsApi.md#MatcherLyricsGetGet) | **Get** /matcher.lyrics.get | 
[**TrackLyricsGetGet**](LyricsApi.md#TrackLyricsGetGet) | **Get** /track.lyrics.get | 


# **MatcherLyricsGetGet**
> InlineResponse2007 MatcherLyricsGetGet($format, $callback, $qTrack, $qArtist)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 
 **qTrack** | **string**| The song title | [optional] 
 **qArtist** | **string**| The song artist | [optional] 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrackLyricsGetGet**
> InlineResponse2007 TrackLyricsGetGet($trackId, $format, $callback)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trackId** | **string**| The musiXmatch track id | 
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

