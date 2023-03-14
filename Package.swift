// swift-tools-version: 5.7
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let packageName = "mycmd"  // <-- Change this to yours
let package = Package(
  name: "mycmd",
  defaultLocalization: "en",
  platforms: [.iOS("16.2")],
  products: [
    .library(name: packageName, targets: [packageName])
  ],
  targets: [
    .target(
      name: packageName,
      path: packageName
    )
  ]
)
