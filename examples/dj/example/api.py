#!/usr/bin/env python
# -*- coding: utf-8 -*-

from rest_framework import (
    viewsets,
    serializers,
)
from example.models import Address
import json


class AddressSerializer(serializers.ModelSerializer):
    info = serializers.SerializerMethodField()
    alt_names = serializers.SerializerMethodField()

    def get_info(self, obj):
        return json.loads(obj.info)

    def get_alt_names(self, obj):
        return obj.alt_names.replace('{', '').replace('}', '').split(',') if obj.alt_names else None

    class Meta:
        model = Address


class AddressView(viewsets.ReadOnlyModelViewSet):
    permission_classes = []
    serializer_class = AddressSerializer
    queryset = Address.objects.all()


from rest_framework import routers
router = routers.DefaultRouter()
router.register(r'address', AddressView, base_name='address')
