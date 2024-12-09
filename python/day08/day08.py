from pprint import pprint


def load(path):
    with open(path, 'r') as file:
        matrix = [list(line.strip()) for line in file]
        return matrix


# part 1
def release_the_antinodes(antinodes, pos):

    last = pos[-1]
    for i in range(len(pos)-1):
        r_diff, c_diff = pos[i][0] - last[0], pos[i][1] - last[1]
        upper_side = (pos[i][0] + r_diff, pos[i][1] + c_diff)
        lower_side = (last[0] - r_diff, last[1] - c_diff)
        if upper_side not in antinodes:
            antinodes.add(upper_side)
        if lower_side not in antinodes:
            antinodes.add(lower_side)


# part 2
def release_the_antinodes_forever(antinodes, pos, ROWS, COLS):

    last = pos[-1]
    if last not in antinodes:
        antinodes.add(last)
    for i in range(len(pos)-1):
        if pos[i] not in antinodes:
            antinodes.add(pos[i])
        r_diff, c_diff = pos[i][0] - last[0], pos[i][1] - last[1]

        lx, ly = pos[i][0] + r_diff, pos[i][1] + c_diff
        while (lx >= 0 and ly >= 0 and
               lx < ROWS and ly < COLS):
            upper_side = (lx, ly)
            if upper_side not in antinodes:
                antinodes.add(upper_side)
            lx += r_diff
            ly += c_diff
        lx, ly = last[0] - r_diff, last[1] - c_diff
        while (lx >= 0 and ly >= 0 and
               lx < ROWS and ly < COLS):
            lower_side = (lx, ly)
            if lower_side not in antinodes:
                antinodes.add(lower_side)
            lx -= r_diff
            ly -= c_diff


def main():
    matrix = load("./day08/input")
    freq_pos = {}
    antinode_set = set()

    ROWS = len(matrix)
    COLS = len(matrix[1])

    # get all freq, and potential #
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            if matrix[i][j] != ".":
                if matrix[i][j] in freq_pos:
                    freq_pos[matrix[i][j]].append((i, j))
                else:
                    freq_pos[matrix[i][j]] = [(i, j)]
                # create antinodes
                if len(freq_pos[matrix[i][j]]) > 1:
                    # do the thing
                    release_the_antinodes_forever(
                        antinode_set, freq_pos[matrix[i][j]], ROWS, COLS)

    count = 0
    # print(f"freq_pos: {freq_pos}")
    # print(f"antinode_set: {antinode_set}")
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            if (i, j) in antinode_set:
                count += 1
                matrix[i][j] = "#"

    print(count)


if __name__ == "__main__":
    main()
