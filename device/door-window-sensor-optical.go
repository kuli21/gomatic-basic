package device

import (
	"github.com/kuli21/gomatic-basic/handler"
	"github.com/kuli21/gomatic-basic/model"
)

type DoorWindowSensorOptical struct {
	Address string
	Error   bool //ERROR
	Lowbat  bool //LOWBAT
	State   bool //STATE
}

func (p *DoorWindowSensorOptical) GetData(ccuUrl string) error {
	pps, err := handler.GetParamset(p.Address, ccuUrl)
	if err != nil {
		return err
	}
	temp := mapDataDoorWindowSensorOptical(pps, p.Address)
	p.Error = temp.Error
	p.Lowbat = temp.Lowbat
	p.State = temp.State

	return nil
}

func mapDataDoorWindowSensorOptical(r model.XResponse, address string) DoorWindowSensorOptical {
	var t DoorWindowSensorOptical
	t.Address = address
	s := len(r.XParams.XParam.XValue.XStruct.XMember)
	for i := 0; i < s; i++ {
		name := r.XParams.XParam.XValue.XStruct.XMember[i].Name
		switch name {
		case "ERROR":
			t.Error = Int2Bool(r.XParams.XParam.XValue.XStruct.XMember[i].Boolean)
		case "LOWBAT":
			t.Lowbat = Int2Bool(r.XParams.XParam.XValue.XStruct.XMember[i].Boolean)
		case "STATE":
			t.State = Int2Bool(r.XParams.XParam.XValue.XStruct.XMember[i].Boolean)
		}
	}
	return t
}
