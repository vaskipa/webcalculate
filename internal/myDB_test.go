package internal

import (
	"testing"
)

func TestCreateTable(t *testing.T) {
	err := createTable()
	if err != nil {
		t.Error(err)
	}
}

func TestCreateRecord(t *testing.T) {
	err := createTable()
	if err != nil {
		t.Error(err)
	}
	_, err = addRecord("2")
	if err != nil {
		t.Error(err)
	}
}

func TestGetRecords(t *testing.T) {
	err := createTable()
	if err != nil {
		t.Error(err)
	}
	_, err = addRecord("2")
	if err != nil {
		t.Error(err)
	}

	_, err = addRecord("2")
	if err != nil {
		t.Error(err)
	}
	data, err := getRecords()
	if err != nil {
		t.Error(err)
	}
	if len(data) == 0 {
		t.Error("No data")
	}
}

func TestGetRecord(t *testing.T) {
	err := createTable()
	if err != nil {
		t.Error(err)
	}
	id, err := addRecord("2")
	if err != nil {
		t.Error(err)
	}

	data, err := getRecord(id)
	if err != nil {
		t.Error(err)
	}
	if data.Task != "2" {
		t.Error("Error ger")
	}
}
