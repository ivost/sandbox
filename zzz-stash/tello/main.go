/*
To automated the DJI Tello Drone using Gobot Framework
*/
package main

import (
	"io" //  It provides basic interfaces to I/O primitives
	"log"
	"os/exec" // To run the external commands.
	"strconv" // Package strconv implements conversions to and from string
	"sync"
	"time" //For time related operation

	"gobot.io/x/gobot"                     // Gobot Framework.
	"gobot.io/x/gobot/platforms/dji/tello" // DJI Tello package.
	"gocv.io/x/gocv"                       // GoCV package to access the OpenCV library.
)

// Frame size constant.
const (
	frameX    = 960
	frameY    = 720
	frameSize = frameX * frameY * 3
)

type MyState struct {
	val int
	m   sync.Mutex
}

func (st *MyState) Get() int {
	st.m.Lock()
	defer st.m.Unlock()
	return st.val
}

func (st *MyState) Set(val int) {
	st.m.Lock()
	defer st.m.Unlock()
	st.val = val
}

func main() {
	var st MyState
	var err error

	drone := tello.NewDriver("8888")

	work1 := func() {
		//drone.StartVideo()
		//drone.SetVideoEncoderRate(tello.VideoBitRateAuto)
		//drone.SetExposure(0)
		log.Println("work1 enter")
		st.Set(2)
		err := drone.TakeOff()
		if err != nil {
			panic(err)
		}
		drone.Hover()
		st.Set(3)
		time.Sleep(15 * time.Second)
		st.Set(4)
		drone.Land()
		st.Set(5)
		log.Println("work1 exit")
	}

	//work := func() {
	//	err := drone.On(tello.ConnectedEvent, func(data interface{}) {
	//		log.Println("Connected to Tello.")
	//		st.Set(1)
	//		work1()
	//	})
	//	if err != nil {
	//		st.Set(-1)
	//		panic(err)
	//	}
	//}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		//work,
	)

	err = drone.On(tello.WifiDataEvent, func(data interface{}) {
		log.Printf("WifiDataEvent %+v ", data)
		st.Set(1)
		work1()
	})
	if err != nil {
		panic(err)
	}

	err = drone.On(tello.ConnectedEvent, func(data interface{}) {
		log.Printf("Connected %+v ", data)
		st.Set(2)
		//work1()
	})
	if err != nil {
		panic(err)
	}

	gobot.Every(2*time.Second, func() {
		log.Printf("TICK...state %d\n", st.Get())
	})

	err = robot.Start(false)
	if err != nil {
		panic(err)
	}

	time.Sleep(20 * time.Second)
	log.Println("stopping...")

	robot.Stop()
	//drone.Land()
	//drone.Halt()
}

func main1() {
	// Driver: Tello Driver
	drone := tello.NewDriver("8890")

	//Robot: Tello Drone
	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
	)

	// calling Start(false) lets the Start routine return immediately without an additional blocking goroutine
	err := robot.Start(true)
	if err != nil {
		panic(err)
	}

	// OpenCV window to watch the live video stream from Tello.
	window := gocv.NewWindow("Tello")

	//FFMPEG command to convert the raw video from the drone.
	ffmpeg := exec.Command("ffmpeg", "-hwaccel", "auto", "-hwaccel_device", "opencl", "-i", "pipe:0",
		"-pix_fmt", "bgr24", "-s", strconv.Itoa(frameX)+"x"+strconv.Itoa(frameY), "-f", "rawvideo", "pipe:1")
	ffmpegIn, _ := ffmpeg.StdinPipe()
	ffmpegOut, _ := ffmpeg.StdoutPipe()

	work := func() {
		//Starting FFMPEG.
		log.Printf("Starting ffmpeg... %+v", ffmpeg.Args)
		if err := ffmpeg.Start(); err != nil {
			log.Println(err)
			return
		}
		// Event: Listening the Tello connect event to start the video streaming.
		drone.On(tello.ConnectedEvent, func(data interface{}) {
			log.Println("Connected to Tello.")
			drone.StartVideo()
			drone.SetVideoEncoderRate(tello.VideoBitRateAuto)
			drone.SetExposure(0)

			//For continued streaming of video.
			gobot.Every(100*time.Millisecond, func() {
				drone.StartVideo()
			})
		})

		//Event: Piping the video data into the FFMPEG function.
		drone.On(tello.VideoFrameEvent, func(data interface{}) {
			pkt := data.([]byte)
			if _, err := ffmpegIn.Write(pkt); err != nil {
				log.Println(err)
			}
		})

		//TakeOff the Drone.
		gobot.After(2*time.Second, func() {
			log.Println("Tello Taking Off...")
			drone.TakeOff()
		})

		//Land the Drone.
		gobot.After(5*time.Second, func() {
			log.Println("Tello Landing...")
			drone.Land()
		})
	}

	robot.Work = work

	////Robot: Tello Drone
	//robot := gobot.NewRobot("tello",
	//	[]gobot.Connection{},
	//	[]gobot.Device{drone},
	//	work,
	//)
	//
	//// calling Start(false) lets the Start routine return immediately without an additional blocking goroutine
	//err := robot.Start(true)
	//if err != nil {
	//	panic(err)
	//}

	// now handle video frames from ffmpeg stream in main thread, to be macOs friendly
	for {
		buf := make([]byte, frameSize)
		if _, err := io.ReadFull(ffmpegOut, buf); err != nil {
			log.Println(err)
			continue
		}
		img, _ := gocv.NewMatFromBytes(frameY, frameX, gocv.MatTypeCV8UC3, buf)
		if img.Empty() {
			continue
		}

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
