package services

import (
	"github.com/ahfrd/grpc/micro-topup/src/client"
	"github.com/ahfrd/grpc/micro-topup/src/proto/topup"
	"github.com/ahfrd/grpc/micro-topup/src/repository"
)

type TopUpService struct {
	topup.UnimplementedTopUpServiceServer
	repository.TopUpRepository
	EmoneySvc  client.EmoneyServiceClient
	HistorySvc client.HistoryServiceClient
}
