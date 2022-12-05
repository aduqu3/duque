#!/bin/bash

import cv2
import numpy as np
import os

folder = './labs/other/'
_img = 'IMG_0753'
jpg = '.JPG'

frame = folder + _img + jpg


try:
      
    # creating a folder named data
    if not os.path.exists('data_img'):
        os.makedirs('data_img')
  
# if not created then raise error
except OSError:
    print ('Error: Creating directory of data')


img = cv2.imread(frame)

name = './data_img/'+ _img + jpg
name_logo = './data_img/'+ _img + '_logo' + jpg
dim_r = (1320, 583) # django app -> galery
dim_logo = (512, 512) # logo


frame_r = cv2.resize(img, dsize=dim_r, interpolation = cv2.INTER_AREA)
cv2.imwrite(name, frame_r) # writing the extracted images

# frame_logo = cv2.resize(img, dsize=dim_logo, interpolation = cv2.INTER_AREA)
# cv2.imwrite(name_logo, frame_logo) # writing the extracted images


print("!!!!! Complete !!!!! ")

# Release all space and windows once done
cv2.destroyAllWindows()
