import argparse


class Galaxy:
    def __init__(self,x,y,id):
        self.id = id
        self.x = x
        self.y = y

    def __str__(self) -> str:
        return f'id: {self.id}, x: {self.x}, y: {self.y} '
        
    def get_steps(self,other):
        xdiff = abs(self.x - other.x)
        ydiff = abs(self.y - other.y)
        if self == other:
            return 0
        else:
            return ydiff + xdiff +1

    def __eq__(self,other):
        return self.x == other.x and self.y == other.y

class Galaxies:
    def __init__(self,file_name):
        self.read_file(file_name)
        for line in self.lines:
            print(line)
        self.expansion()
        for line in self.lines:
            print(line)
        self.find_galaxies()


    def read_file(self,file_name):
        lines = []
        with open(file_name, 'r') as f:
            ls = f.readlines()
            for l in ls:
                lines.append(l.strip('\n'))
            f.close()
        self.lines = lines

    def expansion(self):
        lines = self.lines[:]
        for i in range(len(lines)-1,-1,-1):
            line = lines[i]
            if '#' in line:
                pass
            else:
                self.lines.insert(i,line)
        indexs = []
        for i in range(len(lines[0])-1,-1,-1):
            if '#' in lines[:][i]:
                pass
            else:
                indexs.append(i)
        for i in range(len(self.lines)):
            line = self.lines[i]
            for index in indexs[::-1]:
                line = line[:index] + '.' + line[index:]
            self.lines[i] = line

            
    def get_steps(self):
        steps = {}
        total_steps = 0
        for i in range(len(self.galaxies.values())):
            galaxy = self.galaxies[i+1]
            step = []
            for j in range(len(self.galaxies.values())):
                other = self.galaxies[j+1]
                step.append(galaxy.get_steps(other))
            total_steps += sum(step)
            steps[galaxy.id] = step
        self.steps = steps
        self.total_steps = total_steps/2


    def find_galaxies(self):
        lines = self.lines
        galaxies = {}
        id_count = 1
        for i in range(len(lines)):
            line = lines[i]
            for j in range(len(line)):
                char = line[j]
                if char == '#':
                    galaxy = Galaxy(j,i,id_count)
                    print(Galaxy.__str__(galaxy))
                    galaxies[id_count] = galaxy
                    id_count += 1
        self.galaxies = galaxies


#check EDM, like conda but different.
if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('file_name')
    args = parser.parse_args()
    file_name = args.file_name
    galaxies = Galaxies(file_name)
    galaxies.get_steps()
    [print(galaxies.steps[i+1]) for i in range(len(galaxies.steps))]
    print(galaxies.total_steps)


