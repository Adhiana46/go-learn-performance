package main

import (
	"testing"
)

type testStruct struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
	IsAlive bool    `json:"is_alive"`
}

var testData = testStruct{
	ID:      1,
	Name:    "Adhiana Mastur",
	Age:     25,
	Address: "Reprehenderit minim dolor officia dolore ea ea magna do quis nostrud officia amet.",
	Balance: 25000.55,
	IsAlive: true,
}

var testDataBytes = []byte{
	123, 34, 105, 100, 34, 58, 49, 44, 34, 110, 97, 109, 101, 34, 58, 34, 65, 100, 104, 105, 97, 110, 97, 32, 77, 97, 115, 116, 117, 114, 34, 44, 34, 97, 103, 101, 34, 58, 50, 53, 44, 34, 97, 100, 100, 114, 101, 115, 115, 34, 58, 34, 82, 101, 112, 114, 101, 104, 101, 110, 100, 101, 114, 105, 116, 32, 109, 105, 110, 105, 109, 32, 100, 111, 108, 111, 114, 32, 111, 102, 102, 105, 99, 105, 97, 32, 100, 111, 108, 111, 114, 101, 32, 101, 97, 32, 101, 97, 32, 109, 97, 103, 110, 97, 32, 100, 111, 32, 113, 117, 105, 115, 32, 110, 111, 115, 116, 114, 117, 100, 32, 111, 102, 102, 105, 99, 105, 97, 32, 97, 109, 101, 116, 46, 34, 44, 34, 98, 97, 108, 97, 110, 99, 101, 34, 58, 50, 53, 48, 48, 48, 46, 53, 53, 44, 34, 105, 115, 95, 97, 108, 105, 118, 101, 34, 58, 116, 114, 117, 101, 125,
}

func BenchmarkStdEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdEncode(testData)
	}
}

func BenchmarkJsoniterEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsoniterEncode(testData)
	}
}

func BenchmarkStdDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := testStruct{}
		stdDecode(testDataBytes, &t)
	}
}

func BenchmarkJsoniterDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := testStruct{}
		jsoniterDecode(testDataBytes, &t)
	}
}
