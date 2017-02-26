package main

import (
	"fmt"
	"log"

	"github.com/aerokite/go-google-translate/pkg"
	"github.com/spf13/cobra"
)

func translate(req *pkg.TranslateRequest) {
	translated, err := pkg.Translate(req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(translated)
}

func newCmdGopret() *cobra.Command {
	var req pkg.TranslateRequest
	cmd := &cobra.Command{
		Use:   "translate",
		Short: "text for translation",
		Run: func(cmd *cobra.Command, args []string) {
			translate(&req)
		},
	}
	cmd.Flags().StringVar(&req.SourceLang, "sl", "", "Translate from source language to targer")
	cmd.Flags().StringVar(&req.TargetLang, "tl", "en", "Translate to target language from source")
	cmd.Flags().StringVar(&req.Text, "text", "", "Translate text from source language to target language")
	return cmd
}
