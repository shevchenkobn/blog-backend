package util

import (
	"github.com/shevchenkobn/blog-backend/internal/types"
)

func RemoveCallbackOnOrder(
	s []types.ExitHandlerCallback,
	callback types.ExitHandlerCallback,
) []types.ExitHandlerCallback {
	for i, c := range s {
		if &c == &callback {
			return RemoveINoOrder(s, i)
		}
	}
	return s
}

func RemoveINoOrder(s []types.ExitHandlerCallback, i int) []types.ExitHandlerCallback {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
