

/*附加题1方法是简单的双重循环，如果把最大次数视为N此的话，时间复杂度为n的平方，
当n为题给的123456时，时间复杂度约为10的12次方*/
/*优化方法如下，在判断的时候我们其实没有必要一直遍历判断到小于N的前一位整数为止，
由数学方法我们可知只要判断到它的根号时，仍未出现因数，那么该数就为素数，所以我们只要使用强转改变一下数据类型，
然后使用库函数sqrt进行开方作为判断条件，就能够大大优化该程序，优化后代码如下*/

package main

 import (
"fmt"
"time"
"math"
 )

var flag int

func primeNumber(a float64){
	var i float64
	var j float64
	for j=2;j <= a;j++{
		flag=1
		for i=2;i<math.Sqrt(j);i++{
			if int(j)%int(i)==0{
			flag=0}
		}
		if flag==1{
			fmt.Printf("%d is a primenumber\n",int(j))
		}
	}
}
func main()  {
      var a=123456.0
      go primeNumber(a)
      time.Sleep(5 *time.Second)
}






