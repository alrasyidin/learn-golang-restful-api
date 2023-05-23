package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		HandleIfPanicError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		HandleIfPanicError(errorCommit)
	}
}
