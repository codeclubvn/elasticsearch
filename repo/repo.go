package repo

import (
	"bytes"
	"context"
	"elasticsearch/model"
	"encoding/json"
	"fmt"
	es "github.com/elastic/go-elasticsearch/v7"
)

type Repo struct {
	db *es.Client
}

func NewRepo(db *es.Client) *Repo {
	return &Repo{
		db: db,
	}
}

type IRepo interface {
	Test(ctx context.Context) (interface{}, error)
	Insert(ctx context.Context, ESRequest model.ESRequest) error
}

func (r *Repo) Test(ctx context.Context) (interface{}, error) {
	return r.db.Info()
}

func (r *Repo) Insert(ctx context.Context, ESRequest model.ESRequest) error {
	bodyBytes, err := json.Marshal(ESRequest.Body)
	if err != nil {
		fmt.Println("Error marshaling struct to bytes:", err)
		return err
	}
	_, err = r.db.Index(ESRequest.Index, bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}
	return nil
}
