# TODO: try again with numpy array
# used the seen set for part 1 and 2,
# but added backtracking in part 2 to get all paths


def load(file_path):
    with open(file_path, 'r') as file:
        return [list(map(int, line.strip())) for line in file]


def main():
    matrix = load("./day10/input")
    ROWS = len(matrix)
    COLS = len(matrix[0])

    def dfs(r, c, prev_height, seen):

        if (
            r < 0 or c < 0 or
            r >= ROWS or c >= COLS or
            prev_height != matrix[r][c] - 1 or
                (r, c) in seen
        ):
            return 0

        if matrix[r][c] == 9:
            return 1

        res = 0
        seen.add((r, c))
        res += dfs(r-1, c, matrix[r][c], seen)
        res += dfs(r, c+1, matrix[r][c], seen)
        res += dfs(r+1, c, matrix[r][c], seen)
        res += dfs(r, c-1, matrix[r][c], seen)
        seen.remove((r, c))
        return res

        # look for 0s
    count = 0
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            if matrix[i][j] == 0:
                count += dfs(i, j, -1, set())

    print(count)


if __name__ == "__main__":
    main()
