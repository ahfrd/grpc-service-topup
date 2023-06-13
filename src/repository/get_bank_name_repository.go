package repository

import (
	"database/sql"
	"fmt"

	"github.com/ahfrd/grpc/micro-topup/src/model"
)

func (o TopUpRepository) SelectBankCode() ([]model.BankCodeModel, *sql.DB, error) {
	var ArrayDataBankCode []model.BankCodeModel
	var resultDataBankCode model.BankCodeModel

	db, err := o.ConnectDB()
	defer db.Close()
	var query string = fmt.Sprintf(`select code,bankName from tbl_bank_code`)
	result, err := db.Query(query)
	if err != nil {
		return ArrayDataBankCode, db, fmt.Errorf("failed Select SQL for tbl_bank_code : %v", err)
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(
			&resultDataBankCode.Code,
			&resultDataBankCode.BankName,
		)
		if err != nil {
			return ArrayDataBankCode, db, fmt.Errorf("failed select tbl_bank_code query scan : %v", err)
		}
		ArrayDataBankCode = append(ArrayDataBankCode, resultDataBankCode)
	}
	return ArrayDataBankCode, db, nil
}
