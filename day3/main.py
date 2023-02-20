#!/usr/bin/env python3


from typing import Iterable, List


def read() -> Iterable[str]:
    with open("input.txt") as f:
        for line in f:
            yield line.strip()


def count_bits(lines: List[str]):

    zeros = [0 for i in range(len(lines[0]))]
    ones = zeros[:]

    for bits in lines:
        for i, bit in enumerate(bits):
            if bit == "0":
                zeros[i] += 1
            else:
                ones[i] += 1

    return zeros, ones


def most_and_least(zeros, ones):
    most = ""
    least = ""

    for i in range(len(zeros)):
        if zeros[i] > ones[i]:
            most += "0"
            least += "1"

        elif zeros[i] < ones[i]:
            most += "1"
            least += "0"

        else:
            raise Exception(f"this is an edge case what can we do ?, {zeros}, {ones}, {zeros[i]}, {ones[i]}")

    return most, least


def filter_by_prefix(lines, prefix):
    i = 0
    while len(lines) > 1 and i < len(prefix):
        keep = []
        for line in lines:
            if line[i] == prefix[i]:
                keep.append(line)

        lines = keep
        i += 1

    return lines


def part_1_power(lines: List[str]):
    zeros, ones = count_bits(lines)
    most, least = most_and_least(zeros, ones)
    gamma, eps = int(f"0b{most}", 2), int(f"0b{least}", 2)
    return gamma * eps


def oxygen(lines):
    i = 0
    for i in range(len(lines[0])):
        zeros, ones = count_bits(lines)
        bit = "0" if zeros[i] > ones[i] else "1"
        lines = [line for line in lines if line[i] == bit]

        if len(lines) == 1:
            return lines[0]


def carbon(lines):
    i = 0
    for i in range(len(lines[0])):
        zeros, ones = count_bits(lines)
        bit = "1" if ones[i] < zeros[i] else "0"
        lines = [line for line in lines if line[i] == bit]

        if len(lines) == 1:
            return lines[0]


def part_2_life_support(lines):

    o2 = oxygen(lines)
    co2 = carbon(lines)

    return int(f"0b{o2}", 2) * int(f"0b{co2}", 2)


test = ["00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"]

print(part_1_power(list(read())))
print(part_2_life_support(test))
print(part_2_life_support(list(read())))
