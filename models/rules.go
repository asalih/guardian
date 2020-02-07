package models

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var xDirectives = []string{"SecAction", "SecArgumentSeparator", "SecAuditEngine", "SecAuditLog", "SecAuditLog2", "SecAuditLogDirMode",
	"SecAuditLogFormat", "SecAuditLogFileMode", "SecAuditLogParts", "SecAuditLogRelevantStatus", "SecAuditLogStorageDir",
	"SecAuditLogType", "SecCacheTransformations", "SecChrootDir", "SecCollectionTimeout", "SecComponentSignature",
	"SecConnEngine", "SecContentInjection", "SecCookieFormat", "SecCookieV0Separator", "SecDataDir", "SecDebugLog", "SecDebugLogLevel",
	"SecDefaultAction", "SecDisableBackendCompression", "SecHashEngine", "SecHashKey", "SecHashParam", "SecHashMethodRx", "SecHashMethodPm",
	"SecGeoLookupDb", "SecGsbLookupDb", "SecGuardianLog", "SecHttpBlKey", "SecInterceptOnError", "SecMarker", "SecPcreMatchLimit",
	"SecPcreMatchLimitRecursion", "SecPdfProtect", "SecPdfProtectMethod", "SecPdfProtectSecret", "SecPdfProtectTimeout", "SecPdfProtectTokenName",
	"SecReadStateLimit", "SecConnReadStateLimit", "SecSensorId", "SecWriteStateLimit", "SecConnWriteStateLimit", "SecRemoteRules",
	"SecRemoteRulesFailAction", "SecRequestBodyAccess", "SecRequestBodyInMemoryLimit", "SecRequestBodyLimit", "SecRequestBodyNoFilesLimit",
	"SecRequestBodyLimitAction", "SecResponseBodyLimit", "SecResponseBodyLimitAction", "SecResponseBodyMimeType", "SecResponseBodyMimeTypesClear",
	"SecResponseBodyAccess", "SecRuleInheritance", "SecRuleEngine", "SecRulePerfTime", "SecRuleRemoveById", "SecRuleRemoveByMsg",
	"SecRuleRemoveByTag", "SecRuleScript", "SecRuleUpdateActionById", "SecRuleUpdateTargetById", "SecRuleUpdateTargetByMsg",
	"SecRuleUpdateTargetByTag", "SecServerSignature", "SecStatusEngine", "SecStreamInBodyInspection", "SecStreamOutBodyInspection", "SecTmpDir",
	"SecUnicodeMapFile", "SecUnicodeCodePage", "SecUploadDir", "SecUploadFileLimit", "SecUploadFileMode", "SecUploadKeepFiles",
	"SecWebAppId", "SecXmlExternalEntity"}

//RulesCollection Rules collection
var RulesCollection []*Rule

//InitRulesCollection Rules data initializer
func InitRulesCollection() {
	InitRulesCollectionFile("crs_xss.conf")
	InitRulesCollectionFile("crs_sqli.conf")
}

//InitRulesCollectionFile Rules data initializer
func InitRulesCollectionFile(path string) {
	confFile, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(confFile)
	var plainTextRules []string
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 1 || strings.HasPrefix(line, "#") {
			continue
		}

		readLine := strings.ReplaceAll(strings.TrimSuffix(strings.TrimSpace(line), "\r"), "\n", " ")

		if len(readLine) <= 1 {
			continue
		}

		plainTextRules = append(plainTextRules, readLine)
	}

	//fmt.Println(plainTextRules)

	plainTextRulesLen := len(plainTextRules)

	for i := 0; i < plainTextRulesLen; i++ {

		row := plainTextRules[i]
		if strings.HasPrefix(row, "SecRule") {
			for {
				li := i + 1

				if li >= plainTextRulesLen {
					break
				}

				lr := plainTextRules[li]

				if strings.HasPrefix(lr, "SecRule") {
					break
				}

				isDirective := false
				for _, dir := range xDirectives {
					if strings.HasPrefix(lr, dir) {
						isDirective = true
						break
					}
				}

				if isDirective {
					break
				}

				i = li
				row += lr
			}
		}

		parseRule(row)

	}
}

func parseRule(ruleTxt string) {
	variablesReg := regexp.MustCompile(`SecRule\s(.*?)\s`)
	operatorReg := regexp.MustCompile(`(\"@?.*?\")\s+?`)
	actionReg := regexp.MustCompile(`\"?(.+)\"?`)

	variablesMatch := variablesReg.FindString(ruleTxt)
	operatorMatch := operatorReg.FindString(ruleTxt)
	actionMatch := actionReg.FindString(ruleTxt)

	fmt.Println(variablesMatch)

	variables := parseVariables(variablesMatch)
	operators := parseOperators(operatorMatch)

	RulesCollection = append(RulesCollection, NewRule(variables, operators, actionMatch))
}

func parseVariables(variable string) []Variable {
	variable = strings.ReplaceAll(variable, "SecRule ", "")
	varsSplit := strings.Split(variable, "|")
	var dataVariable []Variable

	for _, vars := range varsSplit {
		varsAndFilter := strings.Split(vars, ":")

		if len(varsAndFilter) > 2 {
			//TODO Malformed rule
			continue
		}

		var v Variable

		isLengthCheck := varsAndFilter[0][0] == '&'
		if len(varsAndFilter) > 1 {
			isNotType := varsAndFilter[0][0] == '!'
			varName := strings.Trim(varsAndFilter[0], " ")

			if isNotType || isLengthCheck {
				varName = varName[1:]
			}

			v = Variable{varName, strings.Split(strings.Trim(varsAndFilter[1], " "), ","), isNotType, isLengthCheck}
		} else {
			varName := strings.Trim(varsAndFilter[0], " ")

			if isLengthCheck {
				varName = varName[1:]
			}

			v = Variable{varName, nil, false, isLengthCheck}
		}

		dataVariable = append(dataVariable, v)
	}

	return dataVariable
}

func parseOperators(operator string) []Operator {
	isNotOperator := strings.HasPrefix(operator, `"!`)
	isOperatorSpec := false

	if isNotOperator {
		isOperatorSpec = strings.HasPrefix(operator, `"!@`)
	} else {
		isOperatorSpec = strings.HasPrefix(operator, `"@`)
	}

	parsedOperator := "rx"
	parsedExpression := ""

	if isOperatorSpec {
		operatorReg := regexp.MustCompile(`@(.*?)\s`)
		opMatch := operatorReg.FindStringSubmatch(operator)

		opr := strings.NewReplacer("\"", "")
		parsedOperator = opr.Replace(opMatch[1])

		r := strings.NewReplacer(parsedOperator, "")
		parsedExpression = strings.TrimLeft(strings.Trim(r.Replace(operator), "\""), "")
		parsedExpression = strings.TrimLeft(parsedExpression, "@! ")

	} else {
		parsedExpression = strings.TrimLeft(strings.Trim(operator, "\""), "")
		parsedExpression = strings.TrimLeft(parsedExpression, "@! ")
	}

	return []Operator{Operator{parsedOperator, parsedExpression, isNotOperator}}
}
