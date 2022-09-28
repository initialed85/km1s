package network

import (
	"github.com/initialed85/km1s/pkg/test"
	"testing"
)

func TestBridge(t *testing.T) {
	var err error

	ipAddress1 := "10.137.137.101"
	netMask := 24
	macVlanInterfaceName := "macvlan-1"
	macVlanParentInterfaceName := "eth0"
	bridgeInterfaceName := "bridge-1"

	t.Run("HappyPath", func(t *testing.T) {
		defer func() {
			DelInterface(macVlanInterfaceName, "")
			DelIPAddress(bridgeInterfaceName, ipAddress1, netMask, "")
			DelInterface(bridgeInterfaceName, "")
		}()

		err = AddBridge(bridgeInterfaceName, "")
		test.FailOnError(t, err)
		test.AssertInterfaceExists(t, bridgeInterfaceName)

		err = AddIPAddress(bridgeInterfaceName, ipAddress1, netMask, "")
		test.FailOnError(t, err)
		test.AssertInterfaceHasIPAndPrefix(t, bridgeInterfaceName, ipAddress1, netMask)
		test.AssertIPAddressRespondsToPings(t, ipAddress1)

		err = AddMACVLAN(macVlanInterfaceName, macVlanParentInterfaceName, "")
		test.FailOnError(t, err)
		test.AssertInterfaceExists(t, macVlanInterfaceName)

		err = AddBridgePort(macVlanInterfaceName, bridgeInterfaceName, "")
		test.FailOnError(t, err)
		test.AssertInterfaceHasBridge(t, macVlanInterfaceName, bridgeInterfaceName)
	})
}
