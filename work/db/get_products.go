package db

import (
	"github.com/webcustomerapi1/dbmodels"
	"github.com/webcustomerapi1/models"
	"strconv"
	"database/sql"		
)


// repository contains the details of a repository
type RepoProduitsSummary struct {
	ID         int
	Ip       string
	DateConnect      string
}

type RepoProducts struct {
	RepoProducts []RepoProduitsSummary
}

// QueryProduct first fetches the RepoProducts data from the db
func QueryProduct(repos *RepoProducts,id string) error {
	db:=dbmodels.GetDB()
	var rows *sql.Rows
	if (id=="null"){
		rows, _ = db.Query("SELECT connect_id,ip,date_connect FROM connection")
	}else{
		rows, _ = db.Query("SELECT connect_id,ip,date_connect FROM connection where connect_id="+id)
	}

	defer rows.Close()
	for rows.Next() {
		repo := RepoProduitsSummary{}
		err := rows.Scan(
			&repo.ID,
			&repo.Ip,
			&repo.DateConnect,
		)
		if err != nil {
			return err
		}
		repos.RepoProducts = append(repos.RepoProducts, repo)
	}
	err := rows.Err()
	if err != nil {
		return err
	}


	return nil
}

var TransformProduit = func(p int64) *models.Connection {
	repos := RepoProducts{}
	var out *models.Connection
	t1:=strconv.Itoa(int(p))
	err := QueryProduct(&repos,t1)
	if err != nil {
		return &models.Connection{}
	}
	
	if len(repos.RepoProducts)>0 {
		out=&models.Connection{ID:int64(repos.RepoProducts[0].ID), DateConnect: repos.RepoProducts[0].DateConnect,IP: repos.RepoProducts[0].Ip}
	}else{
		out=&models.Connection{}
	}
	return out 
}


var TransformProduitall = func() []*models.Connection {
	repos := RepoProducts{}
	t1:="null"
	err := QueryProduct(&repos,t1)
	if err != nil {
		return nil
	}
	data:= make([]*models.Connection,0)
	for i := 0; i < len(repos.RepoProducts); i++ {
		data =append(data, &models.Connection{ID:int64(repos.RepoProducts[i].ID), DateConnect: repos.RepoProducts[i].DateConnect,IP: repos.RepoProducts[i].Ip})
	}
	return data
}