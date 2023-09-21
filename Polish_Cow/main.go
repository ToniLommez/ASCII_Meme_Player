package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/fatih/color"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func loadFrames(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var frames []string
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&frames)
	return frames, err
}

func playAudio(filePath string, play chan bool) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	decoder, err := mp3.NewDecoder(file)
	if err != nil {
		fmt.Println("Erro ao decodificar o arquivo:", err)
		return
	}

	context, err := oto.NewContext(decoder.SampleRate(), 2, 2, 8192)
	if err != nil {
		fmt.Println("Erro ao criar um novo contexto de áudio:", err)
		return
	}
	defer context.Close()

	player := context.NewPlayer()
	defer player.Close()

	// Quando estiver pronto para inicializar frame
	play <- true

	<-play
	if _, err := io.Copy(player, decoder); err != nil {
		fmt.Println("Erro ao reproduzir o áudio:", err)
		return
	}
}

func frame(play chan bool) {
	<-play // Espera o sinal de playAudio

	cow := color.New(color.FgHiWhite)
	clearScreen()
	const cursorReset = "\033[H"
	frames, err := loadFrames("cow.gob")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i, f := range frames {
		frames[i] = cow.SprintFunc()(f)
	}
	time.Sleep(1000 * time.Millisecond)

	play <- true
	time.Sleep(200 * time.Millisecond)
	for {
		for _, f := range frames {
			fmt.Println(cursorReset)
			fmt.Println(f)
			time.Sleep(30 * time.Millisecond)
		}
	}
}

func main() {
	play := make(chan bool)

	go frame(play)
	go playAudio("cow_song.mp3", play)

	fmt.Scanln()
}
