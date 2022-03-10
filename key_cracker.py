
import string
import itertools
import math

nChunks = 5
chunkSize = 5

def main():
    """This is a brute force attempt to discover some valid keys."""
    
    acceptedChunks = generateChunks()

    # Generate a serial key from the accepted chunks. There are around 400000^5 permutations.
    # So we break as soon as we have 100 valid keys.
    num = 0
    for serialKey in itertools.product(acceptedChunks, repeat=chunkSize):
        if validate(serialKey):
            num += 1
        if num == 100:
            break
        
def generateChunks():
    """Generates valid chunks."""
    chars = string.ascii_uppercase  # 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
    acceptedChunks = []

    # First generate all chunks that can be a valid chunk.
    # All permutations that exists with repeating characters, with a length of chunkSize.
    # There are about 11,881,376 permutations. 
    for chunk in itertools.product(chars, repeat=chunkSize):
        total = 0
        for i in range(chunkSize - 1):
            total += ord(chunk[i]) - 64
        
        # Save the accepted chunk with it's total sum as a tuple in a list.
        if ord(chunk[chunkSize - 1]) - 64 == total % 26:
            acceptedChunks.append(("".join(chunk), total))
    return acceptedChunks
    

def validate(serialKey):
    """Validates proposed key. serialKey is a list with tuples of (chunk, sum)."""
    # It is basically the source code to validate a serial key.
    # Returns True if it's a valid key and prints the key, and False if it's not.
    total = 0
    # Add the sum to the total from all chunks except the last one.
    for n in range(nChunks - 1):
        total += serialKey[n][1] # tuple example ('ZIPJI', 61)

    # Last chunk, and checksum validation.
    checkSum = serialKey[nChunks - 1][1]
    if checkSum == total % math.pow(26, chunkSize-1):
        # Print the valid key and return.
        print("-".join([n[0] for n in serialKey]))
        return True
    return False

if __name__ == "__main__":
    main()