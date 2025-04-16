package getwork

import (
	"encoding/json"
	"fmt"
	"hustLog/header"
	withlogin "hustLog/withLogin"
	"log"
	"os"
	"strconv"
	"strings"
)

func MustQueryWork1(wl *withlogin.WithLogin) []byte {
	_, classes, err := wl.Get("https://hard-working.hust.edu.cn/wechat/publicBenLabor/student/queryOptionalCourseList.do?pageNum=1&pageSize=100", nil)
	if err != nil {
		panic(err)
	}
	return classes
}

func FilterAvailible(wl *withlogin.WithLogin, data []byte, savePath string) (message string, isAnyNew bool) {
	BodyText := &strings.Builder{}
	newClasses := make([]string, 0)

	var r R1
	json.Unmarshal(data, &r)

	lastIds := make([]int, 0)
	fp, _ := os.OpenFile(savePath, os.O_RDWR|os.O_CREATE, 0666)
	defer fp.Close()
	json.NewDecoder(fp).Decode(&lastIds)

	ids := make([]int, 0)

	for _, v := range r.ReturnData.List {
		if v.KXKTS == 0 {
			continue
		}
		ids = append(ids, v.KCID)

		BodyText.WriteString(v.KCMC + "\n")
		GetDitail(wl, v.KCID, BodyText)
		isNew := true
		for _, lv := range lastIds {
			if lv == v.KCID {
				isNew = false
				break
			}
		}
		if isNew {
			newClasses = append(newClasses, v.KCMC)
		}
	}
	fp.Truncate(0)
	fp.Seek(0, 0)
	json.NewEncoder(fp).Encode(ids)

	if len(newClasses) == 0 {
		message = "无新增课程\n\n" + BodyText.String()
		isAnyNew = false
	} else {
		message = "新增课程：\n" + strings.Join(newClasses, "\n") + "\n\n" + BodyText.String()
		isAnyNew = true
	}
	return
}

func GetDitail(wl *withlogin.WithLogin, id int, BodyText *strings.Builder) {
	Id := strconv.Itoa(id)
	url := "https://hard-working.hust.edu.cn/wechat/publicBenLabor/student/queryOptionalCLRMList.do?" + "pageNum=1&pageSize=100&kcid=" + Id + "&kksj=&dwbh=&kcmc="
	_, body, _ := wl.Get(url, header.WorkHeaders)
	var raw Detail
	if err := json.Unmarshal(body, &raw); err != nil {
		log.Println("error:", err)
		return
	}
	count := 0
	for _, list := range raw.ReturnData.List {
		for _, crPlan := range list.ClrmPlanList {
			if count >= 4 {
				break
			}
			BodyText.WriteString(fmt.Sprintf("%s,%s,%s", crPlan.Skrq, crPlan.Xqmc, crPlan.Jc))
			count++
		}
		if count >= 4 {
			BodyText.WriteString("\n有省略\n")
			break
		}
		BodyText.WriteString("\n")
	}
	BodyText.WriteString("\n")
}
