package main 

import ( 
    "net/smtp" 
	"text/template"
	"fmt"
	"bytes"
	"os"
	"time"
	"bufio"
	"log"
	"strings"
)

func prepare(tem, subj,Text,Theme,Hello,TextLink, TextForLink, Sub,  Link, type_send, pass, login string, addr_ []string ) {
	
	var body bytes.Buffer
	
	t, _ := template.ParseFiles(tem)
	
    mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
    body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n",subj, mimeHeaders)))
	
	f := func() {
        auth := smtp.PlainAuth("", login, pass, "smtp.mail.ru") 
		smtp.SendMail("smtp.mail.ru:25", auth, login, addr_, body.Bytes()) 
    }
	
    t.Execute(&body, struct {
	 Theme string
	 Hello string
	 Text string
	 Link string
	 TextLink string
	 TextForLink string
	 Sub string
    }{
     Theme:    Theme,
     Hello: Hello,
	 Text: Text, 
	 Link: Link,
	 TextLink: TextLink, 
	 TextForLink: TextForLink, 
	 Sub: Sub, 
    }) 
	
	if type_send == "Да"{
		var date_to_send string
		fmt.Print("Введите дату отправки сообщения в формате 2006-01-31T15:04:05: \n")
		fmt.Scanln(&date_to_send )
		tm, _ := time.Parse("2006-01-02T15:04:05", date_to_send)
		waiting := (tm.Sub(time.Now()))-3*time.Hour
		Timer1 := time.AfterFunc(waiting,  f )
		defer Timer1.Stop()
		time.Sleep(waiting+time.Second)
	}else{
		waiting := 3*time.Second
		Timer1 := time.AfterFunc(waiting,  f )
		defer Timer1.Stop()
		time.Sleep(waiting+time.Second)
	}
}
func Scan1(msg_ string) string {
  fmt.Print(msg_)
  in := bufio.NewScanner(os.Stdin)
  in.Scan()
  if err := in.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
  }
  return in.Text()
}

func main (){
	
	var type_send string
	var Link string
	var type_ int
	var login string
	var pass string
	
	fmt.Print("Введите логин: \n")
	fmt.Scanln(&login)
	
	fmt.Print("Введите пароль: \n")
	fmt.Scanln(&pass)
	
	addr_ := []string{}
	tem := Scan1("Введите путь до шаблона html: ")
    
	subj := Scan1("Введите тему письма: ")
    Text := Scan1("Введите текст письма: ")
    Theme := Scan1("Введите заголовок: ")
    Hello := Scan1("Введите приветствие: ")
	fmt.Print("Введите ссылку: \n")
	fmt.Scanln(&Link )
    TextLink := Scan1("Введите текст сопровождающий ссылку: ")
    TextForLink := Scan1("Введите текст ссылки: ") 
	Sub := Scan1("Введите подпись: ")
	
	fmt.Print("Отлеженное письмо? (Да/Нет) \n")
	fmt.Scanln(&type_send)
	fmt.Print("Рассылка(0) или отправка письма единичному пользователю(1)?: \n")
	fmt.Scanln(&type_)
	
	
	
	// read line by line
	if type_ == 0{
		configFile, err := os.Open("list.txt")
	
		if err != nil {
			log.Fatalf("Error when opening file: %s", err)
		}
		fileScanner := bufio.NewScanner(configFile)
		
		for fileScanner.Scan() {
			str_ := strings.Split(fileScanner.Text(), " ")
			addr_ = append(addr_, str_[2])
		}
		prepare(tem, subj,Text,Theme,Hello, TextForLink,TextLink, Sub,Link,type_send,pass, login,  addr_ )
		configFile.Close()
	}else{
		var Mail string
		Name := Scan1("Введите обращение к адресату : ") 
		fmt.Print("Введите почту: \n")
		fmt.Scanln(&Mail)
		addr_ = append(addr_, Mail)
		prepare(tem, subj,Text,Theme,Hello+", "+Name,TextForLink,TextLink, Sub,Link, type_send,pass, login, addr_ )
	}
}
