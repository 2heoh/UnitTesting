#!/usr/bin/env python
# -*- coding: utf-8 -*-


from django.conf.urls import patterns, url, include
from example.api import router


urlpatterns = patterns(
    '',
    url(r'^api/', include(router.urls)),
)
