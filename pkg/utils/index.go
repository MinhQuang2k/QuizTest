package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"quiztest/pkg/logger"
	"strconv"
)

func Copy(dest interface{}, src interface{}) {
	data, _ := json.Marshal(src)
	_ = json.Unmarshal(data, dest)
}

func FindUint(a []uint, x uint) uint {
	for _, n := range a {
		if x == n {
			return x
		}
	}
	return 0
}

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}

func IsUniqueArray(arr []string) bool {
	m := make(map[string]struct{})

	for _, v := range arr {
		m[v] = struct{}{}
	}

	return len(m) == len(arr)
}

func ReadFileRoot(path string) ([]byte, error) {
	if &path == nil || path == "" {
		return nil, errors.New("no path")
	}
	absPath, _ := filepath.Abs(path)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		logger.Error(err)
	}
	return data, err
}
