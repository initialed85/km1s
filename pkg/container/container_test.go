package container

import (
	"github.com/initialed85/km1s/pkg/network"
	"github.com/initialed85/km1s/pkg/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestContainer(t *testing.T) {
	var c *Container
	var p *Process
	var err error
	var logs string
	containerName := "km1s-test"
	bridgeInterfaceName := "br-km1s-test"

	defer func() {
		network.DelInterface(bridgeInterfaceName, "")
	}()

	err = network.AddBridge(bridgeInterfaceName, "")
	test.FailOnError(t, err)

	err = network.AddIPAddress(bridgeInterfaceName, "10.137.137.101", 24, "")
	test.FailOnError(t, err)

	t.Run("Integration", func(t *testing.T) {
		defer func() {
			if p != nil {
				p.Close()
			}

			if c != nil {
				c.Close()
			}
		}()

		c = NewContainer(
			containerName,
			bridgeInterfaceName,
			"10.137.137.102",
			24,
			"10.137.137.101",
		)

		err = c.Open()
		test.FailOnError(t, err)

		test.AssertIPAddressRespondsToPings(t, "10.137.137.102")

		p, err = c.Run("nc -lu 13337", 128)
		test.FailOnError(t, err)

		err = test.SendUDP("10.137.137.102:13337", "Hello, world.\n")
		test.FailOnError(t, err)

		time.Sleep(time.Millisecond * 100)

		logs = p.Logs()
		assert.Equal(t, "Hello, world.\n", logs)
	})
}
