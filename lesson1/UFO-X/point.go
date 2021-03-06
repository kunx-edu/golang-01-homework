package main
import(
	"fmt"
)
func main(){

	x,y,z:=1,2,3
	swap(&x,&y,&z)
	fmt.Println("x=",x,"y=",y,"z=",z)
}
/*传的参是引用，修改内存地址的值；但为何不是2,2,100？可以理解为，表达式的值已经是定的，所以等号右边的值不会随着等号前变量值变化而变化*/
/*普通传值是copy了一份，实参不会发生变化；大的结构体使用指针，本质上也会减少系统的开销（内存、时间）*/
/*指针的数组、数组的指针、指针的指针；*/
func swap(x,y,z *int){
	*x,*y=*y,*x
	*z=100
}