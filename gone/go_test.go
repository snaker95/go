package gone

import (
	"fmt"
	"testing"
	"time"
)

/*
@Time : 2021/8/29 下午4:27
@Author : snaker95
@File : go_test.gone
@Software: GoLand
*/

func TestGo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Go(func() {
				fmt.Println("panic=aa")
				panic("ss")
			})
			Go(func() {
				fmt.Println("normal=bb")
			})
			time.Sleep(2 * time.Second)
		})
	}
}
