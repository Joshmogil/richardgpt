package videotapes

import (
	"os"
	"errors"


)

func Validate() (error) {
	good, err := checkEnv("OPENAI_KEY")
	if good {
		return nil
	} else {
		return err
	}

}

func checkEnv (key string) (bool, error) {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		return false, errors.New("OpenAi Key Not Set")
	} else {
		return true, nil
	}
}