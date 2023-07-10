/*
 * Musixmatch API
 *
 * Musixmatch lyrics API is a robust service that permits you to search and retrieve lyrics in the simplest possible way. It just works.  Include millions of licensed lyrics on your website or in your application legally.  The fastest, most powerful and legal way to display lyrics on your website or in your application.  #### Read musixmatch API Terms & Conditions and the Privacy Policy: Before getting started, you must take a look at the [API Terms & Conditions](http://musixmatch.com/apiterms/) and the [Privacy Policy](https://developer.musixmatch.com/privacy). We’ve worked hard to make this service completely legal so that we are all protected from any foreseeable liability. Take the time to read this stuff.  #### Register for an API key: All you need to do is [register](https://developer.musixmatch.com/signup) in order to get your API key, a mandatory parameter for most of our API calls. It’s your personal identifier and should be kept secret:  ```   https://api.musixmatch.com/ws/v1.1/track.get?apikey=YOUR_API_KEY ``` #### Integrate the musixmatch service with your web site or application In the most common scenario you only need to implement two API calls:  The first call is to match your catalog to ours using the [track.search](#!/Track/get_track_search) function and the second is to get the lyrics using the [track.lyrics.get](#!/Lyrics/get_track_lyrics_get) api. That’s it!  ## API Methods What does the musiXmatch API do?  The musiXmatch API allows you to read objects from our huge 100% licensed lyrics database.  To make your life easier we are providing you with one or more examples to show you how it could work in the wild. You’ll find both the API request and API response in all the available output formats for each API call. Follow the links below for the details.  The current API version is 1.1, the root URL is located at https://api.musixmatch.com/ws/1.1/  Supported input parameters can be found on the page [Input Parameters](https://developer.musixmatch.com/documentation/input-parameters). Use UTF-8 to encode arguments when calling API methods.  Every response includes a status_code. A list of all status codes can be consulted at [Status Codes](https://developer.musixmatch.com/documentation/status-codes).  ## Music meta data The musiXmatch api is built around lyrics, but there are many other data we provide through the api that can be used to improve every existent music service.  ## Track Inside the track object you can get the following extra information:  ### TRACK RATING  The track rating is a score 0-100 identifying how popular is a song in musixmatch.  You can use this information to sort search results, like the most popular songs of an artist, of a music genre, of a lyrics language.  ### INSTRUMENTAL AND EXPLICIT FLAGS  The instrumental flag identifies songs with music only, no lyrics.  The explicit flag identifies songs with explicit lyrics or explicit title. We're able to identify explicit words and set the flag for the most common languages.  ### FAVOURITES  How many users have this song in their list of favourites.  Can be used to sort tracks by num favourite to identify more popular tracks within a set.  ### MUSIC GENRE  The music genere of the song.  Can be used to group songs by genre, as input for similarity alghorithms, artist genre identification, navigate songs by genere, etc.  ### SONG TITLES TRANSLATIONS  The track title, as translated in different lanauages, can be used to display the right writing for a given user, example:  LIES (Bigbang) becomes 在光化門 in chinese HALLELUJAH (Bigbang) becomes ハレルヤ in japanese   ## Artist Inside the artist object you can get the following nice extra information:  ### COMMENTS AND COUNTRY  An artist comment is a short snippet of text which can be mainly used for disambiguation.  The artist country is the born country of the artist/group  There are two perfect search result if you search by artist with the keyword \"U2\". Indeed there are two distinct music groups with this same name, one is the most known irish group of Bono Vox, the other is a less popular (world wide speaking) group from Japan.  Here's how you can made use of the artist comment in your search result page:  U2 (Irish rock band) U2 (あきやまうに) You can also show the artist country for even better disambiguation:  U2 (Irish rock band) from Ireland U2 (あきやまうに) from Japan ARTIST TRANSLATIONS  When you create a world wide music related service you have to take into consideration to display the artist name in the user's local language. These translation are also used as aliases to improve the search results.  Let's use PSY for this example.  Western people know him as PSY but korean want to see the original name 싸이.  Using the name translations provided by our api you can show to every user the writing they expect to see.  Furthermore, when you search for \"psy gangnam style\" or \"싸이 gangnam style\" with our search/match api you will still be able to find the song.  ### ARTIST RATING  The artist rating is a score 0-100 identifying how popular is an artist in musixmatch.  You can use this information to build charts, for suggestions, to sort search results. In the example above about U2, we use the artist rating to show the irish band before the japanese one in our serp.  ### ARTIST MUSIC GENRE  We provide one or more main artist genre, this information can be used to calculate similar artist, suggestions, or the filter a search by artist genre.    ## Album Inside the album object you can get the following nice extra information:  ### ALBUM RATING  The album rating is a score 0-100 identifying how popular is an album in musixmatch.  You can use this information to sort search results, like the most popular albums of an artist.  ### ALBUM RATING  The album rating is a score 0-100 identifying how popular is an album in musixmatch.  You can use this information to sort search results, like the most popular albums of an artist.  ### ALBUM COPYRIGHT AND LABEL  For most of our albums we can provide extra information like for example:  Label: Universal-Island Records Ltd. Copyright: (P) 2013 Rubyworks, under license to Columbia Records, a Division of Sony Music Entertainment. ALBUM TYPE AND RELEASE DATE  The album official release date can be used to sort an artist's albums view starting by the most recent one.  Album can also be filtered or grouped by type: Single, Album, Compilation, Remix, Live. This can help to build an artist page with a more organized structure.  ### ALBUM MUSIC GENRE  For most of the albums we provide two groups of music genres. Primary and secondary. This information can be used to help user navigate albums by genre.  An example could be:  Primary genere: POP Secondary genre: K-POP or Mandopop
 *
 * OpenAPI spec version: 1.1.0
 * Contact: info@musixmatch.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package musixmatch

import (
	"encoding/json"
	"net/url"
)

type ArtistApi struct {
	Configuration Configuration
}

func NewArtistApi() *ArtistApi {
	configuration := NewConfiguration()
	return &ArtistApi{
		Configuration: *configuration,
	}
}

func NewArtistApiWithBasePath(basePath string) *ArtistApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &ArtistApi{
		Configuration: *configuration,
	}
}

/**
 *
 *
 *
 * @param artistId  The musiXmatch artist id
 * @param format output format: json, jsonp, xml.
 * @param callback jsonp callback
 * @return *InlineResponse2003
 */
func (a ArtistApi) ArtistGetGet(artistId string, format string, callback string) (*InlineResponse2003, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/artist.get"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(key)' required
	// set key with prefix in query string
	queryParams["apikey"] = append(queryParams["apikey"], a.Configuration.GetAPIKeyWithPrefix("apikey"))
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	queryParams.Add("format", a.Configuration.APIClient.ParameterToString(format, ""))
	queryParams.Add("callback", a.Configuration.APIClient.ParameterToString(callback, ""))
	queryParams.Add("artist_id", a.Configuration.APIClient.ParameterToString(artistId, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(InlineResponse2003)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 *
 *
 *
 * @param artistId The musiXmatch artist id
 * @param format output format: json, jsonp, xml.
 * @param callback jsonp callback
 * @param pageSize Define the page size for paginated results.Range is 1 to 100.
 * @param page Define the page number for paginated results
 * @return *InlineResponse2004
 */
func (a ArtistApi) ArtistRelatedGetGet(artistId string, format string, callback string, pageSize float32, page float32) (*InlineResponse2004, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/artist.related.get"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(key)' required
	// set key with prefix in query string
	queryParams["apikey"] = append(queryParams["apikey"], a.Configuration.GetAPIKeyWithPrefix("apikey"))
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	queryParams.Add("format", a.Configuration.APIClient.ParameterToString(format, ""))
	queryParams.Add("callback", a.Configuration.APIClient.ParameterToString(callback, ""))
	queryParams.Add("artist_id", a.Configuration.APIClient.ParameterToString(artistId, ""))
	queryParams.Add("page_size", a.Configuration.APIClient.ParameterToString(pageSize, ""))
	queryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(InlineResponse2004)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 *
 *
 *
 * @param format output format: json, jsonp, xml.
 * @param callback jsonp callback
 * @param qArtist The song artist
 * @param fArtistId When set, filter by this artist id
 * @param page Define the page number for paginated results
 * @param pageSize Define the page size for paginated results.Range is 1 to 100.
 * @return *InlineResponse2004
 */
func (a ArtistApi) ArtistSearchGet(format string, callback string, qArtist string, fArtistId float32, page float32, pageSize float32) (*InlineResponse2004, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/artist.search"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(key)' required
	// set key with prefix in query string
	queryParams["apikey"] = append(queryParams["apikey"], a.Configuration.GetAPIKeyWithPrefix("apikey"))
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	queryParams.Add("format", a.Configuration.APIClient.ParameterToString(format, ""))
	queryParams.Add("callback", a.Configuration.APIClient.ParameterToString(callback, ""))
	queryParams.Add("q_artist", a.Configuration.APIClient.ParameterToString(qArtist, ""))
	queryParams.Add("f_artist_id", a.Configuration.APIClient.ParameterToString(fArtistId, ""))
	queryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))
	queryParams.Add("page_size", a.Configuration.APIClient.ParameterToString(pageSize, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(InlineResponse2004)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 *
 *
 *
 * @param format output format: json, jsonp, xml.
 * @param callback jsonp callback
 * @param page Define the page number for paginated results
 * @param pageSize Define the page size for paginated results.Range is 1 to 100.
 * @param country A valid ISO 3166 country code
 * @return *InlineResponse2005
 */
func (a ArtistApi) ChartArtistsGetGet(format string, callback string, page float32, pageSize float32, country string) (*InlineResponse2005, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/chart.artists.get"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(key)' required
	// set key with prefix in query string
	queryParams["apikey"] = append(queryParams["apikey"], a.Configuration.GetAPIKeyWithPrefix("apikey"))
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	queryParams.Add("format", a.Configuration.APIClient.ParameterToString(format, ""))
	queryParams.Add("callback", a.Configuration.APIClient.ParameterToString(callback, ""))
	queryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))
	queryParams.Add("page_size", a.Configuration.APIClient.ParameterToString(pageSize, ""))
	queryParams.Add("country", a.Configuration.APIClient.ParameterToString(country, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(InlineResponse2005)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}
