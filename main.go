package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Substitua pelo seu token do bot
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
	}

	// Lê o token da variável de ambiente
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		fmt.Println("Erro: DISCORD_TOKEN não está definido")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Erro ao criar sessão Discord:", err)
		return
	}

	// Habilita intents para mensagens
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	// Adiciona handler para mensagens
	dg.AddHandler(messageCreate)

	// Abre a conexão com o Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Erro ao conectar ao Discord:", err)
		return
	}

	fmt.Println("Bot está rodando! Pressione CTRL+C para sair.")

	// Mantém o bot rodando até receber um sinal de interrupção
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Fecha a conexão ao encerrar
	dg.Close()
}

// Função chamada quando uma mensagem é criada
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignora mensagens do próprio bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Responde a mensagens com "!ping"
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}
