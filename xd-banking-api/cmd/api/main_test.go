package main

import "testing"

func Test_test(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		test()
		return
	})
}
