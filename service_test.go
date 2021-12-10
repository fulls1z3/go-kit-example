package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type args struct {
		model *model
	}

	testCases := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "should add items",
			args: args{
				model: &model{
					Name: "test item",
				},
			},
			err: nil,
		},
	}

	s := NewService()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := s.add(tc.args.model.Name)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		id int
	}

	testCases := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "should remove items",
			args: args{
				id: 1,
			},
			err: nil,
		},
	}

	s := NewService()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := s.remove(tc.args.id)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestGetAll(t *testing.T) {
	testCases := []struct {
		name     string
		expected []model
		err      error
	}{
		{
			name:     "should get all items",
			expected: []model{},
			err:      nil,
		},
	}

	s := NewService()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := s.getAll()

			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.err, err)
		})
	}
}
