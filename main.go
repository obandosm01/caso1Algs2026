package main

import (
	"caso1/algs"
	"caso1/utils"
	"fmt"
	"time"
)

func evaluarGreedy(tiempoSegundos int) utils.ReporteEvaluacion {
	duracionLimite := time.Duration(tiempoSegundos) * time.Second
	tiempoInicioGlobal := time.Now()
	var listaCorridas []utils.DetalleCorrida
	totalCorridas := 0

	// Ejecuta el algoritmo Greedy una y otra vez mientras no se agote el tiempo limite
	for time.Since(tiempoInicioGlobal) < duracionLimite {
		totalCorridas++
		inicioRun := time.Now()

		// Ejecuta la funcion de Greedy
		exito, consultas, passHallada := algs.Greedy()

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

func main() {
	tiempoEvaluacionSegundos := 1

	fmt.Println("Iniciando evaluación secuencial de algoritmos...")

	fmt.Printf("\nEvaluando Backtracking durante %d segundos...\n", tiempoEvaluacionSegundos)
	repBack := evaluarBacktrack(tiempoEvaluacionSegundos)
	fmt.Println(" -> Backtracking finalizado.")

	fmt.Printf("\nEvaluando Greedy durante %d segundos...\n", tiempoEvaluacionSegundos)
	repGreedy := evaluarGreedy(tiempoEvaluacionSegundos)
	fmt.Println(" -> Greedy finalizado.")

	err := utils.GuardarExcel(repBack, repGreedy)
	if err != nil {
		fmt.Printf("Error al guardar el archivo Excel: %s\n", err.Error())
	}
}
