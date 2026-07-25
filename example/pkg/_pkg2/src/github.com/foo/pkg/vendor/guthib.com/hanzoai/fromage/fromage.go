package fromage

import (
	"fmt"

	"guthib.com/hanzoai/cheese"
)

func Hello() string {
	return fmt.Sprintf("Fromage %s", cheese.Hello())
}
