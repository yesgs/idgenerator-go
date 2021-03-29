package idgenerator_go

// 版权属于：yitter(yitter@126.com)
// 代码编辑：guoyahao
// 代码修订：yitter
// 开源地址：https://gitee.com/yitter/idgenerator

import (
	"sync"

	"gitee.com/yitter/idgenerator-go/contract"
	"gitee.com/yitter/idgenerator-go/gen"
)

var ins *YitIDHelper
var once sync.Once

// YitIDHelper .
type YitIDHelper struct {
	idGenInstance interface {
		NewLong() uint64
	}
}

// GetIns .
func GetIns() *YitIDHelper {
	once.Do(func() {
		ins = &YitIDHelper{}
	})
	return ins
}

// GetIDGenInstance .
func (yih *YitIDHelper) GetIDGenInstance() interface{} {
	return yih.idGenInstance
}

// SetIDGenerator .
func (yih *YitIDHelper) SetIDGenerator(options *contract.IDGeneratorOptions) {
	yih.idGenInstance = gen.NewDefaultIDGenerator(options)
}

// NextID .
func (yih *YitIDHelper) NextID() uint64 {
	once.Do(func() {
		if yih.idGenInstance == nil {
			options := contract.NewIDGeneratorOptions(1)
			yih.idGenInstance = gen.NewDefaultIDGenerator(options)
		}
	})

	return yih.idGenInstance.NewLong()
}