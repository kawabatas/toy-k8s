/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package cloudcfg is ...
package cloudcfg

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/kawabatas/toy-k8s/pkg/client"
)

func promptForString(field string) string {
	fmt.Printf("Please enter %s: ", field)
	var result string
	fmt.Scan(&result)
	return result
}

// Parse an AuthInfo object from a file path
func LoadAuthInfo(path string) (client.AuthInfo, error) {
	var auth client.AuthInfo
	if _, err := os.Stat(path); os.IsNotExist(err) {
		auth.User = promptForString("Username")
		auth.Password = promptForString("Password")
		data, err := json.Marshal(auth)
		if err != nil {
			return auth, err
		}
		err = os.WriteFile(path, data, 0600)
		return auth, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return auth, err
	}
	err = json.Unmarshal(data, &auth)
	return auth, err
}

// RequestWithBody is a helper method that creates an HTTP request with the specified url, method
// and a body read from 'configFile'
// FIXME: need to be public API?
func RequestWithBody(configFile, url, method string) (*http.Request, error) {
	if len(configFile) == 0 {
		return nil, fmt.Errorf("empty config file")
	}
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	return RequestWithBodyData(data, url, method)
}

// RequestWithBodyData is a helper method that creates an HTTP request with the specified url, method
// and body data
// FIXME: need to be public API?
func RequestWithBodyData(data []byte, url, method string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	request.ContentLength = int64(len(data))
	return request, err
}

// Execute a request, adds authentication, and HTTPS cert ignoring.
// TODO: Make this stuff optional
// FIXME: need to be public API?
func DoRequest(request *http.Request, user, password string) (string, error) {
	request.SetBasicAuth(user, password)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	return string(body), err
}
