package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		rollbackErr := tx.Rollback()
		PanicIfError(rollbackErr)
		panic(err)
	} else {
		commitErr := tx.Commit()
		PanicIfError(commitErr)
	}
}
