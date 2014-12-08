#!/usr/bin/env python
# -*- coding: utf-8 -*-

from rest_framework import (
    viewsets,
    serializers,
)
from example.models import Address
import json


class AddressSerializer(serializers.ModelSerializer):
    info = serializers.SerializerMethodField('get_info2')
    alt_names = serializers.SerializerMethodField('get_alt_names2')

    def get_info2(self, obj):
        return json.loads(obj.info)

    def get_alt_names2(self, obj):
        return obj.alt_names.replace('{', '').replace('}', '').split(',')

    class Meta:
        model = Address
        fields = ('id', 'name', 'src', 'group_type', 'type',
                  'alt_names', 'lat', 'lon', 'info')


class AddressView(viewsets.ReadOnlyModelViewSet):
    permission_classes = []
    serializer_class = AddressSerializer
    queryset = Address.objects.all()


from rest_framework import routers
router = routers.DefaultRouter()
router.register(r'address', AddressView, base_name='address')
