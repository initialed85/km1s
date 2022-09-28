package network

import (
	"fmt"
	"github.com/initialed85/km1s/pkg/common"
)

func AddIPAddress(
	interfaceName string,
	ipAddress string,
	prefix int,
	namespaceName string,
) error {
	var output string
	var err error

	namespacePrefix := getNamespacePrefix(namespaceName)

	output, err = common.RunCommand(
		fmt.Sprintf(
			"%vip addr add %v/%v dev %v",

			namespacePrefix,
			ipAddress,
			prefix,
			interfaceName,
		),
	)
	if err != nil {
		return fmt.Errorf("%v; %v", err, output)
	}

	return nil
}

func DelIPAddress(
	interfaceName string,
	ipAddress string,
	prefix int,
	namespaceName string,
) {
	namespacePrefix := getNamespacePrefix(namespaceName)

	_ = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"%vip addr del %v/%v dev %v",
				namespacePrefix,
				ipAddress,
				prefix,
				interfaceName,
			),
		),
	)
}
