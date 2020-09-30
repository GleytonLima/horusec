// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/ZupIT/horusec/development-kit/pkg/enums/errors"
	httpResponse "github.com/ZupIT/horusec/development-kit/pkg/utils/http-request/response"
	"github.com/ZupIT/horusec/development-kit/pkg/utils/logger"
)

type Interface interface {
	DoRequest(req *http.Request, tlsConfig *tls.Config) (httpResponse.Interface, error)
}

type Client struct {
	timeout int
}

func NewHTTPClient(timeout int) Interface {
	return &Client{
		timeout: timeout,
	}
}

// nolint
func (c *Client) DoRequest(req *http.Request, tlsConfig *tls.Config) (res httpResponse.Interface, err error) {
	client := &http.Client{
		Timeout: time.Duration(c.timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	response, err := client.Do(req)
	if err != nil {
		logger.LogError(errors.ErrDoHTTPClient.Error(), err)
		return res, err
	}
	return httpResponse.NewHTTPResponse(response), nil
}