package todo_test

import (
	"errors"
	"os"
	"plan/todo"
	"reflect"
	"testing"
)

func fakemarshal(v interface{}) ([]byte, error) {
    return []byte{}, errors.New("Marshalling failed")
}

func restoremarshal(replace func(v interface{}) ([]byte, error)) {
    todo.JsonMarshal = replace
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestSaveItems(t *testing.T) {
	var tests = []struct {
		description string
		items       []todo.Item
		result      string
		wantError   bool
	}{
		{
			description: "Save items successfully",
			items:       []todo.Item{{Text: "Task1", Priority: 0}, {Text: "Task2", Priority: 0}, {Text: "Task3", Priority: 0}},
			result:      "[{\"Text\":\"Task1\",\"Priority\":0},{\"Text\":\"Task2\",\"Priority\":0},{\"Text\":\"Task3\",\"Priority\":0}]",
			wantError:   false,
		},
		{
			description: "Save items returns error",
			items:       []todo.Item{{Text: "", Priority: 0}},
			result:      "",
			wantError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			file, err1 := os.CreateTemp("", "test")
			check(err1)
			defer os.Remove(file.Name())

			if tt.wantError {
				storedMarshal := todo.JsonMarshal
    			todo.JsonMarshal = fakemarshal
				err := todo.SaveItems(file.Name(), tt.items)
				if err == nil {
					t.Errorf("Got nil, wanted error")
				}
				defer restoremarshal(storedMarshal)
			} else {
				todo.SaveItems(file.Name(), tt.items)
				actualBytes, err2 := os.ReadFile(file.Name())
				check(err2)
				actual := string(actualBytes)
				if actual != tt.result {
					t.Errorf("Got %s, wanted %s", actual, tt.result)
				}
			}
			
		})
	}
}

func TestReadItems(t *testing.T) {
	var tests = []struct {
		description string
		fileContent string
		result      []todo.Item
		wantError   bool
	} {
		{
			description: "Read items successfully",
			fileContent: "[{\"Text\":\"task1\",\"Priority\":0},{\"Text\":\"task2\",\"Priority\":1}]",
			result: []todo.Item{{Text: "task1", Priority: 0}, {Text: "task2", Priority: 1}},
			wantError: false,
		},
		{
			description: "Read items returns error",
			fileContent: "[{\"Text\":\"task1\",\"Priority\":0},",
			result: []todo.Item{},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			file, err1 := os.CreateTemp("", "test", )
			check(err1)
			defer os.Remove(file.Name())
			_, err1 = file.Write([]byte(tt.fileContent))
			check(err1)

			if tt.wantError {
				_, err2 := todo.ReadItems(file.Name())
				if err2 == nil {
					t.Errorf("Got nil, wanted error")
				}
			} else {
				items, err2 := todo.ReadItems(file.Name())
				check(err2)
				if !reflect.DeepEqual(items, tt.result) {
					t.Errorf("Got %v, wanted %v", items, tt.result)
				}
			}
		})
	}
}


func TestSetPriority(t *testing.T) {
	var tests = []struct {
		description string
		item todo.Item
		priority int
		result int
	} {
		{
			description: "Set priority successfully",
			item: todo.Item{Text: "task"},
			priority: 1,
			result: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			tt.item.SetPriority(tt.priority)
			if tt.item.Priority != tt.result {
				t.Errorf("Got %v, wanted %v", tt.item.Priority, tt.result)
			}
		})
	}
}