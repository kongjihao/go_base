package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("current time: %v\t || \t%T\n", now, now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	nano := now.Nanosecond() //纳秒
	fmt.Printf("year: %v | month: %v | day: %v | hour: %v | minute: %v | second: %v | nano: %v\n", year, month, day, hour, minute, second, nano)

	// 时间戳：1970.1.1 到现在经过了的毫秒数，绝对时间
	timestamp := now.Unix()
	fmt.Printf("timestamp: %v\t || \t%T\n", timestamp, timestamp) // timestamp: 1725803501	 || 	int64

	// 将时间戳转化为具体的时间格式
	time1 := time.Unix(timestamp, 0)
	fmt.Println(time1.Format("2006.01.02 15:04:05")) // format中传入你想要的时间格式的字符串

	// 时间间隔，单位为纳秒
	duration := time.Duration(10)
	fmt.Printf("duration: %v\t || \t%T\n", duration, duration)

	n := 5
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("sleep1 over")
	time.Sleep(5) // sleep 5 纳秒
	fmt.Println("sleep2 over")

	// now + 1小时
	t := now.Add(time.Hour * 1)
	fmt.Println("now + 1h:", t)

	// 计算两个时间的间隔,差值
	d := t.Sub(now)
	fmt.Println("time diff:", d)

	// Equal 判断两个时间是否相等
	fmt.Println(now.Equal(t))

	// 判断是否是闰年
	fmt.Println(time.Date(2019, 2, 28, 5, 5, 5, 0, time.Local).YearDay())

	// Before 判断 t 是否在 now 的之前
	fmt.Println(t.Before(now))

	// After 判断 t 是否在 now 的之后
	fmt.Println(t.After(now))

	// 定时器，每秒执行一次
	// ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	// for i := range ticker {          //for range 遍历通道
	// 	fmt.Println(i) //每秒都会执行的任务
	// }

	now = time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))

}
