package main

import "fmt"

type DCSResponse struct {
	aircraft_name           string
	engine_rpm_left         float64
	engine_rpm_right        float64
	left_gear               float64
	nose_gear               float64
	right_gear              float64
	acc_x                   float64
	acc_y                   float64
	acc_z                   float64
	vector_velocity_x       float64
	vector_velocity_y       float64
	vector_velocity_z       float64
	tas                     float64
	ias                     float64
	vertical_velocity_speed float64
	aoa                     float64
	pitch                   float64
	bank                    float64
	aos                     float64
	flap_pos                float64
	gear_value              float64
	speedbrake_value        float64
	afterburner_1           float64
	afterburner_2           float64
	cannon_shells           int
	mach                    float64
	h_above_sea_level       float64
}

func (dcsResponse *DCSResponse) toString() string {
	return fmt.Sprintf("aircraft_name,P-51D;"+
		"engine_rpm_left, %f;"+
		"engine_rpm_right,%f;"+
		"left_gear,%f;"+
		"nose_gear,%f;"+
		"right_gear,%f;"+
		"acc_x,%f;"+
		"acc_y,%f;"+
		"acc_z,%f;"+
		"vector_velocity_x,%f;"+
		"vector_velocity_y,%f;"+
		"vector_velocity_z,%f;"+
		"tas,%f;"+
		"ias,%f;"+
		"vertical_velocity_speed,%f;"+
		"aoa,%f;"+
		"pitch,%f;"+
		"bank,%f;"+
		"aos,%f;"+
		"flap_pos,%f;"+
		"gear_value,%f;"+
		"speedbrake_value,%f;"+
		"afterburner_1,%f;"+
		"afterburner_2,%f;"+
		"weapon,;"+
		"flare,0;"+
		"chaff,0;"+
		"cannon_shells,%d;"+
		"mach,%f;"+
		"h_above_sea_level,%f;",
		dcsResponse.engine_rpm_left,
		dcsResponse.engine_rpm_right,
		dcsResponse.left_gear,
		dcsResponse.nose_gear,
		dcsResponse.right_gear,
		dcsResponse.acc_x,
		dcsResponse.acc_y,
		dcsResponse.acc_z,
		dcsResponse.vector_velocity_x,
		dcsResponse.vector_velocity_y,
		dcsResponse.vector_velocity_z,
		dcsResponse.tas,
		dcsResponse.ias,
		dcsResponse.vertical_velocity_speed,
		dcsResponse.aoa,
		dcsResponse.pitch,
		dcsResponse.bank,
		dcsResponse.aos,
		dcsResponse.flap_pos,
		dcsResponse.gear_value,
		dcsResponse.speedbrake_value,
		dcsResponse.afterburner_1,
		dcsResponse.afterburner_2,
		dcsResponse.cannon_shells,
		dcsResponse.mach,
		dcsResponse.h_above_sea_level)
}
