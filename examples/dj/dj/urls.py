#!/usr/bin/env python
# -*- coding: utf-8 -*-


from django.conf.urls import patterns, include, url
from django.contrib import admin
admin.autodiscover()

urlpatterns = patterns(
    '',
    # url(r'^$', 'dj.views.home', name='home'),
    # url(r'^dj/', include('dj.foo.urls')),
    url(r'^admin/', include(admin.site.urls)),
    url(r'^', include('example.urls')),
)
