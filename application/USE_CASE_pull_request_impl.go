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

	log.Printf("Evento Pull Request recibido: Acción=%s, Título='%s', Base='%s', Repositorio='%s'",
		eventPayload.Action,
		eventPayload.PullRequest.Title,
		eventPayload.PullRequest.Base.Ref,
		eventPayload.Repository.FullName)

	


	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		log.Printf("Error al procesar el payload: %v", err)
		return 500
	}

	
	log.Printf("Evento Pull Request recibido: Acción=%s, Título='%s', Base='%s', Repositorio='%s'",
		eventPayload.Action,
		eventPayload.PullRequest.Title,
		eventPayload.PullRequest.Base.Ref,
		eventPayload.Repository.FullName)

	switch eventPayload.Action {
	case "opened":
		fmt.Printf("\n Nuevo Pull Request creado por %s en %s\n", eventPayload.PullRequest.User.Login, eventPayload.Repository.FullName)
	case "synchronize":
		fmt.Printf("\n Pull Request actualizado: %s\n", eventPayload.PullRequest.Title)
	case "closed":
		fmt.Printf("\n Pull Request cerrado: %s\n", eventPayload.PullRequest.Title)
	default:
		fmt.Printf("\n  !Acción del PR detectada: %s\n", eventPayload.Action)
	}


	return 200
}
