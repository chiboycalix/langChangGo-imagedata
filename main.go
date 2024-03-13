package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func API_KEY() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("API_KEY")
}

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(API_KEY()))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro-vision")
	imgData1, err := os.ReadFile("./images/nono2.png")
	if err != nil {
		fmt.Println("err", err)
		log.Fatal(err)
	}

	imgData2, err := os.ReadFile("./images/nono2.png")
	if err != nil {
		log.Fatal(err)
	}

	prompt := []genai.Part{
		genai.ImageData("png", imgData1),
		genai.ImageData("png", imgData2),
		genai.Text("Does these pictures belong to the same person?"),
	}

	resp, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		log.Fatal(err)
	}

	bs, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(bs))
}
