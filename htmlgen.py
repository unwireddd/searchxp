import os
import glob
#from bing_image_downloader import downloader


quer = open("/home/metro/searchxp/output.txt", "r")
quer2 = quer.readlines()
quer3 = quer2[0]
quer3 = str(quer3)
#directory = "/home/metro/searchxp/dataset/" + quer3
directory = "dataset/" + quer3
#dirlink = directory + "/"
dirlink = quer3 + "/"
#downloader.download(quer3, limit=100,  output_dir='dataset', adult_filter_off=True, force_replace=False, timeout=60, verbose=True)

def generate_html(image_folder, output_file):
    """
    Generate an HTML page displaying all images in the specified folder.

    Args:
        image_folder (str): Path to the folder containing images.
        output_file (str): Path to the output HTML file.
    """
    # Get a list of image files in the folder
    image_extensions = ['.jpg', '.jpeg', '.png', '.gif', '.bmp']
    image_files = [os.path.join(image_folder, f) for f in os.listdir(image_folder) if os.path.splitext(f)[1].lower() in image_extensions]

    # Create the HTML content
    html_content = '''
    <!DOCTYPE html>
    <html>
    <head>
        <title>Image Gallery</title>
        <style>
            body {
                font-family: Arial, sans-serif;
            }
            .image-gallery {
                display: flex;
                flex-wrap: wrap;
                justify-content: center;
            }
            .image-gallery img {
                width: 200px;
                height: 150px;
                margin: 10px;
                border-radius: 10px;
                box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
            }
        </style>
    </head>
    <body>
        <h1>Image Gallery</h1>
        <p>test</p>
        <div class="image-gallery">
    '''

    # Add image tags to the HTML content
    for image_file in image_files:
        image_name = dirlink + os.path.basename(image_file)
        html_content += f'<img src="{image_name}" alt="{image_name}">\n'

    # Close the HTML content
    html_content += '''
        </div>
    </body>
    </html>
    '''

    # Write the HTML content to the output file

    print(html_content)
    with open("/home/metro/searchxp/dataset/imgdata.html", "w") as file:
        file.write(html_content)
# Example usage
image_folder = directory
print("AAAAA")
print(image_folder)
print("AAA")
output_file = '/home/metro/searchxp/dataset/imgdata.html'
generate_html(image_folder, output_file)