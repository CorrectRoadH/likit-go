package likit

import (
	"context"

	"buf.build/gen/go/likit/likit/grpc/go/api/v1/apiv1grpc"
	v1 "buf.build/gen/go/likit/likit/protocolbuffers/go/api/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type VoteServer struct {
	voteClient apiv1grpc.VoteServiceClient
}

func NewVoteServer(host string, tls bool) *VoteServer {
	var conn *grpc.ClientConn
	var err error
	if tls {
		conn, err = grpc.Dial(host, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
		if err != nil {
			panic(err)
		}
	} else {
		conn, err = grpc.Dial(host, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
	}

	client := apiv1grpc.NewVoteServiceClient(conn)
	return &VoteServer{
		voteClient: client,
	}
}

type VoteResponse struct {
	Status int   `json:"status"`
	Count  int64 `json:"count"`
}

type IsVoteResponse struct {
	Status int  `json:"status"`
	IsVote bool `json:"isVote"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

func (s *VoteServer) Vote(ctx context.Context, businessId string, messageId string, userId string) (int64, error) {
	resp, err := s.voteClient.Vote(ctx, &v1.VoteRequest{
		BusinessId: businessId,
		MessageId:  messageId,
		UserId:     userId,
	})
	if err != nil {
		return 0, err
	}
	return resp.Count, nil
}

func (s *VoteServer) UnVote(ctx context.Context, businessId string, messageId string, userId string) (int64, error) {
	resp, err := s.voteClient.UnVote(ctx, &v1.VoteRequest{
		BusinessId: businessId,
		MessageId:  messageId,
		UserId:     userId,
	})
	if err != nil {
		return 0, err
	}
	return resp.Count, nil
}

func (s *VoteServer) Count(ctx context.Context, businessId string, messageId string) (int64, error) {
	resp, err := s.voteClient.Count(ctx, &v1.CountRequest{
		BusinessId: businessId,
		MessageId:  messageId,
	})
	if err != nil {
		return 0, err
	}
	return resp.Count, nil
}

func (s *VoteServer) IsVote(ctx context.Context, businessId string, messageId string, userId string) (bool, error) {
	resp, err := s.voteClient.IsVoted(ctx, &v1.IsVotedRequest{
		BusinessId: businessId,
		MessageId:  messageId,
		UserId:     userId,
	})
	if err != nil {
		return false, err
	}
	return resp.IsVoted, nil
}

func (s *VoteServer) VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error) {
	resp, err := s.voteClient.VotedUsers(ctx, &v1.VotedUsersRequest{
		BusinessId: businessId,
		MessageId:  messageId,
	})
	if err != nil {
		return nil, err
	}
	return resp.UserIds, nil
}
