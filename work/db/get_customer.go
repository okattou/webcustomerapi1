package db

import (
	"github.com/webcustomerapi1/dbmodels"
	"github.com/webcustomerapi1/models"
	"strconv"	
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
	rows, err := db.Query("SELECT customer_id,customer_name FROM customers where customer_id="+id)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		repo := CustomerSummary{}
		err = rows.Scan(
			&repo.ID,
			&repo.Name,
		)
		if err != nil {
			return err
		}
		repos.Customers = append(repos.Customers, repo)
	}
	err = rows.Err()
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
		out=&models.Customer{ID:int64(repos.Customers[0].ID), NameClient: repos.Customers[0].Name}
	}else{
		out=&models.Customer{}
	}
	return out 
}