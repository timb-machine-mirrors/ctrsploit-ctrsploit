package env

import (
	"encoding/json"
	"github.com/ctrsploit/ctrsploit/env/auto"
	"github.com/ctrsploit/sploit-spec/pkg/upload"
)

var (
	Upload = upload.GenerateUploadCommand(func() (content []byte, err error) {
		env, err := auto.Auto()
		if err != nil {
			return
		}
		content, err = json.Marshal(env)
		if err != nil {
			return
		}
		return
	})
)
