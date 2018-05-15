package module

// 资产信息
type Product struct {
	TxId         string `json:"txId"`
	ProductId    string `json:"productId"`
	ProductName  string `json:"productName"`
	Operation    string `json:"operation"`
	Operator     string `json:"operator"`    // 入栏人
	CreateTime   uint64 `json:"createTime"`  // 资产创建时间
	BatchNumber  string `json:"batchNumber"` // 入栏批次
	Kind         string `json:"kind"`        // 种类
	Type         string `json:"type"`        // 品种
	MapPosition  string `json:"mapPosition"` // 入栏地理位置
	Status       string `json:"status"`      // 当前状态
	PreOwner     string `json:"preOwner"`
	CurrentOwner string `json:"currentOwner"`
	// 出栏
	OutputTxId        string    `json:"outputTxId"`
	OutputOperation   string    `json:"outputOperation"`   //出栏类型
	OutputOperator    string    `json:"outputOperator"`    //出栏人
	OutputTime        uint64    `json:"outputTime"`        //出栏时间
	OutputMapPosition string    `json:"outputMapPosition"` //地理位置
	FeedList          []Feed    `json:"feedList"`          //喂养列表
	VaccineList       []Vaccine `json:"vaccineList"`       //防疫列表
	// 灭尸
	LostTxId        string `json:"lostTxId"`
	LostOperation   string `json:"lostOperation"`   //灭尸类型
	LostOperator    string `json:"lostOperator"`    //灭尸人
	LostTime        uint64 `json:"lostTime"`        //灭尸时间
	LostMapPosition string `json:"lostMapPosition"` //地理位置
	// 检疫
	ExamTxId        string `json:"examTxId"`
	ExamOperation   string `json:"examOperation"`   //检疫类型
	ExamOperator    string `json:"examOperator"`    //检疫人
	ExamTime        string `json:"examTime"`        //防疫时间
	ExamMapPosition string `json:"examMapPosition"` // 地理位置
	// 屠宰
	ButcherTxId        string `json:"butcherTxId"`
	HookNo             string `json:"hookNo"`             //挂钩号码
	ButcherOperation   string `json:"butcherOperation"`   //出栏类型
	ButcherOperator    string `json:"butcherOperator"`    //出栏人
	ButcherTime        uint64 `json:"butcherTime"`        //出栏时间
	ButcherMapPosition string `json:"butcherMapPosition"` //地理位置

}

type Block struct {
	ParentId  string `json:"parentId"`
	BlockId   string `json:"blockId"`
	BlockName string `json:"blockName"`
	BlockTime uint64 `json:"blockTime"`
}

type Feed struct {
	TxId        string `json:"txId"`
	ProductId   string `json:"productId"`
	Operation   string `json:"operation"`
	Operator    string `json:"operator"`
	FeedTime    uint64 `json:"feedTime"`    //喂养的时间
	MapPosition string `json:"mapPosition"` // 喂养地理位置
}

// 防疫结构
type Vaccine struct {
	TxId        string `json:"txId"`
	ProductId   string `json:"productId"`
	Operation   string `json:"operation"` //防疫类型
	Operator    string `json:"operator"`  //防疫人
	VaccineName string `json:"vaccineName"`
	VaccineTime uint64 `json:"vaccineTime"` //防疫时间
	MapPosition string `json:"mapPosition"` // 地理位置
}

type ChangeAssetOwner struct {
	ProductId    string `json:"productId"`
	Operation    string `json:"operation"`
	Operator     string `json:"operator"`
	OperateTime  uint64 `json:"operateTime"`
	PreOwner     string `json:"preOwner"`
	CurrentOwner string `json:"currentOwner"`
}
