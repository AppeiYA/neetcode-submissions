func searchMatrix(matrix [][]int, target int) bool {
    start, rows, columns := 0, len(matrix), len(matrix[0])
    if rows == 0 {
        return false
    }

    end := (rows * columns) - 1

    for start <= end {
        mid := start + (end - start)/2
        row := mid/columns
        column := mid%columns
        current := matrix[row][column]

        if current == target {
            return true
        } else if current > target {
            end = mid - 1
        }else if current < target {
            start = mid + 1
        }
    }

    return false
}
