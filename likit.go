package likit

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type VoteServer struct {
	host string
}

func NewVoteServer(host string) *VoteServer {
	return &VoteServer{
		host: fmt.Sprintf("%s/api/v1", host),
	}
}

type VoteResponse struct {
	Status int   `json:"status"`
	Count  int64 `json:"count"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

func (s *VoteServer) Vote(ctx context.Context, businessId string, messageId string, userId string) (int64, error) {
	body := struct {
		BusinessId string `json:"businessId"`
		MessageId  string `json:"messageId"`
		UserId     string `json:"userId"`
	}{
		BusinessId: businessId,
		MessageId:  messageId,
		UserId:     userId,
	}

	postBody, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

	// http post the host
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/vote", s.host), bytes.NewBuffer(postBody))
	if err != nil {
		return 0, err
	}

	// get response
	client := &http.Client{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)

		var errorResponse ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			return 0, err
		}

		return 0, fmt.Errorf(errorResponse.Message)
	}

	// get value from response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var voteResponse VoteResponse
	err = json.Unmarshal(respBody, &voteResponse)
	if err != nil {
		fmt.Sprintln("result", string(respBody))
		return 0, err
	}

	return voteResponse.Count, nil
}

func (s *VoteServer) UnVote(ctx context.Context, businessId string, messageId string, userId string) (int64, error) {
	body := struct {
		BusinessId string `json:"businessId"`
		MessageId  string `json:"messageId"`
		UserId     string `json:"userId"`
	}{
		BusinessId: businessId,
		MessageId:  messageId,
		UserId:     userId,
	}

	postBody, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}
	// http post the host
	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/unvote", s.host), bytes.NewBuffer(postBody))
	if err != nil {
		return 0, err
	}

	// get response
	client := &http.Client{}
	// add header json content
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)

		var errorResponse ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			return 0, err
		}

		return 0, fmt.Errorf(errorResponse.Message)
	}

	// get value from response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var voteResponse VoteResponse
	err = json.Unmarshal(respBody, &voteResponse)
	if err != nil {
		return 0, err
	}

	return voteResponse.Count, nil
}

func (s *VoteServer) Count(ctx context.Context, businessId string, messageId string) (int64, error) {
	// http post the host
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/count/%s/%s", s.host, businessId, messageId), nil)
	if err != nil {
		return 0, err
	}

	// get response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// get value from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var voteResponse VoteResponse
	err = json.Unmarshal(body, &voteResponse)
	if err != nil {
		return 0, err
	}

	return voteResponse.Count, nil
}

func (s *VoteServer) IsVote(ctx context.Context, businessId string, messageId string, userId string) (bool, error) {
	return true, nil
}
