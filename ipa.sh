cd mycgo
CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 CC="/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/clang -arch arm64 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/iPhoneOS.sdk" go build -buildmode=c-archive -tags ios -o mycgo.a ./main/
cd ..
xcodebuild build CODE_SIGN_IDENTITY="" CODE_SIGNING_REQUIRED=NO -sdk iphoneos16.2 -target mycmd
rm -rf Payload
mkdir Payload
cp -r build/Release-iphoneos/mycmd.app Payload
zip -qr mycmd.zip Payload
mv mycmd.zip mycmd.ipa
