package utils

import (
	"log"
	"os"
	"regexp"
	"strings"
)

func RemoveComment(text string) (string, error) {
	regexpPattern, err := regexp.Compile(`\n.*// `)
	if err != nil {
		return text, err
	}
	match := regexpPattern.ReplaceAllString(text, "")
	rep := strings.Replace(match, "//", "", -1)
	res := strings.Replace(rep, "\n  ", "", -1)
	return res, nil
}

func GetEnvValue(env string) (string, error) {
	// need to parepare setting API key in env
	val, ok := os.LookupEnv(env)
	if !ok {
		log.Fatalf("Not set environment %s", env)
	} else if val == "" {
		log.Fatalf("Env %s is empty", env)
	}

	return os.Getenv(env), nil
}
