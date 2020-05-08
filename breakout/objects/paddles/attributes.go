package paddles

type PaddleColour int

const (
	Blue PaddleColour = iota
	Green
	Red
	Purple
)

func (c PaddleColour) SpriteGroup() string {
	return []string{"paddles-blue", "paddles-green", "paddles-red", "paddles-purple"}[c]
}

type PaddleSize int

const (
	Smallest PaddleSize = iota
	Small
	Large
	Largest
)

func (s PaddleSize) String() string {
	return []string{"smallest", "small", "large", "largest"}[s]
}
