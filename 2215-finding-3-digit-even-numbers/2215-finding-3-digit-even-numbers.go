func findEvenNumbers(digits []int) []int {
    var arr[]int
    for i:=100;i<1000;i++{

        n1:=i%10
        n2:=(i/10)%10
        n3:=i/100
        flag:=0

        for _,num:=range digits{

            if num==n1 && n1!=-1{
                flag++
                n1=-1
            }else if num==n2 && n2!=-1{
                flag++
                n2=-1
            }else if num==n3 && n3!=-1{
                flag++
                n3=-1
            }

        }
        
        if flag==3 && i%2==0{
            arr=append(arr,i)
        }
    }
    return arr
} 