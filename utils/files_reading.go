package utils

import (
  "encoding/csv"
  "os"
  "fmt"
)

func ReadCsvFile(filePath string) [][]string {
  f, err := os.Open(filePath)
  if err != nil {
    fmt.Println("Unable to read csv file: " + filePath, err)
  }
  defer f.Close()

  csvReader := csv.NewReader(f)
  records, err := csvReader.ReadAll()
  if err != nil {
    fmt.Println("Unable to parse csv file: " + filePath, err)
  }

  return records
}
