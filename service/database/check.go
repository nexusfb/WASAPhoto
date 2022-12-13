package database

// function to check if occurrence exists in database
func (db *appdbimpl) Check(table string, column1 string, column2 string, v1 string, v2 string) bool {
	// 1 - build query to check if an occurrence exists or not
	query := `SELECT * FROM ` + table + ` WHERE ` + column1 + `=? AND ` + column2 + `=?`

	// 2 - execute query
	r, err := db.c.Query(query, v1, v2)
	defer func() { _ = r.Close() }()
	if r.Next() == false {
		// no row rows -> false
		return false
	}
	if err == nil {
		// row found means occurrence
		return true
	}

	// 3 - error occurred when executing query -> return false (for sure no potentially harmful operation is performed)
	return false

}
