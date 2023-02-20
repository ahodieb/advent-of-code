#!/usr/bin/env python3


def part_1_count_increases():
    increases = 0
    with open("day1.txt") as f:
        previous_value = int(f.readline())

        for line in f:
            value = int(line)
            if value > previous_value:
                increases += 1

            previous_value = value

    return increases


def part_2_count_window_increases():
    window_size = 3
    increases = 0

    with open("day1.txt") as f:

        previous_value = int(f.readline())
        window = [previous_value]

        for i in range(window_size - 1):
            window.insert(0, int(f.readline()))

        for line in f:
            value = int(line)
            edge = window.pop()
            w = sum(window)

            if value + w > edge + w:
                increases += 1

            window.insert(0, value)

    return increases


print(part_1_count_increases())
print(part_2_count_window_increases())
