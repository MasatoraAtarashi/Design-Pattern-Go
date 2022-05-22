package welcome

import (
	"errors"
	"fmt"
)

type PageMaker interface {
	Make(mailAddr string, fileName string) error
}

type pageMaker struct {
}

func NewPageMaker() PageMaker {
	return pageMaker{}
}

func (p pageMaker) Make(email string, fileName string) error {
	users, err := database{}.getUser(fileName)
	if err != nil {
		return err
	}

	u, err := toUser(users, email)
	if err != nil {
		return err
	}

	writer := newHTMLWriter()
	writer.Title(fmt.Sprintf("Welcome to %s's page!", u.UserName))
	writer.Paragraph("Welcome to " + u.UserName + "'s web page!")
	writer.Paragraph("Nice to meet you!")
	writer.MailTo(u.Email, u.UserName)
	writer.Close()

	return nil
}

func toUser(users []user, email string) (user, error) {
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}
	return user{}, errors.New("user not found")
}
