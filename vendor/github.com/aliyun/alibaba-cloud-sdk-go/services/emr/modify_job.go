package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyJob invokes the emr.ModifyJob API synchronously
func (client *Client) ModifyJob(request *ModifyJobRequest) (response *ModifyJobResponse, err error) {
	response = CreateModifyJobResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyJobWithChan invokes the emr.ModifyJob API asynchronously
func (client *Client) ModifyJobWithChan(request *ModifyJobRequest) (<-chan *ModifyJobResponse, <-chan error) {
	responseChan := make(chan *ModifyJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyJob(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyJobWithCallback invokes the emr.ModifyJob API asynchronously
func (client *Client) ModifyJobWithCallback(request *ModifyJobRequest, callback func(response *ModifyJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyJobResponse
		var err error
		defer close(result)
		response, err = client.ModifyJob(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyJobRequest is the request struct for api ModifyJob
type ModifyJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Type            string           `position:"Query" name:"Type"`
	FailAct         string           `position:"Query" name:"FailAct"`
	RunParameter    string           `position:"Query" name:"RunParameter"`
	RetryInterval   requests.Integer `position:"Query" name:"RetryInterval"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	Name            string           `position:"Query" name:"Name"`
	Id              string           `position:"Query" name:"Id"`
	MaxRetry        requests.Integer `position:"Query" name:"MaxRetry"`
}

// ModifyJobResponse is the response struct for api ModifyJob
type ModifyJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyJobRequest creates a request to invoke ModifyJob API
func CreateModifyJobRequest() (request *ModifyJobRequest) {
	request = &ModifyJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ModifyJob", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyJobResponse creates a response to parse from ModifyJob response
func CreateModifyJobResponse() (response *ModifyJobResponse) {
	response = &ModifyJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}