# Well I think that just scrapping 4get would be lame in my case so Im gonna keep trying to write various scrappers from different search engines
import requests
from bs4 import BeautifulSoup

l=[]
o={}

query = open("/home/metro/searchxp/output.txt", "r")
query = query.readlines()
print(query)
query = query[0]
query2 = query.replace(" ", "+")
print(query2)

for i in range(0,100,10):

    target_url=f"https://www.bing.com/search?q={query2}&rdr=1".format(i+1)

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
with open('/home/metro/searchxp/res_combined.txt.txt', 'w') as f:
    f.write(your_variable)
