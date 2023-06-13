package repository

import (
	"database/sql"
	"fmt"

	"github.com/ahfrd/grpc/micro-topup/src/model/entity"
)

func (o TopUpRepository) SelectFeeTopUp(methodParams string) (entity.FeeTopUpEntity, *sql.DB, error) {
	var entityFeeTopUp entity.FeeTopUpEntity
	var id NullString
	var method NullString
	var fee NullInt
	db, err := o.ConnectDB()
	if err != nil {
		return entityFeeTopUp, db, err
	}
	var query string = fmt.Sprintf(`select id,method,fee from tbl_method_topup where method = "%s"`, methodParams)
	db.QueryRow(query).Scan(
		&id,
		&method,
		&fee,
	)
	fmt.Println(query)
	entityFeeTopUp.Id = id.String
	entityFeeTopUp.Method = method.String
	entityFeeTopUp.Fee = int(fee.Int64)
	defer db.Close()
	if err != nil && err != sql.ErrNoRows {
		return entityFeeTopUp, db, fmt.Errorf("failed Select SQL for tbl_method : %v", err)
	}

	return entityFeeTopUp, db, nil
}
