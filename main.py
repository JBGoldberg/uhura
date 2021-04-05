from pathlib import Path
import yaml

if __name__ == "__main__":

    home = str(Path.home())
    stream =  open(home + "/.uhura.yml", "r")
    config  = yaml.safe_load(stream)
    
    print(config)