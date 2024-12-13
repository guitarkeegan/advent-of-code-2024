from pprint import pprint


def load(path):
    with open(path) as file:
        return [list(lines.strip()) for lines in file]


def main():

    matrix = load("./day12/test-input")
    ROWS = len(matrix)
    COLS = len(matrix[0])
    PERIMETER = 0
    AREA = 1

    seen = set()
    # A: [perimiter, area]
    regions = {}

    def dfs(r, c, seen, regions, target):
        # part 1
        if (r < 0 or c < 0 or r >= ROWS or
                c >= COLS or not matrix[r][c].startswith(target[0])):
            regions[target][PERIMETER] += 1
            return

        if (r, c) in seen:
            return

        regions[target][AREA] += 1

        seen.add((r, c))
        dfs(r+1, c, seen, regions, target)
        dfs(r, c+1, seen, regions, target)
        dfs(r-1, c, seen, regions, target)
        dfs(r, c-1, seen, regions, target)

        return

    # part1: when the regions of the same letter AREA
    # are seperated, they need to be counted seperately
    # so I put a random number to the right of the letter
    # and use str.startswith() in the dfs
    variation = 1
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            if (i, j) not in seen:
                target = matrix[i][j]
                if target not in regions:
                    regions[target] = [0, 0]
                else:
                    target += str(variation)
                    regions[target] = [0, 0]
                    variation += 1

                dfs(i, j, seen, regions, target)

    total = 0
    for measurements in regions.values():
        total += measurements[AREA] * measurements[PERIMETER]

    pprint(matrix)
    print(regions)
    print(f"total: {total}")


if __name__ == "__main__":
    main()
