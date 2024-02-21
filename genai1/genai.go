package main

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
)

func main() {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyD1M0kFB0TrgxTfC4x1G_fyLyBDFyGnSH8"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ModelName := "gemini-pro"
	l := client.ListModels(ctx)
	fmt.Println(l.Next())
	model := client.GenerativeModel(ModelName)
	//var part = []string{"input: Beauty\nShow me some perfumes\nShow me some fragrance\nCould you suggest some beauty products?\n#Beauty #perfume #LipEssentials",
	//	"output: Tease Collector's Edition Eau De Parfum\nTease Mini Fragrance Duo\nDeluxe Mini Fragrance Trio\nTease Eau de Parfum\nTravel Fine Fragrance Mist\nTease Crème Cloud Eau de Parfum",
	//	"input: Show me some lip essentials for Valentine's Day\nLip bomb\n#beauty #LipEssentials",
	//	"output: Flavored Lip Gloss\nColor Shine Lip Gloss\nShine Plumper Extreme Lip Plumper\nLip Oil\nLip Care Kit",
	//	"input: Nightwear\nShow me some trending regular Sleepwear\nSleepwear\nCould you show me some trending regular Nightwear?\nPlease show me some trending regular Nightwear\nNightDress\n#sleepshirt #camisets #pajamas",
	//	"output: Short Cozy Robe\nSatin Jacquard Long Pajama Set\nModal Short Pajama Set\nThermal Sleepshirt\nSatin Long Pajama Set\nShimmer Knit Long Pajama Set",
	//	"input: Show me some trending regular Sleep wear",
	//	"output: Short Cozy Robe\nSatin Jacquard Long Pajama Set\nModal Short Pajama Set\nThermal Sleepshirt\nSatin Long Pajama Set\nShimmer Knit Long Pajama Set",
	//	"input: Sleep wear",
	//	"output: Short Cozy Robe\nSatin Jacquard Long Pajama Set\nModal Short Pajama Set\nThermal Sleepshirt\nSatin Long Pajama Set\nShimmer Knit Long Pajama Set",
	//	"input: #perfume",
	//	"output: ",
	//}
	//var par []genai1.Part

	//Parts := make([]genai1.Part, 0)
	//Parts = append(Parts, part...)
	////var parts genai1.Part
	//model.GenerateContent(ctx, genai1.Content{Parts: make([]genai1.Part, 0)})

	prompt := []genai.Part{
		genai.Text("Show me some trending regular Sleepwear"),
	}

	resp, _ := model.GenerateContent(ctx, prompt...)

	for _, value := range resp.Candidates {
		fmt.Println(*value.Content)
	}

}
