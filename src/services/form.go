package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ahfrd/grpc/micro-topup/src/model/entity"
	"github.com/ahfrd/grpc/micro-topup/src/proto/topup"
	util "github.com/ahfrd/grpc/micro-topup/src/utils"

	"google.golang.org/protobuf/types/known/anypb"
)

func (o TopUpService) Form(ctx context.Context, req *topup.FormRequest) (*topup.GeneralResponse, error) {
	listTransaction, db, err := o.SelectMethodTopUp()
	db.Close()
	if err != nil {
		return &topup.GeneralResponse{
			RessponseCode:   http.StatusBadRequest,
			ResponseMessage: err.Error(),
		}, nil
	}
	if listTransaction[0].Id == "" {
		return &topup.GeneralResponse{
			RessponseCode:   http.StatusBadRequest,
			ResponseMessage: "Data Not Found",
		}, nil
	}
	listBank, db, err := o.SelectBankCode()
	db.Close()
	if err != nil {
		return &topup.GeneralResponse{
			RessponseCode:   http.StatusBadRequest,
			ResponseMessage: err.Error(),
		}, nil
	}
	var dataArr = []entity.FormStructEntity{}
	var data = entity.FormStructEntity{}
	for _, item := range listTransaction {
		data.Id = item.Id
		data.Method = item.Method
		if data.Method == "Bank Transfer" || data.Method == "Virtual Account" {
			for _, details := range listBank {
				data.Details = append(data.Details, entity.BankNameEntity{
					Code:     details.Code,
					BankName: details.BankName,
				})
			}
		} else {
			data.Details = append(data.Details, entity.BankNameEntity{
				Code:     "",
				BankName: "",
			})
		}
		dataArr = append(dataArr, data)
	}
	byteData, err := json.Marshal(dataArr)
	if err != nil {
		fmt.Println(err)
	}

	anyData := &anypb.Any{
		TypeUrl: util.TypeUrlStringProto(),
		Value:   byteData,
	}

	return &topup.GeneralResponse{
		RessponseCode:   http.StatusAccepted,
		ResponseMessage: "Succses",
		ResponseData:    anyData,
	}, nil
}
