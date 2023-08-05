package route

import (
	"elasticsearch/conf"
	"elasticsearch/handler"
	"elasticsearch/repo"
	"elasticsearch/service"
	"fmt"
)

type Service struct {
	*conf.App
}

func NewService() *Service {
	s := Service{
		conf.NewApp(),
	}

	db, err := s.GetDB()
	if err != nil {
		fmt.Errorf("cannot connect to Elastic search: %v", err)
	}
	repo := repo.NewRepo(db)

	esService := service.NewEsService(repo)
	esHandler := handler.NewESHandler(esService)

	router := s.Router

	v1 := router.Group("/v1")
	v1.GET("/test", esHandler.Test)
	v1.POST("/insert", esHandler.Insert)

	return &s
}
