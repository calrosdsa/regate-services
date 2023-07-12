package repository

type SalaPayload struct {
	Id          int    `json:"id"`
	Titulo      string `json:"titulo"`
	GrupoId     int    `json:"grupo_id"`
}
