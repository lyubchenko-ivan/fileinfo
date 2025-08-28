package main

import (
	"os"
	"flag"
	"fmt"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [filename1] [filename2] ...\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Display file size and modification time for each file\n")

		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: you should pass at least one filename\n")
		flag.Usage()
		os.Exit(1)
	}

	exitCode := 0

	for _, filename := range args {
		if err := printFileInfo(filename); err != nil {
			fmt.Fprintf(os.Stderr, "Error processing %s: %v\n", filename, err)
			exitCode = 1
		}
	}

	os.Exit(exitCode)
}

func printFileInfo(filename string) error {
	fileInfo, err := os.Stat(filename)

	if err != nil {
		return fmt.Errorf("file does not existsor cannot be accessed")
	}

	if fileInfo.IsDir() {
		return fmt.Errorf("is a directory")
	}

	size := fileInfo.Size()
	modTime := fileInfo.ModTime()

	fmt.Printf("File: %s\n", filename)
	fmt.Printf("	Size: %d bytes (%s)\n", size, formatFileSize(size))
	fmt.Printf(" 	Modified: %s\n", modTime.Format("2006-01-02 15:04:05"))
	fmt.Printf(" 	Permissions: %s\n", fileInfo.Mode().String())
	fmt.Println("---")
	
	return nil
}

func formatFileSize(size int64) string {
		const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/KB)
	default:
		return fmt.Sprintf("%d bytes", size)
	}
}