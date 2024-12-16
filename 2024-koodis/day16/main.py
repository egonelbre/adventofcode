from estnltk.vabamorf.morf import syllabify_word


def process(fname):
	with open(fname, 'r') as file:
		failed = 0
		for line in file:
			count = 0
			for word in line.strip().split():
				syllables = syllabify_word(word)
				count += len(syllables)
			if count != 16:
				failed += 1
		print(failed)

process("test1.txt")
process("input.txt")