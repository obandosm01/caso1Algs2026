package utils

import "time"

type ReporteEvaluacion struct {
	TiempoLimite  time.Duration
	TiempoTotal   time.Duration
	TotalCorridas int
	Corridas      []DetalleCorrida
}

type RespuestaServer struct {
	Status         string `json:"status"`
	MatchedLetters uint8  `json:"matchedLetters,omitempty"`
	Distance       uint8  `json:"distance,omitempty"`
	Score          uint16 `json:"score,omitempty"`
}

type DetalleCorrida struct {
	NumeroCorrida int
	Password      string
	Consultas     int
	Duracion      time.Duration
	Exito         bool
}
