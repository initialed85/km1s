package network

import (
	"fmt"
	"github.com/initialed85/km1s/pkg/common"
)

func AddMACVLAN(
	interfaceName string,
	parentInterfaceName string,
	namespaceName string,
) error {
	var err error

	namespacePrefix := getNamespacePrefix(namespaceName)

	err = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"%vip link add %v link %v type macvlan mode bridge",
				namespacePrefix,
				interfaceName,
				parentInterfaceName,
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
				interfaceName,
			),
		),
	)
	if err != nil {
		return err
	}

	return nil
}
