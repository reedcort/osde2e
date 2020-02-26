/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// CredentialsServer represents the interface the manages the 'credentials' resource.
type CredentialsServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the credentials of a cluster.
	Get(ctx context.Context, request *CredentialsGetServerRequest, response *CredentialsGetServerResponse) error
}

// CredentialsGetServerRequest is the request for the 'get' method.
type CredentialsGetServerRequest struct {
}

// CredentialsGetServerResponse is the response for the 'get' method.
type CredentialsGetServerResponse struct {
	status int
	err    *errors.Error
	body   *ClusterCredentials
}

// Body sets the value of the 'body' parameter.
//
//
func (r *CredentialsGetServerResponse) Body(value *ClusterCredentials) *CredentialsGetServerResponse {
	r.body = value
	return r
}

// Status sets the status code.
func (r *CredentialsGetServerResponse) Status(value int) *CredentialsGetServerResponse {
	r.status = value
	return r
}

// dispatchCredentials navigates the servers tree rooted at the given server
// till it finds one that matches the given set of path segments, and then invokes
// the corresponding server.
func dispatchCredentials(w http.ResponseWriter, r *http.Request, server CredentialsServer, segments []string) {
	if len(segments) == 0 {
		switch r.Method {
		case "GET":
			adaptCredentialsGetRequest(w, r, server)
			return
		default:
			errors.SendMethodNotAllowed(w, r)
			return
		}
	}
	switch segments[0] {
	default:
		errors.SendNotFound(w, r)
		return
	}
}

// adaptCredentialsGetRequest translates the given HTTP request into a call to
// the corresponding method of the given server. Then it translates the
// results returned by that method into an HTTP response.
func adaptCredentialsGetRequest(w http.ResponseWriter, r *http.Request, server CredentialsServer) {
	request := &CredentialsGetServerRequest{}
	err := readCredentialsGetRequest(request, r)
	if err != nil {
		glog.Errorf(
			"Can't read request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	response := &CredentialsGetServerResponse{}
	response.status = 200
	err = server.Get(r.Context(), request, response)
	if err != nil {
		glog.Errorf(
			"Can't process request for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		errors.SendInternalServerError(w, r)
		return
	}
	err = writeCredentialsGetResponse(response, w)
	if err != nil {
		glog.Errorf(
			"Can't write response for method '%s' and path '%s': %v",
			r.Method, r.URL.Path, err,
		)
		return
	}
}
