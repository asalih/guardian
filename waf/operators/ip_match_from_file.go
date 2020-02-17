package operators

import (
	"net"

	"strings"
)

func (opMap *OperatorMap) loadIPMatchFromFile() {
	fn := func(expression interface{}, variableData interface{}) bool {

		remoteAddressIp := net.ParseIP(variableData.(string))
		fileCache := DataFileCaches[expression.(string)]

		if remoteAddressIp == nil || fileCache == nil {
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
