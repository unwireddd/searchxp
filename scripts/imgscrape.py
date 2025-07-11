from bs4 import BeautifulSoup
import requests
import re
import sys
import os
import http.cookiejar
import json
import urllib.request, urllib.error, urllib.parse

quer = open("/home/metro/searchxp/output.txt", "r")
quer2 = quer.readlines()
quer3 = quer2[0]
quer3 = str(quer3)
print(quer3)

def get_soup(url,header):
    return BeautifulSoup(urllib.request.urlopen(
        urllib.request.Request(url,headers=header)),
        'html.parser')


# okay so this script works as it should but I need to add an additional loop to it so it actually scrapes all the images
# yeah so I think Its gonna be necessary to add more useragents and randomize them
# also I think that bing keeps blocking me from making requests so I may additionally need a vpn to change my ips every while
# next thing I may want to do is to add a copy of the loop that makes it scrape more images      

def bing_image_search(query):
    query= query.split()
    query='+'.join(query)
    url="http://www.bing.com/images/search?q=" + query + "&FORM=HDRSC2"

    #add the directory for your image here
    DIR="Pictures"
    header={'User-Agent':"Mozilla/5.0 (Linux; Linux i686 x86_64; en-US) Gecko/20100101 Firefox/64.7"}
    soup = get_soup(url,header)
    image_result_raw = soup.find("a",{"class":"iusc"})

    m = json.loads(image_result_raw["m"])
    murl, turl = m["murl"],m["turl"]# mobile image, desktop image

    image_name = urllib.parse.urlsplit(murl).path.split("/")[-1]
    return (image_name,murl, turl)



if __name__ == "__main__":
    query = quer3
    results = bing_image_search(query)
    print(results)
    #save the contents to the output file
    with open("/home/metro/searchxp/scripts/images.html", "w") as file:
        #file.write(f'<a href="{results[1]}">Link</a>')
        file.write("AAAAAAAAAAAAAAAAAAAAAA")