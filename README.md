Solution writeup for dev.appsec.nu challenge.
=============================================

Approaching the problem
-----------------------

To figure out what I was dealing with, I initially began reading the source code and inserting some comments on my understanding of what it did.

I understood from the regular expression used that a key is in the format of XXXXX-XXXXX-XXXXX-XXXXX-XXXXX using only letters A-Z.

A key is split into chunks separated by '-'. Each chunk is run through a little computation to get a sum by subtracting 64 from each characters ASCII-value. The sum from the first four chunks are added together and the last chunks sum is used as a checksum.

The full validation of a key is in the checksum computation. If the key passes here, it's a valid key.

Solution
--------

Using this understanding and knowing the computations necessary to validate a key, I can write a little python script to generate some valid chunks first, and then see if I can find some combinations of these chunks that can become a valid key.

*I have made comments in the script that explain this process as well.*

    def generateChunks():
        chars = string.ascii_uppercase  # 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
        acceptedChunks = []

        for chunk in itertools.product(chars, repeat=chunkSize):
            total = 0
            for i in range(chunkSize - 1):
                total += ord(chunk[i]) - 64

            if ord(chunk[chunkSize - 1]) - 64 == total % 26:
                acceptedChunks.append(("".join(chunk), total))

        return acceptedChunks

I'm using the `itertools` library to produce the permutations of chunks I need. Each chunk is validated using the computations from the source code and saved as a tuple like `('ABCDE', 15)`.

    def main():
        acceptedChunks = generateChunks()

        num = 0
        for serialKey in itertools.product(acceptedChunks, repeat=nChunks):
            if validate(serialKey):
                num += 1
            if num == 100:
                break

After the chunks are generated, again using the `itertools` package I produce all permutations of the accepted chunks. Since this part produces a lot of elements, I want this to stop as soon as we have a 100 valid keys.

    def validate(serialKey):
        total = 0
        for n in range(nChunks - 1):
            total += serialKey[n][1]

        checkSum = serialKey[nChunks - 1][1]
        if checkSum == total % math.pow(26, chunkSize-1):
            print("-".join([n[0] for n in serialKey]))
            return True
        return False

Here again using the same computations as in the source code to validate a proposed serial key.
If it is valid, it is printed to console, and returns true to the main function. Else it only returns false to main.
