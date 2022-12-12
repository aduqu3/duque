#!/bin/bash

import cv2
import numpy as np
import os

def scale_img(path, image):

    image = image.split(".")

    _img = image[0]
    format = image[1]

    frame = path + _img + "." + format
    print(frame)


    try:        
        # creating a folder named data
        if not os.path.exists('data_img'):
            os.makedirs('data_img')
    # if not created then raise error
    except OSError:
        print ('Error: Creating directory of data')


    img = cv2.imread(frame)

    # print(img)

    # galeria
    name = './data_img/'+ _img + "." + format
    dim_r = (1320, 583) # django app -> galery

    # logo
    name_logo = './data_img/'+ _img + '_logo' + "." + format
    dim_logo = (512, 512) # logo

    # galeria
    frame_r = cv2.resize(img, dsize=dim_r, interpolation = cv2.INTER_AREA)
    cv2.imwrite(name, frame_r) # writing the extracted images

    # logo
    frame_logo = cv2.resize(img, dsize=dim_logo, interpolation = cv2.INTER_AREA)
    cv2.imwrite(name_logo, frame_logo) # writing the extracted images

    # Release all space and windows once done
    cv2.destroyAllWindows()



# get the path or directory
folder_dir = "./labs/other2/"

for img in os.listdir(folder_dir):
    # check if the image ends with png or jpg or jpeg
    if (img.endswith(".png") or img.endswith(".JPG") or img.endswith(".jpeg")):
        # display
        scale_img(folder_dir, img)

print("!!!!! Complete !!!!! ")
