package data

import (
	"database/sql"

	"github.com/asalih/guardian/models"
	//_ ...
	_ "github.com/lib/pq"
)

var connString = "host=localhost port=5432 user=postgres password=1q2w3e dbname=GuardianDB sslmode=disable"

/*DBHelper The database query helper*/
type DBHelper struct {
}

/*GetTarget Reads the Target from database*/
func (h *DBHelper) GetTarget(domain string) *models.Target {
	conn, err := sql.Open("postgres", connString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	row := conn.QueryRow("SELECT \"Id\", \"Domain\", \"OriginIpAddress\", \"CertKey\", \"CertCrt\", \"UseHttps\", \"WAFEnabled\" FROM public.\"Targets\" where \"Domain\"= $1", domain)

	var target = &models.Target{}
	rerr := row.Scan(&target.ID, &target.Domain, &target.OriginIPAddress, &target.CertKey, &target.CertCrt, &target.UseHTTPS, &target.WAFEnabled)

	if rerr != nil {
		return nil
	}

	return target
}

//GetFirewallRules Gets the firewall rule
func (h *DBHelper) GetFirewallRules(targetID string) []*models.FirewallRule {
	conn, err := sql.Open("postgres", connString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	rows, qerr := conn.Query("SELECT \"Id\", \"Expression\" FROM public.\"FirewallRules\" where \"TargetId\"= $1", targetID)

	if qerr != nil {
		panic(qerr)
	}

	result := make([]*models.FirewallRule, 0)

	for rows.Next() {
		var fwRule = &models.FirewallRule{}
		ferr := rows.Scan(&fwRule.ID, &fwRule.Expression)

		if ferr != nil {
			panic(ferr)
		}

		result = append(result, fwRule)
	}

	return result
}
