#!/bin/bash

if [ ! -f "shape_predictor_5_face_landmarks.dat" ]; then
   wget https://github.com/davisking/dlib-models/raw/master/shape_predictor_5_face_landmarks.dat.bz2
   bunzip2 shape_predictor_5_face_landmarks.dat.bz2
fi

if [ ! -f "shape_predictor_68_face_landmarks.dat" ]; then
   wget https://github.com/davisking/dlib-models/raw/master/shape_predictor_68_face_landmarks.dat.bz2
   bunzip2 shape_predictor_68_face_landmarks.dat.bz2
fi

if [ ! -f "dlib_face_recognition_resnet_model_v1.dat" ]; then
    wget https://github.com/davisking/dlib-models/raw/master/dlib_face_recognition_resnet_model_v1.dat.bz2
    bunzip2 dlib_face_recognition_resnet_model_v1.dat.bz2
fi

if [ ! -f "mmod_human_face_detector.dat" ]; then
   wget https://github.com/davisking/dlib-models/raw/master/mmod_human_face_detector.dat.bz2
   bunzip2 mmod_human_face_detector.dat.bz2
fi

if [ ! -f "dnn_age_predictor_v1.dat" ]; then
   wget https://github.com/davisking/dlib-models/raw/master/age-predictor/dnn_age_predictor_v1.dat.bz2
   bunzip2 dnn_age_predictor_v1.dat.bz2
fi

if [ ! -f "dnn_gender_classifier_v1.dat" ]; then
   wget https://github.com/davisking/dlib-models/raw/master/gender-classifier/dnn_gender_classifier_v1.dat.bz2
   bunzip2 dnn_gender_classifier_v1.dat.bz2
fi
