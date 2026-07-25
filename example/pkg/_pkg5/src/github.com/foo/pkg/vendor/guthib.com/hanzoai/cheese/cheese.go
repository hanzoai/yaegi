package cheese

import (
	"fmt"

	"guthib.com/hanzoai/cheese/vin"
)

func Hello() string {
	return fmt.Sprintf("Cheese %s", vin.Hello())
}
