//go:generate file2byteslice -package=sounds -input=./sounds/explosion.wav -output=./sounds/explosion.generated.go -var=Explosion_wav
//go:generate file2byteslice -package=sounds -input=./sounds/hurt.wav -output=./sounds/hurt.generated.go -var=Hurt_wav
//go:generate file2byteslice -package=sounds -input=./sounds/jump.wav -output=./sounds/jump.generated.go -var=Jump_wav
//go:generate file2byteslice -package=sounds -input=./sounds/pause.wav -output=./sounds/pause.generated.go -var=Pause_wav
//go:generate file2byteslice -package=sounds -input=./sounds/score.wav -output=./sounds/score.generated.go -var=Score_wav
//go:generate file2byteslice -package=sounds -input=./sounds/marios_way.mp3 -output=./sounds/marios_way.generated.go -var=MariosWay_mp3
//go:generate file2byteslice -package=fonts -input=./fonts/flappy.ttf -output=./fonts/flappy.generated.go -var=Flappy_ttf
//go:generate file2byteslice -package=art -input=./art/background.png -output=./art/background.generated.go -var=Background_png
//go:generate file2byteslice -package=art -input=./art/bird.png -output=./art/bird.generated.go -var=Bird_png
//go:generate file2byteslice -package=art -input=./art/bronze_medal.png -output=./art/bronze_medal.generated.go -var=BronzeMedal_png
//go:generate file2byteslice -package=art -input=./art/gold_medal.png -output=./art/gold_medal.generated.go -var=GoldMedal_png
//go:generate file2byteslice -package=art -input=./art/ground.png -output=./art/ground.generated.go -var=Ground_png
//go:generate file2byteslice -package=art -input=./art/pause.png -output=./art/pause.generated.go -var=Pause_png
//go:generate file2byteslice -package=art -input=./art/pipe.png -output=./art/pipe.generated.go -var=Pipe_png
//go:generate file2byteslice -package=art -input=./art/silver_medal.png -output=./art/silver_medal.generated.go -var=SilverMedal_png
//go:generate gofmt -s -w .

package assets
