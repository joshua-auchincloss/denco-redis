from pathlib import Path
import os

# prefix components:
space =  '    '
branch = '│   '
# pointers:
tee =    '├── '
last =   '└── '

# https://stackoverflow.com/questions/9727673/list-directory-tree-structure-in-python
def tree(dir_path: Path, prefix: str=''):
    """A recursive generator, given a directory Path object
    will yield a visual tree structure line by line
    with each line prefixed by the same characters
    """    
    contents = list(dir_path.iterdir())

    contents = [c for c in contents if not '.git' in c.as_posix()]
    # contents each get pointers that are ├── with a final └── :
    pointers = [tee] * (len(contents) - 1) + [last]
    for pointer, path in zip(pointers, contents):
        yield prefix + pointer + path.name
        if path.is_dir(): # extend the prefix and recurse:
            extension = branch if pointer == tee else space 
            # i.e. space because last, └── , above so no more |
            yield from tree(path, prefix=prefix+extension)


def sync():

    with open('README-master.md') as f: 
        txt = f.read()
    
    with open("docker-compose.yaml") as f:
        compose = f.read()

    lines = tree(Path(os.getcwd()))


    txt = txt.replace("{COMPOSE-FILE}", compose)
    txt = txt.replace("{TREE}", '\n'.join(list(lines)))

    with open("README.md", 'w') as f: 
        f.write(txt)



if __name__ == '__main__':
    sync()