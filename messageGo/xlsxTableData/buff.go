package xlsxTable

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// buff效果类型.
type BuffEffectType int32

const (
	//未知类型，不处理
	BuffEffectUnknown BuffEffectType = 0
	//随时间减血
	BuffEffectSubHp BuffEffectType = 1
	//随时间加血
	BuffEffectAddHp BuffEffectType = 2
	//随时间减饥饿度(deprecated)
	BuffEffectSubHungry BuffEffectType = 3
	//随时间加饥饿度(deprecated)
	BuffEffectAddHungry BuffEffectType = 4
	//随时间减饥渴度(deprecated)
	BuffEffectSubThirsty BuffEffectType = 5
	//随时间加饥渴度(deprecated)
	BuffEffectAddThirsty BuffEffectType = 6
	//减速
	BuffEffectSubSpeed BuffEffectType = 7
	//加速
	BuffEffectAddSpeed BuffEffectType = 8
)

// buff模板
type BuffTableRow struct {
	UId             uint           `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	BuffId          int32          `json:"buffId"`
	EffectType      BuffEffectType `json:"effectType"`
	GroupId         int32          `json:"groupId"`
	GroupPriority   int32          `json:"groupPriority"` //优先级(越大优先级越高)
	ParamStr        string         `json:"paramStr"`
	TotalTime       int32          `json:"totalTime"`       //总持续时间
	TriggerInterval int32          `json:"triggerInterval"` //触发间隔

	CreatedAt time.Time `json:"createdAt"` // 过期判断条件
}

func (this *BuffTableRow) SetParams(params []int32) {
	for _, i := range params {
		this.ParamStr += fmt.Sprintf("%d,", i)
	}
	if this.ParamStr != "" {
		this.ParamStr = this.ParamStr[:len(this.ParamStr)-1]
	}
}

func (this *BuffTableRow) GetParams() (params []int32) {
	splitArray := strings.Split(this.ParamStr, ",")
	for _, s := range splitArray {
		num, _ := strconv.ParseInt(s, 10, 32)
		params = append(params, int32(num))
	}
	return
}
