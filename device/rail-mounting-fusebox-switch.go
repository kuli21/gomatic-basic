package device

import (
	"strings"

	"github.com/kuli21/gomatic-basic/handler"
	"github.com/kuli21/gomatic-basic/model"
)

type RailMountingSwitch struct {
	Address       string
	Inhibit       bool
	State         bool
	Working       bool
	Boot          bool
	Current       float64
	EnergyCounter float64
	Frequency     float64
	Power         float64
	Voltage       float64
}

func (p *RailMountingSwitch) GetData(ccuUrl string) error {
	//Check if address contains :1 or :2 ... if not then get data from :1 and :2
	address1 := p.Address
	address2 := p.Address
	if strings.Contains(p.Address, ":") == false {
		address1 = p.Address + ":1"
		address2 = p.Address + ":2"
	}
	pps, err := handler.GetParamset(address1, ccuUrl)
	if err != nil {
		return err
	}
	temp := mapDataRailMountingSwitch(pps, address1)
	p.Inhibit = temp.Inhibit
	p.State = temp.State
	p.Working = temp.Working
	p.Boot = temp.Boot
	p.Current = temp.Current
	p.EnergyCounter = temp.EnergyCounter
	p.Frequency = temp.Frequency
	p.Power = temp.Power
	p.Voltage = temp.Voltage

	if address1 != address2 {
		pps, err := handler.GetParamset(address2, ccuUrl)
		if err != nil {
			return err
		}
		temp := mapDataRailMountingSwitch(pps, address2)
		p.Boot = temp.Boot
		p.Current = temp.Current
		p.EnergyCounter = temp.EnergyCounter
		p.Frequency = temp.Frequency
		p.Power = temp.Power
		p.Voltage = temp.Voltage
	}

	return nil
}

func (p *RailMountingSwitch) ToggleOn(ccuUrl string) error {
	err := handler.SetValue(p.Address, "STATE", "1", ccuUrl)
	if err != nil {
		return err
	}
	return nil
}

func (p *RailMountingSwitch) ToggleOff(ccuUrl string) error {
	err := handler.SetValue(p.Address, "STATE", "0", ccuUrl)
	if err != nil {
		return err
	}
	return nil
}

func mapDataRailMountingSwitch(r model.XResponse, address string) RailMountingSwitch {
	var t RailMountingSwitch
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
		case "BOOT":
			t.Boot = Int2Bool(r.XParams.XParam.XValue.XStruct.XMember[i].Boolean)
		case "CURRENT":
			t.Current = r.XParams.XParam.XValue.XStruct.XMember[i].Double
		case "ENERGY_COUNTER":
			t.EnergyCounter = r.XParams.XParam.XValue.XStruct.XMember[i].Double
		case "FREQUENCY":
			t.Frequency = r.XParams.XParam.XValue.XStruct.XMember[i].Double
		case "POWER":
			t.Power = r.XParams.XParam.XValue.XStruct.XMember[i].Double
		case "VOLTAGE":
			t.Voltage = r.XParams.XParam.XValue.XStruct.XMember[i].Double
		}
	}
	return t
}
