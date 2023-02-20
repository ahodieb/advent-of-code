#!/usr/bin/env python3


def part_1_navigation():
    h = 0
    d = 0

    with open("input.txt") as f:
        for line in f:
            direction, value = line.split()
            value = int(value)

            if direction == "forward":
                h += value
            if direction == "down":
                d += value
            if direction == "up":
                d -= value

    return h * d


def part_2_navigation():
    h = 0
    d = 0
    aim = 0

    with open("input.txt") as f:
        for line in f:
            direction, value = line.split()
            value = int(value)

            if direction == "down":
                aim += value
            if direction == "up":
                aim -= value
            if direction == "forward":
                h += value
                d += aim * value

    return h * d


print(part_1_navigation())
print(part_2_navigation())
