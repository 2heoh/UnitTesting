#!/usr/bin/env python
# -*- coding: utf-8 -*-
# http://www.django-rest-framework.org/api-guide/testing/#apiclient


from django.test import TestCase
from django_any import any_model
from example.models import Address


class ModelTest(TestCase):

    def test_address_creation(self):
        address = any_model(Address)
        self.assertTrue(isinstance(address, Address))
        self.assertEqual(str(address), address.name)
