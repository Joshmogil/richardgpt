package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/joshmogil/richardgpt/videotapes"
	"github.com/urfave/cli/v2"

)

func main() {
	var api_token string


	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "token",
			Aliases: []string{"t"},
			Usage:   "provide a token to use to authenticate with OpenAi API",
			EnvVars: []string{"OPENAI_KEY"},
			Destination: &api_token,
		},

		&cli.StringFlag{
			Name:    "model",
			Aliases: []string{"m"},
			Usage:   "specify which model to use",
		},
	}

    app := &cli.App{
		Name: "rich",
		Usage: "A CLI for richard iovanisci",
		Flags: flags,
        Commands: []*cli.Command{
            {
                Name:    "question",
                Aliases: []string{"q"},
                Usage:   "ask a question to richard",
                Action: func(cCtx *cli.Context) error {
					err := videotapes.Validate()
					
					if err != nil {
						fmt.Println("No api key found, please provide one by setting the OPENAI_KEY environment variable")
						return nil
					}

					fmt.Println("Asking richard a question ...")
					gpt := &videotapes.GptConnection{ApiToken: api_token}
					base_prompt := "You are a great programmer at a large company, people sometimes call you richard, please answer the following question, without mentioning you are an ai language model:\n"
					user_input := cCtx.Args().Slice()[0]
					
					final_prompt := base_prompt + user_input

					response, err := gpt.GetResponse(final_prompt)
					if err != nil {
						fmt.Println("Unable to connect to richard, sorry.")
						return nil
					}


					fmt.Println(response)

					return nil
                },
            },
            {
                Name:    "generate",
                Aliases: []string{"g"},
                Usage:   "generate some code from richard",
                Action: func(cCtx *cli.Context) error {
                    err := videotapes.Validate()
					if err != nil {
						fmt.Println("No api key found, please provide one by setting the OPENAI_KEY environment variable")
						return nil
					}

					fmt.Println("Stealing richard's code ...")
					gpt := &videotapes.GptConnection{ApiToken: api_token}
					base_prompt := "You are a great programmer at a large company who can't speak, please write some code according to the following specifications, please only provide code:\n"
					user_input := cCtx.Args().Slice()[0]
					
					final_prompt := base_prompt + user_input

					response, err := gpt.GetResponse(final_prompt)
					if err != nil {
						fmt.Println("Unable to connect to richard, sorry.")
						return nil
					}


					fmt.Println(response)

					return nil
                },
            },
        },
    }

    sort.Sort(cli.FlagsByName(app.Flags))
    sort.Sort(cli.CommandsByName(app.Commands))

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
