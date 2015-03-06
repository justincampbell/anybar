// Package anybar provides an interface for the AnyBar OS X menubar app:
// https://github.com/tonsky/AnyBar
package anybar

import (
	"fmt"
	"net"
)

const DefaultPort = 1738

func Black()       { _ = Send("black") }
func Blue()        { _ = Send("blue") }
func Cyan()        { _ = Send("cyan") }
func Green()       { _ = Send("green") }
func Orange()      { _ = Send("orange") }
func Purple()      { _ = Send("purple") }
func Red()         { _ = Send("red") }
func White()       { _ = Send("white") }
func Yellow()      { _ = Send("yellow") }
func Question()    { _ = Send("question") }
func Exclamation() { _ = Send("exclamation") }

// Send sets the icon to the string given using the default AnyBar port
func Send(icon string) error {
	return SendTo(icon, DefaultPort)
}

// SendTo sets the icon to the string given using a custom AnyBar port
func SendTo(icon string, port int) error {
	address := fmt.Sprintf("localhost:%v", port)
	conn, err := net.Dial("udp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(icon))
	if err != nil {
		return err
	}

	return nil
}
