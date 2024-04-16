package auto

import (
	"fmt"
	"github.com/ctrsploit/ctrsploit/env/where"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {
	w, err := where.Where()
	assert.NoError(t, err)
	human := []interface{}{
		where.Human(w),
	}
	fmt.Println(printer.Printer.Print(human))
}
