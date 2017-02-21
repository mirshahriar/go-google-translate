package pkg

import (
	"errors"
	"fmt"
	"github.com/aerokite/google-translate/pkg/client"
	"regexp"
	"strings"
)

type TranslateRequest struct {
	SourceLang string
	TargetLang string
	Text       string
}

func Translate(req TranslateRequest) (string, error) {
	config := &client.Config{
		Source: req.SourceLang,
		Target: req.TargetLang,
	}
	resp := client.NewClient(config).Translate(req.Text).Get().Do()
	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Failed. Status:", resp.Status))
	}
	respHtml := string(resp.ResponseBody)

	re := regexp.MustCompile(`class="t0">(.*?)<`)
	match := re.FindStringSubmatch(respHtml)
	if len(match) != 2 {
		return "", errors.New("Failed to translate")
	}

	translated := strings.Replace(match[1], "&quot;", "", -1)
	return translated, nil
}
