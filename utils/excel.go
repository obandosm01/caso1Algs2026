package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/xuri/excelize/v2"
)

func GuardarExcel(repBack ReporteEvaluacion, repGreedy ReporteEvaluacion) error {
	// 1. Crear el directorio "resultados" si no existe
	dirPath := "resultados"
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error al crear la carpeta resultados: %w", err)
	}

	// 2. Crear un nuevo libro de Excel
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	headersDetalle := []string{"N° Corrida", "Contraseña", "Consultas/Peticiones", "Tiempo (ms)", "Éxito"}

	// Función auxiliar interna para rellenar las filas de cualquier hoja de detalle
	escribirHoja := func(sheetName string, corridas []DetalleCorrida) error {
		index, err := f.NewSheet(sheetName)
		if err != nil {
			return err
		}

		// Encabezados
		for colIdx, header := range headersDetalle {
			cell, _ := excelize.CoordinatesToCellName(colIdx+1, 1)
			f.SetCellValue(sheetName, cell, header)
		}

		// Filas de datos
		for rowIdx, c := range corridas {
			numFila := rowIdx + 2 // Fila 1 reservada para los encabezados

			f.SetCellValue(sheetName, fmt.Sprintf("A%d", numFila), c.NumeroCorrida)
			f.SetCellValue(sheetName, fmt.Sprintf("B%d", numFila), c.Password)
			f.SetCellValue(sheetName, fmt.Sprintf("C%d", numFila), c.Consultas)
			f.SetCellValue(sheetName, fmt.Sprintf("D%d", numFila), float64(c.Duracion.Microseconds())/1000.0)
			f.SetCellValue(sheetName, fmt.Sprintf("E%d", numFila), c.Exito)
		}

		f.SetActiveSheet(index)
		return nil
	}

	// 3. Escribir Hoja 1: Backtracking
	if err := escribirHoja("Resultados Backtracking", repBack.Corridas); err != nil {
		return fmt.Errorf("error al escribir hoja de Backtracking: %w", err)
	}

	// 4. Escribir Hoja 2: Greedy
	if err := escribirHoja("Resultados Greedy", repGreedy.Corridas); err != nil {
		return fmt.Errorf("error al escribir hoja de Greedy: %w", err)
	}

	// 5. Escribir Hoja 3: Comparativa Algoritmos
	sheetComparativa := "Comparativa Algoritmos"
	indexComp, err := f.NewSheet(sheetComparativa)
	if err != nil {
		return fmt.Errorf("error al crear hoja de comparativa: %w", err)
	}

	// Función para calcular métricas acumuladas
	calcularMetricas := func(corridas []DetalleCorrida) (float64, float64, float64) {
		if len(corridas) == 0 {
			return 0, 0, 0
		}
		var totalConsultas int
		var totalTiempoMs float64
		var exitos int

		for _, c := range corridas {
			totalConsultas += c.Consultas
			totalTiempoMs += float64(c.Duracion.Microseconds()) / 1000.0
			if c.Exito {
				exitos++
			}
		}
		n := float64(len(corridas))
		return float64(totalConsultas) / n, totalTiempoMs / n, (float64(exitos) / n) * 100.0
	}

	promConsultasBack, promTiempoBack, exitoBack := calcularMetricas(repBack.Corridas)
	promConsultasGreedy, promTiempoGreedy, exitoGreedy := calcularMetricas(repGreedy.Corridas)

	headersComp := []string{"Métrica / Indicador", "Backtracking", "Greedy", "Diferencia"}
	for colIdx, header := range headersComp {
		cell, _ := excelize.CoordinatesToCellName(colIdx+1, 1)
		f.SetCellValue(sheetComparativa, cell, header)
	}

	filasComp := [][]interface{}{
		{"Total Corridas Ejecutadas", repBack.TotalCorridas, repGreedy.TotalCorridas, fmt.Sprintf("%+d corridas", repGreedy.TotalCorridas-repBack.TotalCorridas)},
		{"Promedio Consultas TCP", fmt.Sprintf("%.2f", promConsultasBack), fmt.Sprintf("%.2f", promConsultasGreedy), fmt.Sprintf("%.2f consultas", promConsultasGreedy-promConsultasBack)},
		{"Tiempo Promedio por Corrida (ms)", fmt.Sprintf("%.3f ms", promTiempoBack), fmt.Sprintf("%.3f ms", promTiempoGreedy), fmt.Sprintf("%.3f ms", promTiempoGreedy-promTiempoBack)},
		{"Tasa de Éxito (%)", fmt.Sprintf("%.1f%%", exitoBack), fmt.Sprintf("%.1f%%", exitoGreedy), "Tasa comparada"},
	}

	for rowIdx, fila := range filasComp {
		numFila := rowIdx + 2
		f.SetCellValue(sheetComparativa, fmt.Sprintf("A%d", numFila), fila[0])
		f.SetCellValue(sheetComparativa, fmt.Sprintf("B%d", numFila), fila[1])
		f.SetCellValue(sheetComparativa, fmt.Sprintf("C%d", numFila), fila[2])
		f.SetCellValue(sheetComparativa, fmt.Sprintf("D%d", numFila), fila[3])
	}

	f.SetActiveSheet(indexComp)

	// Eliminar la hoja vacía por defecto que genera excelize ("Sheet1")
	_ = f.DeleteSheet("Sheet1")

	// 6. Nombre único con Timestamp para no sobrescribir ejecuciones anteriores
	nombreArchivo := fmt.Sprintf("reporte_comparativo_%s.xlsx", time.Now().Format("20060102_150405"))
	rutaFinal := filepath.Join(dirPath, nombreArchivo)

	// 7. Guardar el archivo en la subcarpeta "resultados"
	if err := f.SaveAs(rutaFinal); err != nil {
		return fmt.Errorf("error al guardar el archivo Excel: %w", err)
	}

	fmt.Printf("¡Reporte comparativo generado con éxito en: %s!\n", rutaFinal)
	return nil
}
