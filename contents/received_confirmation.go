package contents

import (
	"fmt"
	"strings"
)

func EmailReceivedConfirmation(senderName string, subject string, content string) string {
	return fmt.Sprintf(`
	Hi there, %v!
	<br><br>
	Thanks for your email to me, I've received it safe and sound. Let me run through your thoughts, feedback, or questions and let me get back
	to you soon once I'm done.
	<br>
	Just a friendly reminder that I received a huge load of email everyday and I aim to answer all of them within <b>48 hours</b>. Should I take any longer,
	I apologise in advance.
	<br>
	On top of everything, I wish you a good day and please stay healthy in the midst of the situation.
	<br><br>
	Cheers,
	<br>
	Hubert
	<br><br><br>
	<i>P.S. Here's your email you sent to me, in case you lose it:</i>
	<br><br>
	<b>Subject:</b> <i>%v</i>
	<br>
	<b>Question:</b> %v
	`, strings.Split(senderName, " ")[0], subject, content)
}
