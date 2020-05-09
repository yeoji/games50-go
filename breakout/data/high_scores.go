package data

import (
	"bytes"
	"encoding/csv"
	"games50-go/breakout/objects"
	"log"
	"strconv"

	"github.com/shibukawa/configdir"
)

const HighScoreFile = "high_scores.csv"

var highScores []*objects.HighScore

func LoadHighScores() {
	configDirs := configdir.New("yeoji", "breakout")
	folder := configDirs.QueryFolderContainsFile(HighScoreFile)
	if folder != nil {
		data, _ := folder.ReadFile(HighScoreFile)
		readHighScores(data)
	}
}

func UpdateHighScores(name string, score int) {
	var newHighScores []*objects.HighScore
	newHighScore := &objects.HighScore{
		Name:  name,
		Score: score,
	}

	if len(highScores) == 0 {
		newHighScores = append(highScores, newHighScore)
	}

	for i := 0; i < len(highScores); i++ {
		if highScores[i].Score < score || (highScores[i].Score == score && len(highScores) < 10) {
			if i > 0 {
				newHighScores = highScores[:i-1]
			}
			newHighScores = append(newHighScores, newHighScore)
			newHighScores = append(newHighScores, highScores[i:len(highScores)-1]...)
			break
		}
	}

	if len(newHighScores) > 0 {
		highScores = newHighScores
		saveHighScores()
	}
}

func GetHighScores() []*objects.HighScore {
	return highScores
}

func saveHighScores() {
	configDirs := configdir.New("yeoji", "breakout")

	var data bytes.Buffer
	var records [][]string
	for _, highScore := range highScores {
		records = append(records, []string{highScore.Name, strconv.Itoa(highScore.Score)})
	}
	err := csv.NewWriter(&data).WriteAll(records)
	if err != nil {
		log.Fatalf("Error saving high score data: %v", err)
	}

	folders := configDirs.QueryFolders(configdir.Global)
	err = folders[0].WriteFile(HighScoreFile, data.Bytes())
	if err != nil {
		log.Fatalf("Error saving high score data: %v", err)
	}
}

func readHighScores(data []byte) {
	r := csv.NewReader(bytes.NewReader(data))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error reading high scores file: %v", err)
	}

	for _, line := range records {
		score, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalf("High scores file is corrupted! %v", err)
		}

		highScore := &objects.HighScore{
			Name:  line[0],
			Score: score,
		}

		highScores = append(highScores, highScore)
	}
}
