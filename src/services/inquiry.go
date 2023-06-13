package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ahfrd/grpc/micro-topup/src/model/entity"
	"github.com/ahfrd/grpc/micro-topup/src/proto/topup"
	"github.com/ahfrd/grpc/micro-topup/src/utils"
	"google.golang.org/protobuf/types/known/anypb"
)

func (o TopUpService) Inquiry(ctx context.Context, req *topup.InquiryRequest) (*topup.GeneralResponse, error) {
	randomNumb, err := utils.GenerateRandomNumber(3)
	if err != nil {
		return &topup.GeneralResponse{
			RessponseCode:   http.StatusBadRequest,
			ResponseMessage: err.Error(),
		}, nil
	}
	strRandNumb := strconv.Itoa(randomNumb)
	var strAdd string
	if req.Method == "Bank Transfer" {
		yearNow := time.Now().Year()
		strY := strconv.Itoa(yearNow)
		fmt.Println(strY[2:4])
		intY, _ := strconv.Atoi(strY)
		monthNow := time.Now().Month().String()
		intM, _ := strconv.Atoi(monthNow)
		validCard := strings.Split(req.CardInfo.ValidUntil, "/")
		intVY, _ := strconv.Atoi(validCard[1])
		intVM, _ := strconv.Atoi(validCard[0])
		if intY < intVY || intM < intVM {
			return &topup.GeneralResponse{
				RessponseCode:   http.StatusBadRequest,
				ResponseMessage: "Sorry, Your card not valid",
			}, nil
		}
		req.CodeNumber = req.BankCode + req.PhoneNumber
		strAdd = req.BankCode
	} else {
		req.CodeNumber = strRandNumb + req.PhoneNumber
		strAdd = strRandNumb
	}

	selectFee, db, err := o.SelectFeeTopUp(req.Method)
	db.Close()
	if err != nil {
		return &topup.GeneralResponse{
			RessponseCode:   http.StatusBadRequest,
			ResponseMessage: err.Error(),
		}, nil
	}
	if selectFee.Id == "" {
		return &topup.GeneralResponse{
			RessponseCode:   http.StatusNotFound,
			ResponseMessage: "Sorry, Data not found",
		}, nil
	}
	//time ddMMYY
	timeNow := time.Now().Format("020106")
	roman := utils.IntegerToRoman(timeNow[2:6])
	transactionName := roman + "/" + timeNow + "/" + req.Method + "/" + strAdd
	feeStr := strconv.Itoa(selectFee.Fee)
	var responseData entity.InquiryEntity
	responseData.Method = req.Method
	responseData.Nominal = req.NominalTopUp
	responseData.Fee = feeStr
	responseData.TransactionName = transactionName
	responseData.PhoneNumb = req.PhoneNumber
	responseData.Card.CardNumb = req.CardInfo.NomorKartu
	responseData.Card.SecurityNum = req.CardInfo.SecurityNum
	responseData.Card.Valid = req.CardInfo.ValidUntil
	responseData.CodeNumb = req.CodeNumber

	byteData, err := json.Marshal(responseData)
	if err != nil {
		fmt.Println(err)
	}

	anyData := &anypb.Any{
		TypeUrl: utils.TypeUrlStringProto(),
		Value:   byteData,
	}
	return &topup.GeneralResponse{
		RessponseCode:   http.StatusAccepted,
		ResponseMessage: "Succses",
		ResponseData:    anyData,
	}, nil

}
