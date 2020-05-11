package operators

import (
	"net"

	"strings"
)

func init() {
	fn := func(expression interface{}, variableData interface{}) bool {

		remoteAddressIP := net.ParseIP(variableData.(string))
		fileCache := DataFileCaches[expression.(string)]

		if remoteAddressIP == nil || fileCache == nil {
			return false
		}

		for _, ip := range fileCache.Lines {

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

	OperatorMaps.funcMap["ipMatchF"] = fn
	OperatorMaps.funcMap["ipMatchFromFile"] = fn
}
