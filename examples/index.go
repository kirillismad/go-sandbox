package examples

import (
	"math/rand"
	"reflect"
	"runtime"
	"strings"

	"github.com/google/uuid"
)

const newline = "\n"

func rint(max int) int {
	return rand.Intn(max) + 1
}

func rfloat(max int) float64 {
	return rand.Float64() * float64(max)
}

func rstring() string {
	result := make([]rune, 10)
	result[0] = rune(rand.Intn(26) + 65)
	for i := 1; i < len(result); i++ {
		result[i] = rune(rand.Intn(26) + 97)
	}
	return string(result)
}

func fname(f interface{}) string {
	fullFuncName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	funcName := fullFuncName[strings.LastIndex(fullFuncName, ".")+1:]
	return funcName
}

func uuid4() uuid.UUID {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return uuid
}
