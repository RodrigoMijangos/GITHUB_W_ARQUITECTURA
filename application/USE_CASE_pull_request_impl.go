package application

import (
	"encoding/json"
	domain "github_wb/domain/value_objects"
	"log"
	"fmt"
)

func ProcessPullRequest(payload []byte) int {
	var eventPayload domain.PullRequestEventPayload

	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		return 500
	}

	if eventPayload.Action == "closed" {
		base := eventPayload.PullRequest.Base.Ref
		branch := eventPayload.PullRequest.Head.Ref
		user := eventPayload.PullRequest.User.Login
		pRID := eventPayload.PullRequest.ID

		log.Printf("Pull Request Recibido:\nID:%d\nBase:%s\nHead:%s\nUser:%s", pRID, base, branch, user)
	} else {
		log.Printf("Pull Request Action no es Closed: %s", eventPayload.Action)
	}


	log.Printf("Evento Pull Request recibido: Acción=%s, Título='%s', Base='%s', Repositorio='%s', Cambio='%s', Rama='%s'",
		eventPayload.Action,
		eventPayload.PullRequest.Title,
		eventPayload.PullRequest.Base.Ref,
		eventPayload.Repository.FullName,
		eventPayload.PullRequest.Title, // Aquí puedes incluir la descripción del cambio
		eventPayload.PullRequest.Head.Ref) // Rama del cambio

	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		log.Printf("Error al procesar el payload: %v", err)
		return 500
	}


	switch eventPayload.Action {
	case "opened":
		fmt.Printf("\nNuevo Pull Request creado por %s en %s\n", eventPayload.PullRequest.User.Login, eventPayload.Repository.FullName)
	case "synchronize":
		fmt.Printf("\nPull Request actualizado: %s\n", eventPayload.PullRequest.Title)
	case "closed":
		fmt.Printf("\nPull Request cerrado: %s\n", eventPayload.PullRequest.Title)
	default:
		fmt.Printf("\n¡Acción del PR detectada: %s\n", eventPayload.Action)
	}

	return 200
}
