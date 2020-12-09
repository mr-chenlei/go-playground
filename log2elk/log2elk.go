package main

// Keyword: ELK、iKuai、爱快、日志、Galaxy client、解压

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	src    = flag.String("src", "./upload_files", "Source of upload_files folder.")
	dst    = flag.String("dst", "", "Destination folder which filebeat monitoring.")
	logDir = flag.String("log", "./log", "Directory of log.")
)

var (
	logger *log.Logger
)

func Insert2EachLine(filename, str string) error {
	fi, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fi.Close()

	return nil
}

func ReadIDTxt(txt string) (string, error) {
	logger.Println("read from", txt)
	fi, err := os.Open(txt)
	if err != nil {
		return "", err
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		return string(line), nil
	}
	return "", nil
}

func DeCompress(tarFile, dest string) (string, error) {
	logger.Println("decompress", tarFile)
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return "", err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)

	var filename string
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return "", err
			}
		}
		filename = dest + hdr.Name
		err = os.MkdirAll(string([]rune(filename)[0:strings.LastIndex(filename, "/")]), 0755)
		if err != nil {
			return "", err
		}
		file, err := os.Create(filename)
		if err != nil {
			return "", err
		}
		io.Copy(file, tr)
	}
	return filename, nil
}

func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				return s, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

func GetAllSubDirectories(path string) []string {
	var s []string
	files, _ := ioutil.ReadDir(*src)
	for _, f := range files {
		s = append(s, f.Name())
	}
	return s
}

func main() {
	flag.Parse()

	_ = os.Mkdir(*logDir, 0755)
	logFile, err := os.OpenFile(*logDir+"/log2elk.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	if *src == "" {
		logger.Panic("source folder can not be nil")
	}
	if *dst == "" {
		logger.Panic("destination folder can not be nil")
	}

	logger.Println("start working...")

	devIDList := GetAllSubDirectories(*src)
	for _, v1 := range devIDList {
		// Process logs one folder at a time
		var s []string
		dir := *src + "/" + v1
		s, err = GetAllFile(*src+"/"+v1, s)
		if err != nil {
			logger.Println(dir, "not a directory")
			continue
		}
		// Read from id.txt
		var idTxt string
		for k, v2 := range s {
			if strings.Contains(v2, "id.txt") {
				idTxt, err = ReadIDTxt(v2)
				if err != nil {
					logger.Println(err)
				}
				idTxt = idTxt[1:]
				idTxt = idTxt[:len(idTxt)-1]

				s[k] = s[len(s)-1] // Copy last element to index k.
				s[len(s)-1] = ""   // Erase last element (write zero value).
				s = s[:len(s)-1]   // Truncate slice.
			}
		}
		// DeCompress tar.gz to temp directory
		for _, v3 := range s {
			dstDir := "./temp/" + v1 + "/"
			filename, err := DeCompress(v3, dstDir)
			if err != nil {
				logger.Println(v3, "not a tar.gz file")
				continue
			}
			// TODO: Insert id.txt content to each line of log file
			err = Insert2EachLine(dstDir+filename, idTxt)
			if err != nil {
				logger.Println(err)
				continue
			}
			// TODO: Move log from temp folder to filebeat
			// TODO: Confirm jobs success, then delete *.tar.gz
		}
	}
}
