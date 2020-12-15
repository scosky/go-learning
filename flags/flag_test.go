package flags

import (
	"flag"
	"fmt"
	"testing"
)

func TestFlag(t *testing.T) {
	var ip = flag.Int("flagname", 1234, "help message for flagname")
	fmt.Println(*ip)
}
