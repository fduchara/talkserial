package main

import (
	"flag"
	"fmt"
	"github.com/tarm/serial"
	"strings"
	"time"
)

func main() {

	serialdevice, atcommand, timeout := ReadFlags()

	c := &serial.Config{Name: serialdevice, Baud: 115200, ReadTimeout: timeout}
	s, _ := serial.OpenPort(c)

	buf := make([]byte, 1)

	n, _ := s.Read(buf)
	//	n, _ = s.Write([]byte("AT^CURC=0\r")) // отключение лога модема.
	n, _ = s.Write([]byte(atcommand + "\r"))

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

func ReadFlags() (string, string, time.Duration) { // читаю аргументы вызова программы. Возвращаю строчку с именем файла.
	ConfigDevIn := flag.String("d", "/dev/ttyUSB1", "serial device")        // читаем переданные параметры.
	ConfigCommandIn := flag.String("c", "ATI", "AT command")                // читаем переданные параметры.
	ConfigReadTimeoutIn := flag.Duration("t", 100000000, "read timeout mc") // 100ms
	flag.Parse()                                                            // парсим параметры
	return *ConfigDevIn, *ConfigCommandIn, *ConfigReadTimeoutIn
}
