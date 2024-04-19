func numIslands(grid [][]byte) int {
    g := &graph{island : make(map[[2]int][][]int)}

    for i := 0 ; i < len(grid) ; i++ {

        for j := 0 ; j < len(grid[i]) ; j++ {
            if grid[i][j] == '1' {
                    addVertex(g,i,j)
                    if i > 0 && grid[i-1][j] == '1' {
                        addEdge(g,i,j,i-1,j)
                    }
                       if i < len(grid)-1 && grid[i+1][j] == '1' {
                           addEdge(g,i,j,i+1,j)
                    }
        
            
                    if j > 0 && grid[i][j-1] == '1' {
                       addEdge(g,i,j,i,j-1)
                    }
                       if j < len(grid[i])-1 && grid[i][j+1] == '1' {
                           addEdge(g,i,j,i,j+1)
                    }
            
            }
        }
    }
    count := islands(g)
    return count
    
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

func islands(g *graph) int {
    visited := make(map[[2]int]bool)
    count := 0
    for v,n := range g.island {
        if !visited[v] {
            visited[v] = true
            count++
            stack := [][]int{}
            stack = append(stack,n...)
            for len(stack) > 0 {
                val := stack[0]
                visited[[2]int{val[0],val[1]}] = true
                stack = stack[1:]
                for _,e := range g.island[[2]int{val[0],val[1]}] {
                key := [2]int{e[0], e[1]}
                    if !visited[key] {
                        visited[key] = true
                        stack = append(stack,e)
                    }
                }

            }
        }
    }
    return count
}