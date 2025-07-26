from helium import *
from selenium import webdriver
from selenium.webdriver.common.by import By
import time

def write_arrays_to_file(array1, array2, array3, filename):
    with open(filename, 'a') as file:
        for i in range(len(array1)):
            file.write(f'<a href="{array1[i]}">{array2[i]}</a>\n')
            file.write(f'<p>{array3[i]}</p>\n')


searchPhrase = open("/home/metro/searchxp/output.txt", "r")
searchPhrase = searchPhrase.readlines()
sPhrase = " ".join(searchPhrase)
print(sPhrase)

links_str = []
titles_str = []
descs_str = []

helium.get_driver()
driver = start_firefox('https://yandex.com/', headless=True)
#start_firefox("google.com", headless=True)
write(str(sPhrase), into="Search with Yandex AI")
press(ENTER)
time.sleep(5)
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