from bing_image_downloader import downloader

quer = open("/home/metro/searchxp/output.txt", "r")
quer2 = quer.readlines()
quer3 = quer2[0]
quer3 = str(quer3)
directory = "/home/metro/searchxp/dataset/" + quer3

downloader.download(quer3, limit=100,  output_dir='dataset', adult_filter_off=True, force_replace=False, timeout=60, verbose=True)
# the library seems to be working as it should so what I need to do now is to simply make a script that displays the downloaded images in the html file
# I might try just moving the html generation thing to another file and then execute it in main

