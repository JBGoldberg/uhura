from gql import gql, Client
from gql.transport.aiohttp import AIOHTTPTransport
from pathlib import Path
import yaml, os, feedparser

# https://pypi.org/project/PyRSS2Gen/

def userspace_path(path):
    p = os.path.expanduser(path)
    if not os.path.isdir(p):
        os.mkdir(p)
    return p

def process_RSS(source):

    if source['name'] is None:
        print("Processing RSS: " + source['id'])
    else:
        print("Processing RSS: " + source['name'])

    data = feedparser.parse(source["url"])
    if source['name'] is None:
        query = load_gql("update-rss")
        params = {
            "id": source['id'],
            "name": data.feed.title,
            "description": data.feed.description
            }
        res = alexandria.execute(query, variable_values=params)
        source['name'] = res['updateRss']['rss']['name']
        print(source['name'])
    
    for entry in data.entries:
        line = entry.published

        if hasattr(entry, 'author'):
            line += " ("+ entry.author +")"
         
        line +=":" + entry.title

        print(line)

def load_config():
    global config  

    home = str(Path.home())
    stream =  open(home + "/.uhura.yml", "r")    
    config = yaml.safe_load(stream)

    config['root'] = userspace_path(config["root"])

def load_gql(tag_name):
    with open('graphql/'+tag_name + '.gql', 'r') as file:
        data = file.read().replace('\n', '')
    return gql(data)


def connect_alexandria():    

    global alexandria

    # Create a GraphQL client using the defined transport
    temp_client = Client(transport=AIOHTTPTransport(url=config['alexandria']['endpoint']), fetch_schema_from_transport=True)

    query = load_gql("authenticate")
    params = {
        "email": config['alexandria']['user'],
        "password": config['alexandria']['pwd']
        }

    token = temp_client.execute(query, variable_values=params)['authenticate']['jwt']

    transport = AIOHTTPTransport(
        url=config['alexandria']['endpoint'],
        headers={'Authorization': 'Bearer ' + token})

    alexandria = Client(transport=transport, fetch_schema_from_transport=True)

def personal_rsses():
    query = load_gql('rsses')
    return alexandria.execute(query)['rsses']['nodes']


if __name__ == "__main__":
    print("Uhura 0.0.1")

    load_config()
    connect_alexandria()

    for rss in personal_rsses():
        process_RSS(rss)



