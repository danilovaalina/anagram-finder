package main

import (
	"testing"
)

// equalStringSlices предполагает, что оба среза отсортированы
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  map[string][]string
	}{
		{
			name:  "базовый пример",
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"},
			want: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:  "пустой список",
			words: []string{},
			want:  map[string][]string{},
		},
		{
			name:  "нет анаграмм",
			words: []string{"стол", "дом", "кот"},
			want:  map[string][]string{},
		},
		{
			name:  "одна пара анаграмм с разным регистром",
			words: []string{"Арка", "кАрА"},
			want: map[string][]string{
				"арка": {"арка", "кара"},
			},
		},
		{
			name:  "три одинаковых слова (должны быть в одной группе)",
			words: []string{"мир", "МИР", "Мир"},
			want: map[string][]string{
				"мир": {"мир", "мир", "мир"},
			},
		},
		{
			name:  "слово и его анаграмма, но одна — без анаграмм",
			words: []string{"липа", "пила", "кот"},
			want: map[string][]string{
				"липа": {"липа", "пила"},
			},
		},
		{
			name:  "все слова - анаграммы",
			words: []string{"сон", "нос", "сно"},
			want: map[string][]string{
				"сон": {"нос", "сно", "сон"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := FindAnagrams(tt.words)

			// Проверим количество групп
			if len(got) != len(tt.want) {
				t.Fatalf("количество групп: получено %d, ожидалось %d", len(got), len(tt.want))
			}

			// Проверим каждую группу
			for key, wantWords := range tt.want {
				gotWords, exists := got[key]
				if !exists {
					t.Errorf("ожидался ключ %q, но его нет в результате", key)
					continue
				}

				if !equalStringSlices(wantWords, gotWords) {
					t.Errorf("для ключа %q: получено %q, ожидалось %q", key, gotWords, wantWords)
				}
			}

		})
	}
}
