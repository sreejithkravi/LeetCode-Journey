func setZeroes(matrix [][]int)  {
    var row[]int
    var col[]int
    for i:=0;i<len(matrix);i++{
        for j:=0;j<len(matrix[0]);j++{
            if matrix[i][j]==0{
                row=append(row,i)
                col=append(col,j)
            }
        }
    }
    for _,rl:=range row{
        for j:=0;j<len(matrix[rl]);j++{
            matrix[rl][j]=0
        }
    }
        for _,cl:=range col{
        for j:=0;j<len(matrix);j++{
            matrix[j][cl]=0
        }
    }
}