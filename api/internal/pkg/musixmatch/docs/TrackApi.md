# \TrackApi

All URIs are relative to *https://api.musixmatch.com/ws/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlbumTracksGetGet**](TrackApi.md#AlbumTracksGetGet) | **Get** /album.tracks.get | 
[**ChartTracksGetGet**](TrackApi.md#ChartTracksGetGet) | **Get** /chart.tracks.get | 
[**MatcherTrackGetGet**](TrackApi.md#MatcherTrackGetGet) | **Get** /matcher.track.get | 
[**TrackGetGet**](TrackApi.md#TrackGetGet) | **Get** /track.get | 
[**TrackSearchGet**](TrackApi.md#TrackSearchGet) | **Get** /track.search | 


# **AlbumTracksGetGet**
> InlineResponse2001 AlbumTracksGetGet($albumId, $format, $callback, $fHasLyrics, $page, $pageSize)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **albumId** | **string**| The musiXmatch album id | 
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 
 **fHasLyrics** | **string**| When set, filter only contents with lyrics | [optional] 
 **page** | **float32**| Define the page number for paginated results | [optional] 
 **pageSize** | **float32**| Define the page size for paginated results.Range is 1 to 100. | [optional] 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartTracksGetGet**
> InlineResponse2006 ChartTracksGetGet($format, $callback, $page, $pageSize, $country, $fHasLyrics)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 
 **page** | **float32**| Define the page number for paginated results | [optional] 
 **pageSize** | **float32**| Define the page size for paginated results.Range is 1 to 100. | [optional] 
 **country** | **string**| A valid ISO 3166 country code | [optional] [default to us]
 **fHasLyrics** | **string**| When set, filter only contents with lyrics | [optional] 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MatcherTrackGetGet**
> InlineResponse2009 MatcherTrackGetGet($format, $callback, $qArtist, $qTrack, $fHasLyrics, $fHasSubtitle)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 
 **qArtist** | **string**| The song artist | [optional] 
 **qTrack** | **string**| The song title | [optional] 
 **fHasLyrics** | **float32**| When set, filter only contents with lyrics | [optional] 
 **fHasSubtitle** | **float32**| When set, filter only contents with subtitles | [optional] 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrackGetGet**
> InlineResponse2009 TrackGetGet($trackId, $format, $callback)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trackId** | **string**| The musiXmatch track id | 
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrackSearchGet**
> InlineResponse2006 TrackSearchGet($format, $callback, $qTrack, $qArtist, $qLyrics, $fArtistId, $fMusicGenreId, $fLyricsLanguage, $fHasLyrics, $sArtistRating, $sTrackRating, $quorumFactor, $pageSize, $page)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**| output format: json, jsonp, xml. | [optional] [default to json]
 **callback** | **string**| jsonp callback | [optional] 
 **qTrack** | **string**| The song title | [optional] 
 **qArtist** | **string**| The song artist | [optional] 
 **qLyrics** | **string**| Any word in the lyrics | [optional] 
 **fArtistId** | **float32**| When set, filter by this artist id | [optional] 
 **fMusicGenreId** | **float32**| When set, filter by this music category id | [optional] 
 **fLyricsLanguage** | **float32**| Filter by the lyrics language (en,it,..) | [optional] 
 **fHasLyrics** | **float32**| When set, filter only contents with lyrics | [optional] 
 **sArtistRating** | **string**| Sort by our popularity index for artists (asc|desc) | [optional] 
 **sTrackRating** | **string**| Sort by our popularity index for tracks (asc|desc) | [optional] 
 **quorumFactor** | **float32**| Search only a part of the given query string.Allowed range is (0.1 – 0.9) | [optional] [default to 1]
 **pageSize** | **float32**| Define the page size for paginated results.Range is 1 to 100. | [optional] 
 **page** | **float32**| Define the page number for paginated results | [optional] 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

