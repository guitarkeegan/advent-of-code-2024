from pprint import pprint


class Region:
    def __init__(self):
        # this doesn't quite work because the sides must be contiguous
        self.up_sides = []
        self.right_sides = []
        self.down_sides = []
        self.left_sides = []
        self.area = 0
        self.contiguous_sides = 0
        self.side_counts = {"left": 0, "right": 0, "up": 0, "down": 0}

    def __repr__(self):
        return (f'''
        up: {self.side_counts["up"]} -> {self.up_sides}
        right: {self.side_counts["right"]} -> {self.right_sides},
        down: {self.side_counts["down"]} -> {self.down_sides}
        left: {self.side_counts["left"]} -> {self.left_sides},
        total_sides: {self.contiguous_sides},
        area: {self.area}, calulated area * sides = {self.calculate()}
        ''')

    def calculate(self):
        # clear before calculation
        self.contiguous_sides = 0
        self._calculate_sides()
        return self.contiguous_sides * self.area

    def _sort_up_down(self):
        self.up_sides.sort(key=lambda x: (x[0], x[1]))
        self.down_sides.sort(key=lambda x: (x[0], x[1]))

    def _sort_left_right(self):
        self.left_sides.sort(key=lambda x: (x[1], x[0]))
        self.right_sides.sort(key=lambda x: (x[1], x[0]))

    def _calculate_sides(self):
        self._sort_left_right()
        self._sort_up_down()
        self._calculate_side("up")
        self._calculate_side("right")
        self._calculate_side("down")
        self._calculate_side("left")

    def _calculate_side(self, side):
        if side == "up":
            for i, tup in enumerate(self.up_sides):
                if i == 0:
                    self.contiguous_sides += 1
                    self.side_counts[side] += 1
                    continue
                if (self.up_sides[i-1][1] == tup[1] - 1
                        and self.up_sides[i-1][0] == tup[0]):
                    continue
                else:
                    self.contiguous_sides += 1
                    self.side_counts[side] += 1
        elif side == "down":
            for i, tup in enumerate(self.down_sides):
                if i == 0:
                    self.contiguous_sides += 1
                    self.side_counts[side] += 1
                    continue
                if (self.down_sides[i-1][1] == tup[1] - 1
                        and self.down_sides[i-1][0] == tup[0]):
                    continue
                else:
                    self.contiguous_sides += 1
                    self.side_counts[side] += 1
        elif side == "left":
            for i, tup in enumerate(self.left_sides):
                if i == 0:
                    self.contiguous_sides += 1
                    self.side_counts[side] += 1
                    continue
                if (self.left_sides[i-1][0] == tup[0] - 1
                        and self.left_sides[i-1][1] == tup[1]):
                    continue
                else:
                    self.contiguous_sides += 1
                    self.side_counts[side] += 1
        elif side == "right":
            for i, tup in enumerate(self.right_sides):
                if i == 0:
                    self.contiguous_sides += 1
                    self.side_counts[side] += 1
                    continue
                if (self.right_sides[i-1][0] == tup[0] -
                        1 and self.right_sides[i-1][1] == tup[1]):
                    continue
                else:
                    self.contiguous_sides += 1
                    self.side_counts[side] += 1
        else:
            print("this shouldn't happen")

    def insert(self, side, point):
        if side == "up":
            self.up_sides.append(point)
        elif side == "right":
            self.right_sides.append(point)
        elif side == "down":
            self.down_sides.append(point)
        elif side == "left":
            self.left_sides.append(point)
        else:
            print("invalid side")
            return


def load(path):
    with open(path) as file:
        return [list(lines.strip()) for lines in file]


def main():

    matrix = load("./day12/input")
    ROWS = len(matrix)
    COLS = len(matrix[0])
    UP = "up"
    RIGHT = "right"
    DOWN = "down"
    LEFT = "left"

    seen = set()
    # A: [perimiter, area]
    regions = {}

    def dfs(r, c, seen, regions, target, direction):

        if (r < 0 or c < 0 or r >= ROWS or
                c >= COLS or not matrix[r][c].startswith(target[0])):
            regions[target].insert(direction, (r, c))

            return

        if (r, c) in seen:
            return

        regions[target].area += 1

        seen.add((r, c))
        dfs(r-1, c, seen, regions, target, UP)
        dfs(r, c+1, seen, regions, target, RIGHT)
        dfs(r+1, c, seen, regions, target, DOWN)
        dfs(r, c-1, seen, regions, target, LEFT)

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
                    regions[target] = Region()
                else:
                    target += str(variation)
                    regions[target] = Region()
                    variation += 1

                dfs(i, j, seen, regions, target, UP)

    total = 0
    for key, region in regions.items():
        region_total = region.calculate()
        # print(f"{key}: ", region)
        total += region_total

    # pprint(matrix)
    print(f"total: {total}")


if __name__ == "__main__":
    main()
