package todo_test

import (
	"context"
	"github.com/imjowend/mastering-go-with-goland/internal/db"
	"github.com/imjowend/mastering-go-with-goland/internal/todo"
	"reflect"
	"testing"
)

// tips: option + enter to implement interface

type MockDB struct {
	items []db.Item
}

func (m *MockDB) InsertItem(ctx context.Context, item db.Item) error {
	m.items = append(m.items, item)
	return nil
}

func (m *MockDB) GetAllItems(ctx context.Context) ([]db.Item, error) {
	return m.items, nil
}

func TestService_Search(t *testing.T) {
	tests := []struct {
		name           string
		toDosToAdd     []string
		query          string
		expectedResult []string
	}{
		// TODO: Add test cases.
		// Press control + space to create new Test Cases
		{
			name:           "given a todo of shop and a search of sh, i shoulg get shop back",
			toDosToAdd:     []string{"shop"},
			query:          "sh",
			expectedResult: []string{"shop"},
		},
		{
			name:           "still returns shop, even if the case doesn't match",
			toDosToAdd:     []string{"Shopping"},
			query:          "sh",
			expectedResult: []string{"Shopping"},
		},
		{
			name:           "spaces",
			toDosToAdd:     []string{"go Shopping"},
			query:          "go",
			expectedResult: []string{"go Shopping"},
		},
		{name: "space at start of word", toDosToAdd: []string{" Space at beginning"}, query: "space", expectedResult: []string{" Space at beginning"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDB{}
			svc := todo.NewService(m)
			for _, toAdd := range tt.toDosToAdd {
				err := svc.Add(toAdd)
				if err != nil {
					t.Errorf("todo.Service.Search() error = %v", err)
				}
			}
			got, err := svc.Search(tt.query)
			if err != nil {
				t.Errorf("todo.Service.Search() error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("Search() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}
