package main

import (

	"testing"
)

func Benchmark_seqArchiver_many(b *testing.B) {
	for i:=0; i< b.N; i++{
		seqArchivatore([]string{
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
		})
	}
}

func Benchmark_seqArchiver_one(b *testing.B) {
	for i:=0; i< b.N; i++{
		seqArchivatore([]string{
			"1.txt",
		})
	}
}
func Benchmark_conArchiver_many(b *testing.B) {
	for i:=0; i< b.N; i++{
		seqArchivatore([]string{
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
			"1.txt",
			"2.txt",
		})
	}
}

func Benchmark_conArchiver_one(b *testing.B) {
	for i:=0; i< b.N; i++{
		seqArchivatore([]string{
			"1.txt",
		})
	}
}