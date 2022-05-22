package welcome

import (
	"encoding/json"
	"io"
	"os"
)

type database struct{}

type user struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

func (d database) getUser(fileName string) ([]user, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var us []user
	err = json.Unmarshal(bytes, &us)
	if err != nil {
		return nil, err
	}

	return us, nil
}
