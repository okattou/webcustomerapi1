package db

import (
	"github.com/webcustomerapi1/dbmodels"
	"github.com/webcustomerapi1/models"
	"strconv"
	"database/sql"	
)


// repository contains the details of a repository
type CustomerSummary struct {
	ID         int
	Name       string
}

type Customers struct {
	Customers []CustomerSummary
}

// queryRepos first fetches the repositories data from the db
func QueryRepos1(repos *Customers ,id string) error {
	db:=dbmodels.GetDB()
	var rows *sql.Rows
	if (id=="null"){
		rows, _ = db.Query("SELECT customer_id,customer_name FROM customers")
	}else{
		rows, _ = db.Query("SELECT customer_id,customer_name FROM customers where customer_id="+id)
	}

	defer rows.Close()
	for rows.Next() {
		repo := CustomerSummary{}
		err := rows.Scan(
			&repo.ID,
			&repo.Name,
		)
		if err != nil {
			return err
		}
		repos.Customers = append(repos.Customers, repo)
	}
	err := rows.Err()
	if err != nil {
		return err
	}

	return nil
}

var TransformCustomer = func(p int64) *models.Customer {
	repos := Customers{}
	var out *models.Customer
	t1:=strconv.Itoa(int(p))
	err := QueryRepos1(&repos,t1)
	if err != nil {
		return &models.Customer{}
	}
	
	if len(repos.Customers)>0 {
		out=&models.Customer{ID:int64(repos.Customers[0].ID), Name: repos.Customers[0].Name}
	}else{
		out=&models.Customer{}
	}
	return out 
}


var TransformCustomerall = func() []*models.Customer {
	repos := Customers{}
	t1:="null"
	err := QueryRepos1(&repos,t1)
	if err != nil {
		return nil
	}
	data:= make([]*models.Customer,0)
	for i := 0; i < len(repos.Customers); i++ {
		data =append(data, &models.Customer{ID:int64(repos.Customers[i].ID), Name: repos.Customers[i].Name})
	}
	return data
}