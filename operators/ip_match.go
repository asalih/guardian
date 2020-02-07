package operators

import (
	"net"
	"strings"

	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadIPMatch() {
	opMap.funcMap["ipMatch"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		ipAddresses := strings.Split(expression.(string), ",")
		remoteAddressIp := net.ParseIP(variableData.(string))

		matchResult := matches.NewMatchResult(false)

		if remoteAddressIp == nil {
			return matchResult
		}

		for _, ip := range ipAddresses {
			if ip == "" {
				continue
			}

			isCidrBlock := strings.Contains(ip, "/")

			if isCidrBlock {
				_, subnet, err := net.ParseCIDR(ip)

				//TODO: Add in error log in here
				if err != nil {
					continue
				}

				if subnet.Contains(remoteAddressIp) {
					return matchResult.SetMatch(true)
				}
			} else {
				if net.ParseIP(ip).Equal(remoteAddressIp) {
					return matchResult.SetMatch(true)
				}
			}
		}

		return matchResult
	}
}
