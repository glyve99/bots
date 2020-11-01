package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file: ", err)
	}
	temp := os.Getenv("TOKEN")
	flag.StringVar(&token, "t", temp, "Bot Token")
	flag.Parse()
}

var token string
var bufferMap = make(map[string][][]byte)
var audioMutex = &sync.Mutex{}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		return
	}

	if strings.HasPrefix(m.Content, "!ajuda") {
		s.ChannelMessageSend(m.ChannelID, `Dá pra lançar esses:
!passeio @usuário


COMANDOS DE AUDIO:
!cavalo	!stronda	!rafa		!sofro
!milton	!alek		!zoio		!tchurosbango
!yamete	!rs		!tarde		!fish
!jojofag	!lohane		!miranha	!tripaloski
!potencia	!parmalate	!oba		!gado
!implacavel	!gabriel    !poze	!molina
!djazeitona !azeitonabl !chuva  !tilt
!almir      !douglas    !skate


COMANDOS DE TEXTO:
!medusa	!gabi		!jabes		 !lixo
!adm		!miranha	!amongao !boraamong`)
	}

	if strings.HasPrefix(m.Content, "!passeio") {

		id := strings.Split(m.Content, " ")[1]
		id = strings.Split(strings.Split(id, "!")[1], ">")[0]
		for _, vchannel := range guild.Channels {
			if vchannel.Type == 2 {
				err = s.GuildMemberMove(guild.ID, id, &vchannel.ID)
				if err != nil {
					fmt.Println("Error on moving member: ", err)
				}
				time.Sleep(500 * time.Millisecond)
			}
		}

	}

	if strings.HasPrefix(m.Content, "!jabes") {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		n := r1.Intn(3)
		if n == 0 {
			s.ChannelMessageSend(m.ChannelID, "Alo")
		} else if n == 1 {
			s.ChannelMessageSend(m.ChannelID, "oi")
		} else {
			s.ChannelMessageSend(m.ChannelID, "Sim")
		}
	}

	if strings.HasPrefix(m.Content, "!gabi") {
		s.ChannelMessageSend(m.ChannelID, "SOFRO")
	}

	if strings.HasPrefix(m.Content, "!lixo") {
		s.ChannelMessageSend(m.ChannelID, "PHP É LIXO")
	}

	if strings.HasPrefix(m.Content, "!medusa") {
		s.ChannelMessageSend(m.ChannelID, "parabéns medusa pela resposta da 3")
	}

	if strings.HasPrefix(m.Content, "!adm") {
		s.ChannelMessageSend(m.ChannelID, "Você é corno?")
	}

	if strings.HasPrefix(m.Content, "!amongao") {
		s.ChannelMessageSend(m.ChannelID, "@everyone rinha de astronauta")
	}
	if strings.HasPrefix(m.Content, "!boraamong") {
		s.ChannelMessageSend(m.ChannelID, `⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠿⠛⠛⠿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿ 
⣿⣿⣿⣿⣿⣿⣿⡿⠋⢀⣠⣤⣶⣶⣶⣤⡀⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿ 
⣿⣿⣿⣿⣿⣿⡟⠄⢰⣿⣿⡿⠛⠉⠉⣉⣉⠄⠄⠘⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿ 
⣿⣿⣿⣿⣿⣿⠇⠄⢸⣿⣿⠁⠘⣿⣿⣿⣿⣿⣿⣦⡀⢹⣿⣿⣿⣿⣿⣿⣿⣿ 
⣿⣿⠿⠟⠛⠛⠄⠄⣼⣿⣿⡀⠄⠈⠛⠛⠛⠛⠛⠋⠁⢸⣿⣿⣿⣿⣿⣿⣿⣿ 
⣿⡇⢀⣶⣿⠂⠄⢸⣿⣿⣿⣧⣤⣄⣀⣀⣀⣀⣤⠄⢰⣿⣿⣿⣿⣿⣿⣿⣿⣿ 
⣿⡇⠄⣿⣿⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣄⣈⠙⠛⠿⣿⣿⣿⣿⣿ 
⣿⡇⠄⣿⣿⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣤⣈⠻⣿⣿⣿ 
⣿⡇⠄⠈⠉⠄⣠⣿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⠈⢻⣿ 
⣿⣿⠆⠄⠄⠄⢸⣿⡆⠙⢿⣿⣿⣿⣿⣿⣿⣿⡿⠿⠛⣿⣿⣿⡿⠟⠋⢀⣾⣿ 
⣿⣿⠄⠄⠄⠄⣼⣿⣿⣆⡀⠈⠉⠉⠉⠉⠉⠄⢀⣀⣀⠄⠁⠄⢀⣠⣾⣿⣿⣿ 
⣿⣿⡆⠄⠄⠄⢿⣿⣿⣿⣿⣶⣤⣤⣤⣤⣶⣶⣾⣿⣿⠁⠄⣿⣿⣿⣿⣿⣿⣿ 
⣿⣿⣿⡆⠄⠄⠈⠻⣿⣿⣿⣿⠟⠉⠉⠉⠁⣼⣿⣿⠏⠄⢠⣿⣿⣿⣿⣿⣿⣿ 
⣿⣿⣿⣿⡄⠄⠄⠄⠈⠛⠛⠁⢠⣄⠄⠄⠄⠄⠄⠄⠄⣰⣿⣿⣿⣿⣿⣿⣿⣿ 
⣿⣿⣿⣿⣷⡄⠄⠄⠄⠄⠄⠄⠛⠛⠂⠄⠄⠄⠄⠄⠴⢿⣿⣿⣿⣿⣿⣿⣿⣿`)
	}

	for slug := range bufferMap {
		command := "!" + slug
		if strings.HasPrefix(m.Content, command) {
			err = playSound(s, guild, m.Author.ID, slug)
			if err != nil {
				fmt.Println("Error playing sound: ", err)
			}
		}
	}

}

func ready(s *discordgo.Session, e *discordgo.Ready) {

	s.UpdateStatus(0, "!passeio")

}

func loadSound(filepath string) error {

	slug := strings.Split(strings.Split(filepath, "/")[1], ".")[0]
	buffer := make([][]byte, 0)

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return err
	}

	var opuslen int16

	for {

		err = binary.Read(file, binary.LittleEndian, &opuslen)

		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return err
			}
			break
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		buffer = append(buffer, InBuf)

	}

	bufferMap[slug] = buffer
	fmt.Printf("Audio file %s loaded successfully.\n", filepath)
	return nil

}

func playSound(s *discordgo.Session, guild *discordgo.Guild, authorID, slug string) (err error) {
	audioMutex.Lock()
	defer audioMutex.Unlock()

	for _, vs := range guild.VoiceStates {
		if vs.UserID == authorID {

			vc, err := s.ChannelVoiceJoin(guild.ID, vs.ChannelID, false, true)
			if err != nil {
				return err
			}

			vc.Speaking(true)

			for _, buff := range bufferMap[slug] {
				vc.OpusSend <- buff
			}

			vc.Speaking(false)

			time.Sleep(250 * time.Millisecond)

			vc.Disconnect()

		}
	}

	return nil
}

func main() {

	if token == "" {
		fmt.Println("No token provided.")
		return
	}

	err := filepath.Walk("dca_files/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		err = loadSound(path)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error on walk func: ", err)
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error on creating session: ", err)
		return
	}

	session.AddHandler(messageCreate)

	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates)

	err = session.Open()
	if err != nil {
		fmt.Println("Error opening discord session: ", err)
	}

	fmt.Println("Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	session.Close()

}
