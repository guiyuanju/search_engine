import sys
from collections import defaultdict
from pathlib import Path
from operator import itemgetter, attrgetter

class Item:
    def __init__(self, count=None, lines=None):
        self.count = count or 0
        self.lines = lines or []

def search(dir, kw):
    # {keyword: {filename: Item}}
    m = defaultdict(lambda: defaultdict(Item))

    # collect
    for root, _, files in Path(dir).walk():
        for f in files:
            if Path(f).suffix not in ['.txt', '.html', '.xml']:
                continue
            words = []
            for w in Path(root).joinpath(f).read_text().split():
                words.append(w)
                if len(words) > 10:
                    words = words[1:]
                m[w][f].count += 1
                m[w][f].lines.append(' '.join(words))

    # sort
    sorted_m = {k: sorted(list(v.items()), key=lambda p: p[1].count, reverse=True) for k, v in m.items()}

    for filename, item in sorted_m.get(kw, [("Not Found", Item())]):
        print(f'[{item.count}]    {filename}')
        for line in item.lines:
            print(f"        {line}")

def main():
    args = sys.argv
    if len(args) < 2:
        print("ERROR: Expect searching directory and search keyword.")
    
    search(args[1], args[2])

main()