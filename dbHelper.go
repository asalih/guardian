package main

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

var connString = "server=localhost\\SQLExpress;database=GuardianDB;user id=sa;password=1q2w3e;port=1433"

//The target
type Target struct {
	Domain          string
	OriginIPAddress string
	CertKey         string
	CertCrt         string
	UseHTTPS        bool
}

//The database query helper
type DBHelper struct {
}

//Reads the Target
func (h *DBHelper) GetTarget(domain string) *Target {
	conn, err := sql.Open("mssql", connString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	row := conn.QueryRow("SELECT Domain, OriginIpAddress, CertCrt, CertKey, UseHttps FROM Targets Where Domain=?;", domain)

	var target = &Target{}
	rerr := row.Scan(&target.Domain, &target.OriginIPAddress, &target.CertCrt, &target.CertKey, &target.UseHTTPS)

	if rerr != nil {
		return nil
	}

	return target
}
