import argparse

class Combination:
    def __init__(self,line):
        self.line = line
        self.get_order()
        self.get_unknown()

    def get_order(self):
        line = self.line
        order = []
        values = line.split(',')
        first = values[0].split(' ')[-1]
        order.append(int(first))
        for val in values[1:]:
            order.append(int(val))
        self.order = order

    def get_unkown(self):
        line = self.line
        coords = []
        for i in range(len(line)):
            char = line[i]
            if char == '?':
                coords.append(i)
        self.unkown = coords






class Combinations:
    def __init__(self,file_name):
        self.read_file(file_name)
        self.get_combination()
        
    def read_file(self,file_name):
        lines = []
        print(file_name)
        with open(file_name, 'r') as f:
            ls = f.readlines()
            for l in ls:
                lines.append(l.strip('\n'))
            f.close()
        self.lines = lines
    def get_combination(self):
        lines = self.lines
        lineN = 0
        combinations = {}
        for line in lines:
            combinations[lineN] = Combination(line)
        self.combinations = combinations



if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('file_name')
    args = parser.parse_args()
    file_name = args.file_name
    combo = Combinations(file_name)
