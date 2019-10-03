package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/asalih/guardian/models"
	"github.com/google/uuid"

	//_ ...
	_ "github.com/lib/pq"
)

var connString = "host=localhost port=5432 user=guardian password=1q2w3e dbname=guardiandb sslmode=disable"

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

	row := conn.QueryRow("SELECT \"Id\", \"Domain\", \"OriginIpAddress\", \"CertKey\", \"CertCrt\", \"UseHttps\", \"WAFEnabled\", \"Proto\" FROM public.\"Targets\" where \"State\"=1 and \"Domain\"= $1", domain)

	var target = &models.Target{}
	rerr := row.Scan(&target.ID, &target.Domain, &target.OriginIPAddress, &target.CertKey, &target.CertCrt, &target.UseHTTPS, &target.WAFEnabled, &target.Proto)

	if rerr != nil {
		fmt.Println(rerr)
		return nil
	}

	return target
}

//GetRequestFirewallRules Gets the firewall rule
func (h *DBHelper) GetRequestFirewallRules(targetID string) []*models.FirewallRule {
	conn, err := sql.Open("postgres", connString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	rows, qerr := conn.Query("SELECT \"Id\", \"Expression\" FROM public.\"FirewallRules\" where \"RuleFor\"=0 and \"TargetId\"= $1", targetID)

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

//GetResponseFirewallRules Gets the firewall rule
func (h *DBHelper) GetResponseFirewallRules(targetID string) []*models.FirewallRule {
	conn, err := sql.Open("postgres", connString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	rows, qerr := conn.Query("SELECT \"Id\", \"Expression\" FROM public.\"FirewallRules\" where \"RuleFor\"=1 and \"TargetId\"= $1", targetID)

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

//LogMatchResult ...
func (h *DBHelper) LogMatchResult(matchResult *models.MatchResult, target *models.Target, requestURI string, forResponse bool) {
	conn, err := sql.Open("postgres", connString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	sqlStatement := `
INSERT INTO "RuleLogs" ("Id", "CreatedAt", "TargetId", "IsHitted", "ExecutionMillisecond", "LogType", "Description", "RequestUri", "RuleFor")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	ruleFor := 0
	if forResponse {
		ruleFor = 1
	}

	_, err = conn.Exec(sqlStatement, uuid.New(), time.Now(), target.ID, true, matchResult.Elapsed, 0, matchResult.MatchedPayload.Payload, requestURI, ruleFor)
	if err != nil {
		panic(err)
	}
}

//LogFirewallMatchResult ...
func (h *DBHelper) LogFirewallMatchResult(matchResult *models.FirewallMatchResult, target *models.Target, requestURI string, forResponse bool) {
	conn, err := sql.Open("postgres", connString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	sqlStatement := `
INSERT INTO "RuleLogs" ("Id", "CreatedAt", "TargetId", "IsHitted", "ExecutionMillisecond", "LogType", "FirewallRuleId", "RequestUri", "RuleFor")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	ruleFor := 0
	if forResponse {
		ruleFor = 1
	}

	_, err = conn.Exec(sqlStatement, uuid.New(), time.Now(), target.ID, true, matchResult.Elapsed, 1, matchResult.FirewallRule.ID, requestURI, ruleFor)
	if err != nil {
		panic(err)
	}
}

//LogHTTPRequest ...
func (h *DBHelper) LogHTTPRequest(log *models.HTTPLog) {
	conn, err := sql.Open("postgres", connString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	sqlStatement := `
INSERT INTO "HTTPLogs" ("Id", "CreatedAt", "TargetId", "RequestUri", "StatusCode", "RuleCheckElapsed", "HttpElapsed", "RequestSize", "ResponseSize")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err = conn.Exec(sqlStatement, uuid.New(), time.Now(), log.TargetID, log.RequestURI, log.StatusCode, log.RuleCheckElapsed, log.HTTPElapsed, log.RequestSize, log.ResponseSize)
	if err != nil {
		panic(err)
	}
}
