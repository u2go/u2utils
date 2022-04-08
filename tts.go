package u2utils

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

const (
	TtsZh = "zh"
)

func TtsSpeak(text, dir, lang string) error {

	err := FileMkdirAll(dir)
	if err != nil {
		return err
	}
	file := path.Join(dir, SHA256(text)+".mp3")
	exists, err := FileExists(file)
	if err != nil {
		return err
	}
	if !exists {
		url := fmt.Sprintf("https://tts.youdao.com/fanyivoice?word=%s&le=%s&keyfrom=speaker-target",
			url.QueryEscape(text), lang)
		response, err := http.Get(url)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		output, err := os.Create(file)
		if err != nil {
			return err
		}
		defer output.Close()

		_, err = io.Copy(output, response.Body)
		if err != nil {
			return err
		}
	}

	return TtsSpeakFile(file)
}

func TtsSpeakFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return err
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done

	return nil
}
