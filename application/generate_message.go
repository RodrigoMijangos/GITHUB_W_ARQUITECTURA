package application

import "fmt"

func generateDiscordMessage(base, head, html_url, user, repository_full_name string) string {
	return fmt.Sprintf("Nuevo pull request a la rama %s en el repositorio %s\nRama Inicial: %s\nUsuario: %s\n\nURL del Pull Request: %s", base, repository_full_name, head, user, html_url)
}
