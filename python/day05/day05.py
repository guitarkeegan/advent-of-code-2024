def load(path):

    adj = {}
    updates = []

    with open(path) as file:
        sections = file.read().strip().split("\n\n")

        orders, lines = sections[0].splitlines(), sections[1].splitlines()

        for order in orders:
            first, second = int(order[:2]), int(order[3:])
            if first not in adj:
                adj[first] = [second]
            else:
                adj[first].append(second)

        updates = [list(map(int, line.split(","))) for line in lines]

    return adj, updates


def update_lst_swap(update_list, idx1, idx2):
    update_list[idx1], update_list[idx2] = update_list[idx2], update_list[idx1]


def idx_dict_swap(idx_dict, n1, n2):
    idx_dict[n1], idx_dict[n2] = idx_dict[n2], idx_dict[n1]


def get_mid_value(update_list) -> int:
    return update_list[len(update_list) // 2]


def validate_position(idx_dict, first, second):
    if second in idx_dict:
        return idx_dict[first] < idx_dict[second]
    return True


def process(adj, update):

    og_update = update[:]

    idx_dict = {val: index for index, val in enumerate(update)}
    for n in update:
        if n in adj:
            for child in adj[n]:
                if not validate_position(idx_dict, n, child):
                    update_lst_swap(update, idx_dict[n], idx_dict[child])
                    idx_dict_swap(idx_dict, n, child)

    # part one..
    # return update if og_update == update else 0

    # sort of a machine learning/brute force approach, but
    # take this output and subtract it by the previous total to get part 2
    if og_update != update:
        return process(adj, update)
    else:
        return get_mid_value(update)


def main():

    adj, updates = load("./day05/input")

    res = 0

    for update in updates:
        res += process(adj, update)

    print(res)


if __name__ == "__main__":
    main()
