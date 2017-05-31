package main

import (
	"github.com/tarm/serial"
//	"log"
	"flag"
	"strings"
	"time"
	"fmt"
)


func main() {

	serialdevice, atcommand := ReadFlags()

	c := &serial.Config{Name: serialdevice, Baud: 115200, ReadTimeout: time.Millisecond * 10}
	s, _ := serial.OpenPort(c)

	buf := make([]byte, 1)

	n, _ := s.Write([]byte("AT^CURC=0\r"))
	n, _ = s.Read(buf)
	n, _ = s.Write([]byte(atcommand+"\r"))

	ret := string("")
	for {
		n, _ = s.Read(buf)
		str := string(buf[:n])
		if strings.Count(str, "") > 1 {
			ret = ret + str
		}
		if strings.Count(str, "") == 1 {
			break
		}
	}

	fmt.Printf(ret)

}


func ReadFlags() (comport string, command string) { // читаю аргументы вызова программы. Возвращаю строчку с именем файла.
	ConfigDevIn := flag.String("d", "/dev/ttyUSB1", "serial device") // читаем переданные параметры.
	ConfigCommandIn := flag.String("c", "ATI", "AT command") // читаем переданные параметры.
	flag.Parse()                                                              // парсим параметры
	return *ConfigDevIn, *ConfigCommandIn
}