package testx

type RunHandler func()

func Run(handler RunHandler) {
	handler()
}
