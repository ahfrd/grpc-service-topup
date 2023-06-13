package client

import (
	"context"
	"fmt"

	"github.com/ahfrd/grpc/micro-topup/config"
	"github.com/ahfrd/grpc/micro-topup/src/model/entity"
	proto "github.com/ahfrd/grpc/micro-topup/src/proto/emoney"
	"google.golang.org/grpc"
)

type EmoneyServiceClient struct {
	Client proto.EmoneyServiceClient
}

func InitEmoneyServiceClient(conf *config.Config) EmoneyServiceClient {
	cc, err := grpc.Dial(conf.EmoneySvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	c := EmoneyServiceClient{
		Client: proto.NewEmoneyServiceClient(cc),
	}
	return c
}

func (c *EmoneyServiceClient) GetWalletProfile(phoneNumber string) (*proto.GetWalletProfileResponse, error) {

	bodyReq := &proto.GetWalletProfileRequest{
		PhoneNumber: phoneNumber,
	}
	res, err := c.Client.GetWalletProfile(context.Background(), bodyReq)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (c *EmoneyServiceClient) GetWalletInfo(phoneNumber string) (*proto.GetWalletInfoResponse, error) {
	bodyReq := &proto.GetWalletInfoRequest{
		PhoneNumber: phoneNumber,
	}

	res, err := c.Client.GetWalletInfo(context.Background(), bodyReq)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *EmoneyServiceClient) InsertTransaction(req *entity.UpdateBallance) (*proto.InsertTransactionHistoryResponse, error) {
	bodyReq := &proto.InsertTransactionHistoryRequest{
		PhoneNumber:         req.PhoneNumber,
		CategoryTransaction: req.CategoryTransaction,
		NominalTransaction:  req.NominalTransaction,
		SecurityCode:        req.SecurityCode,
		FeeTransaction:      req.FeeTransaction,
	}
	res, err := c.Client.InsertTransaction(context.Background(), bodyReq)
	if err != nil {
		return nil, err

	}
	return res, nil
}

// func GetWalletInfo(ctx *gin.Context, c proto.EmoneyServiceClient) {
// 	bodyReq := request.GeneralRequestBody{}

// 	if err := ctx.BindJSON(&bodyReq); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}
// 	res, err := c.GetWalletInfo(context.Background(), &proto.GetWalletInfoRequest{
// 		PhoneNumber: bodyReq.PhoneNumber,
// 	})
// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadGateway, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, &res)
// }
