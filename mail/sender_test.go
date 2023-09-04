package mail

import (
	"testing"

	"github.com/MohammadZeyaAhmad/Bank-App/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello there,</h1>
	<p>This is a test message from <a href="http://techschool.guru">Mohammad Zeya Ahmad</a></p>
	`
	to := []string{"mohammadzeyaahmad@gmail.com"}
	// attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, nil)
	require.NoError(t, err)
}