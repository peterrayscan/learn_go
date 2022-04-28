package test

import "learn/tools/console"

func runFunc(f func()) {
	console.Before()
	f()
	console.After()
}
