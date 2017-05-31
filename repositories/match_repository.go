package repositories

import (
	"github.com/rymccue/fifa-tracker/models"
)

func GetMatchByID(id int) (*models.Match, error) {
	return models.MatchByID(db, id)
}

func GetMatchesByUser(userID int) ([]*models.Match, error) {
	return models.GetMatchesByUser(db, userID)
}

func CreateMatch(homeTeam, awayTeam string, homeUserID, awayUserID, homePoints, awayPoints int) (*models.Match, error) {
	m := &models.Match{
		HomeTeam:   homeTeam,
		AwayTeam:   awayTeam,
		HomeUserID: homeUserID,
		AwayUserID: awayUserID,
		HomePoints: homePoints,
		AwayPoints: awayPoints,
	}
	err := m.Insert(db)
	return m, err
}

func DeleteMatch(m *models.Match) error {
	return m.Delete(db)
}

func GetAllMatches(page, resultsPerPage int) ([]*models.Match, error) {
	return models.GetAllMatchs(db, page, resultsPerPage)
}

func SaveMatch(m *models.Match) error {
	return m.Save(db)
}
