# Well I think that just scrapping 4get would be lame in my case so Im gonna keep trying to write various scrappers from different search engines
import sys
sys.path.append("/home/metro/searchxp/internal/data/linux-amd64/requests")
#import requests

#import importlib.util
#import sys
import requests
from bs4 import BeautifulSoup

l=[]
o={}

for i in range(0,100,10):

    target_url=f"https://www.bing.com/search?q=germany&rdr=1".format(i+1)

    print(target_url)

    resp=requests.get(target_url)

    soup = BeautifulSoup(resp.text, 'html.parser')

    completeData = soup.find_all("li",{"class":"b_algo"})

    for i in range(0, len(completeData)):
        o["Title"]=completeData[i].find("a").text
        o["link"]=completeData[i].find("a").get("href")
        o["Description"]=completeData[i].find("div",
        {"class":"b_caption"}).text
        o["Position"]=i+1
        l.append(o)
        o={}

print(l)
