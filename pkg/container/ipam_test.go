package container

import (
	"github.com/initialed85/km1s/pkg/common"
	"github.com/initialed85/km1s/pkg/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPAM(t *testing.T) {
	var i *IPAM
	var err error
	containerID := common.GetRandomUUID()

	t.Run("HappyPath", func(t *testing.T) {
		i, err = NewIPAM("172.16.137.1/29")
		test.FailOnError(t, err)

		ipAddresses := make([]string, 0)
		for j := 0; j < 4; j++ {
			ipAddress, err := i.Allocate(containerID)
			test.FailOnError(t, err)

			ipAddresses = append(ipAddresses, ipAddress)
		}
		assert.Equal(
			t,
			[]string{"172.16.137.0", "172.16.137.1", "172.16.137.2", "172.16.137.3"},
			ipAddresses,
		)

		ipAddresses = make([]string, 0)
		for j := 0; j < 4; j++ {
			ipAddress, err := i.Allocate(containerID)
			test.FailOnError(t, err)

			ipAddresses = append(ipAddresses, ipAddress)
		}
		assert.Equal(
			t,
			[]string{"172.16.137.4", "172.16.137.5", "172.16.137.6", "172.16.137.7"},
			ipAddresses,
		)

		_, err = i.Allocate(containerID)
		assert.NotNil(t, err)
	})
}
