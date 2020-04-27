//go:generate file2byteslice -package=sounds -input=./sounds/paddle_hit.wav -output=./sounds/paddle_hit.generated.go -var=PaddleHit_wav
//go:generate file2byteslice -package=sounds -input=./sounds/wall_hit.wav -output=./sounds/wall_hit.generated.go -var=WallHit_wav
//go:generate file2byteslice -package=sounds -input=./sounds/score.wav -output=./sounds/score.generated.go -var=Score_wav
//go:generate file2byteslice -package=fonts -input=./fonts/font.ttf -output=./fonts/font.generated.go -var=Font_ttf
//go:generate gofmt -s -w .

package assets
