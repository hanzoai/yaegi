package cheese

import (
	"fmt"

	"guthib.com/hanzoai/fromage"
)

func Hello() string {
	return fmt.Sprintf("cheese %s", fromage.Hello())
}
