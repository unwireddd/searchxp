from helium import *
from selenium import webdriver
from selenium.webdriver.common.by import By
import time


searchPhrase = open("/home/metro/searchxp/output.txt", "r")
searchPhrase = searchPhrase.readlines()
sPhrase = " ".join(searchPhrase)
print(sPhrase)
#linkPhrase = sPhrase.replace(" ", "+")

#https://yandex.com/images/search?text=wroclaw
link = "https://yandex.com/images/search?text=" + sPhrase
print(link)

helium.get_driver()
driver = start_firefox(link)
time.sleep(10)
#links = driver.find_elements(By.CLASS_NAME, 'result-link')
links = driver.find_elements(By.CLASS_NAME, 'Button_link')
print(links)
for x in links:
    #Elinks_str.append(x.get_attribute("href"))
    print(x.get_attribute("href"))

#start_firefox("google.com", headless=True)

