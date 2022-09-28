package network

import (
	"fmt"
	"strings"
)

func getNamespacePrefix(namespaceName string) string {
	if strings.TrimSpace(namespaceName) == "" {
		return ""
	}

	return fmt.Sprintf("ip netns exec %v ", namespaceName)
}

func getNamespaceSuffix(namespaceName string) string {
	if strings.TrimSpace(namespaceName) == "" {
		return ""
	}

	return fmt.Sprintf(" netns %v", namespaceName)
}
