func maxAscendingSum(nums []int) int {
    temp:=0
    for i:=0;i<len(nums);i++{
        sum:=nums[i]
        for j:=i+1;j<len(nums);j++{
            if nums[j]>nums[j-1]{
                sum+=nums[j]
            }else{
                break
            }
        }
        if sum>temp{
            temp=sum
        }
    }
    return temp
}