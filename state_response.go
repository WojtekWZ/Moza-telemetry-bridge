package main

type WTState struct {
	Valid                bool    `json:"valid"`
	Aileron              float64 `json:"aileron"`
	Elevator             float64 `json:"elevator"`
	Rudder               float64 `json:"rudder"`
	Flaps                float64 `json:"flaps"`
	Gear                 float64 `json:"gear"`
	Airbrake             float64 `json:"airbrake"`
	HM                   int     `json:"H"`
	TASKmH               int     `json:"TAS"`
	IASKmH               int     `json:"IAS"`
	M                    float64 `json:"M"`
	AoADeg               float64 `json:"AoA"`
	AoSDeg               float64 `json:"AoS"`
	Ny                   float64 `json:"Ny"`
	VyMS                 float64 `json:"Vy"`
	WxDegS               float64 `json:"Wx"`
	MfuelKg              int     `json:"Mfuel"`
	Mfuel0Kg             int     `json:"Mfuel0"`
	Throttle1            int     `json:"throttle 1"`
	RPMThrottle1         int     `json:"RPM throttle 1"`
	Radiator1            int     `json:"radiator 1"`
	Magneto1             int     `json:"magneto 1"`
	Power1Hp             float64 `json:"power 1"`
	RPM1                 int     `json:"RPM 1"`
	ManifoldPressure1Atm float64 `json:"manifold pressure 1"`
	WaterTemp1C          int     `json:"water temp 1"`
	OilTemp1C            int     `json:"oil temp 1"`
	Pitch1Deg            float64 `json:"pitch 1"`
	Thrust1Kgs           int     `json:"thrust 1"`
	Efficiency1          int     `json:"efficiency 1"`
	Throttle2            int     `json:"throttle 2"`
	RPMThrottle2         int     `json:"RPM throttle 2"`
	Radiator2            int     `json:"radiator 2"`
	Magneto2             int     `json:"magneto 2"`
	Power2Hp             float64 `json:"power 2"`
	RPM2                 int     `json:"RPM 2"`
	ManifoldPressure2Atm float64 `json:"manifold pressure 2"`
	WaterTemp2C          int     `json:"water temp 2"`
	OilTemp2C            int     `json:"oil temp 2"`
	Pitch2Deg            float64 `json:"pitch 2"`
	Thrust2Kgs           int     `json:"thrust 2"`
	Efficiency2          int     `json:"efficiency 2"`
}
