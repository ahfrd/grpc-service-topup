package repository

import (
	"database/sql"
	"fmt"

	"github.com/ahfrd/grpc/micro-topup/src/model/entity"
)

func (o TopUpRepository) InsertHistoryTransactionTopUpRepository(req *entity.InsertHistoryTopUpEntity) (int64, *sql.DB, error) {
	var err error
	var res sql.Result
	var prepare *sql.Stmt
	db, err := o.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return 0, nil, fmt.Errorf("%s", err)
	}
	defer db.Close()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	queryInsert := "INSERT INTO tbl_transaksi (refferenceNumber,nominalTopUp,feeTopUp,topUpDestination,topUpSource,historyId) values (?,?,?,?,?,?)"
	prepare, err = db.Prepare(queryInsert)
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert tbl_transaksi SQL : %v", err)
	}
	res, err = prepare.Exec(req.ReffNum, req.Nominal, req.Fee, req.TopUpDestination, req.TopUpSource)

	if err != nil {
		return 0, db, fmt.Errorf("failed to insert error_general on tbl_transaksi SQL : %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, db, fmt.Errorf("failed to populate status inserted : %v", err)
	}
	return count, db, err
}
