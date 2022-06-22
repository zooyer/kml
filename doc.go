/*
幽灵键鼠SDK封装（依赖kml.dll）
支持的键名：
	可见字符：
		大写A-Z
		小写a-z
		数字0-9
		符号：`-=[]\;',./~_+{}|:"<>?!@#$%^&*()空格
		以上字符皆可通过SimulationPressKey接口混合输入，且无需关心键盘大写灯状态
	不可见字符：
		Enter、Esc、Backspace、Tab、CapsLock、
		F1、F2、F3、F4、F5、F6、F7、F8、F9、F10、
		F11、F12、Num0~Num9、NumEnter、NumDot、
		NumAdd、NumDec、NumMul、NumDiv、
		Space、PrtSc、ScrLk、Pause、Insert、Delete、
		Home、End、PageUp、PageDown、
		Right、Left、Down、Up、NumLock、
		Ctrl、Shift、Alt、Win、（本行为左侧键）
		以上按键可通过除SimulationPressKey接口之外的键盘函数输入，键名以字符串形式输入，不区分大小写

特别说明：使用KeyDown、KeyUp、KeyPress、SimulationPressKey（SayString）接口按字母键时区分大小写，如果小写状态输入大写字母会自动按下Shift键以配合输入，反之亦然。输入上档字符时也会按下Shift键。
*/

package kml // Package kml import "github.com/zooyer/kml"
