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

var connString = "host=167.71.46.213 port=5432 user=guardian password=#&Lx&M7c$7E^Zrda dbname=guardiandb sslmode=disable"

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

	row := conn.QueryRow("SELECT \"Id\", \"Domain\", \"OriginIpAddress\", \"CertKey\", \"CertCrt\", \"AutoCert\", \"UseHttps\", \"WAFEnabled\", \"Proto\" FROM public.\"Targets\" where \"State\"=1 and \"Domain\"= $1", domain)

	var target = &models.Target{}
	rerr := row.Scan(&target.ID,
		&target.Domain,
		&target.OriginIPAddress,
		&target.CertKey,
		&target.CertCrt,
		&target.AutoCert,
		&target.UseHTTPS,
		&target.WAFEnabled,
		&target.Proto)

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

	rows, qerr := conn.Query("SELECT \"Id\", \"Expression\", \"Action\" FROM public.\"FirewallRules\" where \"IsActive\"=true and \"RuleFor\"=0 and \"TargetId\"= $1", targetID)

	if qerr != nil {
		panic(qerr)
	}

	result := make([]*models.FirewallRule, 0)

	for rows.Next() {
		var fwRule = &models.FirewallRule{}
		ferr := rows.Scan(&fwRule.ID, &fwRule.Expression, &fwRule.Action)

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

	rows, qerr := conn.Query("SELECT \"Id\", \"Expression\", \"Action\" FROM public.\"FirewallRules\" where \"IsActive\"=true and \"RuleFor\"=1 and \"TargetId\"= $1", targetID)

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
func (h *DBHelper) LogMatchResult(
	matchResult *models.MatchResult,
	payload *models.PayloadData,
	target *models.Target,
	requestURI string,
	forResponse bool) {

	conn, err := sql.Open("postgres", connString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	sqlStatement := `
INSERT INTO "RuleLogs" ("Id", "CreatedAt", "TargetId", "IsHitted", "ExecutionMillisecond", "LogType", "Description", "RequestUri", "RuleFor", "WafAction")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	ruleFor := 0
	if forResponse {
		ruleFor = 1
	}

	wafAction := models.GetWafAction(payload.Action)
	_, err = conn.Exec(sqlStatement,
		uuid.New(),
		time.Now(),
		target.ID, true,
		matchResult.Elapsed,
		models.LogTypeWAF,
		payload.Payload,
		requestURI,
		ruleFor,
		wafAction)

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
INSERT INTO "RuleLogs" ("Id", "CreatedAt", "TargetId", "IsHitted", "ExecutionMillisecond", "LogType", "FirewallRuleId", "RequestUri", "RuleFor", "WafAction")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	ruleFor := 0
	if forResponse {
		ruleFor = 1
	}

	_, err = conn.Exec(sqlStatement,
		uuid.New(),
		time.Now(),
		target.ID,
		true,
		matchResult.Elapsed,
		models.LogTypeFirewall,
		matchResult.FirewallRule.ID,
		requestURI,
		ruleFor,
		matchResult.FirewallRule.Action)

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
INSERT INTO "HTTPLogs" ("Id", "CreatedAt", "TargetId", "RequestUri", "StatusCode", "RequestRulesCheckElapsed", "ResponseRulesCheckElapsed", "HttpElapsed", "RequestSize", "ResponseSize")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err = conn.Exec(sqlStatement,
		uuid.New(),
		time.Now(),
		log.TargetID,
		log.RequestURI,
		log.StatusCode,
		log.RequestRulesCheckElapsed,
		log.ResponseRulesCheckElapsed,
		log.HTTPElapsed,
		log.RequestSize,
		log.ResponseSize)

	if err != nil {
		panic(err)
	}
}
