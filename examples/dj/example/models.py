#!/usr/bin/env python
# -*- coding: utf-8 -*-


from __future__ import unicode_literals
from django.db import models

null = {'null': True, 'blank': True}


class Address(models.Model):
    name = models.TextField(**null)
    src = models.TextField(**null)
    group_type = models.TextField(**null)
    type = models.TextField(**null)
    alt_names = models.TextField(**null)
    lat = models.FloatField(**null)
    lon = models.FloatField(**null)
    info = models.TextField(**null)
