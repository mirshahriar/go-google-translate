# go-google-translate

**Do you want to use google translator API?** API is not free anymore. But you can translate 1000 words per day free. And then may be this project can help you to translate unlimited text from one language to another using Google Translator.

# Overview

`go-google-translate` provides a `go` package to translate using Google Translator by parsing HTML.  This library can be used to translate from any language to any other.

In **addition** this library provides a [cli](#cli) to translate text via console.

# Installation

```sh
$ go get -u -v github.com/aerokite/go-google-translate/...
```

# Usage
```go
import (
	trans "github.com/aerokite/go-google-translate/pkg"
)

// request struct
req := &trans.TranslateRequest{
  SourceLang: "bn",
  TargetLang: "en",
  Text:       "আমি বাংলায় গান গাই",
}
// translate
translatedText, err := trans.Translate(req)
if err != nil {
  os.Stderr.WriteString("Failed to translate text, error" + err.Error())
}

fmt.Println(translatedText) // I sing in Bangla
```

# CLI

Command line interface to translate text using command line.

## Install

```sh
$ go get -u -v github.com/aerokite/go-google-translate/cmd/gopret
$ go install github.com/aerokite/go-google-translate/cmd/gopret
```

## Usages

```sh
$ gopret --sl bn --tl en --text "আমি বাংলায় গান গাই"
I sing in Bangla
```

# Acknowledgement

[Arnaud Aliès](https://github.com/mouuff) for [mtranslate](https://github.com/mouuff/mtranslate) in python.

# License
Copyright (c) 2017 Mir Shahriar

Licensed under [MIT Licence](LICENSE).
