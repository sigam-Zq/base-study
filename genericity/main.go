package main

import (
	"fmt"
	"time"
)

type ImportStruct[T any] interface {
	Convert(item T) error
}

type SupermarketForm struct{}

func (s *SupermarketForm) Convert(item *Supermarket) error {
	fmt.Println("Convert:", item.Name)
	return nil
}

// 验证实现
var _ ImportStruct[*Supermarket] = (*SupermarketForm)(nil)

func main() {
	form := &SupermarketForm{}
	supermarket := &Supermarket{Name: "Test Market"}

	var imp ImportStruct[*Supermarket] = form // 实例化接口
	imp.Convert(supermarket)                  // 调用方法
	fG(imp)                                   // 调用泛型方法
}

func fG[T any](t ImportStruct[T]) {
	fmt.Println("success")
}

// Supermarket 大型商超
type Supermarket struct {
	ID           string    `json:"id" gorm:"size:20;primaryKey;"`  // Unique ID
	Name         string    `json:"name" gorm:"size:255;"`          // 消防站点名称
	Adress       string    `json:"adress" gorm:"size:255;"`        // 消防站点名称
	Linkman      string    `json:"linkman" gorm:"size:255;"`       // 联系人
	ContactPhone string    `json:"contact_phone" gorm:"size:255;"` // 联系人电话
	Remark       string    `json:"remark" gorm:"size:2048;"`       // 备注
	CreatedAt    time.Time `json:"created_at" gorm:"index;"`       // 创建时间
	UpdatedAt    time.Time `json:"updated_at" gorm:"index;"`       // 修改时间
	CreatedBy    string    `json:"created_by" gorm:"size:20;"`     // 创建人
	UpdatedBy    string    `json:"updated_by" gorm:"size:20;"`     // 修改人
}

// Defining the data structure for creating a `Supermarket` struct.
// type SupermarketForm struct {
// 	Name         string `json:"name"`          // 消防站点名称
// 	Adress       string `json:"adress"`        // 消防站点名称
// 	Linkman      string `json:"linkman"`       // 联系人
// 	ContactPhone string `json:"contact_phone"` // 联系人电话
// 	Remark       string `json:"remark"`        // 备注
// }
