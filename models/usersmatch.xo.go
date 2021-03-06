// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

// UsersMatch represents a row from 'public.users_matches'.
type UsersMatch struct {
	UserID  int `json:"user_id"`  // user_id
	MatchID int `json:"match_id"` // match_id
}

// GetAllUsersMatchs returns paginated most recent rows from 'public.users_matches',
// ordered by "created_at" in descending order.
func GetAllUsersMatchs(db XODB, page, resultsPerPage int) ([]*UsersMatch, error) {
	startIndex := (page - 1) * resultsPerPage
	const sqlstr = `SELECT ` +
		`user_id, match_id ` +
		`FROM public.users_matches ` +
		`ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	q, err := db.Query(sqlstr, resultsPerPage, startIndex)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*UsersMatch
	for q.Next() {
		um := UsersMatch{}

		// scan
		err = q.Scan(&um.UserID, &um.MatchID)
		if err != nil {
			return nil, err
		}

		res = append(res, &um)
	}

	return res, nil
}

// GetMostRecentUsersMatchs returns n most recent rows from 'public.users_matches',
// ordered by "created_at" in descending order.
func GetMostRecentUsersMatchs(db XODB, n int) ([]*UsersMatch, error) {
	const sqlstr = `SELECT ` +
		`user_id, match_id ` +
		`FROM public.users_matches ` +
		`ORDER BY created_at DESC LIMIT $1`

	q, err := db.Query(sqlstr, n)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*UsersMatch
	for q.Next() {
		um := UsersMatch{}

		// scan
		err = q.Scan(&um.UserID, &um.MatchID)
		if err != nil {
			return nil, err
		}

		res = append(res, &um)
	}

	return res, nil
}

// Match returns the Match associated with the UsersMatch's MatchID (match_id).
//
// Generated from foreign key 'users_matches_match_id_fk'.
func (um *UsersMatch) Match(db XODB) (*Match, error) {
	return MatchByID(db, um.MatchID)
}

// User returns the User associated with the UsersMatch's UserID (user_id).
//
// Generated from foreign key 'users_matches_user_id_fk'.
func (um *UsersMatch) User(db XODB) (*User, error) {
	return UserByID(db, um.UserID)
}

// UsersMatchesByMatchID retrieves a row from 'public.users_matches' as a UsersMatch.
//
// Generated from index 'users_matches_match_id_idx'.
func UsersMatchesByMatchID(db XODB, matchID int) ([]*UsersMatch, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`user_id, match_id ` +
		`FROM public.users_matches ` +
		`WHERE match_id = $1`

	// run query
	XOLog(sqlstr, matchID)
	q, err := db.Query(sqlstr, matchID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*UsersMatch{}
	for q.Next() {
		um := UsersMatch{}

		// scan
		err = q.Scan(&um.UserID, &um.MatchID)
		if err != nil {
			return nil, err
		}

		res = append(res, &um)
	}

	return res, nil
}

// UsersMatchesByUserID retrieves a row from 'public.users_matches' as a UsersMatch.
//
// Generated from index 'users_matches_user_id_idx'.
func UsersMatchesByUserID(db XODB, userID int) ([]*UsersMatch, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`user_id, match_id ` +
		`FROM public.users_matches ` +
		`WHERE user_id = $1`

	// run query
	XOLog(sqlstr, userID)
	q, err := db.Query(sqlstr, userID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*UsersMatch{}
	for q.Next() {
		um := UsersMatch{}

		// scan
		err = q.Scan(&um.UserID, &um.MatchID)
		if err != nil {
			return nil, err
		}

		res = append(res, &um)
	}

	return res, nil
}
