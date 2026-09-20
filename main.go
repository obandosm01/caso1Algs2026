package main

import (
	"caso1/algs"
	"caso1/utils"
	"fmt"
	"time"
)

func evaluarBacktrack(tiempoSegundos int) utils.ReporteEvaluacion {
	duracionLimite := time.Duration(tiempoSegundos) * time.Second
	tiempoInicioGlobal := time.Now()
	var listaCorridas []utils.DetalleCorrida
	totalCorridas := 0

	// Ejecuta el algoritmo una y otra vez mientras no se agote el tiempo limite
	for time.Since(tiempoInicioGlobal) < duracionLimite {
		totalCorridas++
		inicioRun := time.Now()

		// Ejecuta tu funcion de backtracking modificada
		exito, consultas, passHallada := algs.Backtracking()

		duracionRun := time.Since(inicioRun)

		// Guarda el detalle especifico de esta ejecucion
		corrida := utils.DetalleCorrida{
			NumeroCorrida: totalCorridas,
			Password:      passHallada,
			Consultas:     consultas,
			Duracion:      duracionRun,
			Exito:         exito,
		}

		listaCorridas = append(listaCorridas, corrida)
	}

	reporte := utils.ReporteEvaluacion{
		TiempoLimite:  duracionLimite,
		TiempoTotal:   time.Since(tiempoInicioGlobal),
		TotalCorridas: totalCorridas,
		Corridas:      listaCorridas,
	}
	return reporte
}

func imprimirReporte(rep utils.ReporteEvaluacion) {
	fmt.Printf("\n=== DETALLE DE CORRIDAS (Tiempo Límite: %v) ===\n", rep.TiempoLimite)
	fmt.Printf("Total ejecuciones: %d | Tiempo real transcurrido: %v\n\n", rep.TotalCorridas, rep.TiempoTotal)

	for _, c := range rep.Corridas {
		fmt.Printf("Run #%-3d | Password: %-6s | Consultas: %-3d | Tiempo: %-10v | Éxito: %t\n",
			c.NumeroCorrida,
			fmt.Sprintf(`"%s"`, c.Password),
			c.Consultas,
			c.Duracion,
			c.Exito,
		)
	}
	fmt.Println("================================================\n")
}

func main() {
	reporte := evaluarBacktrack(1)
	imprimirReporte(reporte)
	err := utils.GuardarExcel(reporte)
	if err != nil {
		println(err.Error())
	}
}
