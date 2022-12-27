package device

import (
	"github.com/kuli21/gomatic-basic/handler"
	"github.com/kuli21/gomatic-basic/model"
)

type PowerPlugSwitch struct {
	Address string `hmk:"ADDRESS"`
	Inhibit bool   `hmk:"INHIBIT"`
	State   bool   `hmk:"STATE"`
	Working bool   `hmk:"WORKING"`
}

func (p *PowerPlugSwitch) GetData(ccuUrl string) error {
	pps, err := handler.GetParamset(p.Address, ccuUrl)
	if err != nil {
		return err
	}
	temp := mapDataPowerPlugSwitch(pps, p.Address)
	p.Inhibit = temp.Inhibit
	p.State = temp.State
	p.Working = temp.Working

	return nil
}

func (p *PowerPlugSwitch) ToggleOn(ccuUrl string) error {
	err := handler.SetValue(p.Address, "STATE", "1", ccuUrl)
	if err != nil {
		return err
	}
	return nil
}

func (p *PowerPlugSwitch) ToggleOff(ccuUrl string) error {
	err := handler.SetValue(p.Address, "STATE", "0", ccuUrl)
	if err != nil {
		return err
	}
	return nil
}

func mapDataPowerPlugSwitch(r model.XResponse, address string) PowerPlugSwitch {
	var t PowerPlugSwitch
	t.Address = address
	s := len(r.XParams.XParam.XValue.XStruct.XMember)
	for i := 0; i < s; i++ {
		name := r.XParams.XParam.XValue.XStruct.XMember[i].Name
		switch name {
		case "INHIBIT":
			t.Inhibit = Int2Bool(r.XParams.XParam.XValue.XStruct.XMember[i].Boolean)
		case "STATE":
			t.State = Int2Bool(r.XParams.XParam.XValue.XStruct.XMember[i].Boolean)
		case "WORKING":
			t.Working = Int2Bool(r.XParams.XParam.XValue.XStruct.XMember[i].Boolean)
		}
	}
	return t
}
