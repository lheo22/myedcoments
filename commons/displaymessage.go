package commons

import (
	"CursoGolang/myedcoments/models"
	"encoding/json"
	"log"
	"net/http"
)

//DisplayMessage devuelve un mensaje al cliente
func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error al convertir el mensaje: %s", err)
	}
	w.WriteHeader(m.Code)
	w.Write(j)
}
