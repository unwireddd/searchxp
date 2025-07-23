from helium import *
from selenium import webdriver
from selenium.webdriver.common.by import By
import time

helium.get_driver()
driver = start_firefox('https://www.startpage.com/sp/search', headless=True)
#start_firefox("google.com", headless=True)
write("metro", into="Search privately")
press(ENTER)
time.sleep(10)
#res1 = find_all(S("wgl-title"))
titles = driver.find_elements(By.CLASS_NAME, 'result')
#titles = driver.find_elements
print(titles)
for title in titles:
    print(title.text)
print("a")

with open("/home/metro/searchxp/res_startpage.txt", "w") as file:
    # Write each element of the array to the file
    for element in titles:
        file.write(element + "\n")


