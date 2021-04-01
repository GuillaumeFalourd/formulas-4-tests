#!/usr/bin/env python
# coding: utf-8
from youtube_transcript_api import YouTubeTranscriptApi
from sumy.parsers.html import HtmlParser
from sumy.parsers.plaintext import PlaintextParser
from sumy.nlp.tokenizers import Tokenizer
from sumy.summarizers.lsa import LsaSummarizer as Summarizer
from sumy.nlp.stemmers import Stemmer
from sumy.utils import get_stop_words
import os

LANGUAGE = "english"
SENTENCES_COUNT = 10

def run():
    srt = YouTubeTranscriptApi.get_transcript('vN7pdC-A_nQ') # Youtube Video reference
    with open('srt_file.txt', 'w') as f:
        f.writelines('%s\n' % i for i in srt)
        pwd = os.environ.get("CURRENT_PWD")
        generate_summary(pwd + "/srt_file.txt")

def generate_summary(file):
    parser = PlaintextParser.from_file(file, Tokenizer(LANGUAGE))
    stemmer = Stemmer(LANGUAGE)

    summarizer = Summarizer(stemmer)
    summarizer.stop_words = get_stop_words(LANGUAGE)

    for sentence in summarizer(parser.document, SENTENCES_COUNT):
        print(sentence)