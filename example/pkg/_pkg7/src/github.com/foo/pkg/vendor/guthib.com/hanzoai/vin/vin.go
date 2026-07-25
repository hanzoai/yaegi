package vin

import (
	"fmt"

	"guthib.com/hanzoai/cheese"
)

func Hello() string {
	return fmt.Sprintf("vin %s", cheese.Hello())
}
