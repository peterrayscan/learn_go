package test

import (
	"context"
	"fmt"
	"learn/tools/testx"
	"testing"
)

type TestStructOne struct{}
type TestStructTwo struct{}

func Test_ContextWithValue(t *testing.T) {
	// WithValue(ctx, key, value)
	// use struct as key
	testx.RunFunc(func() {
		ctx := context.Background()

		// set and get use string
		ctx = context.WithValue(ctx, "123", 123)
		fmt.Println("ctx, key: 123, value:", ctx.Value("123"))

		// set and get use another string
		ctx = context.WithValue(ctx, "456", 456)
		fmt.Println("ctx, key: 456, value:", ctx.Value("456"))

		// set and get use struct
		ctx = context.WithValue(ctx, TestStructOne{}, "it's struct one")
		fmt.Println("ctx, key: TestStructOne{}, value:", ctx.Value(TestStructOne{}))
		fmt.Println("ctx, key: TestStructTwo{}, value:", ctx.Value(TestStructTwo{}))
		fmt.Println("ctx, key: struct{}{}, value:     ", ctx.Value(struct{}{}))

		ctx = context.WithValue(ctx, struct{}{}, "empty struct")
		fmt.Println("ctx, key: TestStructOne{}, value:", ctx.Value(TestStructOne{}))
		fmt.Println("ctx, key: TestStructTwo{}, value:", ctx.Value(TestStructTwo{}))
		fmt.Println("ctx, key: struct{}{}, value:     ", ctx.Value(struct{}{}))
	})
}
