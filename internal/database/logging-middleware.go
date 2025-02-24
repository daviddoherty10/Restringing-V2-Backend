package database

import (
	"Restringing-V2/entity"
	"errors"
	"time"
)

func (s *service) CreateLog(l entity.LoggingMiddleware) error {
	query := `INSERT INTO logging (statuscode,useragent,referer,requestmethod,requesturl,headers,responsebody,duration,created_at) VALUES (?, ?, ?, ?, ?, ?,?,?,?)`
	_, err := s.db.Exec(query, l.StatusCode, l.UserAgent, l.Referer, l.RequestMethod, l.RequestURL, l.Hearders, l.ResponseBodyStr, l.Duration, time.Now())
	if err != nil {
		return errors.New("failed to insert user: " + err.Error())
	}

	return nil
}
