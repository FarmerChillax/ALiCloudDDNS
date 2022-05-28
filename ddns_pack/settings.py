# -*- coding: utf-8 -*-
'''
    :file: settings.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2022/05/28 10:55:52
'''

from environs import Env



env = Env()
env.read_env()

accessKeyId = env.str("ACCESS_KEY_ID")
accessSecret = env.str("ACCESS_SECRET")

DomainName = env.str("DOMAIN_NAME")
RR_list = env.list("RR_LIST", ["@", "www"])

