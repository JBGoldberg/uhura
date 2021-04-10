from pathlib import Path
import yaml, os, feedparser

root_path = ""

def warrant_path(path):
    p = os.path.expanduser(path)
    if not os.path.isdir(p):
        os.mkdir(p)
    return p

def process_RSS(source):
    print("Processing RSS: " + source["name"])
    rss_path = warrant_path( root_path + "/" + source["name"] )    
    data = feedparser.parse(source["URL"])
    print(data.feed.title)
    
    for entry in data.entries:
        print("    " + entry.title)


if __name__ == "__main__":

    home = str(Path.home())
    stream =  open(home + "/.uhura.yml", "r")
    config  = yaml.safe_load(stream)

    root_path = warrant_path(config["root"])

    for rss in config["sources"]:
        process_RSS(rss)


