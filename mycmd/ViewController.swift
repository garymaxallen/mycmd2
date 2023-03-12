//
//  ViewController.swift
//  mycmd
//
//  Created by pcl on 3/12/23.
//

import UIKit

class ViewController: UIViewController {
    
    let screenWidth = UIScreen.main.bounds.size.width
    let screenHeight = UIScreen.main.bounds.size.height
    
    let textView = UITextView()
    
    let textField = UITextField()
    
    override func viewDidLoad() {
        NSLog("com.gg.mycmd.log: %@", "fdasfasfasfasfasdffasa")
        super.viewDidLoad()
        view.backgroundColor = UIColor.black
        // Do any additional setup after loading the view.
        // setScrollView1()
        // listFont()
        // setTextView()
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
    
    func setMyView() {
        textField.frame = CGRect(x: 10, y: 100, width: 400.00, height: 50.00)
        textField.backgroundColor = UIColor.black
        textField.textColor = UIColor.white
        textField.font = UIFont(name: "Consolas", size: 24)
        textField.attributedPlaceholder = NSAttributedString(string: "input command here ...",
                                                             attributes: [NSAttributedString.Key.foregroundColor: UIColor.white])
        textField.addTarget(self, action: #selector(inputDone), for: UIControl.Event.editingDidEndOnExit)
        view.addSubview(textField)
        
        textView.frame = CGRect(x: 10, y: 200, width: 400.00, height: 700.00)
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
    
    @objc func inputDone() {
        textField.resignFirstResponder()
        textView.text = textField.text
        //        textView.text = String(cString: getipnum(nil))
        //        textView.text = String(getipnum())
        //        textView.text = String(cString: listVM(100, 0))
        // gotest()
    }
    
}

