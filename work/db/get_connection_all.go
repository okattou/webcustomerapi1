package db

import (
	"github.com/webcustomerapi1/dbmodels"
	"github.com/webcustomerapi1/models"
	"strconv"	
)


// DataConnection contains the details of a DataConnection
type DataConnectionSummary struct {
	ID         int
	Ip       string
	DateConnect      string
}

type DataConnectiers struct {
	DataConnectiers []DataConnectionSummary
}

//QueryRepos3 first fetches the DataConnectiers data from the db
funcQueryRepos3(repos *DataConnectiers,id string) error {
	db:=dbmodels.GetDB()
	rows, err := db.Query("SELECT connect_id,ip,date_connect FROM connection")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		repo := DataConnectionSummary{}
		err = rows.Scan(
			&repo.ID,
			&repo.Ip,
			&repo.DateConnect,
		)
		if err != nil {
			return err
		}
		repos.DataConnectiers = append(repos.DataConnectiers, repo)
	}
	err = rows.Err()
	if err != nil {
		return err
	}


	return nil
}

var TransformConnectionall = func() *models.Connection {
	repos := DataConnectiers{}
	var out *models.Connection
	t1:=strconv.Itoa(int(p))
	err :=QueryRepos3(&repos,t1)
	if err != nil {
		return &models.Connection{}
	}
	
	if len(repos.DataConnectiers)>0 {
		out=&models.Connection{ID:int64(repos.DataConnectiers[0].ID), DateConnect: repos.DataConnectiers[0].DateConnect,IP: repos.DataConnectiers[0].Ip}
	}else{
		out=&models.Connection{}
	}
	return out 
}