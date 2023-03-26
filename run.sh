#log stream --process mycmd --style syslog | grep "com.gg.mycmd.log:"
#xcode-select --switch /Applications/Xcode.app/Contents/Developer
#xcodebuild -showsdks
#xcrun simctl boot "iPhone 12 Pro Max"
xcrun simctl uninstall booted com.gg.mycmd
xcodebuild build CODE_SIGN_IDENTITY="" CODE_SIGNING_REQUIRED=NO -arch x86_64 -sdk iphonesimulator16.2 -target mycmd
#xcodebuild build CODE_SIGN_IDENTITY="" CODE_SIGNING_REQUIRED=NO -sdk iphoneos16.2 -target mycmd
xcrun simctl install booted build/Release-iphonesimulator/mycmd.app
xcrun simctl launch booted com.gg.mycmd

#archive .app to ipa
#cp -r build/Release-iphoneos/mycmd.app Payload
#zip -r mycmd.zip Payload
#mv mycmd.zip mycmd.ipa
#python3 -n http.server 80
