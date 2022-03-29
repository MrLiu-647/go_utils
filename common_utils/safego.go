package common_utils

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/MrLiu-647/go_utils/logs"
)

func SafeGo(ctx context.Context, f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				content := fmt.Sprintf("Safe Go Capture Panic In Go Groutine\n%s", string(debug.Stack()))
				logs.CtxError(ctx, content)
			}
		}()

		f()
	}()
}
