package database

// Get number of rows of a quert
func (db *appdbimpl) CountRows(table string, column string, event string) uint32 {
	// 1 - execute query
	var count int
	query := "SELECT COUNT(*) FROM " + table + " WHERE " + column + "=?"
	err := db.c.QueryRow(query, event).Scan(count)
	if err != nil {
		// queryrow returned error which means no user exists -> return error
		return 0
	}
	// 2 - return result userID
	return count
}
