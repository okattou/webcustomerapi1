package db

import (
	"github.com/webcustomerapi1/dbmodels"
	"github.com/webcustomerapi1/models"
	"strconv"
	"database/sql"		
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
	var rows *sql.Rows
	if (id=="null"){
		rows, _ = db.Query("SELECT connect_id,ip,date_connect FROM connection")
	}else{
		rows, _ = db.Query("SELECT connect_id,ip,date_connect FROM connection where connect_id="+id)
	}

	defer rows.Close()
	for rows.Next() {
		repo := RepositorySummary{}
		err := rows.Scan(
			&repo.ID,
			&repo.Ip,
			&repo.DateConnect,
		)
		if err != nil {
			return err
		}
		repos.Repositories = append(repos.Repositories, repo)
	}
	err := rows.Err()
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


var TransformConnectionall = func() []*models.Connection {
	repos := Repositories{}
	t1:="null"
	err := QueryRepos(&repos,t1)
	if err != nil {
		return nil
	}
	data:= make([]*models.Connection,0)
	for i := 0; i < len(repos.Repositories); i++ {
		data =append(data, &models.Connection{ID:int64(repos.Repositories[i].ID), DateConnect: repos.Repositories[i].DateConnect,IP: repos.Repositories[i].Ip})
	}
	return data
}