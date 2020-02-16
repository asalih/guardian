package operators

import (
	"net"
	"strings"
)

func (opMap *OperatorMap) loadIPMatch() {
	opMap.funcMap["ipMatch"] = func(expression interface{}, variableData interface{}) bool {

		ipAddresses := strings.Split(expression.(string), ",")
		remoteAddressIp := net.ParseIP(variableData.(string))

		if remoteAddressIp == nil {
			return false
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
}
