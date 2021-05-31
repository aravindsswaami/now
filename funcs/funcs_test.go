package funcs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/mitchellh/go-homedir"
)

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

var timeNow = func() string { return time.Now().Format("2006-01-02 15:04:05") }
var db = fmt.Sprintf("/%s.csv", time.Now().Format("2006-01-02"))

func TestDoing(t *testing.T) {
	want := fmt.Sprintf("%s > counting > %s > started", timeNow(), strings.Trim(fmt.Sprint([]string{"school", "assignment", "work"}), "[]"))
	got := Doing("counting", []string{"school", "assignment", "work"})

	assert(t, got, want)
}

func TestAppendEntry(t *testing.T) {
	home, err := homedir.Dir()
	if err != nil {
		log.Panic(err)
	}
	home = home + "/nowData"

	t.Run("single entry append", func(t *testing.T) {
		tags := []string{"coding", "work", "personal"}
		entry := Doing("coding now", tags)
		AppendEntry(entry)
		f, err := os.OpenFile(home+db, os.O_RDONLY, 0644)
		if err != nil {
			log.Panic(err)
		}
		scanner := bufio.NewScanner(f)
		match := false
		got := ""
		for scanner.Scan() {
			if scanner.Text() == entry {
				got = scanner.Text()
				match = true
				break
			}
		}
		if !match {
			t.Errorf("got %q want %q", got, entry)
		}
		defer f.Close()

	})

	t.Run("multiple entry append", func(t *testing.T) {
		tags := []string{"coding", "work", "personal"}
		entry := Doing("coding now", tags)
		entry2 := Doing("hair checking", tags)
		AppendEntry(entry)
		AppendEntry(entry2)
		f, err := os.OpenFile(home+db, os.O_RDONLY, 0644)
		if err != nil {
			log.Panic(err)
		}
		scanner := bufio.NewScanner(f)
		match := false
		match2 := false
		for scanner.Scan() {
			if scanner.Text() == entry {
				match = true
			} else if scanner.Text() == entry2 {
				match2 = true
			}
		}
		if !match && !match2 {
			t.Errorf("Both entries weren't appended")
		}
		if match2 && !match {
			t.Errorf("Entry %q not appeneded", entry)
		}
		if match && !match2 {
			t.Errorf("Entry %q not appeneded", entry2)
		}
		defer f.Close()
	})
}

func TestDone(t *testing.T) {
	want := fmt.Sprintf("%s > counting > school assignment work > completed", timeNow())
	got := Done("counting", []string{"school", "assignment", "work"})

	assert(t, got, want)
}

func TestBreak(t *testing.T) {
	want := fmt.Sprintf("%s > break > school assignment work > bstarted", timeNow())
	got := Break([]string{"school", "assignment", "work"})

	assert(t, got, want)
}

func TestContinue(t *testing.T) {
	want := fmt.Sprintf("%s > break > school assignment work > bcompleted", timeNow())
	got := Continue([]string{"school", "assignment", "work"})

	assert(t, got, want)
}
