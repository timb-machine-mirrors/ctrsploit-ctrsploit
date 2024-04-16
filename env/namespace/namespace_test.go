package namespace

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/colorful"
	"github.com/ctrsploit/sploit-spec/pkg/env/container"
	"testing"
)

func Test_level2result(t *testing.T) {
	colorful.O = colorful.Colorful{}
	r := level2result("user", container.NamespaceLevelHost)
	fmt.Println(r.Colorful())
}
