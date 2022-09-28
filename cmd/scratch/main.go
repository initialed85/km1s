package main

import (
	"github.com/initialed85/km1s/pkg/common"
	"github.com/initialed85/km1s/pkg/container"
	"github.com/initialed85/km1s/pkg/network"
	"log"
)

func main() {
	var err error
	containerName := "km1s-test"
	bridgeInterfaceName := "br-km1s-test"

	defer func() {
		network.DelIPAddress(bridgeInterfaceName, "10.137.137.1", 24, "")
		network.DelInterface(bridgeInterfaceName, "")
	}()

	err = network.AddBridge(bridgeInterfaceName, "")
	if err != nil {
		log.Fatal(err)
	}

	err = network.AddIPAddress(bridgeInterfaceName, "10.137.137.1", 24, "")
	if err != nil {
		log.Fatal(err)
	}

	c := container.NewContainer(
		containerName,
		bridgeInterfaceName,
		"10.137.137.102",
		24,
		"10.137.137.1",
	)

	err = c.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	common.WaitForSIGINT()
}
