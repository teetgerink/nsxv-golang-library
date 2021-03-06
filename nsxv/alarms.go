/*
 * NSX-V
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nsxv

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type AlarmsApiService service

/*
AlarmsApiService servicesAlarmsSourceList
Retrive all alarms from the specified source.   Parameters:  sourceId: ID of the object for which you want to manage alarms. *sourceId* can be the ID of a cluster, host, resource pool, security group, or edge.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param sourceId

*/
func (a *AlarmsApiService) ServicesAlarmsSourceIdGet(ctx context.Context, sourceId string) (*http.Response, error) {
	var (
		lHttpMethod = strings.ToUpper("Get")
		lPostBody   interface{}
		lFileName   string
		lFileBytes  []byte
	)

	// create path and map variables
	lPath := a.client.cfg.BasePath + "/api/2.0/services/alarms/{sourceId}"
	lPath = strings.Replace(lPath, "{"+"sourceId"+"}", fmt.Sprintf("%v", sourceId), -1)

	lHeaderParams := make(map[string]string)
	lQueryParams := url.Values{}
	lFormParams := url.Values{}

	// to determine the Content-Type header
	lHttpContentTypes := []string{}

	// set Content-Type header
	lHttpContentType := selectHeaderContentType(lHttpContentTypes)
	if lHttpContentType != "" {
		lHeaderParams["Content-Type"] = lHttpContentType
	}

	// to determine the Accept header
	lHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	lHttpHeaderAccept := selectHeaderAccept(lHttpHeaderAccepts)
	if lHttpHeaderAccept != "" {
		lHeaderParams["Accept"] = lHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, lPath, lHttpMethod, lPostBody, lHeaderParams, lQueryParams, lFormParams, lFileName, lFileBytes)
	if err != nil {
		return nil, err
	}

	lHttpResponse, err := a.client.callAPI(r)
	if err != nil || lHttpResponse == nil {
		return lHttpResponse, err
	}

	lBody, err := ioutil.ReadAll(lHttpResponse.Body)
	lHttpResponse.Body.Close()
	if err != nil {
		return lHttpResponse, err
	}

	if lHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  lBody,
			error: lHttpResponse.Status,
		}
		return lHttpResponse, newErr
	}

	return lHttpResponse, nil
}

/*
AlarmsApiService servicesAlarmsSourceUpdate
Resolve all alarms for the specified source.  Alarms will resolve automatically when the cause of the alarm is resolved.  For example, if an NSX Edge appliance is powered off, this will trigger an alarm. If you power the NSX Edge appliance back on, the alarm will resolve. If however, you delete the NSX Edge appliance, the alarm will persist, because the alarm cause was never resolved. In this case, you may want to manually resolve the alarm. Resolving the alarms will clear them from the NSX dashboard.  Use &#x60;GET /api/2.0/services/alarms/{sourceId}&#x60; to retrieve the list of alarms for the source. Use this response as the request body for the &#x60;POST&#x60; call.   Parameters:  sourceId: ID of the object for which you want to manage alarms. *sourceId* can be the ID of a cluster, host, resource pool, security group, or edge.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param sourceId
 * @param optional nil or *AlarmsApiServicesAlarmsSourceIdPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of interface{}) -
     * @param "ContentType" (optional.String) -
     * @param "Action" (optional.String) -

*/

type AlarmsApiServicesAlarmsSourceIdPostOpts struct {
	Body        optional.Interface
	ContentType optional.String
	Action      optional.String
}

func (a *AlarmsApiService) ServicesAlarmsSourceIdPost(ctx context.Context, sourceId string, lOptionals *AlarmsApiServicesAlarmsSourceIdPostOpts) (*http.Response, error) {
	var (
		lHttpMethod = strings.ToUpper("Post")
		lPostBody   interface{}
		lFileName   string
		lFileBytes  []byte
	)

	// create path and map variables
	lPath := a.client.cfg.BasePath + "/api/2.0/services/alarms/{sourceId}"
	lPath = strings.Replace(lPath, "{"+"sourceId"+"}", fmt.Sprintf("%v", sourceId), -1)

	lHeaderParams := make(map[string]string)
	lQueryParams := url.Values{}
	lFormParams := url.Values{}

	if lOptionals != nil && lOptionals.Action.IsSet() {
		lQueryParams.Add("action", parameterToString(lOptionals.Action.Value(), ""))
	}
	// to determine the Content-Type header
	lHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	lHttpContentType := selectHeaderContentType(lHttpContentTypes)
	if lHttpContentType != "" {
		lHeaderParams["Content-Type"] = lHttpContentType
	}

	// to determine the Accept header
	lHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	lHttpHeaderAccept := selectHeaderAccept(lHttpHeaderAccepts)
	if lHttpHeaderAccept != "" {
		lHeaderParams["Accept"] = lHttpHeaderAccept
	}
	if lOptionals != nil && lOptionals.ContentType.IsSet() {
		lHeaderParams["Content-Type"] = parameterToString(lOptionals.ContentType.Value(), "")
	}
	// body params
	if lOptionals != nil && lOptionals.Body.IsSet() {

		lOptionalBody := lOptionals.Body.Value()
		lPostBody = &lOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, lPath, lHttpMethod, lPostBody, lHeaderParams, lQueryParams, lFormParams, lFileName, lFileBytes)
	if err != nil {
		return nil, err
	}

	lHttpResponse, err := a.client.callAPI(r)
	if err != nil || lHttpResponse == nil {
		return lHttpResponse, err
	}

	lBody, err := ioutil.ReadAll(lHttpResponse.Body)
	lHttpResponse.Body.Close()
	if err != nil {
		return lHttpResponse, err
	}

	if lHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  lBody,
			error: lHttpResponse.Status,
		}
		return lHttpResponse, newErr
	}

	return lHttpResponse, nil
}
