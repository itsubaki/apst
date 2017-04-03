package ranking

import (
	"encoding/json"
	"fmt"
)

type Feed struct {
	AppList
}

type AppList [](*App)

func (l AppList) Select(keyword string) AppList {
	list := AppList{}
	for i := 0; i < len(l); i++ {
		if l[i].Contains(keyword) {
			list = append(list, l[i])
		}
	}
	return list
}

func (f *Feed) Select(keyword string) AppList {
	list := AppList{}
	for i := 0; i < len(f.AppList); i++ {
		if f.AppList[i].Contains(keyword) {
			list = append(list, f.AppList[i])
		}
	}
	return list
}

func NewFeed(b []byte) *Feed {
	var content interface{}
	err := json.Unmarshal(b, &content)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	feed := content.(map[string]interface{})["feed"]
	entry := feed.(map[string]interface{})["entry"]
	entrylist := entry.([]interface{})

	list := AppList{}
	for i := 0; i < len(entrylist); i++ {
		app := NewApp(entrylist[i], i+1)
		list = append(list, app)
	}

	return &Feed{list}
}
