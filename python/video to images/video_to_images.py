#!/bin/bash

import cv2
import os
import time


folder_name = 'LAB. CONTROL BIOLÓGICO'
path = './labs/'+folder_name+'/'
file_path = path + 'MVI_7031' + '.MOV'

cam = cv2.VideoCapture(file_path)

try:
      
    # creating a folder named data
    if not os.path.exists('data'):
        os.makedirs('data')
  
# if not created then raise error
except OSError:
    print ('Error: Creating directory of data')
  
# frame
currentframe = 0
  
while(True):
      
    # reading from frame
    ret,frame = cam.read()
  
    if ret:
        # if video is still left continue creating images
        name = './data/1920/'+ folder_name +' frame_' + str(currentframe) + " " + str(time.time()) + '.jpg'
        name_r = './data/1320/'+ folder_name +' frame_' + str(currentframe) + " " +  str(time.time()) + '.jpg'
        # name_logo = './data/512/'+ folder_name +' frame_' + str(currentframe) + " " +  str(time.time()) + '.jpg'

        # print ('Creating...' + name)

        dim_r = (1320, 583) # django app -> galery
        dim_logo = (512, 512) # logo     

        frame_r = cv2.resize(frame, dim_r, interpolation = cv2.INTER_AREA)
        # frame_logo = cv2.resize(frame, dim_logo, interpolation = cv2.INTER_AREA)

        # writing the extracted images
        cv2.imwrite(name, frame)
        cv2.imwrite(name_r, frame_r)
        # cv2.imwrite(name_logo, frame_logo)
  
        # increasing counter so that it will
        # show how many frames are created
        currentframe += 1
    else:
        break

print("!!!!! Complete !!!!! ")

# Release all space and windows once done
cam.release()
cv2.destroyAllWindows()
