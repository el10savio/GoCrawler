package database

// InsertLink sends a DB query to insert
// a child link for a given URL
func InsertLink(Parent string, Link string) error {
	db := GetDB()

	// INSERT INTO
	// 	links (parent, link)
	// VALUES
	// 	($1, $2)
	// ON CONFLICT
	// 	DO NOTHING;

	sqlStatement := "INSERT INTO links (parent, link) VALUES ($1, $2) ON CONFLICT DO NOTHING"
	_, err := db.Exec(sqlStatement, Parent, Link)
	if err != nil {
		return err
	}

	return nil
}

// RemoveLink sends a DB query to delete
// a given link row for a given URL
func RemoveLink(Link string) error {
	db := GetDB()

	// DELETE FROM
	// 	links
	// WHERE
	// 	link = $1;

	sqlStatement := "DELETE FROM links WHERE link = $1;"
	_, err := db.Exec(sqlStatement, Link)
	if err != nil {
		return err
	}

	return nil
}

// GetParentCount sends a DB query to obtain
// a count of all the links under a given URL
func GetParentCount(Parent string) (int, error) {
	var count int
	db := GetDB()

	// SELECT
	// 	COUNT(*)
	// FROM
	// 	links
	// WHERE
	// 	Parent = $1;

	sqlStatement := "SELECT COUNT(*) FROM links WHERE Parent = $1;"
	row := db.QueryRow(sqlStatement, Parent)
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
