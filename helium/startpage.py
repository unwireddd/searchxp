from helium import *
from selenium import webdriver
from selenium.webdriver.common.by import By
import time
from selenium.webdriver.firefox.options import Options
from selenium.webdriver.chrome.options import Options
import ua_generator

#I think it would be the best if I just extracted the href attrib from those links instead
def parse_startpage():
    def write_arrays_to_file(array1, array2, array3, filename):
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
                file.write(f'<p>{array3[i]}</p>\n')
            file.write(f'<form action="/metaNext" method="post"><button>Next page</button></form>')


    searchPhrase = open("/home/metro/searchxp/output.txt", "r")
    searchPhrase = searchPhrase.readlines()
    sPhrase = " ".join(searchPhrase)
    print(sPhrase)

    #with open("/home/metro/searchxp/helium/res_spage.html", "w") as file:
    #    file.write(f'<h1>Search results for {sPhrase}')

    links_str = []
    titles_str = []
    descs_str = []

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
    
    go_to('https://www.startpage.com/sp/search')

    # REMOTE WD STOP

    #helium.get_driver()
    #driver.get('https://www.startpage.com/sp/search')
    #driver = start_firefox('https://www.startpage.com/sp/search')
    #start_firefox("google.com", headless=True)
    write(str(sPhrase), into="Search privately")
    press(ENTER)
    #ok so the timesleep is actually needed for it to work properly
    time.sleep(1)
    #res1 = find_all(S("wgl-title"))
    links = driver.find_elements(By.CLASS_NAME, 'result-link')
    #titles = driver.find_elements
    print(links)
    for link in links:
        print(link.text)
    #wgl-title
    titles = driver.find_elements(By.CLASS_NAME, 'wgl-title')
    for title in titles:
        print(title.text)

    descs = driver.find_elements(By.CLASS_NAME, 'description')
    for desc in descs:
        print(desc.text)
    print("a")

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
parse_startpage()



