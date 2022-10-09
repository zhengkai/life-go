package life

import (
	"context"
	"errors"
)

// Stop 用于检测是否退出
var Stop bool

// CTX 可以用来替代 `context.Background`
var CTX context.Context

// BeforeCancel 在宣布退出前可以做些额外操作的 hook
var BeforeCancel func()

// ErrStop 用于判断 err 是退出还是别的原因
var ErrStop = errors.New(`life stop`)

var cancel context.CancelFunc
