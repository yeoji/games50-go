//go:generate file2byteslice -package=fonts -input=./fonts/font.ttf -output=./fonts/font.generated.go -var=Font_ttf
//go:generate file2byteslice -package=graphics -input=./graphics/background.png -output=./graphics/background.generated.go -var=Background_png
//go:generate file2byteslice -package=graphics -input=./graphics/arrows.png -output=./graphics/arrows.generated.go -var=Arrows_png
//go:generate file2byteslice -package=graphics -input=./graphics/breakout.png -output=./graphics/breakout.generated.go -var=Breakout_png
//go:generate gofmt -s -w .

package assets
