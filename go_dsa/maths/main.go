package main

import (
	"fmt"
	"math"
    "slices"
)

func isArmstrong(n int) bool {
	temp1 := n
    temp2 := n
    digitCount := 0
    for n != 0{
        n = n/10
        digitCount++
    }

    armStrongNum := 0.0

    for temp1 != 0{
        rem := temp1%10
        armStrongNum = armStrongNum + math.Pow(float64(rem),float64(digitCount))
        temp1 = temp1/10
    }

    return temp2 == int(armStrongNum)

}

func sortZeroOneTwo(nums []int) {
    countZeros := 0
    countOnes := 0
    countTwos := 0

    for i := 0; i < len(nums); i++{
        if(nums[i] == 0){
            countZeros++
        }else if(nums[i] == 1){
            countOnes++
        }else{
            countTwos++
        }
    }

	j := 0
	for j := 0; j < countZeros; {
		nums[j] = 0
		j++
	}
	for j = countZeros; j < countZeros + countOnes;{
		nums[j] = 1
		j++
	}
	for j = countZeros + countOnes; j < len(nums);{
		nums[j] = 2
		j++
	}
}

func leaders(nums []int) []int {
    leaderSlice := []int{}

    for i := 0; i < len(nums) - 1; i++{
        if(nums[i+1] < nums[i]){
            leaderSlice = append(leaderSlice,nums[i])
        }
    }
    leaderSlice = append(leaderSlice,nums[len(nums)-1])

    return leaderSlice
}



func longestConsecutive(nums []int) int {
	slices.Sort(nums)

	var previousDifference,currDifference,count int

	if(len(nums) == 1){
		previousDifference = nums[0]
		currDifference = previousDifference
		count = 1
	}
		
	previousDifference = nums[1] - nums[0]
	currDifference = previousDifference
	count = 1
		

	for i := 1; i < len(nums); i++{
		previousDifference = nums[i] - nums[i-1]
		if(previousDifference == currDifference){
			count++
            currDifference = previousDifference
		}
	}

	return count
}

type ListNode struct {
    data int
    next *ListNode
}

func LLTraversal(head *ListNode) []int {
    resArray := []int{}
    if head == nil{
        return resArray
    }

    current := head
    for current != nil{
        resArray = append(resArray,current.data)
        current = current.next
    }

    return resArray

}

func main(){
	// fmt.Println(isArmstrong(153))
	arr := []int{5,4,3,2,1,0}
    y1 := &ListNode{data: arr[0], next: nil}
	y2 := &ListNode{data: arr[1], next: nil}
	y3 := &ListNode{data: arr[2], next: nil}
	y4 := &ListNode{data: arr[3], next: nil}
    y5 := &ListNode{data: arr[4], next: nil}
    y6 := &ListNode{data: arr[5], next: nil}


    y1.next =y2
    y2.next = y3
    y3.next = y4
    y4.next = y5
    y5.next = y6
    y6.next = nil

    fmt.Println(LLTraversal(y1))

    //res := longestConsecutive(nums)
	//fmt.Println(res)
}