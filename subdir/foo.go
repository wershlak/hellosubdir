package subdir

import (
	"fmt"
)

func GetMessage(environment string) (string, error) {
	return fmt.Sprintf("Hello from the %s environment", environment), nil
}
