package kml

import (
	"fmt"
	"testing"
	"time"
)

func TestGetVersion(t *testing.T) {
	t.Log("OpenDevice:", OpenDevice())
	t.Log("CheckDevice:", CheckDevice())
	t.Log("GetSN:", GetSN())
	t.Log("GetModel:", GetModel())
	t.Log("GetVersion:", GetVersion())
	t.Log("GetProductionDate:", GetProductionDate())

	data := EncString("abc")
	t.Log("EncString:", data)
	t.Log("DecString:", DecString(data))

	// 打开运行
	t.Log("KeyDown:", KeyDown("Win"))
	t.Log("KeyDown:", KeyDown("r"))
	t.Log("KeyUpAll:", KeyUpAll())

	// 运行记事本
	t.Log("SimulationPressKey:", SimulationPressKey("notepad"))
	t.Log("KeyPress:", KeyPress("Enter", 1))
	time.Sleep(time.Second)

	// 输入到记事本
	t.Log("SimulationPressKey:", SimulationPressKey("123456789,hello,world,zhangzhongyuan"))
	t.Log("KeyPress:", KeyPress("Enter", 1))
	t.Log("KeyPress:", KeyPress("0", 10))

	t.Log("GetCapsLock:", GetCapsLock())
	t.Log("GetCapsLock:", GetNumLock())

	t.Log("Restart:", Restart())
	t.Log("CloseDevice:", CloseDevice())
	t.Log("Disconnect:", Disconnect(0))
}

func TestOpenDevice(t *testing.T) {
	fmt.Println(OpenDevice())

	time.Sleep(time.Second)

	KeyDown("Win")
	KeyDown("r")
	KeyUpAll()

	time.Sleep(time.Second)

	SimulationPressKey("notepad")
	KeyPress("Enter", 1)

	time.Sleep(time.Second)

	SimulationPressKey("123456789")
	SimulationPressKey("hello,world,xiaomage")

	time.Sleep(time.Second)

	KeyDown("Ctrl")
	KeyDown("s")
	KeyUpAll()

	time.Sleep(time.Second)

	SimulationPressKey("test.txt")
	KeyDown("Alt")
	KeyDown("s")
	KeyUpAll()

	time.Sleep(time.Second)

	KeyDown("Alt")
	KeyDown("F4")
	KeyUpAll()
}