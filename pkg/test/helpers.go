package test

import (
	"fmt"
	"github.com/initialed85/km1s/pkg/common"
	"github.com/stretchr/testify/assert"
	"log"
	"net"
	"strings"
	"testing"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix)
}

func FailOnError(t *testing.T, err error) {
	if err == nil {
		return
	}

	t.Fatal(err)
}

func RunCommandOrFailOnError(
	t *testing.T,
	command string,
) string {
	output, err := common.RunCommand(command)
	if err != nil {
		FailOnError(t, err)
	}

	return output
}

func AssertInterfaceExists(
	t *testing.T,
	interfaceName string,
) {
	output := RunCommandOrFailOnError(t, fmt.Sprintf("ip addr show %v", interfaceName))
	log.Printf("\n\n%v\n\n", strings.TrimSpace(output))
	assert.Contains(t, output, interfaceName)
}

func AssertInterfaceHasBridge(
	t *testing.T,
	interfaceName string,
	bridgeInterfaceName string,
) {
	output := RunCommandOrFailOnError(t, fmt.Sprintf("ip addr show %v", interfaceName))
	log.Printf("\n\n%v\n\n", strings.TrimSpace(output))
	assert.Contains(t, output, interfaceName)
	assert.Contains(t, output, fmt.Sprintf("master %v", bridgeInterfaceName))
}

func AssertInterfaceExistsInNamespace(
	t *testing.T,
	interfaceName string,
	namespaceName string,
) {
	output := RunCommandOrFailOnError(
		t,
		fmt.Sprintf("ip netns exec %v ip addr show %v", namespaceName, interfaceName),
	)
	log.Printf("\n\n%v\n\n", strings.TrimSpace(output))
	assert.Contains(t, output, interfaceName)
}

func AssertInterfaceHasIPAndPrefix(
	t *testing.T,
	interfaceName string,
	ipAddress string,
	prefix int,
) {
	output := RunCommandOrFailOnError(t, fmt.Sprintf("ip addr show %v", interfaceName))
	log.Printf("\n\n%v\n\n", strings.TrimSpace(output))
	assert.Contains(t, output, fmt.Sprintf("%v/%v", ipAddress, prefix))
}

func AssertIPAddressRespondsToPings(
	t *testing.T,
	ipAddress string,
) {
	output := RunCommandOrFailOnError(t, fmt.Sprintf("ping -s 16 -i 0.1 -W 100 -c 2 %v", ipAddress))
	log.Printf("\n\n%v\n\n", strings.TrimSpace(output))
	assert.Contains(t, output, "0% packet loss")
}

func AssertNamespaceExists(
	t *testing.T,
	namespaceName string,
) {
	output := RunCommandOrFailOnError(t, fmt.Sprintf("ip netns show %v", namespaceName))
	log.Printf("\n\n%v\n\n", strings.TrimSpace(output))
	assert.Contains(t, output, namespaceName)
}

func AssertRouteExists(
	t *testing.T,
	ipAddress string,
	mask int,
	nextHopIPAddress string,
) {
	output := RunCommandOrFailOnError(t, "ip route show")
	log.Printf("\n\n%v\n\n", strings.TrimSpace(output))
	assert.Contains(t, output, fmt.Sprintf("%v/%v", ipAddress, mask))
	assert.Contains(t, output, nextHopIPAddress)
}

func AssertDefaultRouteExists(
	t *testing.T,
	nextHopIPAddress string,
) {
	output := RunCommandOrFailOnError(t, "ip route show")
	log.Printf("\n\n%v\n\n", strings.TrimSpace(output))
	assert.Contains(t, output, "default")
	assert.Contains(t, output, nextHopIPAddress)
}

func SendUDP(udpAddress string, data string) error {
	addr, err := net.ResolveUDPAddr("udp", udpAddress)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}
