package smtp

import "net/smtp"

func EmailSendMain() {
	println("email func")
	// 메일서버 로그인 정보설정
	auth := smtp.PlainAuth("", "hschae@euclidsoft.co.kr", "ch798240", "smtp.cafe24.com")

	from := "hschae@euclidsoft.co.kr"
	to := []string{"hschae@euclidsoft.co.kr", "prettychs@gmail.com"}

	// 메시지작성 : RFC822(이메일스타일) 헤더+빈칸한줄+메시지내용
	headerSubject := "메일제목입니다.\r\n"
	headerBlank := "\r\n"
	body := "메일 내용입니다. \r\n" //마지막은 꼭 CRLF(new line으로 바꾸는방식)로 끝나야한다.
	msg := []byte(headerSubject + headerBlank + body)

	// 메일보내기
	err := smtp.SendMail("smtp.cafe24.com:587", auth, from, to, msg)
	if err != nil {
		panic(err)
	}
}
