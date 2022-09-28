package network

import (
	"github.com/initialed85/km1s/pkg/test"
	"testing"
)

func TestVETH(t *testing.T) {
	var err error

	ipAddress1 := "10.137.137.101"
	ipAddress2 := "10.137.137.102"
	netMask := 24
	vethInterface1Name := "veth-1a"
	vethInterface2Name := "veth-1b"

	t.Run("HappyPath", func(t *testing.T) {
		defer func() {
			DelIPAddress(vethInterface1Name, ipAddress1, netMask, "")
			DelIPAddress(vethInterface2Name, ipAddress2, netMask, "")
			DelInterface(vethInterface1Name, "")
			DelInterface(vethInterface2Name, "")
		}()

		err = AddVETH(vethInterface1Name, vethInterface2Name, "")
		test.FailOnError(t, err)
		test.AssertInterfaceExists(t, vethInterface1Name)
		test.AssertInterfaceExists(t, vethInterface2Name)

		err = AddIPAddress(vethInterface1Name, ipAddress1, netMask, "")
		test.FailOnError(t, err)
		test.AssertInterfaceHasIPAndPrefix(t, vethInterface1Name, ipAddress1, netMask)
		test.AssertIPAddressRespondsToPings(t, ipAddress1)

		err = AddIPAddress(vethInterface2Name, ipAddress2, netMask, "")
		test.FailOnError(t, err)
		test.AssertInterfaceHasIPAndPrefix(t, vethInterface2Name, ipAddress2, netMask)
		test.AssertIPAddressRespondsToPings(t, ipAddress2)
	})
}
