func threeSum(nums []int) [][]int {
    result := [][]int{}
    sort.Ints(nums)
    for i := 0 ; i < len(nums) ; i++ {
    j :=  i + 1
    k := len(nums) - 1
        for i < len(nums)-1 && nums[i] == nums[i+1] {
            i++
        }
        for j < k {
            v :=  nums[i] + nums[j] + nums[k]  
            if v == 0  {
            
                temp := []int{nums[i],nums[j],nums[k]}
                result = append(result,temp)

                for j < k && nums[j] == nums[j+1] {
                    j++
                }
                for j < k && nums[k] == nums[k-1] {
                    k--
                }
                j++
                k--
            }else if v < 0 {
                j++
            }else{
                k-- 
            }
        }

    }
    return result  
}



                // pairMap := make(map[int]bool)
                // pairMap[nums[i]] = true
                // pairMap[nums[j]] = true
                // pairMap[nums[k]] = true
                // jsonData,_ := json.Marshal(pairMap)
                // jsonString := string(jsonData)
                // if !tripletMap[jsonString] {
                //     tripletMap[jsonString] = true 
                // }else{
                //     break
                // }


        // for j := i+1 ; j < len(nums) ; j++ {
        //     v := 0 - (nums[i]+nums[j])
        //     if numMap[v] {
        //         pairMap := make(map[int]bool)
        //         pairMap[nums[i]] = true
        //         pairMap[nums[j]] = true
        //         pairMap[v] = true
        //         jsonData,_ := json.Marshal(pairMap)
        //         jsonString := string(jsonData)
        //         if !tripletMap[jsonString] {
        //             tripletMap[jsonString] = true 
        //             temp := []int{nums[i],nums[j],v}
        //             result = append(result,temp)
        //         }else{
        //             break
        //         }
        //     }
        // } 