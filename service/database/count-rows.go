package database

// Get number of rows of a query result
func (db *appdbimpl) CountRows(table string, column string, event string) uint32 {
	// 1 - execute query
	var count uint32
	query := "SELECT COUNT(*) FROM " + table + " WHERE " + column + "=?"
	err := db.c.QueryRow(query, event).Scan(count)
	if err != nil {
		// queryrow returned error which means no user exists -> return 0
		return 0
	}
	// 2 - return query resylt
	return count
}
