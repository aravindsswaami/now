package funcs

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
)

func Doing(task string, tags []string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	entry := fmt.Sprintf("%s > %s > %s > started", timestamp, task, strings.Trim(fmt.Sprint(tags), "[]"))
	return entry
}

func Done(task string, tags []string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	entry := fmt.Sprintf("%s > %s > %s > completed", timestamp, task, strings.Trim(fmt.Sprint(tags), "[]"))
	return entry
}

func Break(tags []string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	entry := fmt.Sprintf("%s > break > %s > bstarted", timestamp, strings.Trim(fmt.Sprint(tags), "[]"))
	return entry
}

func Continue(tags []string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	entry := fmt.Sprintf("%s > break > %s > bcompleted", timestamp, strings.Trim(fmt.Sprint(tags), "[]"))
	return entry
}

func AppendEntry(entry string) {
	db := fmt.Sprintf("/%s.csv", time.Now().Format("2006-01-02"))
	home, err := homedir.Dir()
	if err != nil {
		log.Panic(err)
	}
	home = home + "/nowData"
	f, err := os.OpenFile(home+db, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	writebuffer := bufio.NewWriter(f)
	writebuffer.WriteString(fmt.Sprintln(entry))
	writebuffer.Flush()
}

func ReadLastEntry(n int) [][]string {
	db := fmt.Sprintf("/%s.csv", time.Now().Format("2006-01-02"))
	home, err := homedir.Dir()
	if err != nil {
		log.Panic(err)
	}
	home = home + "/nowData"
	f, err := os.OpenFile(home+db, os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	r.Comma = '>'
	record, err := r.ReadAll()

	if err != nil {
		log.Panic(err)
	}
	if n < len(record) {
		return record[len(record)-n:]
	} else {
		return record
	}
}
