package main

import "fmt"

func main() {
    p := Port{ Name: "xe-0/0/0", FormFactor: QSFP28, Speeds: []Speed{ SPEED_40G,SPEED_100G }, Breakouts: map[Speed]int{ SPEED_10G: 4, SPEED_25G: 4, SPEED_50G: 2 }, PortGroup: 0 }
    p1 := Port{ Name: "xe-0/0/1", FormFactor: QSFP28, Speeds: []Speed{ SPEED_40G,SPEED_100G }, Breakouts: map[Speed]int{ SPEED_10G: 4, SPEED_25G: 4, SPEED_50G: 2 }, PortGroup: 0 }
    bpMap := map[Speed]BreakoutPort{}
    ProcessBreakoutPorts(p, bpMap)
    ProcessBreakoutPorts(p1, bpMap)
    fmt.Println(bpMap)

    fmt.Println(GetPortDetails(Speed(10), bpMap))
}

type Port struct {
	Name       string             `yaml:"name"`
	FormFactor FormFactor    `yaml:"form_factor"`
	Speeds     []Speed       `yaml:"speeds"`
	Breakouts  map[Speed]int `yaml:"breakouts"`
	PortGroup  int                `yaml:"port_group"`
}

type BreakoutPort struct {
	Name       string             `yaml:"name"`
    NativeSpeed Speed `yaml:"native_speed"`
    Speed   Speed `yaml:"speed"`
	FormFactor FormFactor    `yaml:"form_factor"`
}

type Speed int

// If a Speed constant s1 is less than another Speed constant s2, then it must
// be true that s1.Gbps() < s2.Gbps().  This allows for comparing the constants
// without conversion to Gbps.
//
// A way to view this constraint in code is:
//   var s1, s2 Speed
//   true == (s1 < s2) && (s1.Gbps() < s2.Gbps())
const (
	UnknownSpeed Speed = 0
	SPEED_1G     Speed = 1
	SPEED_10G    Speed = 10
	SPEED_25G    Speed = 25
	SPEED_40G    Speed = 40
	SPEED_50G    Speed = 50
	SPEED_100G   Speed = 100
	SPEED_400G   Speed = 400
	SPEED_MAX    Speed = SPEED_400G
)

type FormFactor int

const (
	UnknownFormFactor FormFactor = iota
	AnyFormFactor
	AnyPluggableFormFactor
	RJ45
	SFP
	SFPP
	SFP28
	QSFP
	QSFPP
	QSFP28
	QSFP_DD
)


func ProcessBreakoutPorts(port Port, bpMap map[Speed]BreakoutPort) {
    if port.Breakouts != nil && len(port.Breakouts) > 0 {
        for k,v := range port.Breakouts {
            if _, ok := bpMap[k]; !ok {
                fmt.Printf("%v not present, adding this entry \n", k)
                name := fmt.Sprintf("%v x %v Gigabit Ethernet\n", v, int(k))
                bp := BreakoutPort{Name: name, Speed: k, FormFactor: port.FormFactor, NativeSpeed: Speed(int(k)*v)}
                bpMap[k] = bp
            }
        }
    } 
}

func GetPortDetails(speed Speed, bpMap map[Speed]BreakoutPort) BreakoutPort{
    return bpMap[speed]
}


