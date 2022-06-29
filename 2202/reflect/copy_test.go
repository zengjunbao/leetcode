package main

import (
	"fmt"
	"reflect"
	"testing"
)

type AccountOM struct {
	Account string    `json:"account"` // 绑定的账号
	Topic   []TopicV1 `json:"topic"`   // 话题
}

type TopicV1 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PostsDTO 帖子 posts
type AccountIM struct {
	Account string    `json:"account"` // 绑定的账号
	Topic   []TopicV2 `json:"topic"`   // 话题
}

type TopicV2 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestStructCopy(t *testing.T) {
	im := []AccountIM{{
		Account: "123123",
		Topic: []TopicV2{
			{
				ID:   1,
				Name: "t1",
			}, {
				ID:   2,
				Name: "t2",
			},
		},
	},
	}
	om := make([]AccountOM,0)

	// err := StructCopy(&om,&im)
	// assert.NoError(t, err)
	//
	// fmt.Println("om:",om)
	srcV := reflect.ValueOf(im)
	dstV := reflect.ValueOf(om)
	srcT := reflect.TypeOf(im)
	dstT := reflect.TypeOf(om)
	fmt.Println(srcT, srcV)
	fmt.Println(dstT, dstV)

	if srcT.Kind() == reflect.Slice{
		for _, v := range im {
			v1 := reflect.ValueOf(v)
			t1 := reflect.TypeOf(v)
			fmt.Println(v1, t1)
		}
		// fmt.Println(row.Kind())
	}

}
