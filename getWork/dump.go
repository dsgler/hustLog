package getwork

type R1Row struct {
	KCID   int    `json:"KCID"`
	KXKTS  int    `json:"KXKTS"`
	DWMC   string `json:"DWMC"`
	KTS    int    `json:"KTS"`
	KCMC   string `json:"KCMC"`
	KCZXS  int    `json:"KCZXS"`
	ROW_ID int    `json:"ROW_ID"`
}
type R1 struct {
	ReturnCode string `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
	ReturnData struct {
		List []R1Row `json:"list"`
	} `json:"returnData"`
}

type List struct {
	RowID       int    `json:"row_id"`
	Xmid        int    `json:"xmid"`
	Xqh         string `json:"xqh"`
	Xqmc        string `json:"xqmc"`
	Xmmc        string `json:"xmmc"`
	Kcid        int    `json:"kcid"`
	Kcmc        string `json:"kcmc"`
	Kcdm        string `json:"kcdm"`
	Kczxs       int    `json:"kczxs"`
	Sfxk        string `json:"sfxk"`
	Sfpk        string `json:"sfpk"`
	Sfxyqd      string `json:"sfxyqd"`
	Xmrl        int    `json:"xmrl"`
	Ktxh        int    `json:"ktxh"`
	Scbz        string `json:"scbz"`
	Zt          string `json:"zt"`
	Ztmc        string `json:"ztmc"`
	URLKey      string `json:"url_key"`
	Dwbh        string `json:"dwbh"`
	Dwmc        string `json:"dwmc"`
	Ktrs        int    `json:"ktrs"`
	Jsfzxs      string `json:"jsfzxs"`
	Jskdrs      int    `json:"jskdrs"`
	Xmbjrs      string `json:"xmbjrs"`
	Xmkssj      string `json:"xmkssj"`
	Xmjssj      string `json:"xmjssj"`
	Xmsj        string `json:"xmsj"`
	Zzxmly      string `json:"zzxmly"`
	Bz          string `json:"bz"`
	Cjjhzt      string `json:"cjjhzt"`
	Cjrid       string `json:"cjrid"`
	Cjip        string `json:"cjip"`
	Zhxgrid     string `json:"zhxgrid"`
	Zhxgip      string `json:"zhxgip"`
	Qrlrwbsj    string `json:"qrlrwbsj"`
	Cjjhrid     string `json:"cjjhrid"`
	Cjjhip      string `json:"cjjhip"`
	UnitList    string `json:"unitList"`
	ClrmTeaList []struct {
		Xmid    int    `json:"xmid"`
		Xldyid  int    `json:"xldyid"`
		Jhjsid  int    `json:"jhjsid"`
		Jlid    int    `json:"jlid"`
		Zdjsgh  string `json:"zdjsgh"`
		Zdjsxm  string `json:"zdjsxm"`
		Sfylrcj string `json:"sfylrcj"`
		Jskdrs  int    `json:"jskdrs"`
		Yxrs    int    `json:"yxrs"`
		Sqid    int    `json:"sqid"`
		Sfypj   string `json:"sfypj"`
		Pf      string `json:"pf"`
		Zyjszc  string `json:"zyjszc"`
		Zp      string `json:"zp"`
	} `json:"clrmTeaList"`
	ClrmPlanList []crPlan `json:"clrmPlanList"`
	Rowspan      int      `json:"rowspan"`
	Kxrl         int      `json:"kxrl"`
	Sfyxk        string   `json:"sfyxk"`
	Sfyks        string   `json:"sfyks"`
	Sfyjs        string   `json:"sfyjs"`
	Dys          int      `json:"dys"`
	Dcldys       int      `json:"dcldys"`
	Wqddys       int      `json:"wqddys"`
	Sfzxksjd     string   `json:"sfzxksjd"`
	K0           string   `json:"k0"`
	K1           string   `json:"k1"`
	Zdqxktsj     string   `json:"zdqxktsj"`
}
type crPlan struct {
	Jlid        int    `json:"jlid"`
	Xq          int    `json:"xq"`
	Xqmc        string `json:"xqmc"`
	Qszc        int    `json:"qszc"`
	Jszc        int    `json:"jszc"`
	Zc          string `json:"zc"`
	Qsjc        int    `json:"qsjc"`
	Jsjc        int    `json:"jsjc"`
	Jc          string `json:"jc"`
	Jsbh        string `json:"jsbh"`
	Jsmc        string `json:"jsmc"`
	Jsrl        int    `json:"jsrl"`
	Skrq        string `json:"skrq"`
	TeacherList string `json:"teacherList"`
}
type Detail struct {
	ReturnCode string `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
	ReturnData struct {
		PageNum           int    `json:"pageNum"`
		PageSize          int    `json:"pageSize"`
		Size              int    `json:"size"`
		StartRow          int    `json:"startRow"`
		EndRow            int    `json:"endRow"`
		Total             int    `json:"total"`
		Pages             int    `json:"pages"`
		List              []List `json:"list"`
		PrePage           int    `json:"prePage"`
		NextPage          int    `json:"nextPage"`
		IsFirstPage       bool   `json:"isFirstPage"`
		IsLastPage        bool   `json:"isLastPage"`
		HasPreviousPage   bool   `json:"hasPreviousPage"`
		HasNextPage       bool   `json:"hasNextPage"`
		NavigatePages     int    `json:"navigatePages"`
		NavigatepageNums  []int  `json:"navigatepageNums"`
		NavigateFirstPage int    `json:"navigateFirstPage"`
		NavigateLastPage  int    `json:"navigateLastPage"`
		FirstPage         int    `json:"firstPage"`
		LastPage          int    `json:"lastPage"`
	} `json:"returnData"`
}
