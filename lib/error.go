package lib

import "fmt"

func ErrFileParseError(line int, key string, value any) error {
	return fmt.Errorf("file parsing error on line %d, key=%s value=%v", line, key, value)
}
