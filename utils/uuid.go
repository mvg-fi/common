package utils

import "fmt"

func ShortenUUID(uuid string) string {
	return fmt.Sprintf("%s%s%s%s%s", uuid[:8], uuid[9:13], uuid[14:18], uuid[19:23], uuid[24:])
}

func RestoreUUID(s string) string {
	return fmt.Sprintf("%s-%s-%s-%s-%s", s[:8], s[8:12], s[12:16], s[16:20], s[20:])
}
