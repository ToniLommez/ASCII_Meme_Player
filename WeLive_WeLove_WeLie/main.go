package main

import (
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

	clearScreen()
	smurf := color.New(color.FgBlue)
	time.Sleep(1000 * time.Millisecond)

	play <- true

	fmt.Println()
	fmt.Println()

	fmt.Fprint(color.Output, smurf.SprintFunc()("\t    It "))
	time.Sleep(200 * time.Millisecond)
	fmt.Fprint(color.Output, smurf.SprintFunc()("all "))
	time.Sleep(300 * time.Millisecond)
	fmt.Fprint(color.Output, smurf.SprintFunc()("belongs "))
	time.Sleep(700 * time.Millisecond)
	fmt.Fprint(color.Output, smurf.SprintFunc()("to "))
	time.Sleep(300 * time.Millisecond)
	fmt.Fprint(color.Output, smurf.SprintFunc()("the "))
	time.Sleep(300 * time.Millisecond)
	fmt.Fprint(color.Output, smurf.SprintFunc()("other"))
	time.Sleep(700 * time.Millisecond)
	fmt.Fprint(color.Output, smurf.SprintFunc()("side"))
	time.Sleep(300 * time.Millisecond)
	fmt.Fprint(color.Output, smurf.SprintFunc()("."))
	time.Sleep(600 * time.Millisecond)
	fmt.Fprint(color.Output, smurf.SprintFunc()("."))
	time.Sleep(600 * time.Millisecond)
	fmt.Fprint(color.Output, smurf.SprintFunc()(".\n"))
	time.Sleep(200 * time.Millisecond)

	fmt.Fprint(color.Output, smurf.SprintFunc()("\t\t        WE LIVE\n"))
	time.Sleep(900 * time.Millisecond)
	fmt.Fprint(color.Output, smurf.SprintFunc()("\t\t        WE LOVE\n"))
	time.Sleep(900 * time.Millisecond)
	fmt.Fprint(color.Output, smurf.SprintFunc()("\t\t        WE LIE\n"))
	time.Sleep(500 * time.Millisecond)

	fmt.Println()
	fmt.Println()

	fmt.Fprint(color.Output, smurf.SprintFunc()("    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠀⢀⠀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠈⠀⠄⠂⠄⢂⠈⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n     ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⠀⠠⠁⠌⠄⡑⠄⡂⢂⠡⠠⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n      ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠄⠂⠀⡁⠌⠄⠅⢅⢊⠰⡐⠠⠂⠅⡂⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n       ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⠀⠄⠂⠄⡂⠅⠅⡂⠪⡐⡌⠪⠨⢂⠂⠅⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀ \n       ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⠂⠐⠈⠄⢂⠅⡐⠨⠨⡂⡊⢜⠨⠠⠡⢁⠂⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀ \n       ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠠⠀⠠⠁⠐⠠⢁⢂⠊⠌⢌⠐⠄⠅⡊⠨⠨⠐⡈⠄⠡⢀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀ \n       ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⠈⠀⠠⠐⠀⠐⠀⠂⠠⠈⠄⢁⠂⡁⢂⠠⠐⠀⢁⠠⠈⡈⠄⠂⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀ \n       ⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠁⠀⠐⠀⠄⠀⠄⠈⢀⠂⠀⠂⠐⢀⠠⠀⢂⠠⠀⡂⠠⠀⠂⡀⠐⡀⠄⠂⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀ \n       ⠀⠀⠀⠀⠀⠀⠀⠠⠈⠀⢀⢀⠁⡠⢀⢂⠠⢐⢠⢠⢨⢨⡐⡄⡢⢨⢠⠠⡡⡠⠡⡁⡢⠀⠅⠠⠐⢀⠄⢁⠂⢀⠀⠀⠀⠀⠀⠀⠀⠀ \n       ⠀⠀⠀⢀⠠⢐⠨⡠⢌⢆⢖⡲⡸⣘⢦⣓⢭⡣⣇⡧⣳⢵⢕⣧⣫⢮⢮⣪⣎⡎⣎⢖⢜⢬⣘⡐⡅⢄⢂⢐⠠⢀⠀⡀⠀⠀⠀⠀⠀⠀ \n       ⠀⢀⢔⢰⢱⡱⣕⢧⣗⣗⣷⢽⣽⣞⣷⢽⣗⣟⣷⢿⣳⡿⣽⡷⣟⣿⣟⣿⣾⣻⣾⣯⣟⣷⣳⡽⣮⣞⡴⣢⡑⢔⠠⢀⠀⠀⠀⠀⠀⠀ \n       ⠐⢌⡪⣪⢧⣻⣺⢽⣞⣯⡿⣯⣷⢿⣾⣟⣷⡿⣾⣿⣟⣿⣟⣿⣿⢿⣻⣽⣾⣿⣷⡿⣟⣯⣿⢿⣿⣾⢿⣷⣻⣼⡰⣂⢂⠄⠀⠀⠀⠀ \n       ⠈⠔⡚⡮⣗⣟⣾⣻⣽⣷⡿⣿⢾⣿⢿⣽⣟⣿⣿⣯⣿⣿⣻⣟⣿⣿⣟⣿⣻⣽⣾⡿⣟⣿⣻⡯⣷⣿⢿⣯⣿⣞⢷⠵⠥⠀⠀⠀⠀⠀ \n    "))
	fmt.Fprint(color.Output, smurf.SprintFunc()("   ⠀⠀⠈⠊⠓⠛⠮⠯⣿⢾⣟⣿⣻⢽⣟⣗⣿⣽⣷⡿⣗⣿⣿⣿⣷⣿⣽⣟⡿⣯⣿⣻⣿⣽⣷⣿⢿⣾⣻⣽⢾⣻⡫⡫⠁⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠈⣾⣿⢷⣳⢝⠕⡕⡵⣯⣿⣽⣟⣿⣿⣷⣿⣿⣿⣷⣻⣻⣽⣾⣻⣞⣿⣺⣽⣿⢽⡯⡿⡽⡒⠁⠀⠀⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⡝⠎⠂⠀⠄⡝⣞⢮⢗⠯⣝⢍⡟⣟⢟⣝⢮⢯⣟⣞⣗⣟⣞⡾⣽⡞⡯⡏⢎⠃⡑⠀⠀⠀⠀⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠀⡀⢄⢐⠌⡎⡫⡊⢈⠘⡸⡱⣝⢗⡯⡯⣫⢗⡗⣗⢗⢗⡳⣝⢵⠝⠜⠠⠑⠁⠀⠀⠀⢄⠀⡠⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢂⢳⢽⢮⡿⡾⡚⡐⡰⡨⡢⡱⣱⡫⣏⢯⡺⡵⡹⢜⢎⠪⠂⠁⠁⠈⠈⠀⠀⠀⠀⠀⢸⣸⠠⡀⠄⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠨⠢⡍⣗⢯⣓⢜⡜⡼⡸⡪⡺⡪⡪⡪⡪⡪⡪⢊⠊⠄⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⢕⢮⢊⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠁⠘⠈⢇⢎⡗⣝⡜⡝⡪⡣⣋⢎⢮⣪⢪⡢⡁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢔⠔⠁⣸⢝⢆⠀⠄⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠀⠈⠀⠀⠱⡸⣝⣗⣗⢵⢱⡐⡴⣦⢀⡀⠀⠀⡀⠀⣀⠜⠁⠀⠀⣪⣯⠢⡠⠀⠂⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠄⢌⢺⡺⡺⡪⡎⣖⣟⢎⠔⢭⢺⢮⣻⣢⣷⠁⠀⠀⠀⠀⣺⢎⢮⢢⢅⢂⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠌⠠⢑⠕⠕⡕⣝⢞⡕⡕⢕⢍⢢⠡⢑⢾⣻⣺⡥⠀⠀⠀⠀⡢⣯⢾⣹⣪⢖⢌⠌ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⠌⠀⢂⠂⡌⢎⠪⡨⡊⡎⢎⢎⣗⡵⡱⡡⠊⠝⡷⡇⠀⠀⢀⢮⣞⣮⢯⣞⡷⣝⡆⠇ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠔⠌⠀⠌⢔⠪⡈⢆⢕⢜⢜⣜⡵⣿⡾⣿⠙⠪⡣⡕⡨⢑⠀⠀⠨⡯⣿⢾⣻⣮⢟⡎⠎⠂ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⢎⠆⠃⠁⠈⡎⠪⡘⡌⡎⡖⡵⣳⣳⣟⣿⣻⡽⡁⠀⠀⠈⢚⡎⡔⡀⠈⠊⠕⠟⡝⢎⠃⠁⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⡢⠪⠂⠂⠀⢀⠱⡁⠅⠢⡨⡪⡸⡪⣳⢳⣳⣻⣪⡳⡐⠀⠀⠀⠀⠈⢪⡂⡂⠀⠀⠀⠀⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠄⠂⠄⠌⡀⠀⠀⠀⡨⢂⠢⡑⡱⡘⡜⡜⣎⢮⢳⡣⣗⢞⡮⡧⠡⠀⠀⠀⠀⠀⠱⣜⢄⠄⡀⡀⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠄⡠⡡⣊⠔⠡⢂⠆⠀⠀⠀⢊⢐⠅⠕⡌⢎⢪⢪⡪⡎⡧⡳⢕⢏⠎⢎⠨⠂⠀⠀⠀⠀⠨⣳⢗⡵⡡⠢⢥⢀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⢄⠡⣪⠋⠀⠀⠠⡰⡑⠀⠀⠀⠐⠄⠕⡑⡌⡪⡪⡪⠪⡊⠎⢊⠌⡐⡌⡂⡃⠁⠢⡠⡀⠀⠈⠘⡽⣪⢷⢭⡣⡣⡫⠀⠀ \n    ⠀⠀⠀⠀⠈⢆⢞⡎⠀⠀⠀⠈⠊⠀⠀⠀⠀⠀⠁⠅⠊⢌⠊⠌⠌⠨⡠⠨⢀⠊⡀⢂⠠⢀⠡⠀⠑⠀⠀⠀⠀⠈⠘⠹⡳⢧⠣⡊⠔⠀ \n    "))
	fmt.Fprint(color.Output, smurf.SprintFunc()("   ⠀⠀⠀⠀⠀⠠⡳⠍⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠌⠠⠐⡀⢂⠅⡄⠊⢀⠐⡀⠂⠄⡂⠄⡂⠅⠀⠀⠀⠀⠀⠀⠀⠀⠈⠃⠣⣂⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡐⠨⠨⠠⡃⠀⠌⢀⢂⠢⡡⡑⢄⠕⡐⠠⢀⠐⢀⠀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠌⢌⠂⢀⠌⡐⡡⢂⠊⢔⢌⠆⡕⢌⢪⢐⢌⢆⢐⠠⠀⠀⠀⠀⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⢐⢐⠌⡂⢇⠣⡱⢘⠈⠌⠢⡣⡃⡪⡐⢌⢀⠀⠀⠀⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠄⢅⠂⠂⡁⠊⠔⠑⠈⠀⠀⠈⠈⡢⡊⡢⢑⢐⢐⠀⠀⠀⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠨⢐⠈⠄⢂⠡⠀⠀⠀⠀⠀⠀⡐⡐⡐⠌⡒⡔⣐⢐⠀⠀⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠀⡀⢀⠀⡀⢀⠀⡈⠄⠌⠌⠄⠂⠀⠀⠀⠀⢀⢂⢆⢢⢡⢑⢌⢜⢔⡕⡄⠀⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠠⢀⢁⢐⠠⢂⢐⢐⠐⡀⡂⡢⢁⠅⠅⠅⠀⠀⠀⠀⠀⠂⢆⢪⢸⢸⢸⡸⡱⡕⡕⠂⠀⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⠰⢌⣎⢦⡲⡑⡔⡔⡐⢌⠔⡐⢌⠔⡠⢅⠁⠀⠀⠀⠀⠀⠀⠁⠁⠑⠑⠁⠁⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠨⠪⠳⡯⣾⣺⢼⣬⢎⢪⢊⡢⠱⡰⡐⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀ \n    ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠀⠉⠉⠚⠛⠿⢶⢵⠱⣑⢸⠀⠀\n"))
}

func main() {
	play := make(chan bool)

	go frame(play)
	playAudio("www.mp3", play)
	fmt.Printf("Aperte enter para sair")
	fmt.Scanln()
}
