package toasts

import (
	"fmt"
	"strings"
)

func GetAWSError(authErr error) error {
	parts := strings.Split(authErr.Error(), ",")

	lastPart := parts[len(parts)-1]
	finalParts := strings.Split(lastPart, ":")

	return fmt.Errorf("%s", strings.TrimSpace(finalParts[len(finalParts)-1]))
}
