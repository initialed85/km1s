package network

import (
	"fmt"
	"github.com/initialed85/km1s/pkg/common"
)

func AddNamespace(namespaceName string) error {
	var err error

	err = common.HandleRunCommandError(
		common.RunCommand(
			fmt.Sprintf(
				"ip netns add %v",
				namespaceName,
			),
		),
	)
	if err != nil {
		return err
	}

	return nil
}

func DelNamespace(namespaceName string) {
	_, _ = common.RunCommand(
		fmt.Sprintf(
			"ip netns del %v",
			namespaceName,
		),
	)
}
