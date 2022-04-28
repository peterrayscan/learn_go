package testx

import "learn/tools/console"

func RunFunc(f func()) {
	console.Before()
	f()
	console.After()
}
