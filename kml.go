package kml

import (
	"syscall"
	"unsafe"
)

var dll []byte

var (
	// 设备接口
	procOpenDevice        = kml.NewProc("OpenDevice")
	procOpenDeviceEx      = kml.NewProc("OpenDeviceEx")
	procCloseDevice       = kml.NewProc("CloseDevice")
	procCheckDevice       = kml.NewProc("CheckDevice")
	procRestart           = kml.NewProc("Restart")
	procDisconnect        = kml.NewProc("Disconnect")
	procGetSN             = kml.NewProc("GetSN")
	procGetModel          = kml.NewProc("GetModel")
	procGetVersion        = kml.NewProc("GetVersion")
	procGetProductionDate = kml.NewProc("GetProductionDate")
	procSetDeviceID       = kml.NewProc("SetDeviceID")
	procRestoreDeviceID   = kml.NewProc("RestoreDeviceID")

	// 加密锁接口
	procInitLock         = kml.NewProc("InitLock")
	procSetReadPassword  = kml.NewProc("SetReadPassword")
	procSetWritePassword = kml.NewProc("SetWritePassword")
	procReadString       = kml.NewProc("ReadString")
	procWriteString      = kml.NewProc("WriteString")
	procSetStringKey     = kml.NewProc("SetStringKey")
	procEncString        = kml.NewProc("EncString")
	procDecString        = kml.NewProc("DecString")

	// 键盘接口
	procKeyDown             = kml.NewProc("KeyDown")
	procKeyUp               = kml.NewProc("KeyUp")
	procKeyPress            = kml.NewProc("KeyPress")
	procCombinationKeyDown  = kml.NewProc("CombinationKeyDown")
	procCombinationKeyUp    = kml.NewProc("CombinationKeyUp")
	procCombinationKeyPress = kml.NewProc("CombinationKeyPress")
	procSimulationPressKey  = kml.NewProc("SimulationPressKey")
	procKeyUpAll            = kml.NewProc("KeyUpAll")
	procGetCapsLock         = kml.NewProc("GetCapsLock")
	procGetNumLock          = kml.NewProc("GetNumLock")

	// 鼠标接口
	procLeftDown          = kml.NewProc("LeftDown")
	procLeftUp            = kml.NewProc("LeftUp")
	procLeftClick         = kml.NewProc("LeftClick")
	procLeftDoubleClick   = kml.NewProc("LeftDoubleClick")
	procRightDown         = kml.NewProc("RightDown")
	procRightUp           = kml.NewProc("RightUp")
	procRightClick        = kml.NewProc("RightClick")
	procRightDoubleClick  = kml.NewProc("RightDoubleClick")
	procMiddleDown        = kml.NewProc("MiddleDown")
	procMiddleUp          = kml.NewProc("MiddleUp")
	procMiddleClick       = kml.NewProc("MiddleClick")
	procMiddleDoubleClick = kml.NewProc("MiddleDoubleClick")
	procMouseUpAll        = kml.NewProc("MouseUpAll")
	procSimulationMove    = kml.NewProc("SimulationMove")
	procMoveTo            = kml.NewProc("MoveTo")
	procMoveToR           = kml.NewProc("MoveToR")
	procMoveToFrom        = kml.NewProc("MoveToFrom")
	procReMoveTo          = kml.NewProc("ReMoveTo")

	// KM21专用接口
	procGetKM21Mode  = kml.NewProc("GetKM21Mode")
	procSetKM21Mode  = kml.NewProc("SetKM21Mode")

	// TODO 未找到接口文档，但是DLL中有这两个函数
	procSetMoveSpeed = kml.NewProc("SetMoveSpeed")
	procWheelsMove   = kml.NewProc("WheelsMove")

	// TODO 接口文档中辅助函数 改DLL中没有的函数
	procDelay        = kml.NewProc("Delay")
	procRandomDelay  = kml.NewProc("RandomDelay")
	procRunApp       = kml.NewProc("RunApp")
	procGetCursorPos = kml.NewProc("GetCursorPos")
	procMouseX       = kml.NewProc("MouseX")
	procMouseY       = kml.NewProc("MouseY")
	procScreenHeight = kml.NewProc("ScreenHeight")
	procScreenWidth  = kml.NewProc("ScreenWidth")
	procNow          = kml.NewProc("Now")
	procMessageBox   = kml.NewProc("MessageBox")
)

func toBool(ret uintptr) bool {
	return ret == 1
}

func toError(errno syscall.Errno) error {
	if errno == 0 {
		return nil
	}
	return errno
}

func toString(ret uintptr) string {
	var buf = make([]byte, 0, 64)
	for {
		if char := *(*byte)(unsafe.Pointer(ret)); char != 0 && char != 0xFF {
			buf = append(buf, char)
			ret++
			continue
		}
		break
	}
	return string(buf)
}

func fromByte(b byte) uintptr {
	return uintptr(unsafe.Pointer(&b))
}

func fromString(str string) uintptr {
	var buf = make([]byte, len(str)+1)
	copy(buf, str)
	return uintptr(unsafe.Pointer(&buf[0]))
}

func openDevice() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procOpenDevice.Addr(), 0, 0, 0, 0)
	return
}

func openDeviceEx(vid, pid uint16) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procOpenDeviceEx.Addr(), 2, uintptr(vid), uintptr(pid), 0)
	return
}

func closeDevice() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procCloseDevice.Addr(), 0, 0, 0, 0)
	return
}

func checkDevice() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procCheckDevice.Addr(), 0, 0, 0, 0)
	return
}

func restart() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procRestart.Addr(), 0, 0, 0, 0)
	return
}

func disconnect(second uint8) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procDisconnect.Addr(), 1, uintptr(second), 0, 0)
	return
}

func getSN() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procGetSN.Addr(), 0, 0, 0, 0)
	return
}

func getModel() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procGetModel.Addr(), 0, 0, 0, 0)
	return
}

func getVersion() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procGetVersion.Addr(), 0, 0, 0, 0)
	return
}

func getProductionDate() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procGetProductionDate.Addr(), 0, 0, 0, 0)
	return
}

func setDeviceID(pwd string, vid, pid uint16) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procSetDeviceID.Addr(), 3, fromString(pwd), uintptr(vid), uintptr(pid))
	return
}

func restoreDeviceID(pwd string) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procRestoreDeviceID.Addr(), 1, fromString(pwd), 0, 0)
	return
}

func initLock() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procInitLock.Addr(), 0, 0, 0, 0)
	return
}

func setReadPassword(pwd, new string) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procSetReadPassword.Addr(), 2, fromString(pwd), fromString(new), 0)
	return
}

func setWritePassword(pwd, new string) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procSetWritePassword.Addr(), 2, fromString(pwd), fromString(new), 0)
	return
}

func readString(pwd string, address, count int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procReadString.Addr(), 3, fromString(pwd), uintptr(address), uintptr(count))
	return
}

func writeString(pwd string, address int, str string) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procWriteString.Addr(), 3, fromString(pwd), uintptr(address), fromString(str))
	return
}

func setStringKey(pwd, key string) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procSetStringKey.Addr(), 2, fromString(pwd), fromString(key), 0)
	return
}

func encString(str string) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procEncString.Addr(), 1, fromString(str), 0, 0)
	return
}

func decString(str string) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procDecString.Addr(), 1, fromString(str), 0, 0)
	return
}

func keyDown(key string) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procKeyDown.Addr(), 1, fromString(key), 0, 0)
	return
}

func keyUp(key string) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procKeyUp.Addr(), 1, fromString(key), 0, 0)
	return
}

func keyPress(key string, count int) (i int, err error) {
	ret, _, errno := syscall.Syscall(procKeyPress.Addr(), 2, fromString(key), uintptr(count), 0)
	i, err = int(ret), toError(errno)
	return
}

func combinationKeyDown(key1, key2, key3, key4, key5, key6 string) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall6(procCombinationKeyDown.Addr(), 6, fromString(key1), fromString(key2), fromString(key3), fromString(key4), fromString(key5), fromString(key6))
	return
}

func combinationKeyUp(key1, key2, key3, key4, key5, key6 string) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall6(procCombinationKeyUp.Addr(), 6, fromString(key1), fromString(key2), fromString(key3), fromString(key4), fromString(key5), fromString(key6))
	return
}

func combinationKeyPress(key1, key2, key3, key4, key5, key6 string, count int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall9(procCombinationKeyPress.Addr(), 7, fromString(key1), fromString(key2), fromString(key3), fromString(key4), fromString(key5), fromString(key6), uintptr(count), 0, 0)
	return
}

func simulationPressKey(str string) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procSimulationPressKey.Addr(), 1, fromString(str), 0, 0)
	return
}

func keyUpAll() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procKeyUpAll.Addr(), 0, 0, 0, 0)
	return
}

func getCapsLock() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procGetCapsLock.Addr(), 0, 0, 0, 0)
	return
}

func getNumLock() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procGetNumLock.Addr(), 0, 0, 0, 0)
	return
}

func leftDown() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procLeftDown.Addr(), 0, 0, 0, 0)
	return
}

func leftUp() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procLeftUp.Addr(), 0, 0, 0, 0)
	return
}

func leftClick(count int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procLeftClick.Addr(), 1, uintptr(count), 0, 0)
	return
}

func leftDoubleClick(count int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procLeftDoubleClick.Addr(), 1, uintptr(count), 0, 0)
	return
}

func rightDown() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procRightDown.Addr(), 0, 0, 0, 0)
	return
}

func rightUp() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procRightUp.Addr(), 0, 0, 0, 0)
	return
}

func rightClick(count int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procRightClick.Addr(), 1, uintptr(count), 0, 0)
	return
}

func rightDoubleClick(count int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procRightDoubleClick.Addr(), 1, uintptr(count), 0, 0)
	return
}

func middleDown() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procMiddleDown.Addr(), 0, 0, 0, 0)
	return
}

func middleUp() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procMiddleUp.Addr(), 0, 0, 0, 0)
	return
}

func middleClick(count int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procMiddleClick.Addr(), 1, uintptr(count), 0, 0)
	return
}

func middleDoubleClick(count int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procMiddleDoubleClick.Addr(), 1, uintptr(count), 0, 0)
	return
}

func mouseUpAll() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procMouseUpAll.Addr(), 0, 0, 0, 0)
	return
}

func simulationMove(x, y int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procSimulationMove.Addr(), 2, uintptr(x), uintptr(y), 0)
	return
}

func moveTo(x, y int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procMoveTo.Addr(), 2, uintptr(x), uintptr(y), 0)
	return
}

func moveToR(x, y int8) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procMoveToR.Addr(), 2, uintptr(x), uintptr(y), 0)
	return
}

func getKM21Mode() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procGetKM21Mode.Addr(), 0, 0, 0, 0)
	return
}

func setKM21Mode(model int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procSetKM21Mode.Addr(), 1, uintptr(model), 0, 0)
	return
}

func moveToFrom(fx, fy, tx, ty int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall6(procMoveToFrom.Addr(), 4, uintptr(fx), uintptr(fy), uintptr(tx), uintptr(ty), 0, 0)
	return
}

func reMoveTo(x, y int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procReMoveTo.Addr(), 2, uintptr(x), uintptr(y), 0)
	return
}

func delay(ms int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procDelay.Addr(), 1, uintptr(ms), 0, 0)
	return
}

func randomDelay(min, max int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procRandomDelay.Addr(), 2, uintptr(min), uintptr(max), 0)
	return
}

func runApp(filepath string) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procRunApp.Addr(), 1, fromString(filepath), 0, 0)
	return
}

func getCursorPos(x, y *int) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procGetCursorPos.Addr(), 2, uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(y)), 0)
	return
}

func mouseX() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procMouseX.Addr(), 0, 0, 0, 0)
	return
}

func mouseY() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procMouseY.Addr(), 0, 0, 0, 0)
	return
}

func screenHeight() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procScreenHeight.Addr(), 0, 0, 0, 0)
	return
}

func screenWidth() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procScreenWidth.Addr(), 0, 0, 0, 0)
	return
}

func now() (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procNow.Addr(), 0, 0, 0, 0)
	return
}

func messageBox(msg string) (ret uintptr, errno syscall.Errno) {
	ret, _, errno = syscall.Syscall(procMessageBox.Addr(), 1, fromString(msg), 0, 0)
	return
}

// OpenDevice 打开设备
func OpenDevice() bool {
	ret, _ := openDevice()
	return toBool(ret)
}

// OpenDeviceEx 打开指定设备
// 入参:
//     vid: 设备的厂商ID
//     pid: 产品ID
func OpenDeviceEx(vid, pid uint16) bool {
	ret, _ := openDeviceEx(vid, pid)
	return toBool(ret)
}

// CloseDevice 关闭设备
func CloseDevice() bool {
	ret, _ := closeDevice()
	return toBool(ret)
}

// CheckDevice 检测设备是否连接
func CheckDevice() bool {
	ret, _ := checkDevice()
	return toBool(ret)
}

// Restart 重启设备
func Restart() bool {
	ret, _ := restart()
	return toBool(ret)
}

// Disconnect 断开设备延时连接
func Disconnect(second uint8) bool {
	ret, _ := disconnect(second)
	return toBool(ret)
}

// GetSN 获取设备系列号
func GetSN() string {
	ret, _ := getSN()
	return toString(ret)
}

// GetModel 获取设备型号
func GetModel() string {
	ret, _ := getModel()
	return toString(ret)
}

// GetVersion 获取设备版本号
func GetVersion() string {
	ret, _ := getVersion()
	return toString(ret)
}

// GetProductionDate 获取设备出厂日期
func GetProductionDate() string {
	ret, _ := getProductionDate()
	return toString(ret)
}

// SetDeviceID 修改设备ID
// 入参:
//     pwd: 写密码，该密码由SetWritePassword函数设定
//     vid: 设备的厂商ID
//     pid: 产品ID
// 返回值:
//     1: 成功
//     0: 失败
//    -1: 密码错误
func SetDeviceID(pwd string, vid, pid uint16) int {
	ret, _ := setDeviceID(pwd, vid, pid)
	return int(ret)
}

// RestoreDeviceID 恢复设备出厂ID，InitLock函数也会导致恢复设备ID
// 恢复ID将会在设备重启后生效，可使用Restart重启设备
// 入参:
//     pwd: 写密码，该密码由SetWritePassword函数设定
// 返回值:
//     1: 成功
//     0: 失败
//    -1: 密码错误
func RestoreDeviceID(pwd string) int {
	ret, _ := restoreDeviceID(pwd)
	return int(ret)
}

// InitLock 初始化加密锁，初始化后所有用户存储器内容及配置将丢失，恢复到出厂状态
func InitLock() bool {
	ret, _ := initLock()
	return toBool(ret)
}

// SetReadPassword 设置读存储器密码，初始密码为空，设置后读存储器将需要该密码，密码长度最大为8个字节
// 入参:
//     pwd: 写密码，该密码由SetWritePassword函数设定
//     new: 新读密码
// 返回值:
//     1: 成功
//     0: 失败
//    -1: 密码错误
//    -3: 设备不支持该接口（AP-KM01）
func SetReadPassword(pwd, new string) int {
	ret, _ := setReadPassword(pwd, new)
	return int(ret)
}

// SetWritePassword 设置写存储器密码，初始密码为空，，设置后写存储器将需要该密码，密码长度最大为8个字节
// 入参:
//     pwd: 写密码，该密码由SetWritePassword函数设定
//     new: 新读密码
// 返回值:
//     1: 成功
//     0: 失败
//    -1: 密码错误
//    -3: 设备不支持该接口（AP-KM01）
func SetWritePassword(pwd, new string) int {
	ret, _ := setWritePassword(pwd, new)
	return int(ret)
}

// ReadString 从存储器读字符串
// 入参:
//     pwd: 读密码，该密码由SetReadPassword函数设定
//     address: 存储器地址（0-511），共512字节
//     count: 要读取的字节数
func ReadString(pwd string, address, count int) string {
	ret, _ := readString(pwd, address, count)
	return toString(ret)
}

// WriteString 写字符串到存储器，只需指定地址和要写入的字符串，写入长度以字符串长度为准
// 入参:
//     pwd: 写密码，该密码由SetWritePassword函数设定
//     address: 存储器地址（0-511），共512个字节
//     str: 要写入的字符串
// 返回值:
//     1: 成功
//     0: 失败
//    -1: 密码错误
//    -2: 地址无效
//    -3: 设备不支持该接口（AP-KM01）
func WriteString(pwd string, address int, str string) int {
	ret, _ := writeString(pwd, address, str)
	return int(ret)
}

// SetStringKey 设置算法密钥，用于TEA算法的加密、解密
// 入参:
//     pwd: 写密码，该密码由SetWritePassword函数设定
//     key: 密钥，最大长度8字节
// 返回值:
//     1: 成功
//     0: 失败
//    -1: 密码错误
//    -3: 设备不支持该接口（AP-KM01）
func SetStringKey(pwd, key string) int {
	ret, _ := setStringKey(pwd, key)
	return int(ret)
}

// EncString 加密字符串，返回加密后的字符串，待加密的字符串长度最大为8字节，加密后的字符串长度为16个字节
func EncString(str string) string {
	ret, _ := encString(str)
	return toString(ret)
}

// DecString 解密字符串，待解密字符串为用EncString函数加密后的串，长度固定为16字节（十六进制）
func DecString(str string) string {
	ret, _ := decString(str)
	return toString(ret)
}

// KeyDown 键盘单个键按下，键码表请参考1.4 USB键盘
func KeyDown(key string) bool {
	ret, _ := keyDown(key)
	return toBool(ret)
}

// KeyUp 键盘键弹起，键码表请参考1.4 USB键盘
func KeyUp(key string) bool {
	ret, _ := keyUp(key)
	return toBool(ret)
}

// KeyPress 键盘单次按键，键码表请参考1.4 USB键盘
func KeyPress(key string, count int) bool {
	ret, _ := keyPress(key, count)
	return ret > 0
}

// CombinationKeyDown 键盘组合键按下
func CombinationKeyDown(key1, key2, key3, key4, key5, key6 string) bool {
	ret, _ := combinationKeyDown(key1, key2, key3, key4, key5, key6)
	return toBool(ret)
}

// CombinationKeyUp 键盘组合键弹起
func CombinationKeyUp(key1, key2, key3, key4, key5, key6 string) bool {
	ret, _ := combinationKeyUp(key1, key2, key3, key4, key5, key6)
	return toBool(ret)
}

// CombinationKeyPress 键盘组合键弹起
func CombinationKeyPress(key1, key2, key3, key4, key5, key6 string, count int) bool {
	ret, _ := combinationKeyPress(key1, key2, key3, key4, key5, key6, count)
	return ret > 0
}

// SimulationPressKey 模拟人工输入字符串
func SimulationPressKey(keys string) bool {
	ret, _ := simulationPressKey(keys)
	return ret > 0
}

// KeyUpAll 所有键盘按键弹起，在复杂的代码中容易忘记哪些键已经按下，该函数可将所有按键都弹起
func KeyUpAll() bool {
	ret, _ := keyUpAll()
	return toBool(ret)
}

// GetCapsLock 获取键盘大定灯状态
// 返回值:
//     1: CapsLock灯熄灭
//     2: CapsLock灯亮
//     0: 失败
func GetCapsLock() int {
	ret, _ := getCapsLock()
	return int(ret)
}

// GetNumLock 获取小键盘灯状态
// 返回值:
//     1: NumLock灯熄灭
//     2: NumLock灯亮
//     0: 失败
func GetNumLock() int {
	ret, _ := getNumLock()
	return int(ret)
}

// LeftDown 鼠标左键按下
func LeftDown() bool {
	ret, _ := leftDown()
	return toBool(ret)
}

// LeftUp 鼠标左键弹起
func LeftUp() bool {
	ret, _ := leftUp()
	return toBool(ret)
}

// LeftClick 鼠标左键单击
func LeftClick(count int) bool {
	ret, _ := leftClick(count)
	return ret > 0
}

// LeftDoubleClick 鼠标左键双击
func LeftDoubleClick(count int) bool {
	ret, _ := leftDoubleClick(count)
	return ret > 0
}

// RightDown 鼠标右键按下
func RightDown() bool {
	ret, _ := rightDown()
	return toBool(ret)
}

// RightUp 鼠标右键弹起
func RightUp() bool {
	ret, _ := rightUp()
	return toBool(ret)
}

// RightClick 鼠标右键单击
func RightClick(count int) bool {
	ret, _ := rightClick(count)
	return ret > 0
}

// RightDoubleClick 鼠标右键双击
func RightDoubleClick(count int) bool {
	ret, _ := rightDoubleClick(count)
	return ret > 0
}

// MiddleDown 鼠标中键按下
func MiddleDown() bool {
	ret, _ := middleDown()
	return toBool(ret)
}

// MiddleUp 鼠标中键弹起
func MiddleUp() bool {
	ret, _ := middleUp()
	return toBool(ret)
}

// MiddleClick 鼠标中键单击
func MiddleClick(count int) bool {
	ret, _ := middleClick(count)
	return ret > 0
}

// MiddleDoubleClick 鼠标中键双击
func MiddleDoubleClick(count int) bool {
	ret, _ := middleDoubleClick(count)
	return ret > 0
}

// MouseUpAll 所有鼠标按键弹起
func MouseUpAll() bool {
	ret, _ := mouseUpAll()
	return toBool(ret)
}

// SimulationMove 模拟人工移动鼠标
func SimulationMove(x, y int) bool {
	ret, _ := simulationMove(x, y)
	return ret > 0
}

// MoveTo 模拟人工移动鼠标
// 在幽灵键鼠集成开发环境里支持使用别名MoveTo
func MoveTo(x, y int) bool {
	ret, _ := moveTo(x, y)
	return ret > 0
}

// MoveToR 相对移动鼠标
func MoveToR(x, y int8) bool {
	ret, _ := moveToR(x, y)
	return ret > 0
}

// GetKM21Mode 获取KM21工作模式，KM21出厂时默认单头模式
// 返回值:
//     1: 单头模式
//     0: 双头模式
//    -1: 获取模式失败（可能没有连接设备）
func GetKM21Mode() int {
	ret, _ := getKM21Mode()
	return int(ret)
}

// SetKM21Mode 设置KM21工作模式
// 入参:
//     model: 工作模式，1表示单头，0表示双头
func SetKM21Mode(model int) bool {
	ret, _ := setKM21Mode(model)
	return toBool(ret)
}

// MoveToFrom 移动鼠标（指定鼠标坐标）
// 入参:
//     fx, fy : 受控端鼠标的当前坐标
//     tx, ty : 要移到的目标坐标
func MoveToFrom(fx, fy, tx, ty int) bool {
	ret, _ := moveToFrom(fx, fy, tx, ty)
	return ret > 0
}

// ReMoveTo 复位移动鼠标
// 入参:
//     x, y : 要移到的目标坐标
func ReMoveTo(x, y int) bool {
	ret, _ := reMoveTo(x, y)
	return ret > 0
}

// Delay 延时等待
// 入参:
//     ms: 延时的毫秒数（1秒=1000毫秒）
func Delay(ms int) {
	_, _ = delay(ms)
}

// RandomDelay 随机延时
// 入参:
//     min: 延时的最小时间，单位：毫秒
//     max: 延时的最大时间，单位：毫秒
// 返回值:
//     实际延时的毫秒数
func RandomDelay(min, max int) int {
	ret, _ := randomDelay(min, max)
	return int(ret)
}

// RunApp 运行指定的程序
// 入参:
//     filepath: 要运行的程序路径和名称
// 返回值:
//   >31: 成功
//     0: 内存不足
//     2: 文件不存在
//     3: 路径不存在
//    11: 文件格式错误
func RunApp(filepath string) int {
	ret, _ := runApp(filepath)
	return int(ret)
}

// GetCursorPos 获取鼠标当前位置
// 入参:
//     x: 返回鼠标的水平坐标
//     y: 返回鼠标的垂直坐标
func GetCursorPos(x, y *int) bool {
	ret, _ := getCursorPos(x, y)
	return toBool(ret)
}

// MouseX 获取鼠标水平坐标
func MouseX() int {
	ret, _ := mouseX()
	return int(ret)
}

// MouseY 获取鼠标垂直坐标
func MouseY() int {
	ret, _ := mouseY()
	return int(ret)
}

// ScreenHeight 获取屏幕高度
func ScreenHeight() int {
	ret, _ := screenHeight()
	return int(ret)
}

// ScreenWidth 获取屏幕宽度
func ScreenWidth() int {
	ret, _ := screenWidth()
	return int(ret)
}

// Now 获取系统时间
func Now() string {
	ret, _ := now()
	return toString(ret)
}

// MessageBox 消息对话框
func MessageBox(msg string) bool {
	ret, _ := messageBox(msg)
	return toBool(ret)
}
