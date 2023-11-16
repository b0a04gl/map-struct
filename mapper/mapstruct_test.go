package mapper

import (
	"reflect"
	"testing"
)

func TestMapToStruct(t *testing.T) {
	tests := []struct {
		name     string
		data     map[string]interface{}
		target   interface{}
		expected interface{}
		wantErr  bool
	}{
		{
			name: "ValidMapping",
			data: map[string]interface{}{
				"Name":    "John Doe",
				"Age":     30,
				"Address": "123 Main St",
			},
			target: &Person{},
			expected: &Person{
				Name:    "John Doe",
				Age:     30,
				Address: "123 Main St",
			},
			wantErr: false,
		},
		{
			name: "InvalidTarget",
			data: map[string]interface{}{
				"Name":    "Jane Doe",
				"Age":     25,
				"Address": "456 Second St",
			},
			target:   nil,
			expected: nil,
			wantErr:  true,
		},
		{
			name: "TypeMismatch",
			data: map[string]interface{}{
				"Name":    "Bob",
				"Age":     "not an int", 
				"Address": "789 Third St",
			},
			target:   &Person{},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MapToStruct(tt.data, tt.target)

			if (err != nil) != tt.wantErr {
				t.Errorf("MapToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(tt.target, tt.expected) {
					t.Errorf("MapToStruct() got = %v, want %v", tt.target, tt.expected)
				}
			}
		})
	}
}


