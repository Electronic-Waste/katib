# coding: utf-8

"""
    Katib

    Swagger description for Katib  # noqa: E501

    The version of the OpenAPI document: v1beta1-0.1
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from kubeflow.katib.configuration import Configuration


class V1beta1FeasibleSpace(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'distribution': 'str',
        'list': 'list[str]',
        'max': 'str',
        'min': 'str',
        'step': 'str'
    }

    attribute_map = {
        'distribution': 'distribution',
        'list': 'list',
        'max': 'max',
        'min': 'min',
        'step': 'step'
    }

    def __init__(self, distribution=None, list=None, max=None, min=None, step=None, local_vars_configuration=None):  # noqa: E501
        """V1beta1FeasibleSpace - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._distribution = None
        self._list = None
        self._max = None
        self._min = None
        self._step = None
        self.discriminator = None

        if distribution is not None:
            self.distribution = distribution
        if list is not None:
            self.list = list
        if max is not None:
            self.max = max
        if min is not None:
            self.min = min
        if step is not None:
            self.step = step

    @property
    def distribution(self):
        """Gets the distribution of this V1beta1FeasibleSpace.  # noqa: E501


        :return: The distribution of this V1beta1FeasibleSpace.  # noqa: E501
        :rtype: str
        """
        return self._distribution

    @distribution.setter
    def distribution(self, distribution):
        """Sets the distribution of this V1beta1FeasibleSpace.


        :param distribution: The distribution of this V1beta1FeasibleSpace.  # noqa: E501
        :type: str
        """

        self._distribution = distribution

    @property
    def list(self):
        """Gets the list of this V1beta1FeasibleSpace.  # noqa: E501


        :return: The list of this V1beta1FeasibleSpace.  # noqa: E501
        :rtype: list[str]
        """
        return self._list

    @list.setter
    def list(self, list):
        """Sets the list of this V1beta1FeasibleSpace.


        :param list: The list of this V1beta1FeasibleSpace.  # noqa: E501
        :type: list[str]
        """

        self._list = list

    @property
    def max(self):
        """Gets the max of this V1beta1FeasibleSpace.  # noqa: E501


        :return: The max of this V1beta1FeasibleSpace.  # noqa: E501
        :rtype: str
        """
        return self._max

    @max.setter
    def max(self, max):
        """Sets the max of this V1beta1FeasibleSpace.


        :param max: The max of this V1beta1FeasibleSpace.  # noqa: E501
        :type: str
        """

        self._max = max

    @property
    def min(self):
        """Gets the min of this V1beta1FeasibleSpace.  # noqa: E501


        :return: The min of this V1beta1FeasibleSpace.  # noqa: E501
        :rtype: str
        """
        return self._min

    @min.setter
    def min(self, min):
        """Sets the min of this V1beta1FeasibleSpace.


        :param min: The min of this V1beta1FeasibleSpace.  # noqa: E501
        :type: str
        """

        self._min = min

    @property
    def step(self):
        """Gets the step of this V1beta1FeasibleSpace.  # noqa: E501


        :return: The step of this V1beta1FeasibleSpace.  # noqa: E501
        :rtype: str
        """
        return self._step

    @step.setter
    def step(self, step):
        """Sets the step of this V1beta1FeasibleSpace.


        :param step: The step of this V1beta1FeasibleSpace.  # noqa: E501
        :type: str
        """

        self._step = step

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1beta1FeasibleSpace):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1beta1FeasibleSpace):
            return True

        return self.to_dict() != other.to_dict()
