package db

import (
	"github.com/webcustomerapi1/dbmodels"
	"github.com/webcustomerapi1/models"
	"strconv"	
)


// repository contains the details of a repository
type RepositorySummary struct {
	ID         int
	Ip       string
	DateConnect      string
}

type Repositories struct {
	Repositories []RepositorySummary
}

// queryRepos first fetches the repositories data from the db
func QueryRepos(repos *Repositories,id string) error {
	db:=dbmodels.GetDB()
	rows, err := db.Query("SELECT connect_id,ip,date_connect FROM connection where connect_id="+id)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		repo := RepositorySummary{}
		err = rows.Scan(
			&repo.ID,
			&repo.Ip,
			&repo.DateConnect,
		)
		if err != nil {
			return err
		}
		repos.Repositories = append(repos.Repositories, repo)
	}
	err = rows.Err()
	if err != nil {
		return err
	}


	return nil
}

var TransformConnection = func(p int64) *models.Connection {
	repos := Repositories{}
	var out *models.Connection
	t1:=strconv.Itoa(int(p))
	err := QueryRepos(&repos,t1)
	if err != nil {
		return &models.Connection{}
	}
	
	if len(repos.Repositories)>0 {
		out=&models.Connection{ID:int64(repos.Repositories[0].ID), DateConnect: repos.Repositories[0].DateConnect,IP: repos.Repositories[0].Ip}
	}else{
		out=&models.Connection{}
	}
	return out 
}