#!/usr/bin/env python
# -*- coding: utf-8 -*-
# http://www.django-rest-framework.org/api-guide/testing/#apiclient


from django.test import TestCase
from django.contrib.auth.models import User
from rest_framework.reverse import reverse
from rest_framework.test import APIClient
from django_any import any_model
from example.models import Address
from django_nose.tools import assert_code


class TestAPIViews(TestCase):
    def setUp(self):
        self.client = APIClient()
        self.user = any_model(User, name='testuser', password='testing')
        kwargs = {
            'info': {},
            'type': 'country',
        }
        [any_model(Address, id=i, **kwargs) for i in range(0, 10)]

    def test_address_api_is_working(self):
        response = self.client.get(reverse('address-list'))
        self.assertEqual(response.status_code, 200, "get address list failed")

    def test_address_api_post(self):
        response = self.client.post(reverse('address-list'))
        assert_code(response, 405,
                    "post address failed (address is not read only)")

    def test_address_is_country(self):
        response = self.client.get(reverse('address-list'), {'pk': 1})
        self.assertIn('country', response.content, response.content)
