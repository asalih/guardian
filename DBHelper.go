package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var connString = "host=localhost port=5432 user=postgres password=1q2w3e dbname=GuardianDB sslmode=disable"

/*Target The target type*/
type Target struct {
	Domain          string
	OriginIPAddress string
	CertKey         string
	CertCrt         string
	UseHTTPS        bool
}

/*DBHelper The database query helper*/
type DBHelper struct {
}

/*GetTarget Reads the Target from database*/
func (h *DBHelper) GetTarget(domain string) *Target {
	conn, err := sql.Open("postgres", connString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	row := conn.QueryRow("SELECT \"Domain\", \"OriginIpAddress\", \"CertKey\", \"CertCrt\", \"UseHttps\" FROM public.\"Targets\" where \"Domain\"= $1", domain)

	var target = &Target{}
	rerr := row.Scan(&target.Domain, &target.OriginIPAddress, &target.CertKey, &target.CertCrt, &target.UseHTTPS)

	if rerr != nil {
		return nil
	}

	return target
}
