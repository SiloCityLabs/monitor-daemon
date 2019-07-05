package memory

import "fmt"

// FormatSizeUnits returns a string representation of bytes to the highest possible unit as a string
func FormatSizeUnits(byteCount int64) string {
	bytes := float64(byteCount)

	if bytes >= 1180591620717411303424 {
		return fmt.Sprintf("%.2f", (bytes/1180591620717411303424)) + " ZB"
	} else if bytes >= 1152921504606846976 {
		return fmt.Sprintf("%.2f", (bytes/1152921504606846976)) + " EB"
	} else if bytes >= 1125899906842624 {
		return fmt.Sprintf("%.2f", (bytes/1125899906842624)) + " PB"
	} else if bytes >= 1099511627776 {
		return fmt.Sprintf("%.2f", (bytes/1099511627776)) + " TB"
	} else if bytes >= 1073741824 {
		return fmt.Sprintf("%.2f", (bytes/1073741824)) + " GB"
	} else if bytes >= 1048576 {
		return fmt.Sprintf("%.2f", (bytes/1048576)) + " MB"
	} else if bytes >= 1024 {
		return fmt.Sprintf("%.2f", (bytes/1024)) + " KB"
	} else if bytes > 1 {
		return fmt.Sprintf("%.2f", bytes) + " bytes"
	} else if bytes == 1 {
		return "1 byte"
	}

	return "0 bytes"
}
