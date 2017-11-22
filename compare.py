#!/usr/bin/python2
#
# A modified version of:
# http://cmusatyalab.github.io/openface/demo-2-comparison/

import sys
import argparse
import cv2
import itertools
import os

import numpy as np
np.set_printoptions(precision=2)

import openface

fileDir = os.path.dirname(os.path.realpath(__file__))
modelDir = os.path.join(fileDir, '../openface', 'models')
dlibModelDir = os.path.join(modelDir, 'dlib')
openfaceModelDir = os.path.join(modelDir, 'openface')

align = openface.AlignDlib(os.path.join(dlibModelDir, "shape_predictor_68_face_landmarks.dat"))
img_dim = 96
net = openface.TorchNeuralNet(os.path.join(openfaceModelDir, 'nn4.small2.v1.t7'), img_dim)

def get_rep(b64_string):
    bgrImg = data_uri_to_cv2_img(b64_string)
    if bgrImg is None:
        raise Exception("Unable to load image: {}".format(imgPath))
    rgbImg = cv2.cvtColor(bgrImg, cv2.COLOR_BGR2RGB)

    bb = align.getLargestFaceBoundingBox(rgbImg)
    if bb is None:
        raise Exception("Unable to find a face: {}".format(imgPath))
    alignedFace = align.align(img_dim, rgbImg, bb,
                              landmarkIndices=openface.AlignDlib.OUTER_EYES_AND_NOSE)
    if alignedFace is None:
        raise Exception("Unable to align image: {}".format(imgPath))
    rep = net.forward(alignedFace)
    return rep

def data_uri_to_cv2_img(b64_string):
    nparr = np.fromstring(b64_string.decode('base64'), np.uint8)
    img = cv2.imdecode(nparr, cv2.IMREAD_COLOR)
    return img

def ensure_padding(b64_string):
    value = b64_string
    if len(value) % 4:
        # not a multiple of 4, add padding:
        value += '=' * (4 - len(value) % 4) 
    return value

def main():
    img1 = raw_input('reading image one from stdin!')
    img2 =  raw_input('reading image two from stdin!')
    img1 = ensure_padding(img1)
    img2 = ensure_padding(img2)
    d = get_rep(img1) - get_rep(img2)
    diff = np.dot(d, d)
    return ((4.0 - diff)/4.0) * 100

if __name__ == "__main__":
    res = main()
    print res
