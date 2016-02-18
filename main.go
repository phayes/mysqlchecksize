package main

import (
	"flag"
	"fmt"
	"github.com/phayes/errors"
	"github.com/vaughan0/go-ini"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	ErrDataDir    = errors.New("Could not determine mysql data directory")
	ErrInvalidDB  = errors.New("Invalid database name")
	ErrInvalidDir = errors.New("Invalid database directory")
	ValidDBName   = regexp.MustCompile("[0-9a-zA-Z$_]+")
	DataDir       string
)

func main() {
	flag.Parse()

	c, err := ini.LoadFile("/etc/my.cnf")
	if err != nil {
		// Syntax errors are OK because my.cnf contains weird syntax
		if _, ok := err.(ini.ErrSyntax); !ok {
			log.Fatal(err)
		}
	}
	var ok bool
	DataDir, ok = c.Get("mysqld", "datadir")
	if !ok {
		log.Fatal(ErrDataDir)
	}

	var databases []string
	if flag.Arg(0) != "" {
		db := flag.Arg(0)
		if !ValidDBName.MatchString(db) {
			log.Fatal(ErrInvalidDB)
		}
		databases = []string{flag.Arg(0)}
	} else {
		databases = make([]string, 0)
		files, err := ioutil.ReadDir(DataDir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if file.IsDir() && ValidDBName.MatchString(file.Name()) {
				databases = append(databases, file.Name())
			}
		}
	}

	for _, db := range databases {
		size, err := GetDBSize(db)
		if err != nil {
			log.Fatal(err)
		}
		if flag.Arg(0) == "" {
			fmt.Println(db + "	" + size)
		} else {
			fmt.Println(size)
		}
	}
}

func GetDBSize(dbname string) (size string, err error) {
	info, err := os.Stat(DataDir + "/" + dbname)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return "", ErrInvalidDir
	}

	cmd := exec.Command("du", DataDir+"/"+dbname, "-s", "-k")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Append(err, errors.New(string(output)))
	}
	size = strings.TrimSpace(strings.Split(string(output), "\t")[0])

	return size, nil
}
