package network

import (
	"fmt"
	"github.com/initialed85/km1s/pkg/common"
)

func AddRoute(
	ipAddress string,
	mask int,
	nextHopIPAddress string,
	namespaceName string,
) error {
	var err error

	namespacePrefix := getNamespacePrefix(namespaceName)

	err = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"%vip route add %v/%v via %v",
				namespacePrefix,
				ipAddress,
				mask,
				nextHopIPAddress,
			),
		),
	)
	if err != nil {
		return err
	}

	return nil
}

func AddDefaultRoute(
	nextHopIPAddress string,
	namespaceName string,
) error {
	var err error

	namespacePrefix := getNamespacePrefix(namespaceName)

	err = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"%vip route add default via %v",
				namespacePrefix,
				nextHopIPAddress,
			),
		),
	)
	if err != nil {
		return err
	}

	return nil
}

func DelRoute(
	ipAddress string,
	mask int,
	nextHopIPAddress string,
	namespaceName string,
) {
	namespacePrefix := getNamespacePrefix(namespaceName)

	_ = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"%vip route del %v/%v via %v",
				namespacePrefix,
				ipAddress,
				mask,
				nextHopIPAddress,
			),
		),
	)
}

func DelDefaultRoute(
	namespaceName string,
) {
	namespacePrefix := getNamespacePrefix(namespaceName)

	_ = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"%vip route del default",
				namespacePrefix,
			),
		),
	)
}
