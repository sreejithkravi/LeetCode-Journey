func getLongestSubsequence(words []string, groups []int) []string {
    var arr []string
    temp:=groups[0]
    arr=append(arr,words[0])
    for i:=1;i<len(words);i++{
        if temp!=groups[i]{
            arr=append(arr,words[i])
            temp=groups[i]
        }
    }
    fmt.Println(arr)
    return arr
}