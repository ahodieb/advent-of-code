import re
with open("data.txt", "r") as file:
    corrupted_memory = file.read()

    corrupted_memory = re.sub(r"don't\(\).*?(?=do\(\)|$)", "", corrupted_memory, flags=re.DOTALL)
    occurences = re.findall("mul\(\d+,\d+\)", corrupted_memory)
    for o in occurences:
        print(o)
    # digits = [re.findall("\d+", occurence) for occurence in occurences]
    
    #print(sum([int(pair[0]) * int(pair[1]) for pair in digits]))

