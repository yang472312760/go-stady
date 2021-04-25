package main

import (
	"fmt"
	"regexp"
)

func main() {

	context1 := "3.14 123123 .68 haha 1.0 abc 6.66 123."

	exp1 := regexp.MustCompile(`\d+\.\d+`)

	result1 := exp1.FindAllStringSubmatch(context1, -1)

	fmt.Printf("%v\n", result1)

	fmt.Printf("\n------------------------------------\n\n")

	context2 := `
				<title>这是一个标题</title>
				<div>这是一个div</div>
				<div>hello world</div>
				<div>这是第二个div</div>
				<body>这是一个body</body>
				`
	exp2 := regexp.MustCompile(`<div>(.*?)</div>`)

	result2 := exp2.FindAllStringSubmatch(context2, -1)

	fmt.Printf("%v\n", result2)
	fmt.Printf("\n------------------------------------\n\n")

	context3 := `
				<title>这是一个标题</title>
				<div>这是一个div</div>
				<div>hello 
				world
				</div>
				<div>这是第二个div</div>
				<body>这是一个body</body>
				`
	exp3 := regexp.MustCompile(`<div>(.*?)</div>`)
	result3 := exp3.FindAllStringSubmatch(context3, -1)

	fmt.Printf("%v\n", result3)
	fmt.Printf("\n------------------------------------\n\n")

	context4 := `
				<title>这是一个标题</title>
				<div>这是一个div</div>
				<div>hello 
				world
				</div>
				<div>这是第二个div</div>
				<body>这是一个body</body>
	`

	exp4 := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	result4 := exp4.FindAllStringSubmatch(context4, -1)

	fmt.Printf("%v\n", result4)
	fmt.Printf("\n------------------------------------\n\n")

	for _, text := range result4 {
		fmt.Println(text[0]) //带有div
		fmt.Println(text[1]) //不带带有div
		fmt.Println("================\n")
	}

}
