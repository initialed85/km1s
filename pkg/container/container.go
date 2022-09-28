package container

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/initialed85/km1s/pkg/common"
	"github.com/initialed85/km1s/pkg/network"
	"strings"
	"sync"
)

type Container struct {
	containerName       string
	bridgeInterfaceName string
	ipAddress           string
	prefix              int
	nextHopIPAddress    string
	containerID         uuid.UUID
	vethInterface1Name  string
	vethInterface2Name  string
	namespaceName       string
	mu                  sync.Mutex
	processByID         map[uuid.UUID]*Process
}

func NewContainer(
	containerName string,
	bridgeInterfaceName string,
	ipAddress string,
	prefix int,
	nextHopIPAddress string,
) *Container {
	c := Container{
		containerName:       containerName,
		bridgeInterfaceName: bridgeInterfaceName,
		ipAddress:           ipAddress,
		prefix:              prefix,
		nextHopIPAddress:    nextHopIPAddress,
		containerID:         common.GetRandomUUID(),
		processByID:         make(map[uuid.UUID]*Process),
	}

	containerID := strings.Replace(c.containerID.String(), "-", "", -1)
	partContainerID := containerID[len(containerID)-11:]

	c.vethInterface1Name = fmt.Sprintf("vb-%v", partContainerID)
	c.vethInterface2Name = fmt.Sprintf("vc-%v", partContainerID)
	c.namespaceName = fmt.Sprintf("ns-%v", partContainerID)

	return &c
}

func (c *Container) Open() error {
	err := network.AddNamespace(c.namespaceName)
	if err != nil {
		return err
	}

	err = network.AddVETH(
		c.vethInterface1Name,
		c.vethInterface2Name,
		c.namespaceName,
	)
	if err != nil {
		return err
	}

	err = network.AddIPAddress(
		c.vethInterface2Name,
		c.ipAddress,
		c.prefix,
		c.namespaceName,
	)
	if err != nil {
		return err
	}

	err = network.AddBridgePort(
		c.vethInterface1Name,
		c.bridgeInterfaceName,
		"",
	)
	if err != nil {
		return err
	}

	err = network.AddDefaultRoute(
		c.nextHopIPAddress,
		c.namespaceName,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Container) Run(command string, logBufferSize int) (*Process, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	p := NewProcess(
		fmt.Sprintf("ip netns exec %v %v", c.namespaceName, command),
		logBufferSize,
		make(map[string]string),
	)

	c.processByID[p.ProcessID()] = p

	err := p.Open()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (c *Container) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	processes := make([]*Process, 0)
	for _, process := range c.processByID {
		processes = append(processes, process)
	}

	for _, process := range processes {
		process.Close()
		delete(c.processByID, process.ProcessID())
	}

	network.DelBridgePort(
		c.vethInterface1Name,
		"",
	)

	network.DelInterface(c.vethInterface2Name, c.namespaceName)
	network.DelInterface(c.vethInterface1Name, "")

	network.DelNamespace(c.namespaceName)
}
