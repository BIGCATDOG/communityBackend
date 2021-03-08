module backend

go 1.16

require github.com/valyala/gorpc v0.0.0-20160519171614-908281bef774
require github.com/untils v0.0.0-incompatible
replace (
	github.com/untils  => ../utils
)
