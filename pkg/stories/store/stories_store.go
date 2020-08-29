package store

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"stories/pkg/stories/domain"
	"stories/pkg/stories/utils"
)

const (
	zero           = 0
	insertNewStory = "insert into stories (title, body) values ($1, $2) returning id"
	searchByTitle  = "select * from stories where title like $1"
	updateStory    = "update stories set title=$1, body=$2, updatedat=now() where id=$3"
	deleteStory    = "delete from stories where id=$1"
)

type StoriesStore interface {
	Add(story *domain.Story) (string, error)
	Search(query string) ([]domain.Story, error)
	Update(story *domain.Story) (int, error)
	Delete(id string) (int, error)
}

type defaultStoriesStore struct {
	db  *sql.DB
	lgr *zap.Logger
}

func (dss *defaultStoriesStore) Add(story *domain.Story) (string, error) {
	var id string
	err := dss.db.QueryRow(insertNewStory, story.GetTitle(), story.GetBody()).Scan(&id)
	if err != nil {
		dss.lgr.Error(err.Error())
		return utils.NA, err
	}

	return id, nil
}

func (dss *defaultStoriesStore) Search(query string) ([]domain.Story, error) {
	var stories []domain.Story

	rows, err := dss.db.Query(searchByTitle, fmt.Sprintf("%s%%", query))
	if err != nil {
		dss.lgr.Error(err.Error())
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var story domain.Story
		if err := rows.Scan(&story.ID, &story.Title, &story.Body, &story.CreatedAt, &story.UpdatedAt); err != nil {
			dss.lgr.Error(err.Error())
			return nil, err
		}

		stories = append(stories, story)
	}

	err = rows.Err()
	if err != nil {
		dss.lgr.Error(err.Error())
		return nil, err
	}

	return stories, nil
}

func (dss *defaultStoriesStore) Update(story *domain.Story) (int, error) {
	return execQuery(dss.lgr, dss.db, updateStory, story.GetTitle(), story.GetBody(), story.GetID())
}

func (dss *defaultStoriesStore) Delete(id string) (int, error) {
	return execQuery(dss.lgr, dss.db, deleteStory, id)
}

func execQuery(lgr *zap.Logger, db *sql.DB, query string, args ...interface{}) (int, error) {
	res, err := db.Exec(query, args...)
	if err != nil {
		lgr.Error(err.Error())
		return zero, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		lgr.Error(err.Error())
		return zero, err
	}

	return int(count), nil
}

func NewStoriesStore(db *sql.DB, lgr *zap.Logger) StoriesStore {
	return &defaultStoriesStore{
		db:  db,
		lgr: lgr,
	}
}
