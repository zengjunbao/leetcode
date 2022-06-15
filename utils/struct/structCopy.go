package _struct

import (
	"encoding/json"
)

type OrderA struct {
	Id      uint64 `json:"id"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Phone   string `json:"phone,omitempty"`
}

type OrderB struct {
	Id      uint64 `json:"id"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Code    string `json:"code"`
}

// CopyStruct oldObj 和 newObj 需要传入指针，通过json tag对应
func CopyStruct(oldObj interface{}, newObj interface{}) error {
	data,err := json.Marshal(oldObj)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data,newObj); err != nil {
		return err
	}
	return nil
}
