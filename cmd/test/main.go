package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/replicate/replicate-go"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)

	}
	ctx := context.Background()

	// You can also provide a token directly with
	// `replicate.NewClient(replicate.WithToken("r8_..."))`
	r8, err := replicate.NewClient(replicate.WithTokenFromEnv())
	if err != nil {
		fmt.Println(err)
	}

	// https://replicate.com/stability-ai/stable-diffusion
	// vers ion := "stability-ai/stable-diffusion:ac732df83cea7fff18b8472768c88ad041fa750ff7682a21affe81863cbe77e4"
	predictionVersion := "ac732df83cea7fff18b8472768c88ad041fa750ff7682a21affe81863cbe77e4"

	input := replicate.PredictionInput{
		"prompt": "a big green rat using a croossbow and walking down to the sewage",
	}

	webhook := replicate.Webhook{
		URL:    "https://webhook.site/bad3a1d6-b971-4cde-ba03-24ad4a1e4c3c",
		Events: []replicate.WebhookEventType{"start", "completed"},
	}

	// Run a model and wait for its output
	// output, err := r8.Run(ctx, version, input, &webhook)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("output: ", output)

	// Create a prediction
	prediction, err := r8.CreatePrediction(ctx, predictionVersion, input, &webhook, false)
	if err != nil {
		fmt.Println(err)
	}

	// // Wait for the prediction to finish
	err = r8.Wait(ctx, prediction)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("output: ", prediction.Output)

	// // The `Run` method is a convenience method that
	// // creates a prediction, waits for it to finish, and returns the output.
	// // If you want a reference to the prediction, you can call `CreatePrediction`,
	// // call `Wait` on the prediction, and access its `Output` field.
	// prediction, err = r8.CreatePrediction(ctx, version, input, &webhook, false)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // Wait for the prediction to finish
	// err = r8.Wait(ctx, prediction)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("output: ", prediction.Output)
}
