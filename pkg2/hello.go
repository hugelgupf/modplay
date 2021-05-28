package hello

import (
	"github.com/hugelgupf/modplay/pkg1"
)

func Hello() string {
	return "pkg2 + " + pkg1.Hello()
}
