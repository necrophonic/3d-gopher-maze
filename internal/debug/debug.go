package debug

import (
	"log"
	"os"
	"strings"
)

func init() {
	d := os.Getenv("DEBUG")
	if strings.ToLower(d) == "true" {
		Debug = true
	}
}

// Debug denotes if debugging is on or off
var Debug = false

// Println delegates to log.Println if Debug==true
func Println(msg ...interface{}) {
	if !Debug {
		return
	}
	log.Println(msg...)
}

// Printf delegates to log.Printf if Debug==true
func Printf(format string, msg ...interface{}) {
	if !Debug {
		return
	}
	log.Printf(format, msg...)
}
