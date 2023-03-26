//
//  Utility.swift
//  mycmd
//
//  Created by pcl on 3/25/23.
//

import Foundation

func totp_mytotp(_ secret: String) -> String {
  //   let str = "AICRSHHFUHB2XGSHLO6QSNDMJYPIUKQC" //coinex secret
  let cStr = mytotp(UnsafeMutablePointer<Int8>(mutating: (secret as NSString).utf8String))
  return String(cString: cStr!, encoding: String.Encoding.utf8)!
}

func ucloud_listVM(_ limit:Int, _ offset:Int) -> String {
  let cStr = listVM(GoInt(limit), GoInt(offset))
  return String(cString: cStr!, encoding: String.Encoding.utf8)!
}
