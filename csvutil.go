package main

func KillEmptyRows(rows []string) []string {
	var filteredRows []string
	for _, row := range rows {
		if len(row) <= 0 {
			continue
		}
		filteredRows = append(filteredRows, row)
	}
	return filteredRows
}
