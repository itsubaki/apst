package format

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func JSON(in interface{}, pretty bool) string {
	if pretty {
		return JSONPretty(in)
	}

	b, err := json.Marshal(in)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func JSONPretty(in interface{}) string {
	b, err := json.Marshal(in)
	if err != nil {
		return err.Error()
	}

	var pretty bytes.Buffer
	err = json.Indent(&pretty, b, "", " ")
	if err != nil {
		return err.Error()
	}
	return string(pretty.Bytes())
}

func ColorPrintln(rating int, message string) {
	ColorPrint(rating, message)
	fmt.Println("")
}

func ColorPrint(rating int, message string) {
	color := "\x1b[30m%s\x1b[0m"
	switch rating {
	case 5:
		color = "\x1b[32m%s\x1b[0m"
	case 4:
		color = "\x1b[34m%s\x1b[0m"
	case 2:
		color = "\x1b[33m%s\x1b[0m"
	case 1:
		color = "\x1b[31m%s\x1b[0m"
	}

	fmt.Printf(color, message)
}
