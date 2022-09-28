package network

import (
	"github.com/initialed85/km1s/pkg/test"
	"testing"
)

func TestMACVLAN(t *testing.T) {
	var err error

	ipAddress1 := "10.137.137.101"
	netMask := 24
	macVlanInterfaceName := "macvlan-1"
	macVlanParentInterfaceName := "eth0"

	t.Run("HappyPath", func(t *testing.T) {
		defer func() {
			DelIPAddress(macVlanInterfaceName, ipAddress1, netMask, "")
			DelInterface(macVlanInterfaceName, "")
		}()

		err = AddMACVLAN(macVlanInterfaceName, macVlanParentInterfaceName, "")
		test.FailOnError(t, err)
		test.AssertInterfaceExists(t, macVlanInterfaceName)

		err = AddIPAddress(macVlanInterfaceName, ipAddress1, netMask, "")
		test.FailOnError(t, err)
		test.AssertInterfaceHasIPAndPrefix(t, macVlanInterfaceName, ipAddress1, netMask)
		test.AssertIPAddressRespondsToPings(t, ipAddress1)
	})
}
