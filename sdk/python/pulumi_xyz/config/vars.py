# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('xyz')


class _ExportableConfig(types.ModuleType):
    @property
    def itsasecret(self) -> Optional[bool]:
        return __config__.get_bool('itsasecret')

