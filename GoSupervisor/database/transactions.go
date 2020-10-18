package database

// GetChildLinks sends a DB query to obtain
// all the child links for a given URL
func GetChildLinks(Parent string) ([]string, error) {
	db := GetDB()

	// SELECT
	// 	parent AS parent,
	// 	link AS link
	// FROM
	// 	links
	// WHERE
	// 	parent = $1;

	sqlStatement := "SELECT parent AS parent, link AS link FROM links WHERE parent = $1;"
	rows, err := db.Query(sqlStatement, Parent)
	if err != nil {
		return []string{}, err
	}

	defer rows.Close()

	links := make([]string, 0)

	for rows.Next() {
		var link string
		err = rows.Scan(&Parent, &link)
		if err != nil {
			return []string{}, err
		}
		links = append(links, link)
	}
	err = rows.Err()
	if err != nil {
		return []string{}, err
	}

	return links, nil
}
