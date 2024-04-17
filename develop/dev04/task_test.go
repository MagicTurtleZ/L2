package main

import (
	"reflect"
	"testing"
)

func TestMakeAnagram(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "ehllo"},
		{"world", "dlorw"},
		{"Listen", "eilnst"},
		{"silent", "eilnst"},
	}

	for _, test := range tests {
		result := makeAnagram(test.input)
		if result != test.expected {
			t.Errorf("makeAnagram(%s) returned %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestFormatMap(t *testing.T) {
	tests := []struct {
		input    map[string][]string
		expected map[string][]string
	}{
		{
			map[string][]string{"a": {"world", "hello"}, "b": {}, "c": {"example"}},
			map[string][]string{"a": {"hello", "world"}, "c": {"example"}},
		},
		{
			map[string][]string{"x": {}, "y": {"test"}},
			map[string][]string{"y": {"test"}},
		},
		{
			map[string][]string{},
			map[string][]string{},
		},
	}

	for _, test := range tests {
		formatMap(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("formatMap(%v) produced %v, expected %v", test.input, test.input, test.expected)
		}
	}
}

func TestFindSets(t *testing.T) {
	tests := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			[]string{"аскет", "секта", "сетка", "стека", "тесак", "теска"},
			map[string][]string{"аскет": {"секта", "сетка", "стека", "тесак", "теска"}},
		},
		{
			[]string{"Носок", "Коса", "Бокал", "Колба", "Мошка", "Кошка", "Камыш", "Мышка", "Крик", "Икар", "Порка", "Корка", "Тополь", 
			"Полоть", "Трал", "Аборт", "Табло", "Боль", "Обь", "Гном", "Мог", "Догма", "Мода", "Гаер", 
			"Герой", "Лазер", "Зарево", "Вокал", "Овал", "Ловкач", "Волька", "Накал", "Канал", "Луна", "Нуль", "Насос", "Сосна",},
			map[string][]string{"бокал": {"колба"}, "камыш": {"мышка"}, "накал": {"канал"}, "насос": {"сосна"}, "тополь": {"полоть"}},
		},
	}

	for _, test := range tests {
		result := findSets(test.input)
		
		if !reflect.DeepEqual(*result, test.expected) {
			t.Errorf("Incorrect result. Expect %s, got %s", test.expected, *result)
		}
	}
}