from helium import *
from selenium import webdriver
from selenium.webdriver.common.by import By
import time
from selenium.webdriver.chrome.options import Options
import ua_generator


def parse_yandex():
    def write_arrays_to_file(array1, array2, array3, filename):
        iter = 1
        with open(filename, 'w') as file:
            file.write(f'<h1>Search results for {sPhrase}</h1>')
            file.write(f'<form action="/output" method="post"><input type="text" name="phrase" id="fer">')
            file.write(f'<input type="radio" id="html" name="Engine" value="Startpage">')
            file.write(f'<label for="html">Startpage</label><br>')
            file.write(f'<input type="radio" id="css" name="Engine" value="Yandex">')
            file.write(f'<label for="css">Yandex</label><br>')
            file.write(f'<input type="radio" id="javascript" name="Engine" value="Metasearch">')
            file.write(f'<label for="javascript">Metasearch</label>')
            file.write(f'</form>')
            #file.write(f'<form action="/displayImages" method="post"><input type="text" name="phrase" id="fer"></form>')
            file.write(f'<form action="/displayImages" method="post"><button>Display images</button></form>')
            #file.write(f'<a href="https://localhost:6060/displayImages">[Images]</a>')
            for i in range(len(array1)):
                file.write(f'<a href="{array1[i]}">{array2[i]}</a>\n')
                if iter <= len(array3):
                    file.write(f'<p>{array3[i]}</p>\n')
                iter += 1
            file.write(f'<form action="/yandexNext" method="post"><button>Next page</button></form>')


    searchPhrase = open("/home/metro/searchxp/output.txt", "r")
    searchPhrase = searchPhrase.readlines()
    sPhrase = " ".join(searchPhrase)
    print(sPhrase)

    links_str = []
    titles_str = []
    descs_str = []

    #helium.get_driver()
    #driver = start_firefox('https://yandex.com/', headless=True)
    ua = ua_generator.generate()
    options = Options()
    options.add_argument(f'user-agent={ua}')
    #options.add_argument("--headless=new")
    #options.add_argument("--headless")
    driver = webdriver.Remote(
    command_executor='http://localhost:9515',
    options=options
    #options.add_argument("headless=True")
    )
    set_driver(driver) 
    go_to('https://yandex.com/')
    #start_firefox("google.com", headless=True)
    searchBox = driver.find_element(By.CLASS_NAME, "search3__label")
    searchBox.send_keys(str(sPhrase))
    press(ENTER)
    time.sleep(1)
    #res1 = find_all(S("wgl-title"))
    titles = driver.find_elements(By.CLASS_NAME, 'OrganicTitleContentSpan')

    #ok so the class link is what I was looking for but I need to extract the href attribute from it
    links = driver.find_elements(By.CLASS_NAME, 'OrganicTitle-Link')

    descs = driver.find_elements(By.CLASS_NAME, 'OrganicTextContentSpan')

    #for x in links:
    #    print(x.get_attribute("href"))

    for x in links:
        links_str.append(x.get_attribute("href"))
    for x in titles:
        titles_str.append(x.text)
    for x in descs:
        descs_str.append(x.text)

    print(links_str)
    print(titles_str)
    print(descs_str)

    write_arrays_to_file(links_str, titles_str, descs_str, '/home/metro/searchxp/helium/res_spage.html')
parse_yandex()