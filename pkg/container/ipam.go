package container

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/korylprince/ipnetgen"
	"log"
	"sync"
)

type IPAM struct {
	mu                     sync.Mutex
	ipNetGenerator         *ipnetgen.IPNetGenerator
	containerIDByIPAddress map[string]uuid.UUID
	ipAddressByContainerID map[uuid.UUID]string
}

func NewIPAM(cidr string) (*IPAM, error) {
	ipNetGenerator, err := ipnetgen.New(cidr)
	if err != nil {
		log.Fatal(err)
	}

	i := IPAM{
		ipNetGenerator:         ipNetGenerator,
		containerIDByIPAddress: make(map[string]uuid.UUID),
		ipAddressByContainerID: make(map[uuid.UUID]string),
	}

	return &i, nil
}

func (i *IPAM) Allocate(containerID uuid.UUID) (string, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	var ipAddress string

	for j := 0; j < 65536; j++ {
		ip := i.ipNetGenerator.Next()

		if ip == nil {
			continue
		}

		ipAddress = ip.String()

		_, ok := i.containerIDByIPAddress[ipAddress]
		if ok {
			continue
		}

		i.containerIDByIPAddress[ipAddress] = containerID
		i.ipAddressByContainerID[containerID] = ipAddress

		return ipAddress, nil
	}

	return "", fmt.Errorf("failed to allocate IP address; pool %#+v exhausted", i.ipNetGenerator.String())
}

func (i *IPAM) Release(containerID uuid.UUID) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	ipAddress, ok := i.ipAddressByContainerID[containerID]
	if !ok {
		return fmt.Errorf("containerID=%#+v is unknown", containerID)
	}

	delete(i.ipAddressByContainerID, containerID)
	delete(i.containerIDByIPAddress, ipAddress)

	return nil
}
