#log stream --process mycmd --style syslog | grep "com.gg.mycmd.log:"
#xcode-select --switch /Applications/Xcode.app/Contents/Developer
#xcodebuild -showsdks
#xcrun simctl boot "iPhone 12 Pro Max"
cd mycgo
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 CC="/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/clang -arch x86_64 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator.sdk" go build -buildmode=c-archive -tags ios -o mycgo.a ./main/
cd ..
rm -rf build
xcrun simctl uninstall booted com.gg.mycmd
xcodebuild build CODE_SIGN_IDENTITY="" CODE_SIGNING_REQUIRED=NO -arch x86_64 -sdk iphonesimulator16.2 -target mycmd
#xcodebuild build CODE_SIGN_IDENTITY="" CODE_SIGNING_REQUIRED=NO -sdk iphoneos16.2 -target mycmd
xcrun simctl install booted build/Release-iphonesimulator/mycmd.app
xcrun simctl launch booted com.gg.mycmd

#archive .app to ipa
#cp -r build/Release-iphoneos/mycmd.app Payload
#zip -r mycmd.zip Payload
#mv mycmd.zip mycmd.ipa
#python3 -m http.server 80
