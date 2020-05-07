package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/asalih/guardian/models"
	"github.com/google/uuid"

	//_ ...
	_ "github.com/lib/pq"
)

/*DBHelper The database query helper*/
type DBHelper struct {
}

/*NewDBHelper Inits new db helper*/
func NewDBHelper() *DBHelper {
	return new(DBHelper)
}

/*GetTarget Reads the Target from database*/
func (h *DBHelper) GetTarget(domain string) *models.Target {
	target := h.getTarget(domain)

	if target == nil {
		if strings.HasPrefix(domain, "www.") {
			return h.getTarget(strings.ReplaceAll(domain, "www.", ""))
		}

		return h.getTarget("www." + domain)
	}

	return target
}

func (h *DBHelper) getTarget(domain string) *models.Target {
	conn, err := sql.Open("postgres", models.Configuration.ConnectionString)
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
func (h *DBHelper) GetRequestFirewallRules(targetID string) []*models.Rule {
	conn, err := sql.Open("postgres", models.Configuration.ConnectionString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	rows, qerr := conn.Query("SELECT \"SerializedExpression\" FROM public.\"FirewallRules\" where \"IsActive\"=true and \"RuleFor\"=0 and \"TargetId\"= $1", targetID)

	if qerr != nil {
		panic(qerr)
	}

	result := []*models.Rule{}

	for rows.Next() {
		var fwRule string
		ferr := rows.Scan(&fwRule)

		if ferr != nil {
			panic(ferr)
		}

		var rules = []models.Rule{}
		if err = json.Unmarshal([]byte(fwRule), &rules); err == nil {
			for _, r := range rules {
				result = append(result, &r)
			}

		}
	}

	return result
}

//GetResponseFirewallRules Gets the firewall rule
func (h *DBHelper) GetResponseFirewallRules(targetID string) []*models.Rule {
	conn, err := sql.Open("postgres", models.Configuration.ConnectionString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	rows, qerr := conn.Query("SELECT \"SerializedExpression\" FROM public.\"FirewallRules\" where \"IsActive\"=true and \"RuleFor\"=1 and \"TargetId\"= $1", targetID)

	if qerr != nil {
		panic(qerr)
	}

	result := []*models.Rule{}

	for rows.Next() {
		var fwRule string
		ferr := rows.Scan(&fwRule)

		if ferr != nil {
			panic(ferr)
		}

		var rules = []models.Rule{}
		if err = json.Unmarshal([]byte(fwRule), &rules); err == nil {
			for _, r := range rules {
				result = append(result, &r)
			}

		}
	}

	return result
}

//LogMatchResult ...
func (h *DBHelper) LogMatchResult(
	ruleExecutionResult *models.RuleExecutionResult,
	ruleID string,
	target *models.Target,
	requestURI string,
	forResponse bool) {

	conn, err := sql.Open("postgres", models.Configuration.ConnectionString)
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

	//wafAction := models.GetWafAction(payload.Action)
	_, err = conn.Exec(sqlStatement,
		uuid.New(),
		time.Now(),
		target.ID, true,
		ruleExecutionResult.MatchResult.Elapsed,
		0,
		ruleID,
		requestURI,
		ruleFor,
		0 /*wafAction*/)

	if err != nil {
		panic(err)
	}

}

//LogHTTPRequest ...
func (h *DBHelper) LogHTTPRequest(log *models.HTTPLog) {
	conn, err := sql.Open("postgres", models.Configuration.ConnectionString)
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

//LogThrottleRequest ...
func (h *DBHelper) LogThrottleRequest(ipAddress string) {
	conn, err := sql.Open("postgres", models.Configuration.ConnectionString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	sqlStatement := `
INSERT INTO "ThrottleLogs" ("Id", "CreatedAt", "IPAddress", "ThrottleType")
VALUES ($1, $2, $3, $4)`

	_, err = conn.Exec(sqlStatement,
		uuid.New(),
		time.Now(),
		ipAddress,
		1)

	if err != nil {
		panic(err)
	}
}
