package db

import (
	"github.com/webcustomerapi1/dbmodels"
	"github.com/webcustomerapi1/models"
	"database/sql"	
)


// repository contains the details of a repository
type CustomerSummary1 struct {
	ID         int
	Name       string
}

type Customers1 struct {
	Customers1 []CustomerSummary1
}

// queryRepos first fetches the repositories data from the db
func Querycustomer(repos *Customers1 ,id string) error {
	db:=dbmodels.GetDB()
	var rows *sql.Rows
	if (id=="null"){
		rows, _ = db.Query("SELECT customer_id,customer_name FROM customers")
	}else{
		rows, _ = db.Query("SELECT customer_id,customer_name FROM customers where customer_id="+id)
	}

	defer rows.Close()
	for rows.Next() {
		repo := CustomerSummary1{}
		err := rows.Scan(
			&repo.ID,
			&repo.Name,
		)
		if err != nil {
			return err
		}
		repos.Customers1 = append(repos.Customers1, repo)
	}
	err := rows.Err()
	if err != nil {
		return err
	}

	return nil
}

var TransformCustomerPost = func() *models.Customer {
	repos := Customers1{}
	var out *models.Customer
	t1:="dod"
	err := Querycustomer(&repos,t1)
	if err != nil {
		return &models.Customer{}
	}
	
	if len(repos.Customers1)>0 {
		out=&models.Customer{ID:int64(repos.Customers1[0].ID), Name: repos.Customers1[0].Name}
	}else{
		out=&models.Customer{}
	}
	return out 
}

var TransformCustomerPut = func() *models.Customer {
	repos := Customers1{}
	var out *models.Customer
	t1:="null"
	err := Querycustomer(&repos,t1)
	if err != nil {
		return &models.Customer{}
	}
	
	if len(repos.Customers1)>0 {
		out=&models.Customer{ID:int64(repos.Customers1[0].ID), Name: repos.Customers1[0].Name}
	}else{
		out=&models.Customer{}
	}
	return out 
}
