package gitLog

import (
	"bufio"
	"bytes"
	"strings"
	"time"
)

type Commit  struct{
	Hash    string
	Date    time.Time
	Email   string
	Author  string
	Message string
}


func Parse(log string) ([]Commit, error) {
	commitsArray := []Commit{}
	scan := bufio.NewScanner(bytes.NewBufferString(log))
	scan.Split(scanSeparator)
	for scan.Scan() {
		res := scan.Text()
		commit := Commit{}
		// 7 is begin of hash
		// 47 is end of hash
		commit.Hash = res[7:47]
		beginEmail := strings.Index(res, "<")
		endEmail := strings.Index(res, ">")
		// 56 is begin of Author
		commit.Author = res[56:beginEmail - 1]
		commit.Email = res[beginEmail + 1:endEmail]
		date, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", strings.TrimSpace(res[endEmail + 10:endEmail + 41]))
		if err != nil {
			println(err.Error())
			return []Commit{}, err
		}
		commit.Date = date
		commit.Message = res[endEmail + 41:]
		commitsArray = append(commitsArray, commit)
	}
	return commitsArray, nil
}

/**
It's just full copy bufio.ScanLines except bytes.IndexByte was replaced bytes.Index with SEPARATOR
 */
func scanSeparator(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	sourceData := data
	offset := 6 // 6 is length of "commit"
	again:
	if i := bytes.Index(data[offset:], []byte("commit")); i >= 0 {
		offset += i
		//if string(data[offset + 47:offset + 48]) == "\n" {
		if doHasCommitHash(data[offset:offset + 48]) {
			return offset, sourceData[0: offset], nil
		}
		offset += 6
		goto again
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

func doHasCommitHash (str_ []byte) bool {
	str := string(str_)
	//7 - 47 is position "c8365645e8733898d831203e140084034a8d0e6d"
	c := !strings.Contains(str[7:47], " ")
	c = c && !strings.Contains(str[7:47], "\n")
	//47 - 48 is position "\n" in "commit c8365645e8733898d831203e140084034a8d0e6d"
	return c && str[47:48] == "\n"
}