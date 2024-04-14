package utils

func PrintWithEmoji[T any](msg T) {
	println("🚀 ", msg)
}

func PrintWithEmojiOnlyStringNum[T string | int](msg T) {
	println("🚀 ", msg)
}

type NumFloat interface {
	int | float64
}

func PrintWithEmojiOnlyNumFloat[T NumFloat](msg T) {
	println("🚀 ", msg)
}
