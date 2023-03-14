//
//  ViewController.swift
//  mycmd
//
//  Created by pcl on 3/12/23.
//

import UIKit
import WebKit

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
    NSLog("com.gg.mycmd.log: %@", "fdasfasfasfasfasdffasa")
    //        log stream --process mycmd --style syslog | grep "com.gg.mycmd.log:"
    //        view.backgroundColor = UIColor.black
    //        view.backgroundColor = UIColor.systemBackground
    self.view.backgroundColor = UIColor.systemCyan
    // setScrollView1()
    //         listFont()
    // setTextView()
    sampleData()
    setMyView()
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
    textField.font = UIFont(name: "Consolas", size: 24)
    textField.attributedPlaceholder = NSAttributedString(
      string: "input command here ...",
      attributes: [NSAttributedString.Key.foregroundColor: UIColor.systemCyan])
    textField.addTarget(
      self, action: #selector(resignFirstResponder), for: UIControl.Event.editingDidEndOnExit)
    self.view.addSubview(textField)

    runButton.frame = CGRect(x: screenWidth - 70, y: 50, width: 60, height: 50)
    runButton.backgroundColor = UIColor.systemBlue
    runButton.setTitle("RUN", for: UIControl.State.normal)
    runButton.setTitleColor(UIColor.black, for: UIControl.State.normal)
    runButton.backgroundColor = UIColor.systemBackground
    runButton.addTarget(self, action: #selector(pressRun), for: UIControl.Event.touchUpInside)
    runButton.layer.borderWidth = 1
    runButton.layer.borderColor = UIColor.black.cgColor
    runButton.layer.cornerRadius = 8
    self.view.addSubview(runButton)

    saveButton.frame = CGRect(x: 10, y: 110, width: 60, height: 50)
    saveButton.backgroundColor = UIColor.systemBlue
    saveButton.setTitle("SAVE", for: UIControl.State.normal)
    saveButton.setTitleColor(UIColor.black, for: UIControl.State.normal)
    saveButton.backgroundColor = UIColor.systemBackground
    saveButton.addTarget(self, action: #selector(pressSave), for: UIControl.Event.touchUpInside)
    saveButton.layer.borderWidth = 1
    saveButton.layer.borderColor = UIColor.black.cgColor
    saveButton.layer.cornerRadius = 8
    self.view.addSubview(saveButton)

    historyButton.frame = CGRect(x: 80, y: 110, width: 60, height: 50)
    historyButton.backgroundColor = UIColor.systemBlue
    historyButton.setTitle("HISTORY", for: UIControl.State.normal)
    historyButton.setTitleColor(UIColor.black, for: UIControl.State.normal)
    historyButton.titleLabel?.font = UIFont(name: "Consolas", size: 12)
    historyButton.backgroundColor = UIColor.systemBackground
    historyButton.addTarget(
      self, action: #selector(pressHistory), for: UIControl.Event.touchUpInside)
    historyButton.layer.borderWidth = 1
    historyButton.layer.borderColor = UIColor.black.cgColor
    historyButton.layer.cornerRadius = 8
    self.view.addSubview(historyButton)

    textView.frame = CGRect(x: 10, y: 200, width: screenWidth - 20, height: 700)
    textView.backgroundColor = UIColor.black
    textView.textColor = UIColor.white
    textView.font = UIFont(name: "Consolas", size: 24)
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
    //        textField.resignFirstResponder()
    textView.text = textView.text + "\n" + textField.text!
    //        textView.text = String(cString: getipnum(nil))
    //        textView.text = String(getipnum())
    //        textView.text = String(cString: listVM(100, 0))
    // gotest()
    //        textView.text = getData()
  }

  @objc func pressSave() {
    //        saveData(textField.text!)
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
    historyVC.delegate = self  // used for historyVCDelegate
    //        let navController = UINavigationController(rootViewController: historyVC)
    //        historyVC.navigationItem.rightBarButtonItem = UIBarButtonItem(barButtonSystemItem: UIBarButtonItem.SystemItem.done, target: self, action: #selector(dismissHistory))
    //        historyVC.navigationItem.leftBarButtonItem = UIBarButtonItem(barButtonSystemItem: UIBarButtonItem.SystemItem.cancel, target: self, action: #selector(dismissHistory))

    self.present(historyVC, animated: true, completion: nil)
  }

  @objc func dismissHistory() {
    self.dismiss(animated: true, completion: nil)
  }

  func sampleData() {
    var cmdArray0: [String] = []
    for i in 0...99 {
      cmdArray0.append("value ucloud" + String(i))
    }
    UserDefaults.standard.set(cmdArray0, forKey: "mycmdKey0")

    var cmdArray1: [String] = []
    for i in 0...50 {
      cmdArray1.append("value google drive" + String(i))
    }
    UserDefaults.standard.set(cmdArray1, forKey: "mycmdKey1")

    var cmdArray2: [String] = []
    for i in 0...60 {
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

  weak var delegate: historyVCDelegate? = nil

  //    private var myArray: NSArray = ["First","Second","Third"]
  var myTableView: UITableView!
  //    var cmdArray = ["one", "two", "three"]
  var scIndex = 0

  override func viewDidLoad() {
    super.viewDidLoad()
    self.view.backgroundColor = UIColor.systemBackground

    let mySC = UISegmentedControl(items: ["UCloud", "Google Drive", "IMAP"])
    mySC.frame = CGRect(x: 50, y: 20, width: 300, height: 30)
    mySC.selectedSegmentIndex = scIndex
    mySC.tintColor = UIColor.yellow
    mySC.backgroundColor = UIColor.systemGray
    mySC.addTarget(self, action: #selector(self.segmentedValueChanged(_:)), for: .valueChanged)
    self.view.addSubview(mySC)

    myTableView = UITableView(
      frame: CGRect(x: 0, y: 60, width: self.view.frame.width, height: self.view.frame.height))
    myTableView.register(UITableViewCell.self, forCellReuseIdentifier: "MyCell")
    myTableView.dataSource = self
    myTableView.delegate = self
    self.view.addSubview(myTableView)
  }

  @objc func segmentedValueChanged(_ sender: UISegmentedControl!) {
    print("Selected Segment Index is : \(sender.selectedSegmentIndex)")
    scIndex = sender.selectedSegmentIndex
    myTableView.reloadData()
  }

  func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
    NSLog("com.gg.mycmd.log: didSelectRowAt %d", indexPath.row)
    //        print("Num: \(indexPath.row)")
    //        print("Value: \(myArray[indexPath.row])")
    delegate?.userDidEnterInformation(
      info: (UserDefaults.standard.array(forKey: "mycmdKey" + String(scIndex))![indexPath.row]
        as? String)!)
    self.dismiss(animated: true)
  }

  func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
    //        return myArray.count
    NSLog(
      "com.gg.mycmd.log: keys count: %d",
      UserDefaults.standard.dictionaryRepresentation().keys.count)
    //        return cmdArray.count
    return UserDefaults.standard.array(forKey: "mycmdKey" + String(scIndex))!.count
  }

  func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
    let cell = tableView.dequeueReusableCell(withIdentifier: "MyCell", for: indexPath as IndexPath)
    //                cell.textLabel!.text = "\(myArray[indexPath.row])"
    cell.textLabel!.text =
      UserDefaults.standard.array(forKey: "mycmdKey" + String(scIndex))![indexPath.row] as? String
    return cell
  }

  func tableView(_ tableView: UITableView, canEditRowAt indexPath: IndexPath) -> Bool {
    return true
  }

  func tableView(
    _ tableView: UITableView, commit editingStyle: UITableViewCell.EditingStyle,
    forRowAt indexPath: IndexPath
  ) {
    if editingStyle == UITableViewCell.EditingStyle.delete {
      NSLog("com.gg.mycmd.log: delete row: %d", indexPath.row)

      let alert = UIAlertController(
        title: nil, message: "Are you sure you'd like to delete this cell",
        preferredStyle: UIAlertController.Style.alert)

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
}
