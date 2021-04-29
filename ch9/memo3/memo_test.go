package memo_test

import (
	memo "GolangTurtorial/ch9/memo3"
	"GolangTurtorial/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

// Note: not concurrency-safe! Test fails
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
