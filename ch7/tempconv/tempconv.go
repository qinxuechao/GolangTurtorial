package tempconv

// *celsiusFlag 满足 flag.Value 接口
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var
}