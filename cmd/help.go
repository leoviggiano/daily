package cmd

import (
	"fmt"
)

func Help() Command {
	return &help{}
}

type help struct {
}

func (h *help) Name() []string {
	return []string{"help"}
}

func (h *help) Help() string {
	return `
	Ao adicionar o sufixo -h ou --help, você pode ver a ajuda de um comando específico.

	Exemplo:
		daily add -h
		daily delete -h
		daily show -h
		daily history -h
		daily help -h

	Comandos disponíveis:
	init - Inicia uma nova daily
	add {a} - Adiciona um item à daily
	show {s} - Mostra a daily atual
	delete {d} - Deleta um item da daily
	history {h} - Mostra o histórico de dailys
	help - Mostra a ajuda

	Exemplo:
		daily add "Fazer o exercício"
		daily delete 1
		daily show
		daily history
		daily help
	`
}

func (h *help) Exec(args ...string) error {
	fmt.Println(h.Help())
	return nil
}
