package test

func Equal(a []byte, b []byte) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// TODO: 在自己家里那台电脑上弄，or入职后的新电脑上下载最新版本
// func Fuzz_x(f *testing.F) {

// }
