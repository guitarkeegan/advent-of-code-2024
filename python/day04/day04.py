class Trie:
    def __init__(self, children=None, end=False):
        self.children = children if children is not None else {}
        self.end = end

    def search_word(self, word) -> bool:
        cur = self
        for letter in word:
            if letter not in cur.children:
                return False
            cur = cur.children[letter]
        return cur.end

    def starts_with(self, letters) -> bool:
        cur = self
        for letter in letters:
            if letter not in cur.children:
                return False
            cur = cur.children[letter]
        return True

    def add(self, word):
        cur = self
        for letter in word:
            if letter not in cur.children:
                cur.children[letter] = Trie()
            cur = cur.children[letter]
        cur.end = True


def main():

    #    DIRS = [
    #        [-1, 0],
    #        [-1, 1],
    #        [0, 1],
    #        [1, 1],
    #        [1, 0],
    #        [1, -1],
    #        [0, -1],
    #        [-1, -1]
    #    ]

    DIRS = [
        [-1, 1],
        [1, 1],
        [1, -1],
        [-1, -1]
    ]

    with open("./day04/test-data") as file:
        lines = file.readlines()

        ROWS = len(lines)
        COLS = len(lines[0])
        bag = set()
        trie = Trie()
        trie.add("MAS")
        trie.add("SAM")

        def search(r, c, dir_idx, cur_word):
            nxt_r = r + dir_idx[0]
            nxt_c = c + dir_idx[1]

            if trie.search_word(cur_word):
                return True
            if (r >= ROWS or r < 0
                or c >= COLS or c < 0
                or nxt_r >= ROWS or nxt_c >= COLS
                or nxt_r < 0 or nxt_c < 0
                    or not trie.starts_with(cur_word)):
                return False
            # call again for current direction
            cur_word += lines[nxt_r][nxt_c]
            return search(nxt_r, nxt_c, dir_idx, cur_word)

        def get_coords(r, c, dir) -> set:
            tmp = set()
            tmp.add((r, c))
            for _ in range(3):
                nxt_r = r + dir[0]
                nxt_c = c + dir[1]
                tmp.add((nxt_r, nxt_c))
            return tmp

        def direct(r, c, cur_word):
            for dir in DIRS:
                if search(r, c, dir, cur_word):
                    line = get_coords(r, c, dir)
                    f_line = frozenset(line)
                    if f_line not in bag:
                        bag.add(f_line)

        for i in range(len(lines)):
            for j in range(len(lines[0])):
                if lines[i][j] == "X" or lines[i][j] == "S":
                    direct(i, j, lines[i][j])

        # print(len(bag)/2)
        print(len(bag))


if __name__ == "__main__":
    main()
