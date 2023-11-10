// Code generated by "requestgen -method GET -url /rates/history -type RatesHistoryRequest -responseType []Rate"; DO NOT EDIT.

package wise

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
)

func (r *RatesHistoryRequest) Source(source string) *RatesHistoryRequest {
	r.source = source
	return r
}

func (r *RatesHistoryRequest) Target(target string) *RatesHistoryRequest {
	r.target = target
	return r
}

func (r *RatesHistoryRequest) Length(length int) *RatesHistoryRequest {
	r.length = length
	return r
}

func (r *RatesHistoryRequest) Resolution(resolution Resolution) *RatesHistoryRequest {
	r.resolution = resolution
	return r
}

func (r *RatesHistoryRequest) Unit(unit Unit) *RatesHistoryRequest {
	r.unit = unit
	return r
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (r *RatesHistoryRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (r *RatesHistoryRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check source field -> json key source
	source := r.source

	// assign parameter of source
	params["source"] = source
	// check target field -> json key target
	target := r.target

	// assign parameter of target
	params["target"] = target
	// check length field -> json key length
	length := r.length

	// assign parameter of length
	params["length"] = length
	// check resolution field -> json key resolution
	resolution := r.resolution

	// assign parameter of resolution
	params["resolution"] = resolution
	// check unit field -> json key unit
	unit := r.unit

	// assign parameter of unit
	params["unit"] = unit

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (r *RatesHistoryRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := r.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if r.isVarSlice(_v) {
			r.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (r *RatesHistoryRequest) GetParametersJSON() ([]byte, error) {
	params, err := r.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (r *RatesHistoryRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (r *RatesHistoryRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (r *RatesHistoryRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (r *RatesHistoryRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (r *RatesHistoryRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := r.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (r *RatesHistoryRequest) Do(ctx context.Context) ([]Rate, error) {

	// empty params for GET operation
	var params interface{}
	query, err := r.GetParametersQuery()
	if err != nil {
		return nil, err
	}

	apiURL := "/rates/history"

	req, err := r.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := r.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse []Rate
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	return apiResponse, nil
}
