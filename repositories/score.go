package repositories

import (
	"github.com/gugadev/triviago/database"
	"github.com/gugadev/triviago/models"
	"github.com/gugadev/triviago/responses"
)

// ScoreRepo repository for score
type ScoreRepo struct{}

/*
Insert insert a new score
*/
func (r *ScoreRepo) Insert(score models.Score) {
	db := database.Connect()
	defer db.Close() // nolint: gas,errcheck

	db.Create(&score)
}

/*
GetStats get stats from database
*/
func (r *ScoreRepo) GetStats() []responses.Stat {
	db := database.Connect()
	defer db.Close() // nolint: gas,errcheck

	stats := []responses.Stat{}

	db.Table("scores").Select("categories.name as 'category', count(scores.id) as 'total', sum(case when scores.hit = 1 then 1 else 0 end) as 'hits'").Joins("inner join questions ON questions.id = scores.question_id").Joins("inner join categories ON categories.id = questions.category_id").Group("categories.name").Scan(&stats)

	return stats
}
