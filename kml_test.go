package kml

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	t.Log("CheckDevice:", CheckDevice())
	t.Log("OpenDevice:", OpenDevice())
	t.Log("GetSN:", GetSN())
	t.Log("GetModel:", GetModel())
	t.Log("GetVersion:", GetVersion())
	t.Log("GetProductionDate:", GetProductionDate())
	data := EncString("abc")
	t.Log("EncString:", data)
	t.Log("DecString:", DecString(data))
	t.Log("SimulationPressKey:", SimulationPressKey("123456789"))
	t.Log("SayString:", SayString("hello,world,zhangzhongyuan"))
	t.Log("GetCapsLock:", GetCapsLock())
	t.Log("GetCapsLock:", GetNumLock())
	t.Log("KeyDown:", KeyDown('n'))

	t.Log("Restart:", Restart())
	t.Log("CloseDevice:", CloseDevice())
	t.Log("Disconnect:", Disconnect(0))
}
