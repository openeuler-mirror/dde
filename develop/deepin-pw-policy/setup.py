# !/usr/bin/env python
# -*- coding:utf-8 -*-

#from importlib.metadata import entry_points
import setuptools

desktop_path = '/usr/share/applications/'
icon_path = '/usr/share/icons/hicolor/scalable/apps/'

setuptools.setup(
    name='pw-policy',
    version='1.0',
    description='pam config tool',
    author='mh',
    author_email='584406942@qq.com',
    url='https://gitee.com/openeuler/dde/tree/master/develop/deepin-pw-policy',
    packages=['pw'],
    #启用清单文件MANIFEST.in
    include_package_data=True,
    
    #install_requires=['PyQt5'],
    entry_points={
        'console_scripts':[
            'pw_policy = pw.pam_main:run'
        ]
    },

    data_files=[
        (desktop_path,['pw/pw-policy.desktop']),
        (icon_path,['pw/pw-policy.png']),
    ],
)
