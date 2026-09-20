package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/xuri/excelize/v2"
)

func GuardarExcel(rep ReporteEvaluacion) error {
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

	sheetName := "Resultados Backtracking"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return err
	}
	f.SetActiveSheet(index)
	_ = f.DeleteSheet("Sheet1") // Eliminar la hoja por defecto

	// 3. Escribir Encabezados de la Tabla
	headers := []string{"N° Corrida", "Contraseña", "Consultas/Peticiones", "Tiempo (ms)", "Éxito"}
	for colIdx, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(colIdx+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	// 4. Escribir las Filas con la información de cada corrida
	for rowIdx, c := range rep.Corridas {
		numFila := rowIdx + 2 // Fila 1 reservada para los encabezados

		f.SetCellValue(sheetName, fmt.Sprintf("A%d", numFila), c.NumeroCorrida)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", numFila), c.Password)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", numFila), c.Consultas)
		// Convertimos la duración a milisegundos flotantes para facilitar fórmulas/gráficos en Excel
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", numFila), float64(c.Duracion.Microseconds())/1000.0)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", numFila), c.Exito)
	}

	// 5. Nombre único con Timestamp para no sobrescribir ejecuciones anteriores
	nombreArchivo := fmt.Sprintf("reporte_backtracking_%s.xlsx", time.Now().Format("20060102_150405"))
	rutaFinal := filepath.Join(dirPath, nombreArchivo)

	// 6. Guardar el archivo en la subcarpeta "resultados"
	if err := f.SaveAs(rutaFinal); err != nil {
		return fmt.Errorf("error al guardar el archivo Excel: %w", err)
	}

	fmt.Printf("¡Reporte generado con éxito en: %s!\n", rutaFinal)
	return nil
}
