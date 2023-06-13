package client

import (
	"context"
	"fmt"

	"github.com/ahfrd/grpc/micro-topup/config"
	"github.com/ahfrd/grpc/micro-topup/src/model/request"
	proto "github.com/ahfrd/grpc/micro-topup/src/proto/history"
	"google.golang.org/grpc"
)

type HistoryServiceClient struct {
	Client proto.HistoryServiceClient
}

func InitHistoryServiceClient(conf *config.Config) HistoryServiceClient {
	cc, err := grpc.Dial(conf.HistorySvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	c := HistoryServiceClient{
		Client: proto.NewHistoryServiceClient(cc),
	}
	return c
}

func (c *HistoryServiceClient) InsertHistoryLog(params *request.InsertHistoryTransaction) (*proto.GeneralResponse, error) {

	var bodyReq proto.InsertHistoryLogRequest
	bodyReq.WalletId = params.WalletId
	bodyReq.TransactionName = params.TransactionName
	bodyReq.NominalTransaction = params.NominalTransaction
	bodyReq.FeeTransaction = params.FeeTransaction
	bodyReq.CategoryTransaction = params.CategoryTransaction
	bodyReq.Status = params.Status
	bodyReq.TransactionService = params.TransactionService
	bodyReq.RefferenceNumber = params.RefferenceNumber

	res, err := c.Client.InsertHistoryLog(context.Background(), &bodyReq)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (c *HistoryServiceClient) UpdateStatusLog(params *request.UpdateStatusRequest) (*proto.GeneralResponse, error) {

	var bodyReq proto.UpdateStatusLogRequest
	bodyReq.LastId = params.LastId
	bodyReq.Status = params.Status

	res, err := c.Client.UpdateStatusLog(context.Background(), &bodyReq)
	if err != nil {
		return nil, err
	}
	return res, nil
}
