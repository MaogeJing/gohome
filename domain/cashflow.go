package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// 微信交易流水记录
// 交易时间,交易类型,交易对方,商品,收/支,金额(元),支付方式,当前状态,交易单号,商户单号,备注
// 2024-08-31 19:08:24,商户消费,物米智能科技（上海）有限公司,"NL0113570设备管理费:￥100.00",支出,¥100.00,招商银行储蓄卡(5677),支付成功,4200002400202408314730622090	,20240831110113130266217798213778	,"/"
type WechatTransaction struct {
	gorm.Model
	TransactionDate  time.Time `gorm:"type:datetime"`       // 交易时间
	TransactionType  string    `gorm:"size:255"`            // 交易类型
	TransactionParty string    `gorm:"size:255"`            // 交易对方
	ProductID        string    `gorm:"size:255"`            // 商品ID，这里假设商品ID是字符串类型
	Amount           float64   `gorm:"type:decimal(10,2);"` // 金额，假设最多有10位整数和2位小数
	PaymentMethod    string    `gorm:"size:255"`            // 支付方式
	Status           string    `gorm:"size:255"`            // 当前状态
	TransactionNo    string    `gorm:"size:255"`            // 交易单号
	MerchantNo       string    `gorm:"size:255"`            // 商户单号
	Remarks          string    `gorm:"type:text"`           // 备注
}

type WechatTransactionRepository interface {
	Create(c context.Context, transaction *WechatTransaction) error
	Fetch(c context.Context) ([]WechatTransaction, error)
}
