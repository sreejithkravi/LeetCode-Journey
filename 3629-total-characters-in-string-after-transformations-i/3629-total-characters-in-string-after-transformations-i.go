// func lengthAfterTransformations(s string, t int) int {
//     arr := make([]int,26)
//     for _,char:=range s{
//         arr[char-'a']++
//     }
//             fmt.Println(arr)
//     for j := 0; j < t; j++ {
//         newArr := make([]int,26)
//         for i:=0;i<26;i++{
//             if i==25{
//                 newArr[0]+=arr[i]
//                 newArr[1]+=arr[i]
//             }else{
//                 newArr[i+1]+=arr[i]
//             }
//         }
//                             fmt.Println(arr)

        
//         arr=newArr

//     }
//     total:=0
//     for _,count:=range arr{
//         total+=count
//     }
//     return total
// }

func lengthAfterTransformations(s string, t int) int {
    mod := int(1e9 + 7)

    arr := make([]int, 26)
    for _, char := range s {
        arr[char-'a']++
    }

    for i := 0; i < t; i++ {
        newArr := make([]int, 26)
        for j := 0; j < 26; j++ {
            if j == 25 {
                newArr[0] = (newArr[0] + arr[25]) % mod
                newArr[1] = (newArr[1] + arr[25]) % mod
            } else {
                newArr[j+1] = (newArr[j+1] + arr[j]) % mod
            }
        }
        arr = newArr
    }

    total := 0
    for _, count := range arr {
        total = (total + count) % mod
    }

    return total
}
