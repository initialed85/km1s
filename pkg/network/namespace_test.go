package network

import (
	"github.com/initialed85/km1s/pkg/test"
	"testing"
)

func TestNamespace(t *testing.T) {
	var err error

	bridgeInterfaceName := "bridge-1"
	namespaceName := "netns-1"

	t.Run("HappyPath", func(t *testing.T) {
		defer func() {
			DelInterface(bridgeInterfaceName, namespaceName)
			DelNamespace(namespaceName)
		}()

		err = AddNamespace(namespaceName)
		test.FailOnError(t, err)
		test.AssertNamespaceExists(t, namespaceName)

		err = AddBridge(bridgeInterfaceName, namespaceName)
		test.FailOnError(t, err)
		test.AssertInterfaceExistsInNamespace(t, bridgeInterfaceName, namespaceName)
	})
}
