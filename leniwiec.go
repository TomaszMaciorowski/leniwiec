package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

type Printer struct{}

func (p Printer) Walk(name exif.FieldName, tag *tiff.Tag) error {
	fmt.Printf("%40s: %s\n", name, tag)
	return nil
}

func MoveFile(source, destination string) (err error) {
	src, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer src.Close()
	fi, err := src.Stat()
	if err != nil {
		log.Fatal(err)
		return err
	}
	flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	perm := fi.Mode() & os.ModePerm
	dst, err := os.OpenFile(destination, flag, perm)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		dst.Close()
		os.Remove(destination)
		log.Fatal(err)
		return err
	}
	err = dst.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = src.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = os.Remove(source)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func main() {

	var files []string

	if len(os.Args) != 3 {
		fmt.Println("leniwiec  requires argument source_directory destination_directory")
		fmt.Println("leniwiec  /source/photo /destination/photo")
		fmt.Println("autor: tomek.maciorowski@gmail.com")
		log.Fatal("try again")
	}
	fmt.Printf("%v", len(os.Args))
	dirPath := os.Args[1]
	My_Dest := os.Args[2]
	
	
	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && filepath.Ext(path) == ".jpg" {
				files = append(files, path)
				fmt.Println(path)
			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	for index, element := range files {

		fmt.Println(index)
		fmt.Println(element)

		f, err := os.Open(element)
		if err != nil {
			log.Fatal(err)
		}

		x, err := exif.Decode(f)
		if err != nil {
			//	log.Fatal(err)
			continue
		}

		f.Close()

		tm, _ := x.DateTime()
		format_name := fmt.Sprintf("%v%v%v_%v%v%v%v", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), index)
		dname := fmt.Sprintf("%v\\%v.jpg", My_Dest, format_name)
		if tm.Year() == 1 {
			now := time.Now()
			format_name = fmt.Sprintf("%v%v%v_%v%v%v%v", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), index)
			dname = fmt.Sprintf("%v\\%v.jpg", My_Dest, format_name)
			fmt.Println(dname)
			MoveFile(element, dname)
		} else {
			fmt.Println(dname)
			MoveFile(element, dname)
		}

		//var p Printer
		//x.Walk(p)

	}

}
