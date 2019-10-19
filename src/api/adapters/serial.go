package adapters

import (
	"log"

	"github.com/tarm/serial"

	"github.com/na3alf6/LynxmotionSerialServo/src/api/domain"
)

func initSerial(settings domain.ConfigSerial) (*serial.Port, error) {
	s, err := serial.OpenPort(
		&serial.Config{
			Name: settings.Port,
			Baud: settings.BaundRate,
		})
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("%q", buf[:n])
	return s, nil
}
