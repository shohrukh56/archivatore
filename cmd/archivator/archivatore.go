package main

import (
	zip2 "archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	files := os.Args[1:]
	seqArchivatore(files)
	conArchivatore(files)
}

func seqArchivatore(files []string) {
	for _, file := range files {
		fileName := "cmd/archivator/files/" + "fileseqarchivatore/" + file + ".zip"

		archiveInZip(fileName, file)
	}
}

func conArchivatore(files []string) {
	waitGroup := sync.WaitGroup{}
	mutex := sync.Mutex{}
	for _, file := range files {
		fileName := "cmd/archivator/files/"+ "fileconarchivatore/"+ file + ".zip"
		waitGroup.Add(1)
		go func(wg *sync.WaitGroup, flN, fl string, mu *sync.Mutex) {
			defer func() {
				mu.Lock()
				waitGroup.Done()
				mu.Unlock()
			}()
			archiveInZip( flN, fl)
		}(&waitGroup, fileName, file, &mutex)
	}
	waitGroup.Wait()
	fmt.Print("done")
}

func archiveInZip(fileName, file string){

	newZipFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Can't create file: %v", err)
	}
	defer func() {
		err = newZipFile.Close()
		if err != nil {
			log.Fatalf("Can't close zip: %v", err)
		}
	}()
	zipWriter := zip2.NewWriter(newZipFile)
	defer func() {
		err = zipWriter.Close()
		if err != nil {
			log.Fatalf("Can't close zip writer: %v", err)
		}
	}()
	zipfile, err := os.Open("cmd/archivator/files/" + file)
	if err != nil {
		log.Fatalf("Can't open file: %v", err)
	}
	defer func() {
		err = zipfile.Close()
		if err != nil {
			log.Fatalf("Can't close zip: %v", err)
		}
	}()
	info, err := zipfile.Stat()
	if err != nil {
		log.Fatalf("Can't stat: %v", err)
	}
	header, err := zip2.FileInfoHeader(info)
	if err != nil {
		log.Fatalf("Can't info file: %v", err)
	}
	header.Name = file
	header.Method = zip2.Deflate
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		log.Fatalf("Can't creat header: %v", err)
	}
	if _, err = io.Copy(writer, zipfile); err != nil {
		log.Fatalf("Can't copy: %v", err)
	}
	//fmt.Println("Zipped File: " + fileName)
}