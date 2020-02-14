package models

import (
	"bufio"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/asalih/guardian/operators"
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
	files, _ := ioutil.ReadDir(operators.RulesAndDatasPath)

	for _, v := range files {
		if v.IsDir() || !strings.HasSuffix(v.Name(), ".conf") {
			continue
		}

		InitRulesCollectionFile(operators.RulesAndDatasPath + v.Name())
	}
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

	plainTextRulesLen := len(plainTextRules)

	for i := 0; i < plainTextRulesLen; i++ {
		row := plainTextRules[i]
		if strings.HasPrefix(row, "SecRule") {
			var rule *Rule
			rule, i = walk(plainTextRules, i, plainTextRulesLen)
			RulesCollection = append(RulesCollection, rule)
		}
	}
}

func walk(plainTextRules []string, i int, plainTextRulesLen int) (*Rule, int) {
	row := plainTextRules[i]
	var chainRule *Rule
	for {
		li := i + 1

		if li >= plainTextRulesLen {
			break
		}

		lr := plainTextRules[li]

		if strings.HasPrefix(lr, "SecRule") {
			if strings.HasPrefix(plainTextRules[li-1], "chain") || strings.HasPrefix(plainTextRules[li-1], "\"chain") {
				chainRule, i = walk(plainTextRules, li, plainTextRulesLen)
			}
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

	rule := parseRule(row)

	if chainRule != nil {
		rule.Chain = chainRule
	}

	return rule, i
}

func parseRule(ruleTxt string) *Rule {

	variablesReg := regexp.MustCompile(`SecRule\s(.*?)\s`)
	operatorReg := regexp.MustCompile(`(\"@?.*?\")\s+?`)

	variablesMatch := variablesReg.FindString(ruleTxt)
	operatorMatch := operatorReg.FindString(ruleTxt)

	if variablesMatch == "" {
		return nil
	}

	variables := parseVariables(variablesMatch)
	operators := parseOperators(operatorMatch)
	action := parseAction(operatorReg.ReplaceAllString(variablesReg.ReplaceAllString(ruleTxt, ""), ""))

	return NewRule(variables, operators, action, nil)
}

func parseVariables(variable string) []*Variable {
	variable = strings.ReplaceAll(variable, "SecRule ", "")
	varsSplit := strings.Split(variable, "|")
	var dataVariable []*Variable

	for _, vars := range varsSplit {
		varsAndFilter := strings.Split(vars, ":")

		if len(varsAndFilter) > 2 {
			//TODO Malformed rule
			continue
		}

		var v *Variable

		isLengthCheck := varsAndFilter[0][0] == '&'
		if len(varsAndFilter) > 1 {
			isNotType := varsAndFilter[0][0] == '!'
			varName := strings.Trim(varsAndFilter[0], " ")

			if isNotType || isLengthCheck {
				varName = varName[1:]
			}

			v = &Variable{varName, strings.Split(strings.Trim(varsAndFilter[1], " "), ","), isNotType, isLengthCheck}
		} else {
			varName := strings.Trim(varsAndFilter[0], " ")

			if isLengthCheck {
				varName = varName[1:]
			}

			v = &Variable{varName, nil, false, isLengthCheck}
		}

		dataVariable = append(dataVariable, v)
	}

	return dataVariable
}

func parseOperators(operator string) *Operator {
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
		parsedExpression = r.Replace(operator)

		parsedExpression = strings.Trim(parsedExpression, "\" ")
		parsedExpression = strings.TrimLeft(parsedExpression, "@! ")

	} else {
		parsedExpression = strings.TrimLeft(strings.Trim(operator, "\""), "")
		parsedExpression = strings.TrimLeft(parsedExpression, "@! ")
	}

	return &Operator{parsedOperator, parsedExpression, isNotOperator}
}

func parseAction(action string) *Action {
	idReg := regexp.MustCompile(`id:(.*?),`)
	phaseReg := regexp.MustCompile(`phase:(.*?),`)

	idRegMatch := idReg.FindStringSubmatch(action)
	idRegIdentified := "-1"

	if len(idRegMatch) > 1 {
		idRegIdentified = idRegMatch[1]
	}

	phaseRegMatch := phaseReg.FindStringSubmatch(action)
	phaseRegIdentified := 1

	if len(phaseRegMatch) > 1 {
		phaseRegIdentified, _ = strconv.Atoi(phaseRegMatch[1])
	}

	disrupAct := DisruptiveActionBlock

	if strings.Contains(action, DisruptiveActionPass.ToString()+",") {
		disrupAct = DisruptiveActionPass
	} else if strings.Contains(action, DisruptiveActionDrop.ToString()+",") {
		disrupAct = DisruptiveActionDrop
	} else if strings.Contains(action, DisruptiveActionDeny.ToString()+",") {
		disrupAct = DisruptiveActionDeny
	} else if strings.Contains(action, DisruptiveActionProxy.ToString()+",") {
		disrupAct = DisruptiveActionProxy
	}

	return &Action{idRegIdentified, phaseRegIdentified, disrupAct, LogActionLog}
}
