
/*
 * NSX-V
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type Class20SystemalarmsApiService service
/*
Class20SystemalarmsApiService servicesSystemAlarmsIdRead
Retrieve information about the specified alarm.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  alarmId: The alarm ID you want to manage. Find the alarm ID using the &#x60;GET /api/2.0/services/alarms/{source-Id}&#x60; method.   
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param alarmId

*/
func (a *Class20SystemalarmsApiService) Api20ServicesSystemalarmsAlarmIdGet(ctx context.Context, alarmId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/services/systemalarms/{alarmId}"
	localVarPath = strings.Replace(localVarPath, "{"+"alarmId"+"}", fmt.Sprintf("%v", alarmId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
Class20SystemalarmsApiService servicesSystemAlarmsIdAction
Resolve the specified alarm.  Alarms will resolve automatically when the cause of the alarm is resolved.  For example, if an NSX Edge appliance is powered off, this will trigger an alarm. If you power the NSX Edge appliance back on, the alarm will resolve. If however, you delete the NSX Edge appliance, the alarm will persist, because the alarm cause was never resolved. In this case, you may want to manually resolve the alarm. Resolving the alarm  will clear it from the NSX dashboard.  **Method history:**  Release | Modification --------|------------- 6.3.0 | Method introduced.   Parameters:  alarmId: The alarm ID you want to manage. Find the alarm ID using the &#x60;GET /api/2.0/services/alarms/{source-Id}&#x60; method.   
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param alarmId
 * @param optional nil or *Class20SystemalarmsApiApi20ServicesSystemalarmsAlarmIdPostOpts - Optional Parameters:
     * @param "Action" (optional.String) - 

*/

type Class20SystemalarmsApiApi20ServicesSystemalarmsAlarmIdPostOpts struct {
    Action optional.String
}

func (a *Class20SystemalarmsApiService) Api20ServicesSystemalarmsAlarmIdPost(ctx context.Context, alarmId string, localVarOptionals *Class20SystemalarmsApiApi20ServicesSystemalarmsAlarmIdPostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/2.0/services/systemalarms/{alarmId}"
	localVarPath = strings.Replace(localVarPath, "{"+"alarmId"+"}", fmt.Sprintf("%v", alarmId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Action.IsSet() {
		localVarQueryParams.Add("action", parameterToString(localVarOptionals.Action.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}