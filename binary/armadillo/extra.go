package armadillo

import (
	"github.com/amiruldev20/waSocket/binary/armadillo/waArmadilloApplication"
	"github.com/amiruldev20/waSocket/binary/armadillo/waCommon"
	"github.com/amiruldev20/waSocket/binary/armadillo/waConsumerApplication"
	"github.com/amiruldev20/waSocket/binary/armadillo/waMultiDevice"
)

type MessageApplicationSub interface {
	IsMessageApplicationSub()
}

type Unsupported_BusinessApplication waCommon.SubProtocol
type Unsupported_PaymentApplication waCommon.SubProtocol
type Unsupported_Voip waCommon.SubProtocol

var (
	_ MessageApplicationSub = (*waConsumerApplication.ConsumerApplication)(nil) // 2
	_ MessageApplicationSub = (*Unsupported_BusinessApplication)(nil)           // 3
	_ MessageApplicationSub = (*Unsupported_PaymentApplication)(nil)            // 4
	_ MessageApplicationSub = (*waMultiDevice.MultiDevice)(nil)                 // 5
	_ MessageApplicationSub = (*Unsupported_Voip)(nil)                          // 6
	_ MessageApplicationSub = (*waArmadilloApplication.Armadillo)(nil)          // 7
)

func (*Unsupported_BusinessApplication) IsMessageApplicationSub() {}
func (*Unsupported_PaymentApplication) IsMessageApplicationSub()  {}
func (*Unsupported_Voip) IsMessageApplicationSub()                {}
