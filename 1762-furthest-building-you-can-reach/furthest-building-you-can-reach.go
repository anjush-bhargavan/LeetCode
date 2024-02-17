func furthestBuilding(heights []int, bricks int, ladders int) int {
    arr := make([]int,len(heights))

    for i := 0 ; i < len(arr)-1 ; i++ {
        if heights[i] < heights[i+1] {
            arr[i] = heights[i+1] - heights[i]
        }else{
            arr[i] = 0
        }
    }

    h := &MaxHeap{}
    for i := 0 ; i < len(arr) ; i++ {
        if arr[i] != 0 {
            h.Insert(arr[i])
            if bricks-arr[i] >= 0 {
                bricks -=arr[i]
            }else if ladders > 0{
                ladders --
                val := h.Remove()
                bricks += val
                bricks -= arr[i]
            }else{
                return i
            }
        }
    }
    return len(arr)-1

}

type MaxHeap struct {
    arr []int
}

func (h *MaxHeap) Insert(val int) {
    h.arr = append(h.arr, val)
    h.ShiftUp(len(h.arr)-1)
}

func (h *MaxHeap) Remove() int {
    val := h.arr[0]
    h.arr[0],h.arr[len(h.arr)-1] = h.arr[len(h.arr)-1],h.arr[0]
    h.arr = h.arr[:len(h.arr)-1]
    h.ShiftDown(0)
    return val
}

func (h *MaxHeap) ShiftUp(idx int) {
    parent := parentIdx(idx)
    for idx > 0 && h.arr[idx] > h.arr[parent] {
        h.arr[idx],h.arr[parent] = h.arr[parent],h.arr[idx]
        idx = parent
        parent = parentIdx(idx)
    }
}

func (h *MaxHeap) ShiftDown(idx int) {
    endIdx := len(h.arr)-1

    for leftIdx := leftChild(idx) ; leftIdx <= endIdx ; leftIdx = leftChild(idx) {
        rightIdx := rightChild(idx)
        idxToShift := leftIdx
        if rightIdx <= endIdx && h.arr[rightIdx] > h.arr[idxToShift] {
            idxToShift = rightIdx
        } 
        if h.arr[idxToShift] > h.arr[idx] {
            h.arr[idxToShift], h.arr[idx] =h.arr[idx],h.arr[idxToShift] 
            idx = idxToShift
        }else{
            return
        }

    }
}

func parentIdx (i int) int {
    return (i-1)/2
}

func leftChild(i int) int {
    return (2*i) + 1
}
func rightChild(i int) int {
    return (2 * i) + 2
}