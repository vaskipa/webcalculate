package repositories

import (
	"testing"
)

func TestInitTables(t *testing.T) {
	err := InitTables()
	if err != nil {
		t.Error(err)
	}
}

func TestCreateRecord(t *testing.T) {
	err := InitTables()
	if err != nil {
		t.Error(err)
	}
	_, err = AddRecord("2")
	if err != nil {
		t.Error(err)
	}
}

func TestGetRecords(t *testing.T) {
	err := InitTables()
	if err != nil {
		t.Error(err)
	}
	_, err = AddRecord("2")
	if err != nil {
		t.Error(err)
	}

	_, err = AddRecord("2")
	if err != nil {
		t.Error(err)
	}
	data, err := GetRecords()
	if err != nil {
		t.Error(err)
	}
	if len(data) == 0 {
		t.Error("No data")
	}
}

func TestGetRecord(t *testing.T) {
	err := InitTables()
	if err != nil {
		t.Error(err)
	}
	id, err := AddRecord("2")
	if err != nil {
		t.Error(err)
	}

	data, err := GetRecord(id)
	if err != nil {
		t.Error(err)
	}
	if data.Task != "2" {
		t.Error("Error ger")
	}
}
