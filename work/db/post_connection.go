package db

import (
	"github.com/webcustomerapi1/dbmodels"
	"github.com/webcustomerapi1/models"
	"database/sql"
	"fmt"	
)


// repository contains the details of a repository
type PostRespostSummary struct {
	ID         int
	Ip       string
	DateConnect      string
}

type Reposts struct {
	Reposts []PostRespostSummary
}

// QueryInsert first fetches the Reposts data from the db
func QueryInsert(repos *Reposts,id string,Ip string,date string) error {
	db:=dbmodels.GetDB()
	var rows *sql.Rows
	if (id=="null"){
		rows, _ = db.Query("SELECT connect_id,ip,date_connect FROM connection")
	}else{
		_, er := db.Exec(`INSERT INTO connection(ip,date_connect) VALUES($1, NOW())`,Ip)
		fmt.Println(er)
		rows, _ = db.Query("select connect_id,ip,date_connect from connection order by connect_id desc")
	}
	
	defer rows.Close()
	for rows.Next() {
		repo := PostRespostSummary{}
		err := rows.Scan(
			&repo.ID,
			&repo.Ip,
			&repo.DateConnect,
		)
		if err != nil {
			return err
		}
		repos.Reposts = append(repos.Reposts, repo)
	}
	err := rows.Err()
	if err != nil {
		return err
	}


	return nil
}

var TransformConnectionPost = func(ip *string,date *string) *models.Connection {
	repos := Reposts{}
	var out *models.Connection
	t1:= *ip
	t2:= *date
	err := QueryInsert(&repos,"post",t1,t2)
	if err != nil {
		return &models.Connection{}
	}
	
	if len(repos.Reposts)>0 {
		out=&models.Connection{ID:int64(repos.Reposts[0].ID), DateConnect: repos.Reposts[0].DateConnect,IP: repos.Reposts[0].Ip}
	}else{
		out=&models.Connection{}
	}
	return out 
}

var TransformConnectionPut = func(ip *string,date *string) *models.Connection {
	repos := Reposts{}
	var out *models.Connection
	t1:= *ip
	t2:= *date
	err := QueryInsert(&repos,"post",t1,t2)
	if err != nil {
		return &models.Connection{}
	}
	
	if len(repos.Reposts)>0 {
		out=&models.Connection{ID:int64(repos.Reposts[0].ID), DateConnect: repos.Reposts[0].DateConnect,IP: repos.Reposts[0].Ip}
	}else{
		out=&models.Connection{}
	}
	return out 
}