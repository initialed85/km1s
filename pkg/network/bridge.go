package network

import (
	"fmt"
	"github.com/initialed85/km1s/pkg/common"
)

func AddBridge(
	bridgeInterfaceName string,
	namespaceName string,
) error {
	var err error

	namespacePrefix := getNamespacePrefix(namespaceName)

	err = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"%vip link add name %v type bridge",
				namespacePrefix,
				bridgeInterfaceName,
			),
		),
	)
	if err != nil {
		return err
	}

	err = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"%vip link set %v up",
				namespacePrefix,
				bridgeInterfaceName,
			),
		))
	if err != nil {
		return err
	}

	return nil
}

func AddBridgePort(
	interfaceName string,
	bridgeInterfaceName string,
	namespaceName string,
) error {
	var err error

	namespacePrefix := getNamespacePrefix(namespaceName)

	err = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"%vip link set dev %v master %v",
				namespacePrefix,
				interfaceName,
				bridgeInterfaceName,
			),
		),
	)
	if err != nil {
		return err
	}

	return nil
}

func DelBridgePort(
	interfaceName string,
	namespaceName string,
) {
	namespacePrefix := getNamespacePrefix(namespaceName)

	_ = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"%vip link set dev %v nomaster",
				namespacePrefix,
				interfaceName,
			),
		),
	)
}
