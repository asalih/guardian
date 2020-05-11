package operators

import (
	"net"
	"strings"
)

func init() {
	OperatorMaps.funcMap["ipMatch"] = func(expression interface{}, variableData interface{}) bool {

		ipAddresses := strings.Split(expression.(string), ",")
		remoteAddressIP := net.ParseIP(variableData.(string))

		if remoteAddressIP == nil {
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

				if subnet.Contains(remoteAddressIP) {
					return true
				}
			} else {
				if net.ParseIP(ip).Equal(remoteAddressIP) {
					return true
				}
			}
		}

		return false
	}
}
