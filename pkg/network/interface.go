package network

import (
	"fmt"
	"github.com/initialed85/km1s/pkg/common"
)

func DelInterface(
	interfaceName string,
	namespaceName string,
) {
	namespacePrefix := getNamespacePrefix(namespaceName)

	_ = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"%vip link set %v down",
				namespacePrefix,
				interfaceName,
			),
		),
	)

	_ = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"%vip link del %v",
				namespacePrefix,
				interfaceName,
			),
		),
	)
}
