package models

func GetMatchesByUser(db XODB, userID int) ([]*Match, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, home_team, away_team, home_user_id, away_user_id, home_points, away_points ` +
		`FROM public.matches ` +
		`WHERE home_user_id = $1 OR away_user_id = $1`

	// run query
	XOLog(sqlstr, userID)
	q, err := db.Query(sqlstr, userID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Match{}
	for q.Next() {
		m := Match{
			_exists: true,
		}

		// scan
		err = q.Scan(&m.ID, &m.HomeTeam, &m.AwayTeam, &m.HomeUserID, &m.AwayUserID, &m.HomePoints, &m.AwayPoints)
		if err != nil {
			return nil, err
		}

		res = append(res, &m)
	}

	return res, nil
}
