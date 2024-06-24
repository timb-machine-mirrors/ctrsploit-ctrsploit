package mountinfo

import (
	"encoding/json"
	"fmt"
	"github.com/moby/sys/mountinfo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMountInfo(t *testing.T) {
	info, err := mountinfo.GetMounts(nil)
	assert.NoError(t, err)
	marshaled, err := json.Marshal(info)
	assert.NoError(t, err)
	fmt.Println(string(marshaled))
}
