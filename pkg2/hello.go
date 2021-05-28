package pkg2

import (
	"github.com/hugelgupf/modplay/v2/pkg1"
)

func Hello() string {
	return "pkg2 + " + pkg1.Hello()
}
