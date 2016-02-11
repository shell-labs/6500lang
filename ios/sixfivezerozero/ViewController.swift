//
//  ViewController.swift
//  sixfivezerozero
//
//  Created by Timothy on 16/2/11.
//  Copyright © 2016年 Shell-Labs. All rights reserved.
//

import UIKit
import FontAwesome

class ViewController: UITabBarController {

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view, typically from a nib.
        // 背景颜色
        self.view.backgroundColor = UIColor.whiteColor()
        
        // 实例化ViewController
        let first = FirstViewController()
        let firstBar = RootNavigationBar()
        firstBar.viewControllers = [first]
        
        let second = SecondViewController()
        let secondBar = RootNavigationBar()
        secondBar.viewControllers = [second]
        
        let third = ThirdViewController()
        let thirdBar = RootNavigationBar()
        thirdBar.viewControllers = [third]
        
        let fourth = FourthViewController()
        let fourthBar = RootNavigationBar()
        fourthBar.viewControllers = [fourth]
        
        self.viewControllers = [firstBar, secondBar, thirdBar, fourthBar]
        
        // Tabbar颜色
        UITabBar.appearance().tintColor = UIColor.init(colorLiteralRed: 0, green: 1, blue: 2, alpha: 1);
        
        //设置tabbar的不同按钮的显示方式,这里通过tabBar的属性items来获取不同的UITabBarItem,且对应到主面板的排列顺序.
        let firstItemBar =  self.tabBar.items![0]
        firstItemBar.title = "Home"
        firstItemBar.image = UIImage.fontAwesomeIconWithName(FontAwesome.Home, textColor: UIColor.blackColor(), size: CGSizeMake(30,30))
        
        let secondbar = self.tabBar.items![1]
        secondbar.title = "Memory"
        secondbar.image = UIImage.fontAwesomeIconWithName(FontAwesome.Tasks, textColor: UIColor.blackColor(), size: CGSizeMake(30,30))
        
        let thirdbar = self.tabBar.items![2]
        thirdbar.title = "Discovery"
        thirdbar.image = UIImage.fontAwesomeIconWithName(FontAwesome.Globe, textColor: UIColor.blackColor(), size: CGSizeMake(30,30))
        
        let fourthbar = self.tabBar.items![3]
        fourthbar.title = "Profile"
        fourthbar.image = UIImage.fontAwesomeIconWithName(FontAwesome.User, textColor: UIColor.blackColor(), size: CGSizeMake(30,30))
        
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }


}

