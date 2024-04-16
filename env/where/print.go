package where

import (
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

type Result struct {
	Name  result.Title `json:"name"`
	Rules []item.Bool  `json:"rules"`
	In    item.Bool    `json:"in"`
}

func Human() {

}

func Print() (err error) {
	return
}
