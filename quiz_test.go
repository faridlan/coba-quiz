package test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

type Word struct {
	Id        int    `json:"id,omitempty"`
	Word      string `json:"word,omitempty"`
	Mean      string `json:"mean,omitempty"`
	Sentence  string `json:"sentence,omitempty"`
	ImageUrl  string `json:"image_url,omitempty"`
	AnswerKey string `json:"answer_key,omitempty"`
}

type Answer struct {
	AnswerId string `json:"answer_id,omitempty"`
	Word     string `json:"word,omitempty"`
}

type Response struct {
	Id        int      `json:"id,omitempty"`
	Word      string   `json:"word,omitempty"`
	Mean      string   `json:"mean,omitempty"`
	Answer    []Answer `json:"answer,omitempty"`
	Sentence  string   `json:"sentence,omitempty"`
	ImageUrl  string   `json:"image_url,omitempty"`
	AnswerKey string   `json:"answer_key,omitempty"`
}

var word = []string{
	"Pintu",
	"Langit",
	"Kursi",
	"Jendela",
	"Kelapa",
	"Botol",
}

func TestWord(t *testing.T) {

	banana := Word{
		Id:        1,
		Word:      "Banana",
		Sentence:  "I Have Banana",
		ImageUrl:  "http://google.com/image/banana.png",
		AnswerKey: "a",
	}

	fmt.Println(banana.AnswerKey + banana.Word)

}

func toChar(i int) rune {
	return rune('a' - 1 + i)
}

func TestRandom(t *testing.T) {

	answer := []Answer{}

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 4; i++ {
		foo := Answer{
			AnswerId: string(toChar(i)),
			Word:     word[rand.Intn(len(word))],
		}

		answer = append(answer, foo)
	}

	fmt.Println(answer)

}

func Quiz() Response {

	banana := Word{
		Id:        1,
		Word:      "Banana",
		Mean:      "Pisang",
		Sentence:  "I Have Banana",
		ImageUrl:  "http://google.com/image/banana.png",
		AnswerKey: "c",
	}
	answer := []Answer{}

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 4; i++ {
		foo := Answer{
			AnswerId: string(toChar(i)),
			Word:     word[rand.Intn(len(word))],
		}

		if foo.AnswerId == banana.AnswerKey {
			foo.Word = banana.Word
		}

		answer = append(answer, foo)
	}

	response := Response{
		Id:        banana.Id,
		Word:      banana.Word,
		Mean:      banana.Mean,
		Answer:    answer,
		Sentence:  banana.Sentence,
		ImageUrl:  banana.ImageUrl,
		AnswerKey: banana.AnswerKey,
	}

	return response

}

func TestQuiz(t *testing.T) {
	file, err := os.Create("../json/response.json")
	if err != nil {
		panic(err)
	}
	responseQuiz := Quiz()

	encode := json.NewEncoder(file)
	err = encode.Encode(responseQuiz)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes Create A Response")
}
