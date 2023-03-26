//
//  ViewController.swift
//  mycmd
//
//  Created by pcl on 3/12/23.
//

import UIKit

class ViewController: UIViewController, historyVCDelegate {
  let screenWidth = UIScreen.main.bounds.size.width
  let screenHeight = UIScreen.main.bounds.size.height

  let textView = UITextView()
  let textField = UITextField()
  let runButton = UIButton(type: UIButton.ButtonType.system)
  let saveButton = UIButton(type: UIButton.ButtonType.system)
  let historyButton = UIButton(type: UIButton.ButtonType.system)

  override func viewDidLoad() {
    super.viewDidLoad()
    NSLog("com.gg.mycmd.log: %@", "viewDidLoad()")
    //        log stream --process mycmd --style syslog | grep "com.gg.mycmd.log:"
    //        view.backgroundColor = UIColor.black
    //        view.backgroundColor = UIColor.systemBackground
    view.backgroundColor = UIColor.systemCyan
    setMyView()

    // sampleData()
    initData()
    createFile()
    createFolder()
    readFile()
  }

  func listFiles() {
    do {
      let urls = try FileManager.default.contentsOfDirectory(at: FileManager.default.urls(for: FileManager.SearchPathDirectory.documentDirectory, in: FileManager.SearchPathDomainMask.userDomainMask)[0], includingPropertiesForKeys: nil)
      for url in urls {
        NSLog("com.gg.mycmd.log: list files: %@", url.absoluteString)
        NSLog("com.gg.mycmd.log: lastPathComponent: %@", url.lastPathComponent)
      }
    } catch {
      NSLog("com.gg.mycmd.log: error: %@", error.localizedDescription)
    }
  }

  func readFile() {
    do {
      let urls = try FileManager.default.contentsOfDirectory(at: FileManager.default.urls(for: FileManager.SearchPathDirectory.documentDirectory, in: FileManager.SearchPathDomainMask.userDomainMask)[0], includingPropertiesForKeys: nil)
      for url in urls {
        NSLog("com.gg.mycmd.log: list files: %@", url.absoluteString)
        if url.lastPathComponent == "output.txt" {
          let output = try String(contentsOf: url, encoding: String.Encoding.utf8)
          textView.text = output
          NSLog("com.gg.mycmd.log: list files: %@", output)
        }
      }
    } catch {
      NSLog("com.gg.mycmd.log: error: %@", error.localizedDescription)
    }
  }

  func createFile() {
    let str = "Super long string here"
    let filename = FileManager.default.urls(for: FileManager.SearchPathDirectory.documentDirectory, in: FileManager.SearchPathDomainMask.userDomainMask)[0].appendingPathComponent("output.txt")
    NSLog("com.gg.mycmd.log: %@", filename.absoluteString)
    do {
      try str.write(to: filename, atomically: true, encoding: String.Encoding.utf8)
    } catch {
      NSLog("com.gg.mycmd.log: error: %@", error.localizedDescription)
      // failed to write file – bad permissions, bad filename, missing permissions, or more likely it can't be converted to the encoding
    }
  }

  func createFolder() {
    let manager = FileManager.default
    let DecomentFolder = manager.urls(for: FileManager.SearchPathDirectory.documentDirectory, in: FileManager.SearchPathDomainMask.userDomainMask).last
    let Folder = DecomentFolder?.appendingPathComponent("xxx")
    do {
      try manager.createDirectory(at: Folder!, withIntermediateDirectories: true, attributes: [:])
    } catch {
      NSLog("com.gg.mycmd.log: %@", error.localizedDescription)
      print(error.localizedDescription)
    }
  }

  override var prefersStatusBarHidden: Bool {
    return false
  }

  func gotest() {
    // let str1 = reverse(UnsafeMutablePointer<Int8>(mutating: (textField.text! as NSString).utf8String))
    // textView.text = String(cString: str1!, encoding: .utf8)!
    //        getipnum(nil)
  }

  func listFont() {
    for family in UIFont.familyNames {
      for font in UIFont.fontNames(forFamilyName: family) {
        NSLog("com.gg.mycmd.log: family: %@ font: %@", family, font)
      }
    }
  }

  func setMyView() {
    textField.frame = CGRect(x: 10, y: 50, width: screenWidth - 20 - 60 - 10, height: 50)
    textField.backgroundColor = UIColor.black
    //        textField.backgroundColor = UIColor.systemCyan
    textField.textColor = UIColor.white
    textField.font = UIFont(name: "Consolas", size: 20)
    textField.attributedPlaceholder = NSAttributedString(
      string: "input command here ...",
      attributes: [NSAttributedString.Key.foregroundColor: UIColor.systemCyan]
    )
    textField.addTarget(
      self, action: #selector(resignFirstResponder), for: UIControl.Event.editingDidEndOnExit
    )
    view.addSubview(textField)

    runButton.frame = CGRect(x: screenWidth - 70, y: 50, width: 60, height: 50)
    runButton.backgroundColor = UIColor.systemBlue
    runButton.setTitle("RUN", for: UIControl.State.normal)
    runButton.setTitleColor(UIColor.black, for: UIControl.State.normal)
    runButton.addTarget(self, action: #selector(pressRun), for: UIControl.Event.touchUpInside)
    runButton.layer.borderWidth = 1
    runButton.layer.borderColor = UIColor.black.cgColor
    runButton.layer.cornerRadius = 8
    view.addSubview(runButton)

    saveButton.frame = CGRect(x: 10, y: 110, width: 60, height: 50)
    saveButton.backgroundColor = UIColor.systemBlue
    saveButton.setTitle("SAVE", for: UIControl.State.normal)
    saveButton.setTitleColor(UIColor.black, for: UIControl.State.normal)
    saveButton.addTarget(self, action: #selector(pressSave), for: UIControl.Event.touchUpInside)
    saveButton.layer.borderWidth = 1
    saveButton.layer.borderColor = UIColor.black.cgColor
    saveButton.layer.cornerRadius = 8
    view.addSubview(saveButton)

    historyButton.frame = CGRect(x: 80, y: 110, width: 60, height: 50)
    historyButton.backgroundColor = UIColor.systemBlue
    historyButton.setTitle("HISTORY", for: UIControl.State.normal)
    historyButton.setTitleColor(UIColor.black, for: UIControl.State.normal)
    historyButton.titleLabel?.font = UIFont(name: "Consolas", size: 12)
    historyButton.addTarget(
      self, action: #selector(pressHistory), for: UIControl.Event.touchUpInside
    )
    historyButton.layer.borderWidth = 1
    historyButton.layer.borderColor = UIColor.black.cgColor
    historyButton.layer.cornerRadius = 8
    view.addSubview(historyButton)

    textView.frame = CGRect(x: 10, y: 200, width: screenWidth - 20, height: 700)
    textView.backgroundColor = UIColor.black
    textView.textColor = UIColor.white
    textView.font = UIFont(name: "Consolas", size: 20)
    textView.isEditable = false
    textView.isScrollEnabled = true
    // textView.text = """
    // This goes
    // over multiple
    // lines
    // This goes
    // over multiple
    // """
    view.addSubview(textView)
  }

  @objc func pressRun() {
    textView.text = ""
    let cmds = textField.text?.components(separatedBy: " ")
    switch cmds![0] {
    case "totp":
      textView.text = totp_mytotp(cmds![1])
    case "ping":
//          textView.text = myicmp_myping(cmds![1])
      let once = try? SwiftyPing(host: cmds![1], configuration: PingConfiguration(interval: 0.5, with: 5), queue: DispatchQueue.global())
      once?.observer = { response in
        self.textView.text += String(response.byteCount!) + " bytes" + " from "
        self.textView.text += response.ipAddress!
        self.textView.text += ": time=" + String(format: "%.3f", response.duration * 1000) + " ms\n"
      }
      once?.targetCount = 4
      try? once?.startPinging()
    default:
      print("no cmds")
    }
  }

  @objc func pressSave() {
//            saveData(textField.text!)
    var cmdArray = UserDefaults.standard.array(forKey: "mycmdKey2")
    cmdArray?.insert(textField.text as Any, at: 0)
    UserDefaults.standard.set(cmdArray, forKey: "mycmdKey2")

//      NSLog("com.gg.mycmd.log: saved cmd: %@", UserDefaults.standard.array(forKey: "mycmdKey2")![0] as! String)
  }

  //    @objc func pressHistory() {
  //        let mywebViewController = UIViewController()
  //        let webView = WKWebView(frame: CGRect.zero, configuration: WKWebViewConfiguration())
  //        webView.load(URLRequest(url: URL(string: "https://www.apple.com")!)) // url from the row a user taps
  //        mywebViewController.view = webView
  //        //        self.present(mywebViewController, animated: true, completion: nil)
  //
  //        let navController = UINavigationController(rootViewController: mywebViewController)
  //        mywebViewController.navigationItem.rightBarButtonItem = UIBarButtonItem(barButtonSystemItem: UIBarButtonItem.SystemItem.done, target: self, action: #selector(dismissHistory))
  //        mywebViewController.navigationItem.leftBarButtonItem = UIBarButtonItem(barButtonSystemItem: UIBarButtonItem.SystemItem.cancel, target: self, action: #selector(dismissHistory))
  //
  //        self.present(navController, animated: true, completion: nil)
  //    }

  @objc func pressHistory() {
    let historyVC = historyVC()
    historyVC.delegate = self // used for historyVCDelegate
    //        let navController = UINavigationController(rootViewController: historyVC)
    //        historyVC.navigationItem.rightBarButtonItem = UIBarButtonItem(barButtonSystemItem: UIBarButtonItem.SystemItem.done, target: self, action: #selector(dismissHistory))
    //        historyVC.navigationItem.leftBarButtonItem = UIBarButtonItem(barButtonSystemItem: UIBarButtonItem.SystemItem.cancel, target: self, action: #selector(dismissHistory))

    present(historyVC, animated: true, completion: nil)
  }

  @objc func dismissHistory() {
    dismiss(animated: true, completion: nil)
  }

  func initData() {
    let cmdArray0: [String] = []
    UserDefaults.standard.set(cmdArray0, forKey: "mycmdKey0")

    let cmdArray1: [String] = []
    UserDefaults.standard.set(cmdArray1, forKey: "mycmdKey1")

    let cmdArray2: [String] = []
    UserDefaults.standard.set(cmdArray2, forKey: "mycmdKey2")
  }

  func sampleData() {
    var cmdArray0: [String] = []
    for i in 0 ... 99 {
      cmdArray0.append("value ucloud" + String(i))
    }
    UserDefaults.standard.set(cmdArray0, forKey: "mycmdKey0")

    var cmdArray1: [String] = []
    for i in 0 ... 50 {
      cmdArray1.append("value google drive" + String(i))
    }
    UserDefaults.standard.set(cmdArray1, forKey: "mycmdKey1")

    var cmdArray2: [String] = []
    for i in 0 ... 60 {
      cmdArray2.append("value imap" + String(i))
    }
    UserDefaults.standard.set(cmdArray2, forKey: "mycmdKey2")
  }

  //    func saveData(_ value: String) {
  //        UserDefaults.standard.set(value, forKey: "key1")
  //    }
  //
  //    func getData() -> String{
  //        return UserDefaults.standard.string(forKey: "key1")!
  //    }

  func userDidEnterInformation(info: String) {
    NSLog("com.gg.mycmd.log: userDidEnterInformation")
    textField.text = info
  }
}

protocol historyVCDelegate: AnyObject {
  func userDidEnterInformation(info: String)
}

class historyVC: UIViewController, UITableViewDelegate, UITableViewDataSource {
  weak var delegate: historyVCDelegate?

  //    private var myArray: NSArray = ["First","Second","Third"]
  var myTableView: UITableView!
  //    var cmdArray = ["one", "two", "three"]
  var scIndex = 0

  override func viewDidLoad() {
    super.viewDidLoad()
    view.backgroundColor = UIColor.systemBackground

    let mySC = UISegmentedControl(items: ["UCloud", "Google Drive", "IMAP"])
    mySC.frame = CGRect(x: 50, y: 20, width: 300, height: 30)
    mySC.selectedSegmentIndex = scIndex
    mySC.tintColor = UIColor.yellow
    mySC.backgroundColor = UIColor.systemGray
    mySC.addTarget(self, action: #selector(segmentedValueChanged(_:)), for: .valueChanged)
    view.addSubview(mySC)

    myTableView = UITableView(
      frame: CGRect(x: 0, y: 60, width: view.frame.width, height: view.frame.height))
    myTableView.register(UITableViewCell.self, forCellReuseIdentifier: "MyCell")
    myTableView.dataSource = self
    myTableView.delegate = self
    // self.myTableView.isEditing = true
    view.addSubview(myTableView)
  }

  @objc func segmentedValueChanged(_ sender: UISegmentedControl!) {
    print("Selected Segment Index is : \(sender.selectedSegmentIndex)")
    scIndex = sender.selectedSegmentIndex
    myTableView.reloadData()
  }

  func tableView(_: UITableView, didSelectRowAt indexPath: IndexPath) {
    NSLog("com.gg.mycmd.log: didSelectRowAt %d", indexPath.row)
    //        print("Num: \(indexPath.row)")
    //        print("Value: \(myArray[indexPath.row])")
    delegate?.userDidEnterInformation(
      info: (UserDefaults.standard.array(forKey: "mycmdKey" + String(scIndex))![indexPath.row]
        as? String)!)
    dismiss(animated: true)
  }

  func tableView(_: UITableView, numberOfRowsInSection _: Int) -> Int {
    //        return myArray.count
    NSLog(
      "com.gg.mycmd.log: keys count: %d",
      UserDefaults.standard.dictionaryRepresentation().keys.count
    )
    //        return cmdArray.count
    return UserDefaults.standard.array(forKey: "mycmdKey" + String(scIndex))?.count ?? 0
  }

  func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
    let cell = tableView.dequeueReusableCell(withIdentifier: "MyCell", for: indexPath as IndexPath)
    //                cell.textLabel!.text = "\(myArray[indexPath.row])"
    cell.textLabel!.text =
      UserDefaults.standard.array(forKey: "mycmdKey" + String(scIndex))![indexPath.row] as? String
    return cell
  }

  func tableView(_: UITableView, canEditRowAt _: IndexPath) -> Bool {
    return true
  }

  func tableView(
    _: UITableView, commit editingStyle: UITableViewCell.EditingStyle,
    forRowAt indexPath: IndexPath
  ) {
    if editingStyle == UITableViewCell.EditingStyle.delete {
      NSLog("com.gg.mycmd.log: delete row: %d", indexPath.row)

      let alert = UIAlertController(
        title: nil, message: "Are you sure you'd like to delete this cell",
        preferredStyle: UIAlertController.Style.alert
      )

      alert.addAction(
        UIAlertAction(title: "Yes", style: UIAlertAction.Style.default) { _ in
          var cmdArray = UserDefaults.standard.array(forKey: "mycmdKey" + String(self.scIndex))
          cmdArray?.remove(at: indexPath.row)
          UserDefaults.standard.set(cmdArray, forKey: "mycmdKey" + String(self.scIndex))
          self.myTableView.beginUpdates()
          self.myTableView.deleteRows(at: [indexPath], with: UITableView.RowAnimation.automatic)
          self.myTableView.endUpdates()
        })
      alert.addAction(
        UIAlertAction(title: "Cancel", style: UIAlertAction.Style.cancel, handler: nil))
      present(alert, animated: true, completion: nil)
    }
  }

  // func tableView(_ tableView: UITableView, canMoveRowAt indexPath: IndexPath) -> Bool {
  //   return true
  // }

  // func tableView(_ tableView: UITableView, editingStyleForRowAt indexPath: IndexPath)
  //   -> UITableViewCell.EditingStyle
  // {
  //   return UITableViewCell.EditingStyle.none
  //   // return UITableViewCell.EditingStyle.delete
  // }

  func tableView(_: UITableView, shouldIndentWhileEditingRowAt _: IndexPath)
    -> Bool
  {
    return false
  }

  func tableView(
    _: UITableView, moveRowAt sourceIndexPath: IndexPath,
    to destinationIndexPath: IndexPath
  ) {
    var cmdArray = UserDefaults.standard.array(forKey: "mycmdKey" + String(scIndex))
    let movedRow = cmdArray?[sourceIndexPath.row]
    cmdArray?.remove(at: sourceIndexPath.row)
    cmdArray?.insert(movedRow!, at: destinationIndexPath.row)
    UserDefaults.standard.set(cmdArray, forKey: "mycmdKey" + String(scIndex))
  }
}
