package main

import (
    "fmt"
    "time"
)

func main() {

    ticker := time.NewTicker(10 * time.Second)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(60 * time.Minute)

    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}

/*
package main

import (
    "flag"
    "fmt"
    "time"

    "syscall"
)

// #include <unistd.h>
// //#include <errno.h>
// //int usleep(useconds_t usec);
import "C"

func usleep() {
    for {
        C.usleep(1000)
    }
}

func sysns() {
    for {
        v := syscall.Timespec{
            Sec:  0,
            Nsec: 1000,
        }
        syscall.Nanosleep(&v, &v)
    }
}

func ticker() {

    ticker := time.NewTicker(time.Millisecond)
    defer ticker.Stop()
    for range ticker.C {
    }
}

func timer() {
    t := time.NewTimer(time.Millisecond)
    for range t.C {
        t.Reset(time.Millisecond)
    }
}

func sleep() {
    for {
        time.Sleep(time.Millisecond)
    }
}

func main() {
    t := flag.String("t", "timer", "use timer")
    flag.Parse()
    switch *t {
    case "timer":
        timer()
    case "ticker":
        ticker()
    case "sleep":
        sleep()
    case "cgo":
        usleep()
    case "sys":
        sysns()
    default:
        fmt.Println("use  timer, ticker, sys, cgo or sleep")
    }
}
*/
