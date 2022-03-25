package common_utils

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
)

func SafeGo(ctx context.Context, f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				content := fmt.Sprintf("Safe Go Capture Panic In Go Groutine\n%s", string(debug.Stack()))
				log.Fatal(ctx, content)
			}
		}()

		f()
	}()
}
