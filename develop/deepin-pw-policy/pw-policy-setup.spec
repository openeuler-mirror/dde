Name:               pw-policy
Version:            1.0
Release:            1
Summary:            A pam config tool
License:            MulanPSL-2.0
URL:                https://gitee.com/openeuler/dde/tree/master/develop/deepin-pw-policy
Source0:            %{name}-%{version}.tar.gz
BuildArch:          noarch

BuildRequires:      python3-devel python3-setuptools
Requires:           python3 python3-PyQt5-base


%description
A pam config tool

%prep
%autosetup  -p1 -n %{name}-%{version}

%build
%py3_build


%install
%py3_install


%files
%{_bindir}/pw_policy
%{_datadir}/icons/hicolor/scalable/apps/pw-policy.png
%{_datadir}/applications/pw-policy.desktop
%{python3_sitelib}/pw/
%{python3_sitelib}/pw_policy-*.egg-info/


%changelog
* Mon Sep 26 2022 Miao Hao <584406942@qq.com> - 1.0
- Package init
