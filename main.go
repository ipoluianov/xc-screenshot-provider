package main

import (
	"bytes"
	"fmt"
	"image/png"

	"github.com/ipoluianov/xchg/xchg_samples"
	"github.com/kbinani/screenshot"
)

func makeScreenshot() []byte {
	var result []byte
	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		buf := bytes.NewBuffer(make([]byte, 0))
		img, err := screenshot.CaptureRect(bounds)
		if err == nil {
			png.Encode(buf, img)
			result = buf.Bytes()
		}
		break
	}
	return result
}

func main() {
	fmt.Println("main")
	s := xchg_samples.StartServerFast("pass", func(function string, parameter []byte) (response []byte, err error) {
		bs := makeScreenshot()
		return bs, nil
	})
	fmt.Println("Started:", s.Address())
	fmt.Scanln()
}
