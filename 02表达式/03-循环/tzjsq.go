package main

import "fmt"

func main() {
	for	{
		var name string
		fmt.Print("姓名：")
		fmt.Scanln(&name)

		var age int
		fmt.Print("年龄：")
		fmt.Scanln(&age)

		var weight float64
		fmt.Print("体重（kg）：")
		fmt.Scanln(&weight)

		var tall float64
		fmt.Print("身高（m）：")
		fmt.Scanln(&tall)

		var sexWeight int
		var sex string
		fmt.Print("性别（男/女）：")
		fmt.Scanln(&sex)
		if sex == "男" {
			sexWeight = 1
		} else if  sex == "女"{
			sexWeight = 0
		} else {
			fmt.Println("输入有误")
		}

		var bmi float64 = weight / (tall * tall)
		var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
		fmt.Println("您的体脂率是：", fatRate)

		//建议
		if sex == "男" {
			//todo 男性体脂率建议表
			if age >= 18 && age <= 39 {
				if fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦")
				}else if fatRate > 0.1 && fatRate <= 0.16{
					fmt.Println("目前是：正常")
				}else if fatRate > 0.16 && fatRate <= 0.26{
					fmt.Println("目前是：偏胖")
				}else {
					fmt.Println("目前是：不健康")
				}
			}else if age >= 40 && age <= 59{
				//todo
			}else if age >= 60{
				//todo
			}else{
				fmt.Println("未成年人，无法评测")
				return
			}
		}else {
			//todo 女性体脂率建议表
		}

		var whethContinue string
		fmt.Println("是否要录入下一个（y/n）: ")
		fmt.Scanln(&whethContinue)
		if whethContinue == "y" {
			//继续录入
		} else {
			break
		}
		
	}
}