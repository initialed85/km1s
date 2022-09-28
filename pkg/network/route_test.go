package network

import (
	"github.com/initialed85/km1s/pkg/test"
	"testing"
)

func TestRoute(t *testing.T) {
	var err error

	bridgeInterfaceName := "bridge-1"
	ipAddress1 := "10.137.137.101"
	ipAddress2 := "10.137.137.1"
	ipAddress3 := "10.136.136.0"
	netMask := 24

	t.Run("SpecificRoute", func(t *testing.T) {
		defer func() {
			DelRoute(ipAddress3, netMask, ipAddress2, "")
			DelIPAddress(bridgeInterfaceName, "10.137.137.102", 24, "")
			DelInterface(bridgeInterfaceName, "")
		}()

		err = AddBridge(bridgeInterfaceName, "")
		test.FailOnError(t, err)
		test.AssertInterfaceExists(t, bridgeInterfaceName)

		err = AddIPAddress(bridgeInterfaceName, ipAddress1, netMask, "")
		test.FailOnError(t, err)
		test.AssertInterfaceHasIPAndPrefix(t, bridgeInterfaceName, ipAddress1, netMask)
		test.AssertIPAddressRespondsToPings(t, ipAddress1)

		err = AddRoute(ipAddress3, netMask, ipAddress2, "")
		test.FailOnError(t, err)

		test.AssertRouteExists(t, ipAddress3, netMask, ipAddress1)
	})

	t.Run("DefaultRoute", func(t *testing.T) {
		defer func() {
			DelDefaultRoute("")
			DelIPAddress(bridgeInterfaceName, "10.137.137.102", 24, "")
			DelInterface(bridgeInterfaceName, "")
		}()

		err = AddBridge(bridgeInterfaceName, "")
		test.FailOnError(t, err)
		test.AssertInterfaceExists(t, bridgeInterfaceName)

		err = AddIPAddress(bridgeInterfaceName, ipAddress1, netMask, "")
		test.FailOnError(t, err)
		test.AssertInterfaceHasIPAndPrefix(t, bridgeInterfaceName, ipAddress1, netMask)
		test.AssertIPAddressRespondsToPings(t, ipAddress1)

		err = AddDefaultRoute(ipAddress2, "")
		test.FailOnError(t, err)

		test.AssertDefaultRouteExists(t, ipAddress1)
	})
}
