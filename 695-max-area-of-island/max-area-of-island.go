func maxAreaOfIsland(grid [][]int) int {
    
    g := &graph{island : make(map[[2]int][][]int)}
    max,flag := 0,true

    for i := 0 ; i < len(grid) ; i++ {

        for j := 0 ; j < len(grid[i]) ; j++ {
            if grid[i][j] == 1 {
                count++
                    addVertex(g,i,j)
                    if i > 0 && grid[i-1][j] == 1 {
                        addEdge(g,i,j,i-1,j)
                    }
                       if i < len(grid)-1 && grid[i+1][j] == 1 {
                           addEdge(g,i,j,i+1,j)
                    }
            
                    if j > 0 && grid[i][j-1] == 1 {
                       addEdge(g,i,j,i,j-1)
                    }
                       if j < len(grid[i])-1 && grid[i][j+1] == 1 {
                           addEdge(g,i,j,i,j+1)
                    }
            
            }else{
                flag = false
                count = 0
            }
        }
    }
    if flag {
        return count
    }
    count = 0
    for k,_ := range g.island {
        visited := make(map[[2]int]bool)
        g.DFS(k,visited)
        if max < count {
            max = count
        }
        count = 0
    }
    return max
}

type graph struct {
    island map[[2]int][][]int
}

func addEdge(g *graph,v1,v2,e1,e2 int) {
    key := [2]int{v1, v2}
    neighborKey := [2]int{e1, e2}
    if _, ok := g.island[key]; !ok {
        g.island[key] = make([][]int, 0)
    }
    g.island[key] = append(g.island[key], []int{neighborKey[0], neighborKey[1]})
}

func addVertex(g *graph,v1,v2 int) {
     key := [2]int{v1, v2}
       if _, ok := g.island[key]; !ok {
        g.island[key] = make([][]int, 0)
    }
}
var count int
func (g *graph) DFS(val [2]int, visited map[[2]int]bool) {
    visited[val] = true
    count++
    for _,v := range g.island[val] {
        arr := [2]int{}
        copy(arr[:],v) 
        if !visited[arr] {
            g.DFS(arr,visited)
        }
    }
}

