package operators

import (
	"bufio"
	"net"
	"os"
	"strings"
)

func (opMap *OperatorMap) loadIPMatchFromFile() {
	fn := func(expression interface{}, variableData interface{}) bool {

		remoteAddressIp := net.ParseIP(variableData.(string))
		dataFile, err := os.Open(RulesAndDatasPath + expression.(string))

		if remoteAddressIp == nil || err != nil {
			return false
		}

		scanner := bufio.NewScanner(dataFile)
		for scanner.Scan() {
			ip := scanner.Text()

			if ip == "" || strings.HasPrefix(ip, "#") {
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
					return true
				}
			} else {
				if net.ParseIP(ip).Equal(remoteAddressIp) {
					return true
				}
			}
		}

		return false
	}

	opMap.funcMap["ipMatchF"] = fn
	opMap.funcMap["ipMatchFromFile"] = fn
}
