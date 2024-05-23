package main

import (
	"github.com/feranwq/ppt/pkg/clipboard"
	"github.com/feranwq/ppt/pkg/translate"
	"github.com/feranwq/ppt/pkg/utils"
	"log"
)

func main() {

	appID, err := utils.GetEnvValue("BD_APP_ID")
	if err != nil {
		log.Fatal(err)
	}

	secretKey, err := utils.GetEnvValue("BD_SECRET_KEY")
	if err != nil {
		log.Fatal(err)
	}

	pasteText, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}
	processPasteText, err := utils.FormatComment(pasteText)
	log.Println("\n", processPasteText)
	if err != nil {
		panic(err)
	}
	err = clipboard.WriteAll(processPasteText)
	if err != nil {
		panic(err)
	}

	bi := translate.BaiduInfo{AppID: appID, Salt: translate.Salt(5), SecretKey: secretKey, From: "auto", To: "zh"}
	bi.Text = processPasteText
	log.Println(bi.Translate())
}
