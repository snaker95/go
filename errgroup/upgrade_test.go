package errgroup

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
@Time : 2021/11/20 下午10:05
@Author : snaker95
@File : upgrade_test.go
@Software: GoLand
*/

func TestGroup_Run(t *testing.T) {
	ctx := context.Background()
	g, egCtx := WithContext(ctx)
	f := func(ctx context.Context) error {
		time.Sleep(4 * time.Second)
		fmt.Println("f func")
		return nil
	}
	fe := func(ctx context.Context) error {
		fmt.Println("f func err")
		return fmt.Errorf("fe err")
	}
	type args struct {
		f []func() error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty",
			args: args{

			},
			wantErr: false,
		},
		{
			name: "one",
			args: args{
				f: []func() error{
					func() error {
						fmt.Println(1)
						return nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "two",
			args: args{
				f: []func() error{
					func() error {
						fmt.Println(1)
						return nil
					},
					func() error {
						f(egCtx)
						return nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "tree-err",
			args: args{
				f: []func() error{
					func() error {
						time.Sleep(5 * time.Second)
						fmt.Println(1)
						return nil
					},
					func() error {
						f(egCtx)
						return nil
					},
					func() error {
						return fe(egCtx)
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := g.Run(tt.args.f...); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
