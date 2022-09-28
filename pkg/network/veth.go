package network

import (
	"fmt"
	"github.com/initialed85/km1s/pkg/common"
)

func AddVETH(
	interface1Name string,
	interface2Name string,
	namespaceName string,
) error {
	var err error

	namespaceSuffix := getNamespaceSuffix(namespaceName)

	err = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"%vip link add name %v type veth peer name %v%v",
				"", // intentional
				interface1Name,
				interface2Name,
				namespaceSuffix,
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
				"", // intentional
				interface1Name,
			),
		),
	)
	if err != nil {
		return err
	}

	namespacePrefix := getNamespacePrefix(namespaceName)

	err = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"%vip link set %v up",
				namespacePrefix,
				interface2Name,
			),
		),
	)
	if err != nil {
		return err
	}

	return nil
}
