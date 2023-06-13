package repository

import (
	"database/sql"
	"fmt"

	"github.com/ahfrd/grpc/micro-topup/src/model"
)

func (o TopUpRepository) SelectMethodTopUp() ([]model.MethodTopUpModel, *sql.DB, error) {
	var ArrayDataMethodTopUp []model.MethodTopUpModel
	var resultDataTopUp model.MethodTopUpModel

	db, err := o.ConnectDB()
	defer db.Close()
	var query string = fmt.Sprintf(`select id,method from tbl_method_topup`)
	result, err := db.Query(query)
	fmt.Println(result)
	if err != nil {
		return ArrayDataMethodTopUp, db, fmt.Errorf("failed Select SQL for tbl_method_topup : %v", err)
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(
			&resultDataTopUp.Id,
			&resultDataTopUp.Method,
		)
		if err != nil {
			return ArrayDataMethodTopUp, db, fmt.Errorf("failed select tbl_method_topup query scan : %v", err.Error())
		}
		ArrayDataMethodTopUp = append(ArrayDataMethodTopUp, resultDataTopUp)
	}
	return ArrayDataMethodTopUp, db, nil
}
