package domain

import (
	"database/sql"

	"github.com/bhaktiutama/banking/errs"
	"github.com/bhaktiutama/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// REPOSITORY - CONTROLLER
// ADAPTER - Implementation of the port
// No business process happen here. Only plain function to Database.

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql :=
			"select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
		// rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql :=
			"select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
		// rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while scanning customers " + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected database error")
	// }

	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateofBirth, &c.Status)

	// 	if err != nil {
	// 		logger.Error("Error while scanning customers " + err.Error())
	// 		return nil, errs.NewUnexpectedError("Unexpected database error")
	// 	}
	// 	customers = append(customers, c)
	// }
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql :=
		"select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	// row := d.client.QueryRow(customerSql, id)
	var c Customer

	err := d.client.Get(&c, customerSql, id)
	// err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	// dbUser := os.Getenv("DB_USER")
	// dbPasswd := os.Getenv("DB_PASSWD")
	// dbAddr := os.Getenv("DB_ADDR")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")
	// dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	// client, err := sqlx.Open("mysql", dataSource)
	// if err != nil {
	// 	panic(err)
	// }
	// // See "Important settings" section.
	// client.SetConnMaxLifetime(time.Minute * 3)
	// client.SetMaxOpenConns(10)
	// client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{dbClient}
}
