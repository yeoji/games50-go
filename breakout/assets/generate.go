//go:generate file2byteslice -package=fonts -input=./fonts/font.ttf -output=./fonts/font.ttf.generated.go -var=Font_ttf
//go:generate file2byteslice -package=sounds -input=./sounds/brick_hit_1.wav -output=./sounds/brick_hit_1.wav.generated.go -var=Brick_hit_1_wav
//go:generate file2byteslice -package=sounds -input=./sounds/brick_hit_2.wav -output=./sounds/brick_hit_2.wav.generated.go -var=Brick_hit_2_wav
//go:generate file2byteslice -package=sounds -input=./sounds/confirm.wav -output=./sounds/confirm.wav.generated.go -var=Confirm_wav
//go:generate file2byteslice -package=sounds -input=./sounds/high_score.wav -output=./sounds/high_score.wav.generated.go -var=High_score_wav
//go:generate file2byteslice -package=sounds -input=./sounds/hurt.wav -output=./sounds/hurt.wav.generated.go -var=Hurt_wav
//go:generate file2byteslice -package=sounds -input=./sounds/music.wav -output=./sounds/music.wav.generated.go -var=Music_wav
//go:generate file2byteslice -package=sounds -input=./sounds/no_select.wav -output=./sounds/no_select.wav.generated.go -var=No_select_wav
//go:generate file2byteslice -package=sounds -input=./sounds/paddle_hit.wav -output=./sounds/paddle_hit.wav.generated.go -var=Paddle_hit_wav
//go:generate file2byteslice -package=sounds -input=./sounds/pause.wav -output=./sounds/pause.wav.generated.go -var=Pause_wav
//go:generate file2byteslice -package=sounds -input=./sounds/recover.wav -output=./sounds/recover.wav.generated.go -var=Recover_wav
//go:generate file2byteslice -package=sounds -input=./sounds/select.wav -output=./sounds/select.wav.generated.go -var=Select_wav
//go:generate file2byteslice -package=sounds -input=./sounds/victory.wav -output=./sounds/victory.wav.generated.go -var=Victory_wav
//go:generate file2byteslice -package=sounds -input=./sounds/wall_hit.wav -output=./sounds/wall_hit.wav.generated.go -var=Wall_hit_wav
//go:generate file2byteslice -package=graphics -input=./graphics/arrows.png -output=./graphics/arrows.png.generated.go -var=Arrows_png
//go:generate file2byteslice -package=graphics -input=./graphics/background.png -output=./graphics/background.png.generated.go -var=Background_png
//go:generate file2byteslice -package=graphics -input=./graphics/breakout.png -output=./graphics/breakout.png.generated.go -var=Breakout_png
//go:generate gofmt -s -w .

package assets
