package services

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ahfrd/grpc/micro-topup/src/model/entity"
	"github.com/ahfrd/grpc/micro-topup/src/model/request"
	"github.com/ahfrd/grpc/micro-topup/src/proto/topup"
	"github.com/ahfrd/grpc/micro-topup/src/utils"
)

func (o TopUpService) Payment(ctx context.Context, req *topup.PaymentRequest) (*topup.GeneralResponse, error) {
	randomNumb, err := utils.GenerateRandomNumber(9)
	if err != nil {
		return &topup.GeneralResponse{
			RessponseCode:   http.StatusBadRequest,
			ResponseMessage: err.Error(),
		}, nil
	}
	strRanNumber := strconv.Itoa(randomNumb)
	timeNow := time.Now().Format("010206")
	timeSec := time.Now().Format("15:04:05")
	sanitzieSecond := utils.SanitizeString(timeSec)
	reffNumber := strRanNumber + timeNow + sanitzieSecond + strRanNumber

	var entityInsert request.InsertHistoryTransaction
	entityInsert.WalletId = req.WalletId
	entityInsert.TransactionName = req.TransactionName
	entityInsert.NominalTransaction = req.NominalTopUp
	entityInsert.FeeTransaction = req.FeeTopUp
	entityInsert.CategoryTransaction = "TopUp"
	entityInsert.Status = "1"
	entityInsert.TransactionService = "TopUp - " + req.Method
	entityInsert.RefferenceNumber = reffNumber

	lastIdHistory, err := o.HistorySvc.InsertHistoryLog(&entityInsert)
	if err != nil {
		return &topup.GeneralResponse{
			RessponseCode:   http.StatusBadRequest,
			ResponseMessage: err.Error(),
		}, nil
	}
	fmt.Println(lastIdHistory)
	entityUpdateWallet := &entity.UpdateBallance{
		PhoneNumber:         req.PhoneNumber,
		CategoryTransaction: "TopUp",
		NominalTransaction:  req.NominalTopUp,
		SecurityCode:        req.SecurityCode,
		FeeTransaction:      req.FeeTopUp,
	}
	_, err = o.EmoneySvc.InsertTransaction(entityUpdateWallet)
	if err != nil {
		return &topup.GeneralResponse{
			RessponseCode:   http.StatusBadRequest,
			ResponseMessage: err.Error(),
		}, nil
	}

	entityInsertHistoryTopup := &entity.InsertHistoryTopUpEntity{
		ReffNum:          reffNumber,
		Nominal:          req.NominalTopUp,
		Fee:              req.FeeTopUp,
		TopUpDestination: req.PhoneNumber,
		TopUpSource:      req.Method,
		HistoryId:        lastIdHistory.String(),
	}
	_, db, err := o.InsertHistoryTransactionTopUpRepository(entityInsertHistoryTopup)
	db.Close()
	if err != nil {
		return &topup.GeneralResponse{
			RessponseCode:   http.StatusBadRequest,
			ResponseMessage: err.Error(),
		}, nil
	}

	requestUpdateLog := &request.UpdateStatusRequest{
		Status: "2",
		LastId: lastIdHistory.String(),
	}
	_, err = o.HistorySvc.UpdateStatusLog(requestUpdateLog)
	if err != nil {
		return &topup.GeneralResponse{
			RessponseCode:   http.StatusBadRequest,
			ResponseMessage: err.Error(),
		}, nil
	}
	return &topup.GeneralResponse{
		RessponseCode:   http.StatusAccepted,
		ResponseMessage: "Succses",
	}, nil
}
