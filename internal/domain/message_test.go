package domain

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MessageTestSuite struct {
	suite.Suite
}

func (s *MessageTestSuite) SetupTest() {}

func (s *MessageTestSuite) TestNewMessageValidInput() {
	content := "Hello world"
	userName := "user123"
	msg, err := NewMessage(content, userName)

	s.NoError(err)
	s.NotNil(msg)

	s.Equal(content, msg.Content())
	s.Equal(userName, msg.UserName())
	s.NotEmpty(msg.CreatedAt())
}

func (s *MessageTestSuite) TestNewMessageEmptyContent() {
	msg, err := NewMessage("", "user123")

	s.Error(err)
	s.Nil(msg)
	s.Equal(ErrContentEmpty, err)
}

func (s *MessageTestSuite) TestNewMessageContentExceedsLimit() {
	longContent := string(make([]byte, 281))
	msg, err := NewMessage(longContent, "user123")

	s.Error(err)
	s.Nil(msg)
	s.Equal(ErrContentTooLong, err)
}

func (s *MessageTestSuite) TestNewMessageEmptyUserName() {
	msg, err := NewMessage("Hola, mundo!", "")

	s.Error(err)
	s.Nil(msg)
	s.Equal(ErrUserNameEmpty, err)
}

func TestMessageSuite(t *testing.T) {
	suite.Run(t, new(MessageTestSuite))
}
