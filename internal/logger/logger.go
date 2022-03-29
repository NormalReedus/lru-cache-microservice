package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

//TODO: take a log file from config to output to if provided, otherwise use Stdout

var (
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
)

func init() {
	clrInfo := color.New(color.FgBlue, color.Bold)
	infoLog = log.New(os.Stdout, clrInfo.Sprint("ℹ️ INFO - "), log.Ldate|log.Ltime|log.Lmsgprefix)

	clrWarn := color.New(color.FgYellow, color.Bold)
	warningLog = log.New(os.Stdout, clrWarn.Sprint("⚠️ WARN - "), log.Ldate|log.Ltime|log.Lmsgprefix)

	clrErr := color.New(color.FgRed, color.Bold)
	errorLog = log.New(os.Stdout, clrErr.Sprint("⛔ ERROR - "), log.Ldate|log.Ltime|log.Lmsgprefix)
}

func CacheWrite(key string) {
	clr := color.New(color.FgBlue, color.Bold)
	infoLog.Println(clr.Sprint("CACHE WRITE: " + key))
}

func CacheRead(key string) {
	clr := color.New(color.FgGreen, color.Bold)
	infoLog.Println(clr.Sprint("CACHE READ: " + key))
}

func CacheEvict(key string) {
	clr := color.New(color.FgRed, color.Bold)
	infoLog.Println(clr.Sprint("CACHE EVICT: " + key))
}

func CacheBust(key string) {
	clr := color.New(color.FgRed, color.Bold)
	infoLog.Println(clr.Sprint("CACHE BUST: " + key))
}

func CacheSkip(key string) {
	clr := color.New(color.FgYellow, color.Bold)
	infoLog.Println(clr.Sprint("CACHE SKIP: " + key))
}

func Warn(msg string) {
	warningLog.Println(msg)
}

func Fatal(err error) {
	errorLog.Fatalln(err)
}

func HiMom(apiUrl string, port string) {
	cacheColor := color.New(color.FgBlue, color.Bold)
	urlColor := color.New(color.FgHiGreen, color.Underline)

	fmt.Print("Your ")
	cacheColor.Print("LRU cache microservice ")
	fmt.Printf("is being served on http://localhost:%v.\n", port)
	fmt.Print("All requests will be proxied to ")
	urlColor.Println(apiUrl + "\n")
}
