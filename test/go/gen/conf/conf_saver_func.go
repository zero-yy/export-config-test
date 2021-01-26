// Code generated by export-config. DO NOT EDIT.
package conf

import (
	"fmt"
	"github.com/zero-yy/export-config/sheet"
	"path"
)

func initSaveFunc() {
	DataSaveFunc["test"] = func(s *sheet.Sheet, outputPath string) {
		m := &Test{
			Records: make(map[int32]*Test_Record),
			Crc32:   0,
		}

		// load from s to m
		for rowIndex := sheet.KDataRowStart; rowIndex < s.Xs.MaxRow; rowIndex++ {
			rowData := s.Xs.Row(rowIndex)

			if s.IdColIndex >= len(rowData.Cells) {
				break
			}

			// Id data is nil then skip this row.
			if len(rowData.Cells[s.IdColIndex].Value) == 0 {
				continue
			}

			record := &Test_Record{}

			s.ParseRecordData(rowIndex, rowData, record)

			if _, ok := m.Records[record.Id]; ok {
				panicIfErr(fmt.Errorf("repeat id(%v) in sheet(%v) cell(%v)",
					record.Id, s.Name, sheet.CellName(int32(rowIndex), int32(s.IdColIndex))))
			}
			m.Records[record.Id] = record
		}

		fullName := path.Join(outputPath, "test.bytes")
		saveBytes(m, fullName)

		fullName = path.Join(outputPath, "test.json")
		saveJson(m, fullName)
	}

	DataSaveFunc["test2"] = func(s *sheet.Sheet, outputPath string) {
		m := &Test2{
			Records: make(map[string]*Test2_Record),
			Crc32:   0,
		}

		// load from s to m
		for rowIndex := sheet.KDataRowStart; rowIndex < s.Xs.MaxRow; rowIndex++ {
			rowData := s.Xs.Row(rowIndex)

			if s.IdColIndex >= len(rowData.Cells) {
				break
			}

			// Id data is nil then skip this row.
			if len(rowData.Cells[s.IdColIndex].Value) == 0 {
				continue
			}

			record := &Test2_Record{}

			s.ParseRecordData(rowIndex, rowData, record)

			if _, ok := m.Records[record.Id]; ok {
				panicIfErr(fmt.Errorf("repeat id(%v) in sheet(%v) cell(%v)",
					record.Id, s.Name, sheet.CellName(int32(rowIndex), int32(s.IdColIndex))))
			}
			m.Records[record.Id] = record
		}

		fullName := path.Join(outputPath, "test2.bytes")
		saveBytes(m, fullName)

		fullName = path.Join(outputPath, "test2.json")
		saveJson(m, fullName)
	}

	DataSaveFunc["test3"] = func(s *sheet.Sheet, outputPath string) {
		m := &Test3{
			Records: make(map[int32]*Test3_Record),
			Crc32:   0,
		}

		// load from s to m
		for rowIndex := sheet.KDataRowStart; rowIndex < s.Xs.MaxRow; rowIndex++ {
			rowData := s.Xs.Row(rowIndex)

			if s.IdColIndex >= len(rowData.Cells) {
				break
			}

			// Id data is nil then skip this row.
			if len(rowData.Cells[s.IdColIndex].Value) == 0 {
				continue
			}

			record := &Test3_Record{}

			s.ParseRecordData(rowIndex, rowData, record)

			if _, ok := m.Records[record.Id]; ok {
				panicIfErr(fmt.Errorf("repeat id(%v) in sheet(%v) cell(%v)",
					record.Id, s.Name, sheet.CellName(int32(rowIndex), int32(s.IdColIndex))))
			}
			m.Records[record.Id] = record
		}

		fullName := path.Join(outputPath, "test3.bytes")
		saveBytes(m, fullName)

		fullName = path.Join(outputPath, "test3.json")
		saveJson(m, fullName)
	}

}