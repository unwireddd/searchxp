from helium import *
from selenium import webdriver
from selenium.webdriver.common.by import By
import time
import re
import urllib
from selenium.webdriver.chrome.options import Options


def parse_yandex_imgs():
    def write_arrays_to_file(array1, array2, filename):
        with open(filename, 'w') as file:
            file.write(f'<h1>Images for {sPhrase}</h1>')
            file.write(f'<form action="/output" method="post"><input type="text" name="phrase" id="fer">')
            file.write(f'<input type="radio" id="html" name="Engine" value="Startpage">')
            file.write(f'<label for="html">Startpage</label><br>')
            file.write(f'<input type="radio" id="css" name="Engine" value="Yandex">')
            file.write(f'<label for="css">Yandex</label><br>')
            file.write(f'<input type="radio" id="javascript" name="Engine" value="Metasearch">')
            file.write(f'<label for="javascript">Metasearch</label>')
            file.write(f'</form>')

            #file.write(f'<a href="https://localhost:8000/imgdata.html">[Images]</a>')
            for i in range(len(array1)):
                file.write(f'<a href="{array2[i]}"><img src="{array1[i]}"/></a>\n')
                #file.write(f'<p>{array3[i]}</p>\n')

    def extract_image_url(url):
        match = re.search(r'img_url=([^&]+)', url)
        if match:
            return urllib.parse.unquote(match.group(1))
        return None


    thumbnails = []
    links_href_arr = []
    links_extracted = []
    searchPhrase = open("/home/metro/searchxp/output.txt", "r")
    searchPhrase = searchPhrase.readlines()
    sPhrase = " ".join(searchPhrase)
    print(sPhrase)
    #linkPhrase = sPhrase.replace(" ", "+")

    #https://yandex.com/images/search?text=wroclaw
    link = "https://yandex.com/images/search?text=" + sPhrase
    print(link)

    # REMOTE WD START
    options = Options()
    options.add_argument("--headless=new")
    #options.add_argument("--headless")
    driver = webdriver.Remote(
    command_executor='http://localhost:9515',
    options=options
    #options.add_argument("headless=True")
    )
    set_driver(driver) 
    go_to(link)

    # REMOTE WD STOP
    #helium.get_driver()
    #driver = start_firefox(link, headless=True)
    #time.sleep(1)
    #links = driver.find_elements(By.CLASS_NAME, 'result-link')
    links = driver.find_elements(By.CLASS_NAME, 'ImagesContentImage-Image_clickable')
    link_hrefs = driver.find_elements(By.CLASS_NAME, 'ImagesContentImage-Cover')
    #so I can just take those encoded links and then parse them into normal links and use them as hrefs to those thumbnails
    print(links)

    for x in link_hrefs:
        print(x.get_attribute("href"))
        links_href_arr.append(x.get_attribute("href"))
    for x in links_href_arr:
        links_extracted.append(extract_image_url(x))

    print("TESTING!")
    print(links_extracted)
    print("FINISH")
    for x in links:
        #Elinks_str.append(x.get_attribute("href"))
        print(x.get_attribute("src"))
        thumbnails.append(x.get_attribute("src"))

    write_arrays_to_file(thumbnails, links_extracted, "/home/metro/searchxp/helium/res_images.html")

    #start_firefox("google.com", headless=True)
parse_yandex_imgs()

