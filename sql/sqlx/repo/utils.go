package repo

import (
	"fmt"
	"strings"
)

func prfx(prefix, name string) string {
	return fmt.Sprintf("%s.%s", prefix, name)
}

func returning(col ...string) string {
	return fmt.Sprintf("RETURNING %s", strings.Join(col, ", "))
}
