package main

import "C"

// other imports should be seperate from the special Cgo import
import (
	"fmt"
	"mycgo/myicmp"
	"mycgo/totp"
	"mycgo/ucloud"
)

//export reverse
// func reverse(in *C.char) *C.char {
// 	return C.CString(foo.Reverse(C.GoString(in)))
// }

//export listVM
func listVM(limit int, offset int) *C.char {
	return C.CString(ucloud.ListVM(limit, offset))
}

//export listImage
func listImage(limit int, offset int) *C.char {
	return C.CString(ucloud.ListImage(limit, offset))
}

//export listEIP
func listEIP(limit int, offset int) *C.char {
	return C.CString(ucloud.ListEIP(limit, offset))
}

//export showImage
func showImage(id *C.char) *C.char {
	return C.CString(ucloud.ShowImage(C.GoString(id)))
}

//export deleteEIP
func deleteEIP(id *C.char) *C.char {
	return C.CString(ucloud.DeleteEIP(C.GoString(id)))
}

//export getipnum
func getipnum() int {
	return ucloud.Getipnum()
}

//export mytotp
func mytotp(secret *C.char) *C.char {
	return C.CString(totp.Totp(C.GoString(secret)))
}

//export myping
func myping(address *C.char) *C.char {
	return C.CString(myicmp.Myping(C.GoString(address)))
}

func main() {
	fmt.Println("ggggg")
	// fmt.Println(myicmp.Myping("www.sohu.com"))
	// fmt.Println("getipnum(): ", getipnum())
	// totp.Totp("AICRSHHFUHB2XGSHLO6QSNDMJYPIUKQC")
}

//the "export getipnum" comment must be there, can you believe it?

// https://rogchap.com/2020/09/14/running-go-code-on-ios-and-android/

// build for ios simulator
// CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 CC="/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/clang -arch x86_64 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator.sdk" go build -buildmode=c-archive -tags ios -o mycgo.a ./main/

// build for linux
// CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildmode=c-archive -o mycgo.a ./main/

// build for ios
// CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 CC="/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/clang -arch arm64 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/iPhoneOS.sdk" go build -buildmode=c-archive -tags ios -o mycgo.a ./main/
