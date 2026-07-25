package fromage

import (
	"fmt"

	"guthib.com/hanzoai/cheese"
	"guthib.com/hanzoai/fromage/couteau"
)

func Hello() string {
	return fmt.Sprintf("Fromage %s %s", cheese.Hello(), couteau.Hello())
}
