//
//  HomeViewController.swift
//  sixfivezerozero
//
//  Created by Timothy on 16/2/11.
//  Copyright © 2016年 Shell-Labs. All rights reserved.
//

import UIKit

class FirstViewController: UIViewController {
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.navigationItem.title = "Home"
        
        let label = UILabel(frame: CGRect(x: 0, y: 80, width: 300, height: 35));
        label.text = "First Tabbar"
        self.view.addSubview(label)
        
    }
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }
}
