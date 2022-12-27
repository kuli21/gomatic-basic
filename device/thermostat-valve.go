package device

import (
	"github.com/kuli21/gomatic-basic/handler"
	"github.com/kuli21/gomatic-basic/model"
)

type ThermostatValve struct {
	Address           string  `hmk:"ADDRESS"`
	ActualTemperature float64 `hmk:"ACTUAL_TEMPERATURE"`
	BatteryState      float64 `hmk:"BATTERY_STATE"`
	BoostState        int     `hmk:"BOOST_STATE"`
	ControlMode       int     `hmk:"CONTROL_MODE"`
	FaultReporting    int     `hmk:"FAULT_REPORTING"`
	PartyStartDay     int     `hmk:"PARTY_START_DAY"`
	PartyStartMonth   int     `hmk:"PARTY_START_MONTH"`
	PartyStartTime    int     `hmk:"PARTY_START_TIME"`
	PartyStartYear    int     `hmk:"PARTY_START_YEAR"`
	PartyStopDay      int     `hmk:"PARTY_STOP_DAY"`
	PartyStopMonth    int     `hmk:"PARTY_STOP_MONTH"`
	PartyStopTime     int     `hmk:"PARTY_STOP_TIME"`
	PartyStopYear     int     `hmk:"PARTY_STOP_YEAR"`
	PartyTemperature  float64 `hmk:"PARTY_TEMPERATURE"`
	SetTemperature    float64 `hmk:"SET_TEMPERATURE"`
	ValveState        int     `hmk:"VALVE_STATE"`
}

func (p *ThermostatValve) GetData(ccuUrl string) error {
	pps, err := handler.GetParamset(p.Address, ccuUrl)
	if err != nil {
		return err
	}
	address := p.Address
	temp := mapDataThermostatValve(pps, address)
	p.ActualTemperature = temp.ActualTemperature
	p.BatteryState = temp.BatteryState
	p.BoostState = temp.BoostState
	p.ControlMode = temp.ControlMode
	p.FaultReporting = temp.FaultReporting
	p.PartyStartDay = temp.PartyStartDay
	p.PartyStartMonth = temp.PartyStartMonth
	p.PartyStartTime = temp.PartyStartTime
	p.PartyStartYear = temp.PartyStartYear
	p.PartyStopDay = temp.PartyStopDay
	p.PartyStopMonth = temp.PartyStopMonth
	p.PartyStopTime = temp.PartyStopTime
	p.PartyStopYear = temp.PartyStopYear
	p.PartyTemperature = temp.PartyTemperature
	p.SetTemperature = temp.SetTemperature
	p.ValveState = temp.ValveState
	return nil
}

func mapDataThermostatValve(r model.XResponse, address string) ThermostatValve {
	var t ThermostatValve
	t.Address = address
	s := len(r.XParams.XParam.XValue.XStruct.XMember)
	for i := 0; i < s; i++ {
		name := r.XParams.XParam.XValue.XStruct.XMember[i].Name
		switch name {
		case "ACTUAL_TEMPERATURE":
			t.ActualTemperature = r.XParams.XParam.XValue.XStruct.XMember[i].Double
		case "BATTERY_STATE":
			t.BatteryState = r.XParams.XParam.XValue.XStruct.XMember[i].Double
		case "BOOST_STATE":
			t.BoostState = r.XParams.XParam.XValue.XStruct.XMember[i].I4
		case "CONTROL_MODE":
			t.ControlMode = r.XParams.XParam.XValue.XStruct.XMember[i].I4
		case "FAULT_REPORTING":
			t.FaultReporting = r.XParams.XParam.XValue.XStruct.XMember[i].I4
		case "PARTY_START_DAY":
			t.PartyStartDay = r.XParams.XParam.XValue.XStruct.XMember[i].I4
		case "PARTY_START_MONTH":
			t.PartyStartMonth = r.XParams.XParam.XValue.XStruct.XMember[i].I4
		case "PARTY_START_TIME":
			t.PartyStartTime = r.XParams.XParam.XValue.XStruct.XMember[i].I4
		case "PARTY_START_YEAR":
			t.PartyStartYear = r.XParams.XParam.XValue.XStruct.XMember[i].I4
		case "PARTY_STOP_DAY":
			t.PartyStopDay = r.XParams.XParam.XValue.XStruct.XMember[i].I4
		case "PARTY_STOP_MONTH":
			t.PartyStopMonth = r.XParams.XParam.XValue.XStruct.XMember[i].I4
		case "PARTY_STOP_TIME":
			t.PartyStopTime = r.XParams.XParam.XValue.XStruct.XMember[i].I4
		case "PARTY_STOP_YEAR":
			t.PartyStopYear = r.XParams.XParam.XValue.XStruct.XMember[i].I4
		case "PARTY_TEMPERATURE":
			t.PartyTemperature = r.XParams.XParam.XValue.XStruct.XMember[i].Double
		case "SET_TEMPERATURE":
			t.SetTemperature = r.XParams.XParam.XValue.XStruct.XMember[i].Double
		case "VALVE_STATE":
			t.ValveState = r.XParams.XParam.XValue.XStruct.XMember[i].I4
		}
	}
	return t
}
