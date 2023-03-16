#log stream --process mycmd --style syslog | grep "com.gg.mycmd.log:"
#xcode-select --switch /Applications/Xcode.app/Contents/Developer
#xcodebuild -showsdks
xcrun simctl uninstall booted com.gg.mycmd
xcodebuild build CODE_SIGN_IDENTITY="" CODE_SIGNING_REQUIRED=NO -arch x86_64 -sdk iphonesimulator16.2 -target mycmd
xcrun simctl install booted build/Release-iphonesimulator/mycmd.app
xcrun simctl launch booted com.gg.mycmd
